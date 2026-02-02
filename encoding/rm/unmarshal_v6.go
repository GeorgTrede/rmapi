// V6 format decoder for reMarkable software version 3+.
// This is a scene-based format that is completely different from the
// V3/V5 binary stroke format.
package rm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// Block types
const (
	BlockTypeSceneInfo           = 0x0D
	BlockTypeAuthorIds           = 0x09
	BlockTypeMigrationInfo       = 0x0A
	BlockTypePageInfo            = 0x0B
	BlockTypeSceneTree           = 0x01
	BlockTypeTreeNode            = 0x02
	BlockTypeSceneGroupItem      = 0x03
	BlockTypeSceneLineItem       = 0x05
	BlockTypeSceneGlyphItem      = 0x04
	BlockTypeRootText            = 0x07
	BlockTypeSceneTombstoneItem  = 0x08
)

// Tag types for tagged data
const (
	TagByte       = 0x00
	TagLength4    = 0x01
	TagLength8    = 0x02
	TagDouble     = 0x03
	TagBool       = 0x04
	TagVarUint    = 0x05
	TagID1        = 0x06
	TagID2        = 0x07
	TagInt        = 0x08
	TagFloat      = 0x09
	TagString     = 0x0A
	TagIDPair     = 0x0B
)

// Pen tool types
const (
	PenBallpoint      = 2
	PenMarker         = 3
	PenFineliner      = 4
	PenSharpPencil    = 7
	PenTiltPencil     = 1
	PenBrush          = 0
	PenHighlighter    = 5
	PenEraser         = 6
	PenEraseArea      = 8
	PenCalligraphy    = 21
	// V6 additions (renumbered types)
	PenBrushV5        = 12
	PenMechanicalPencil = 13  // Same as PenSharpPencilV5
	PenPencil         = 14    // Same as PenTiltPencilV5
	PenBallpointV5    = 15
	PenMarkerV5       = 16
	PenFinelinerV5    = 17
	PenHighlighterV5  = 18
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
)

// Reader wraps an io.Reader and provides methods to read V6 format data
type Reader struct {
	r io.Reader
}

// NewReader creates a new V6 reader
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// ReadVarUint reads a variable-length unsigned integer
func (r *Reader) ReadVarUint() (uint64, error) {
	var result uint64
	var shift uint
	for {
		var b byte
		err := binary.Read(r.r, binary.LittleEndian, &b)
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

// ReadBytes reads n bytes
func (r *Reader) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r.r, buf)
	return buf, err
}

