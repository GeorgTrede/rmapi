// render_v6_svg.go - Test tool to render V6 .rm files to SVG
// Usage: go run render_v6_svg.go <input.rm> <output.svg>
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/juruen/rmapi/encoding/rm"
)

const (
	DeviceWidth  = 1404
	DeviceHeight = 1872
	ScreenDPI    = 226.0
	TargetDPI    = 72.0
)

// Scale factor to convert screen units to points
var Scale = TargetDPI / ScreenDPI

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run render_v6_svg.go <input.rm> <output.svg>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	rmData := &rm.Rm{}
	err = rmData.UnmarshalBinary(data)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		os.Exit(1)
	}

	// Generate SVG
	svg := renderToSVG(rmData)

	// Write to file
	if err := os.WriteFile(outputPath, []byte(svg), 0644); err != nil {
		fmt.Printf("Error writing SVG: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s (version: %d, lines: %d)\n", outputPath, rmData.Version, countLines(rmData))
}

func countLines(rmData *rm.Rm) int {
	total := 0
	for _, layer := range rmData.Layers {
		total += len(layer.Lines)
	}
	return total
}

// scale converts screen units to SVG points
func scale(screenUnit float64) float64 {
	return screenUnit * Scale
}

// Get bounding box of all points
func getBoundingBox(rmData *rm.Rm) (xMin, xMax, yMin, yMax float64) {
	// Default bounding box (same as Python)
	xMin = float64(-DeviceWidth) / 2
	xMax = float64(DeviceWidth) / 2
	yMin = 0
	yMax = float64(DeviceHeight)

	for _, layer := range rmData.Layers {
		for _, line := range layer.Lines {
			for _, point := range line.Points {
				if float64(point.X) < xMin {
					xMin = float64(point.X)
				}
				if float64(point.X) > xMax {
					xMax = float64(point.X)
				}
				if float64(point.Y) < yMin {
					yMin = float64(point.Y)
				}
				if float64(point.Y) > yMax {
					yMax = float64(point.Y)
				}
			}
		}
	}
	return
}

// Color palette matching Python RM_PALETTE
var colorPalette = map[rm.BrushColor][3]int{
	rm.Black:       {0, 0, 0},
	rm.Grey:        {144, 144, 144},
	rm.White:       {255, 255, 255},
	rm.Yellow:      {251, 247, 25},
	rm.Green:       {0, 255, 0},
	rm.Pink:        {255, 192, 203},
	rm.Blue:        {78, 105, 201},
	rm.Red:         {179, 62, 57},
	rm.GreyOverlap: {125, 125, 125},
	rm.Highlight:   {251, 247, 25}, // Same as Yellow
	rm.Green2:      {161, 216, 125},
	rm.Cyan:        {139, 208, 229},
	rm.Magenta:     {183, 130, 205},
	rm.Yellow2:     {247, 232, 81},
}

func getColor(brushColor rm.BrushColor) string {
	if rgb, ok := colorPalette[brushColor]; ok {
		return fmt.Sprintf("rgb(%d,%d,%d)", rgb[0], rgb[1], rgb[2])
	}
	return "rgb(0,0,0)"
}

func renderToSVG(rmData *rm.Rm) string {
	var sb strings.Builder

	// Get bounding box
	xMin, xMax, yMin, yMax := getBoundingBox(rmData)
	widthPt := scale(xMax - xMin + 1)
	heightPt := scale(yMax - yMin + 1)

	// SVG header with viewBox matching content (like Python)
	sb.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" 
     height="%.1f" width="%.1f" 
     viewBox="%.1f %.1f %.1f %.1f"
     style="background-color: white;">
`, heightPt, widthPt, scale(xMin), scale(yMin), widthPt, heightPt))

	sb.WriteString(`	<g id="p1" style="display:inline">
`)

	// Render each layer
	for _, layer := range rmData.Layers {
		for _, line := range layer.Lines {
			if len(line.Points) < 2 {
				continue
			}

			// Skip eraser
			if line.BrushType == rm.Eraser || line.BrushType == rm.EraseArea {
				continue
			}

			// Get color
			color := getColor(line.BrushColor)

			// Get stroke width
			strokeWidth := float64(line.BrushSize) * 1.8 // Match Fineliner scaling
			if strokeWidth < 0.5 {
				strokeWidth = 0.5
			}

			// Check if highlighter
			isHighlighter := line.BrushType == rm.Highlighter || line.BrushType == rm.HighlighterV5
			opacity := 1.0
			if isHighlighter {
				// Use the actual line color for highlighters, don't override to yellow
				// Python highlighter uses the color_id from the line data
				opacity = 0.3
				strokeWidth = 15.0
			}

			// Build polyline (like Python)
			sb.WriteString(`		<polyline style="fill:none; stroke:`)
			sb.WriteString(color)
			sb.WriteString(fmt.Sprintf(`; stroke-width:%.3f`, scale(strokeWidth)))
			if opacity < 1.0 {
				sb.WriteString(fmt.Sprintf(`; opacity:%.1f`, opacity))
			}
			sb.WriteString(`" stroke-linecap="round" points="`)

			for _, point := range line.Points {
				x := scale(float64(point.X))
				y := scale(float64(point.Y))
				sb.WriteString(fmt.Sprintf("%.3f,%.3f ", x, y))
			}

			sb.WriteString(`"/>
`)
		}
	}

	sb.WriteString(`	</g>
</svg>
`)
	return sb.String()
}

