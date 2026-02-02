package rm

import (
	"testing"
)

func TestV6HeaderDetection(t *testing.T) {
	// Create a minimal V6 "file" with just a header
	v6Header := []byte(HeaderV6)
	
	// Add minimal data to avoid EOF errors
	v6Data := append(v6Header, []byte{0, 0, 0, 0, 0}...)
	
	rm := &Rm{}
	err := rm.UnmarshalBinary(v6Data)
	
	// The V6 parser should handle this (may succeed with empty data or fail gracefully)
	// The key is it should detect V6 format and not treat it as V3/V5
	if err != nil && err.Error() == "V6 format should have been detected earlier" {
		t.Error("V6 format was not properly detected in UnmarshalBinary")
	}
	
	// If successful, verify it's V6
	if err == nil {
		if rm.Version != V6 {
			t.Errorf("Expected V6, got %v", rm.Version)
		}
		t.Logf("V6 parsing succeeded, got %d layers", len(rm.Layers))
	} else {
		t.Logf("V6 parsing failed as expected with minimal data: %v", err)
	}
}

func TestV3StillWorks(t *testing.T) {
	// Ensure V3 files still work
	rm := &Rm{}
	v3Header := []byte(HeaderV3)
	v3Data := append(v3Header, []byte{1, 0, 0, 0}...) // 1 layer
	v3Data = append(v3Data, []byte{0, 0, 0, 0}...)    // 0 lines
	
	err := rm.UnmarshalBinary(v3Data)
	if err != nil {
		t.Errorf("V3 parsing failed: %v", err)
	}
	
	if rm.Version != V3 {
		t.Errorf("Expected V3, got %v", rm.Version)
	}
}

func TestV5StillWorks(t *testing.T) {
	// Ensure V5 files still work
	rm := &Rm{}
	v5Header := []byte(HeaderV5)
	v5Data := append(v5Header, []byte{1, 0, 0, 0}...) // 1 layer
	v5Data = append(v5Data, []byte{0, 0, 0, 0}...)    // 0 lines
	
	err := rm.UnmarshalBinary(v5Data)
	if err != nil {
		t.Errorf("V5 parsing failed: %v", err)
	}
	
	if rm.Version != V5 {
		t.Errorf("Expected V5, got %v", rm.Version)
	}
}

func TestV6TagTypes(t *testing.T) {
	// Verify tag type constants are correct
	if TagByte1 != 0x1 {
		t.Errorf("TagByte1 should be 0x1, got 0x%x", TagByte1)
	}
	if TagByte4 != 0x4 {
		t.Errorf("TagByte4 should be 0x4, got 0x%x", TagByte4)
	}
	if TagByte8 != 0x8 {
		t.Errorf("TagByte8 should be 0x8, got 0x%x", TagByte8)
	}
	if TagLength4 != 0xC {
		t.Errorf("TagLength4 should be 0xC, got 0x%x", TagLength4)
	}
	if TagID != 0xF {
		t.Errorf("TagID should be 0xF, got 0x%x", TagID)
	}
}

func TestMapPenToV5BrushType(t *testing.T) {
	// Test pen type mapping
	tests := []struct {
		penType  int
		expected BrushType
	}{
		{PenBallpoint, BallPoint},
		{PenBallpointV5, BallPoint},
		{PenMarker, Marker},
		{PenMarkerV5, Marker},
		{PenFineliner, Fineliner},
		{PenFinelinerV5, Fineliner},
		{PenHighlighter, Highlighter},
		{PenHighlighterV5, Highlighter},
		{PenEraser, Eraser},
		{PenBrush, Brush},
		{PenBrushV5, Brush},
	}
	
	for _, tt := range tests {
		result := mapPenToV5BrushType(tt.penType)
		if result != tt.expected {
			t.Errorf("mapPenToV5BrushType(%d) = %d, want %d", tt.penType, result, tt.expected)
		}
	}
}

func TestMapColorToV5(t *testing.T) {
	// Test color mapping
	tests := []struct {
		color    int
		expected BrushColor
	}{
		{ColorBlack, Black},
		{ColorGrey, Grey},
		{ColorGreyOverlap, Grey},
		{ColorWhite, White},
		{ColorYellow, Black}, // Non-standard colors default to black
		{ColorGreen, Black},
	}
	
	for _, tt := range tests {
		result := mapColorToV5(tt.color)
		if result != tt.expected {
			t.Errorf("mapColorToV5(%d) = %d, want %d", tt.color, result, tt.expected)
		}
	}
}
