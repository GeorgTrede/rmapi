package rm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// UnmarshalBinary implements encoding.UnmarshalBinary for
// transforming bytes into a Rm page
func (rm *Rm) UnmarshalBinary(data []byte) error {
	// Check if this is a V6 file by looking at the header
	if len(data) >= HeaderLen {
		header := string(data[:HeaderLen])
		if header == HeaderV6 {
			// Use V6 parser
			v6Data, err := UnmarshalV6(data)
			if err != nil {
				return fmt.Errorf("failed to parse V6 format: %w", err)
			}
			*rm = *v6Data
			return nil
		}
	}

	// Use V3/V5 parser
	r := newReader(data)
	if err := r.checkHeader(); err != nil {
		return err
	}
	rm.Version = r.version

	nbLayers, err := r.readNumber()
	if err != nil {
		return err
	}

	rm.Layers = make([]Layer, nbLayers)
	for i := uint32(0); i < nbLayers; i++ {
		nbLines, err := r.readNumber()
		if err != nil {
			return err
		}

		rm.Layers[i].Lines = make([]Line, nbLines)
		for j := uint32(0); j < nbLines; j++ {
			line, err := r.readLine()
			if err != nil {
				return err
			}
			rm.Layers[i].Lines[j] = line
		}
	}

	return nil
}

type reader struct {
	bytes.Reader
	version Version
}

func newReader(data []byte) reader {
	br := bytes.NewReader(data)

	// we set V5 as default but the real value is
	// analysed when checking the header
	return reader{*br, V5}
}

func (r *reader) checkHeader() error {
	buf := make([]byte, HeaderLen)

	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	if n != HeaderLen {
		return fmt.Errorf("Wrong header size")
	}

	switch string(buf) {
	case HeaderV6:
		// V6 should be handled in UnmarshalBinary before reaching here
		return fmt.Errorf("V6 format should have been detected earlier")
	case HeaderV5:
		r.version = V5
	case HeaderV3:
		r.version = V3
	default:
		return fmt.Errorf("Unknown header: %q", string(buf))
	}

	return nil
}

func (r *reader) readNumber() (uint32, error) {
	var nb uint32
	if err := binary.Read(r, binary.LittleEndian, &nb); err != nil {
		return 0, fmt.Errorf("Wrong number read")
	}
	return nb, nil
}

func (r *reader) readLine() (Line, error) {
	var line Line

	if err := binary.Read(r, binary.LittleEndian, &line.BrushType); err != nil {
		return line, fmt.Errorf("Failed to read line")
	}

	if err := binary.Read(r, binary.LittleEndian, &line.BrushColor); err != nil {
		return line, fmt.Errorf("Failed to read line")
	}

	if err := binary.Read(r, binary.LittleEndian, &line.Padding); err != nil {
		return line, fmt.Errorf("Failed to read line")
	}

	if err := binary.Read(r, binary.LittleEndian, &line.BrushSize); err != nil {
		return line, fmt.Errorf("Failed to read line")
	}

	// this new attribute has been added in v5
	// Note: V6 is a completely different format and is not supported here
	if r.version == V5 {
		if err := binary.Read(r, binary.LittleEndian, &line.Unknown); err != nil {
			return line, fmt.Errorf("Failed to read line")
		}
	}

	nbPoints, err := r.readNumber()
	if err != nil {
		return line, err
	}

	if nbPoints == 0 {
		return line, nil
	}

	line.Points = make([]Point, nbPoints)

	for i := uint32(0); i < nbPoints; i++ {
		p, err := r.readPoint()
		if err != nil {
			return line, err
		}

		line.Points[i] = p
	}

	return line, nil
}

func (r *reader) readPoint() (Point, error) {
	var point Point

	if err := binary.Read(r, binary.LittleEndian, &point.X); err != nil {
		return point, fmt.Errorf("Failed to read point")
	}
	if err := binary.Read(r, binary.LittleEndian, &point.Y); err != nil {
		return point, fmt.Errorf("Failed to read point")
	}
	if err := binary.Read(r, binary.LittleEndian, &point.Speed); err != nil {
		return point, fmt.Errorf("Failed to read point")
	}
	if err := binary.Read(r, binary.LittleEndian, &point.Direction); err != nil {
		return point, fmt.Errorf("Failed to read point")
	}
	if err := binary.Read(r, binary.LittleEndian, &point.Width); err != nil {
		return point, fmt.Errorf("Failed to read point")
	}
	if err := binary.Read(r, binary.LittleEndian, &point.Pressure); err != nil {
		return point, fmt.Errorf("Failed to read point")
	}

	return point, nil
}
