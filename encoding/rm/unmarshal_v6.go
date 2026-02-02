// V6 format decoder for reMarkable software version 3+.
// This is a scene-based format that is completely different from the
// V3/V5 binary stroke format.
//
// Based on the rmscene Python library by Rick Lupton:
// https://github.com/ricklupton/rmscene
package rm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

// Block types (based on rmscene Python library)
const (
	BlockTypeMigrationInfo      = 0x00
	BlockTypeSceneTree          = 0x01
	BlockTypeTreeNode           = 0x02
	BlockTypeSceneGlyphItem     = 0x03 // Glyph/text highlight regions
	BlockTypeSceneGroupItem     = 0x04 // Group nodes in scene tree
	BlockTypeSceneLineItem      = 0x05 // Stroke/line data
	BlockTypeSceneTextItem      = 0x06 // Text item (different from RootText)
	BlockTypeRootText           = 0x07 // Root text block
	BlockTypeSceneTombstoneItem = 0x08 // Deleted items
	BlockTypeAuthorIds          = 0x09
	BlockTypePageInfo           = 0x0A
	BlockTypeSceneInfo          = 0x0D
)

// Tag types for tagged data (rmscene format)
const (
	TagByte1   = 0x1 // 1 byte (uint8, bool)
	TagByte4   = 0x4 // 4 bytes (uint32, float32)
	TagByte8   = 0x8 // 8 bytes (float64)
	TagLength4 = 0xC // subblock with uint32 length prefix
	TagID      = 0xF // CRDT ID (uint8 + varuint)
)

// Pen tool types
const (
	PenBrush            = 0
	PenTiltPencil       = 1
	PenBallpoint        = 2
	PenMarker           = 3
	PenFineliner        = 4
	PenHighlighter      = 5
	PenEraser           = 6
	PenSharpPencil      = 7
	PenEraseArea        = 8
	PenBrushV5          = 12
	PenMechanicalPencil = 13
	PenPencil           = 14
	PenBallpointV5      = 15
	PenMarkerV5         = 16
	PenFinelinerV5      = 17
	PenHighlighterV5    = 18
	PenCalligraphy      = 21
	PenShader           = 23
)

// PenColor types
const (
	ColorBlack       = 0
	ColorGrey        = 1
	ColorWhite       = 2
	ColorYellow      = 3
	ColorGreen       = 4
	ColorPink        = 5
	ColorBlue        = 6
	ColorRed         = 7
	ColorGreyOverlap = 8
	ColorHighlight   = 9
	ColorGreen2      = 10
	ColorCyan        = 11
	ColorMagenta     = 12
	ColorYellow2     = 13
)

// V6Reader wraps a bytes.Reader and provides methods to read V6 format data.
type V6Reader struct {
	r *bytes.Reader
}

// NewV6Reader creates a new V6 reader
func NewV6Reader(data []byte) *V6Reader {
	return &V6Reader{r: bytes.NewReader(data)}
}

// Remaining returns the number of bytes remaining
func (r *V6Reader) Remaining() int {
	return r.r.Len()
}

// ReadBytes reads n bytes
func (r *V6Reader) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r.r, buf)
	return buf, err
}

// ReadVarUint reads a variable-length unsigned integer
func (r *V6Reader) ReadVarUint() (uint64, error) {
	var result uint64
	var shift uint
	for {
		b, err := r.r.ReadByte()
		if err != nil {
			return 0, err
		}
		result |= uint64(b&0x7F) << shift
		if b&0x80 == 0 {
			break
		}
		shift += 7
	}
	return result, nil
}

// ReadUint8 reads a uint8
func (r *V6Reader) ReadUint8() (uint8, error) {
	return r.r.ReadByte()
}