// ReadUint8 reads a uint8
func (r *Reader) ReadUint8() (uint8, error) {
	var v uint8
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadUint16 reads a uint16
func (r *Reader) ReadUint16() (uint16, error) {
	var v uint16
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadFloat32 reads a float32
func (r *Reader) ReadFloat32() (float32, error) {
	var v float32
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadFloat64 reads a float64
func (r *Reader) ReadFloat64() (float64, error) {
	var v float64
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// ReadInt32 reads an int32
func (r *Reader) ReadInt32() (int32, error) {
	var v int32
	err := binary.Read(r.r, binary.LittleEndian, &v)
	return v, err
}

// UnmarshalV6 reads a V6 .rm file and converts it to the common Rm structure
func UnmarshalV6(data []byte) (*Rm, error) {
	r := NewReader(bytes.NewReader(data))
	
	// Read and verify header
	header, err := r.ReadBytes(43)
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}
	if string(header) != HeaderV6 {
		return nil, fmt.Errorf("invalid V6 header: %q", string(header))
	}

	// Parse blocks and extract lines
	lines, err := parseBlocks(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse blocks: %w", err)
	}

	// Convert to Rm structure
	// V6 doesn't have explicit layers, so we put all lines in one layer
	result := &Rm{
		Version: V6,
		Layers: []Layer{
			{Lines: lines},
		},
	}

	return result, nil
}

// parseBlocks reads all blocks from the stream and extracts line items
func parseBlocks(r *Reader) ([]Line, error) {
	var lines []Line

	for {
		// Try to read next block
		blockType, blockSize, version, err := readBlockHeader(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Read block data
		blockData, err := r.ReadBytes(int(blockSize))
		if err != nil {
			return nil, fmt.Errorf("failed to read block data: %w", err)
		}

		// Parse line items
		if blockType == BlockTypeSceneLineItem {
			line, err := parseLineBlock(blockData, version)
			if err != nil {
				// Log but continue - partial extraction is better than nothing
				continue
			}
			lines = append(lines, line)
		}
		// Ignore other block types for now
	}

	return lines, nil
}

// readBlockHeader reads a block header
func readBlockHeader(r *Reader) (blockType byte, size uint64, version byte, err error) {
	// Read block header tag
	tag, err := r.ReadVarUint()
	if err != nil {
		return 0, 0, 0, err
	}

	blockType = byte(tag & 0x0F)
	
	// Read min version
	minVer, err := r.ReadVarUint()
	if err != nil {
		return 0, 0, 0, err
	}
	
	// Read current version
	curVer, err := r.ReadVarUint()
	if err != nil {
		return 0, 0, 0, err
	}
	version = byte(curVer)
	
	// Read block size
	size, err = r.ReadVarUint()
	if err != nil {
		return 0, 0, 0, err
	}

	_ = minVer // unused
	
	return blockType, size, version, nil
}

// parseLineBlock parses a SceneLineItem block
func parseLineBlock(data []byte, blockVersion byte) (Line, error) {
	r := NewReader(bytes.NewReader(data))
	
	var line Line
	
	// Read tagged fields
	for {
		// Read tag
		tagValue, err := r.ReadVarUint()
		if err == io.EOF {
			break
		}
		if err != nil {
			return line, err
		}
		
		index := tagValue >> 4
		tagType := tagValue & 0x0F
		
		switch index {
		case 1: // tool/pen type
			if tagType == TagInt {
				toolID, err := r.ReadInt32()
				if err != nil {
					return line, err
				}
				line.BrushType = mapPenToV5BrushType(int(toolID))
			}
		case 2: // color
			if tagType == TagInt {
				colorID, err := r.ReadInt32()
				if err != nil {
					return line, err
				}
				line.BrushColor = mapColorToV5(int(colorID))
			}
		case 3: // thickness_scale
			if tagType == TagDouble {
				scale, err := r.ReadFloat64()
				if err != nil {
					return line, err
				}
				line.BrushSize = BrushSize(scale * 2.0) // Approximate mapping
			}
		case 5: // points data (subblock)
			if tagType == TagLength4 {
				// Read subblock size
				subSize, err := r.ReadVarUint()
				if err != nil {
					return line, err
				}
				// Read points
				points, err := parsePoints(r, int(subSize), blockVersion)
				if err != nil {
					return line, err
				}
				line.Points = points
			}
		}
	}
	
	return line, nil
}

// parsePoints reads point data from the stream
func parsePoints(r *Reader, dataSize int, version byte) ([]Point, error) {
	pointSize := 0x0E // Version 2 point size (14 bytes)
	if version == 1 {
		pointSize = 0x18 // Version 1 point size (24 bytes)
	}
	
	numPoints := dataSize / pointSize
	points := make([]Point, numPoints)
	
	for i := 0; i < numPoints; i++ {
		p, err := readPoint(r, version)
		if err != nil {
			return nil, err
		}
		points[i] = p
	}
	
	return points, nil
}

// readPoint reads a single point
func readPoint(r *Reader, version byte) (Point, error) {
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
		// Version 1: float-based encoding
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
		// Version 2: integer-based encoding
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
		return Marker // Map to closest equivalent
	default:
		return BallPoint // Default fallback
	}
}

// mapColorToV5 maps V6 color types to V5 colors
func mapColorToV5(color int) BrushColor {
	switch color {
	case ColorBlack:
		return Black
	case ColorGrey, ColorGreyOverlap:
		return Grey
	case ColorWhite:
		return White
	default:
		return Black // Default to black for other colors
	}
}