// ReadUint16 reads a uint16
func (r *V6Reader) ReadUint16() (uint16, error) {
	var v uint16
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadUint32 reads a uint32
func (r *V6Reader) ReadUint32() (uint32, error) {
	var v uint32
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadFloat32 reads a float32
func (r *V6Reader) ReadFloat32() (float32, error) {
	var v float32
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadFloat64 reads a float64
func (r *V6Reader) ReadFloat64() (float64, error) {
	var v float64
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadCrdtID reads a CRDT ID (uint8 + varuint)
func (r *V6Reader) ReadCrdtID() error {
	_, err := r.ReadUint8()
	if err != nil {
		return err
	}
	_, err = r.ReadVarUint()
	return err
}

// ReadTag reads a tag and returns the index and tag type
func (r *V6Reader) ReadTag() (index uint64, tagType uint64, err error) {
	tagValue, err := r.ReadVarUint()
	if err != nil {
		return 0, 0, err
	}
	index = tagValue >> 4
	tagType = tagValue & 0x0F
	return index, tagType, nil
}

// SkipTagValue skips the value of a tag based on its type
func (r *V6Reader) SkipTagValue(tagType uint64) error {
	switch tagType {
	case TagID:
		return r.ReadCrdtID()
	case TagByte4:
		_, err := r.ReadUint32()
		return err
	case TagByte8:
		_, err := r.ReadFloat64()
		return err
	case TagByte1:
		_, err := r.ReadUint8()
		return err
	case TagLength4:
		size, err := r.ReadUint32()
		if err != nil {
			return err
		}
		_, err = r.ReadBytes(int(size))
		return err
	default:
		return fmt.Errorf("unknown tag type: 0x%x", tagType)
	}
}

// UnmarshalV6 reads a V6 .rm file and converts it to the common Rm structure
func UnmarshalV6(data []byte) (*Rm, error) {
	r := NewV6Reader(data)

	// Read and verify header
	header, err := r.ReadBytes(43)
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}
	if string(header) != HeaderV6 {
		return nil, fmt.Errorf("invalid V6 header: %q", string(header))
	}

	// Parse blocks and extract lines and text
	lines, textItems, err := parseV6Blocks(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse blocks: %w", err)
	}

	// Convert to Rm structure
	// V6 doesn't have explicit layers, so we put all lines in one layer
	result := &Rm{
		Version: V6,
		Layers: []Layer{
			{Lines: lines, Text: textItems},
		},
	}

	return result, nil
}

// parseV6Blocks reads all blocks from the stream and extracts line and text items
func parseV6Blocks(r *V6Reader) ([]Line, []TextItem, error) {
	var lines []Line
	var textItems []TextItem

	for r.Remaining() > 0 {
		// Read block header
		blockType, blockData, version, err := readV6BlockHeader(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return lines, textItems, fmt.Errorf("failed to read block header: %w", err)
		}

		// Parse line items
		if blockType == BlockTypeSceneLineItem {
			line, err := parseV6LineBlock(blockData, version)
			if err != nil {
				// Skip failed blocks and continue - partial extraction is better than nothing
				continue
			}
			if len(line.Points) > 0 {
				lines = append(lines, line)
			}
		}

		// Parse root text block
		if blockType == BlockTypeRootText {
			textItem, err := parseV6RootTextBlock(blockData)
			if err != nil {
				// Skip failed blocks and continue
				continue
			}
			if len(textItem.Paragraphs) > 0 {
				textItems = append(textItems, textItem)
			}
		}
	}

	return lines, textItems, nil
}

// readV6BlockHeader reads a block header and returns the block data
// Block format: uint32 length, uint8 unknown, uint8 min_ver, uint8 cur_ver, uint8 block_type, then data
func readV6BlockHeader(r *V6Reader) (blockType byte, data []byte, version byte, err error) {
	// Read block length
	blockLength, err := r.ReadUint32()
	if err != nil {
		return 0, nil, 0, err
	}

	// Read unknown byte (should be 0)
	_, err = r.ReadUint8()
	if err != nil {
		return 0, nil, 0, err
	}

	// Read min version
	_, err = r.ReadUint8()
	if err != nil {
		return 0, nil, 0, err
	}

	// Read current version
	curVer, err := r.ReadUint8()
	if err != nil {
		return 0, nil, 0, err
	}

	// Read block type
	blockType, err = r.ReadUint8()
	if err != nil {
		return 0, nil, 0, err
	}

	// Read block data
	// The block_length field specifies the size of data AFTER the 4-byte header
	// Structure: <4-byte length><4-byte header><length bytes of data>
	dataSize := int(blockLength)
	if dataSize > 0 {
		data, err = r.ReadBytes(dataSize)
		if err != nil {
			return 0, nil, 0, fmt.Errorf("failed to read block data of size %d: %w", dataSize, err)
		}
	}

	return blockType, data, curVer, nil
}

// parseV6LineBlock parses a SceneLineItem block
// SceneItemBlock structure: parent_id(1), item_id(2), left_id(3), right_id(4), deleted_length(5), value subblock(6)
func parseV6LineBlock(data []byte, blockVersion byte) (Line, error) {
	r := NewV6Reader(data)
	var line Line

	// Read tagged fields
	for r.Remaining() > 0 {
		index, tagType, err := r.ReadTag()
		if err == io.EOF {
			break
		}
		if err != nil {
			return line, err
		}

		// We only care about index 6 (value subblock) for lines
		if index == 6 && tagType == TagLength4 {
			// Read subblock size
			subSize, err := r.ReadUint32()
			if err != nil {
				return line, err
			}

			// Read item type (should be 0x03 for Line)
			itemType, err := r.ReadUint8()
			if err != nil {
				return line, err
			}
			if itemType != 0x03 {
				// Not a line item
				return line, fmt.Errorf("unexpected item type: 0x%02x", itemType)
			}

			// Read the actual Line data
			lineData, err := r.ReadBytes(int(subSize) - 1)
			if err != nil {
				return line, err
			}

			// Parse the Line data
			return parseV6LineData(lineData, blockVersion)
		}

		// Skip other fields
		if err := r.SkipTagValue(tagType); err != nil {
			return line, err
		}
	}

	return line, nil
}

// parseV6LineData parses the actual Line data from within the value subblock
func parseV6LineData(data []byte, blockVersion byte) (Line, error) {
	r := NewV6Reader(data)
	var line Line

	// Read tagged fields for the Line
	for r.Remaining() > 0 {
		index, tagType, err := r.ReadTag()
		if err == io.EOF {
			break
		}
		if err != nil {
			return line, err
		}

		switch index {
		case 1: // tool/pen type (uint32)
			if tagType == TagByte4 {
				toolID, err := r.ReadUint32()
				if err != nil {
					return line, err
				}
				line.BrushType = mapPenToV5BrushType(int(toolID))
			} else {
				if err := r.SkipTagValue(tagType); err != nil {
					return line, err
				}
			}
		case 2: // color (uint32)
			if tagType == TagByte4 {
				colorID, err := r.ReadUint32()
				if err != nil {
					return line, err
				}
				line.BrushColor = mapColorToV5(int(colorID))
			} else {
				if err := r.SkipTagValue(tagType); err != nil {
					return line, err
				}
			}
		case 3: // thickness_scale (float64)
			if tagType == TagByte8 {
				scale, err := r.ReadFloat64()
				if err != nil {
					return line, err
				}
				line.BrushSize = BrushSize(scale * 2.0)
			} else {
				if err := r.SkipTagValue(tagType); err != nil {
					return line, err
				}
			}
		case 4: // starting_length (float32) - not used in rendering
			if tagType == TagByte4 {
				_, err := r.ReadFloat32()
				if err != nil {
					return line, err
				}
			} else {
				if err := r.SkipTagValue(tagType); err != nil {
					return line, err
				}
			}
		case 5: // points data (subblock)
			if tagType == TagLength4 {
				subSize, err := r.ReadUint32()
				if err != nil {
					return line, err
				}
				pointData, err := r.ReadBytes(int(subSize))
				if err != nil {
					return line, err
				}
				points, err := parseV6Points(pointData, blockVersion)
				if err != nil {
					return line, err
				}
				line.Points = points
			} else {
				if err := r.SkipTagValue(tagType); err != nil {
					return line, err
				}
			}
		default:
			// Skip unknown fields
			if err := r.SkipTagValue(tagType); err != nil {
				return line, err
			}
		}
	}

	return line, nil
}

// parseV6Points reads point data from byte slice
func parseV6Points(data []byte, version byte) ([]Point, error) {
	pointSize := 14 // Version 2 point size (14 bytes)
	if version == 1 {
		pointSize = 24 // Version 1 point size (24 bytes)
	}

	if len(data)%pointSize != 0 {
		return nil, fmt.Errorf("point data size %d is not a multiple of point size %d", len(data), pointSize)
	}

	numPoints := len(data) / pointSize
	points := make([]Point, numPoints)
	r := NewV6Reader(data)

	for i := 0; i < numPoints; i++ {
		p, err := readV6Point(r, version)
		if err != nil {
			return nil, err
		}
		points[i] = p
	}

	return points, nil
}

// readV6Point reads a single point
func readV6Point(r *V6Reader, version byte) (Point, error) {
	var p Point

	x, err := r.ReadFloat32()
	if err != nil {
		return p, err
	}
	y, err := r.ReadFloat32()
	if err != nil {
		return p, err
	}

	p.X = x
	p.Y = y

	if version == 1 {
		// Version 1: float-based encoding (24 bytes per point)
		speed, err := r.ReadFloat32()
		if err != nil {
			return p, err
		}
		direction, err := r.ReadFloat32()
		if err != nil {
			return p, err
		}
		width, err := r.ReadFloat32()
		if err != nil {
			return p, err
		}
		pressure, err := r.ReadFloat32()
		if err != nil {
			return p, err
		}

		p.Speed = speed * 4
		p.Direction = 255 * direction / (2 * math.Pi)
		p.Width = width * 4
		p.Pressure = pressure * 255
	} else {
		// Version 2: integer-based encoding (14 bytes per point)
		speed, err := r.ReadUint16()
		if err != nil {
			return p, err
		}
		width, err := r.ReadUint16()
		if err != nil {
			return p, err
		}
		direction, err := r.ReadUint8()
		if err != nil {
			return p, err
		}
		pressure, err := r.ReadUint8()
		if err != nil {
			return p, err
		}

		p.Speed = float32(speed)
		p.Width = float32(width)
		p.Direction = float32(direction)
		p.Pressure = float32(pressure)
	}

	return p, nil
}

// mapPenToV5BrushType maps V6 pen types to V5 brush types
func mapPenToV5BrushType(penType int) BrushType {
	switch penType {
	case PenBallpoint, PenBallpointV5:
		return BallPoint
	case PenMarker, PenMarkerV5:
		return Marker
	case PenFineliner, PenFinelinerV5:
		return Fineliner
	case PenSharpPencil, PenMechanicalPencil:
		return SharpPencil
	case PenTiltPencil, PenPencil:
		return TiltPencil
	case PenBrush, PenBrushV5:
		return Brush
	case PenHighlighter, PenHighlighterV5:
		return Highlighter
	case PenEraser:
		return Eraser
	case PenEraseArea:
		return EraseArea
	case PenCalligraphy:
		return Marker
	case PenShader:
		return Brush
	default:
		return BallPoint
	}
}

// mapColorToV5 maps V6 color types to BrushColor
func mapColorToV5(color int) BrushColor {
	switch color {
	case ColorBlack:
		return Black
	case ColorGrey:
		return Grey
	case ColorWhite:
		return White
	case ColorYellow:
		return Yellow
	case ColorGreen:
		return Green
	case ColorPink:
		return Pink
	case ColorBlue:
		return Blue
	case ColorRed:
		return Red
	case ColorGreyOverlap:
		return GreyOverlap
	case ColorHighlight:
		return Highlight
	case ColorGreen2:
		return Green2
	case ColorCyan:
		return Cyan
	case ColorMagenta:
		return Magenta
	case ColorYellow2:
		return Yellow2
	default:
		return Black
	}
}

// parseV6RootTextBlock parses a RootText block and extracts text content
// RootText structure: contains tagged fields for text items and styles
func parseV6RootTextBlock(data []byte) (TextItem, error) {
r := NewV6Reader(data)
var textItem TextItem

// Read tagged fields
for r.Remaining() > 0 {
index, tagType, err := r.ReadTag()
if err == io.EOF {
break
}
if err != nil {
return textItem, err
}

switch index {
case 1: // block_id (CRDT ID) - skip
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
case 2: // pos_x (float64)
if tagType == TagByte8 {
posX, err := r.ReadFloat64()
if err != nil {
return textItem, err
}
textItem.PosX = posX
} else {
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
}
case 3: // pos_y (float64)
if tagType == TagByte8 {
posY, err := r.ReadFloat64()
if err != nil {
return textItem, err
}
textItem.PosY = posY
} else {
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
}
case 4: // width (float64)
if tagType == TagByte8 {
width, err := r.ReadFloat64()
if err != nil {
return textItem, err
}
textItem.Width = width
} else {
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
}
case 5: // text items subblock
if tagType == TagLength4 {
subSize, err := r.ReadUint32()
if err != nil {
return textItem, err
}
textData, err := r.ReadBytes(int(subSize))
if err != nil {
return textItem, err
}
paragraphs := parseV6TextContent(textData)
textItem.Paragraphs = paragraphs
} else {
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
}
case 6: // styles subblock
if tagType == TagLength4 {
subSize, err := r.ReadUint32()
if err != nil {
return textItem, err
}
// Skip styles for now - we extract text content directly
_, err = r.ReadBytes(int(subSize))
if err != nil {
return textItem, err
}
} else {
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
}
default:
if err := r.SkipTagValue(tagType); err != nil {
return textItem, err
}
}
}

return textItem, nil
}

// parseV6TextContent extracts text content from the text items subblock
// This is a simplified parser that extracts readable text
func parseV6TextContent(data []byte) []TextParagraph {
var paragraphs []TextParagraph

// The text content is stored as CRDT sequence items
// Each item has: item_id, left_id, right_id, deleted_length, value
// For simplicity, we extract text bytes directly

r := NewV6Reader(data)
var currentText strings.Builder
currentStyle := TextStylePlain

for r.Remaining() > 0 {
// Read tag
index, tagType, err := r.ReadTag()
if err != nil {
break
}

switch index {
case 1: // item_id - CRDT ID
if err := r.SkipTagValue(tagType); err != nil {
break
}
case 2: // left_id - CRDT ID  
if err := r.SkipTagValue(tagType); err != nil {
break
}
case 3: // right_id - CRDT ID
if err := r.SkipTagValue(tagType); err != nil {
break
}
case 4: // deleted_length (uint32)
if tagType == TagByte4 {
_, err := r.ReadUint32()
if err != nil {
break
}
} else {
if err := r.SkipTagValue(tagType); err != nil {
break
}
}
case 5: // value - the actual text content
if tagType == TagLength4 {
subSize, err := r.ReadUint32()
if err != nil {
break
}
textBytes, err := r.ReadBytes(int(subSize))
if err != nil {
break
}
text := string(textBytes)
// Handle newlines as paragraph breaks
for _, char := range text {
if char == '\n' {
if currentText.Len() > 0 {
paragraphs = append(paragraphs, TextParagraph{
Style: currentStyle,
Text:  currentText.String(),
})
currentText.Reset()
}
} else {
currentText.WriteRune(char)
}
}
} else if tagType == TagByte4 {
// Format code (bold, italic, etc) - skip
_, err := r.ReadUint32()
if err != nil {
break
}
} else {
if err := r.SkipTagValue(tagType); err != nil {
break
}
}
default:
if err := r.SkipTagValue(tagType); err != nil {
break
}
}
}

// Add remaining text as final paragraph
if currentText.Len() > 0 {
paragraphs = append(paragraphs, TextParagraph{
Style: currentStyle,
Text:  currentText.String(),
})
}

return paragraphs
}
