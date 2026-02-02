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
)

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

func renderToSVG(rmData *rm.Rm) string {
	var sb strings.Builder

	// SVG header with viewBox matching device dimensions
	sb.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" 
     viewBox="0 0 %d %d" 
     width="%d" height="%d"
     style="background-color: white;">
`, DeviceWidth, DeviceHeight, DeviceWidth, DeviceHeight))

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
			color := "black"
			switch line.BrushColor {
			case rm.Black:
				color = "black"
			case rm.Grey:
				color = "gray"
			case rm.White:
				color = "white"
			}

			// Get stroke width
			strokeWidth := float64(line.BrushSize) * 2.0
			if strokeWidth < 1.0 {
				strokeWidth = 1.0
			}

			// Check if highlighter
			isHighlighter := line.BrushType == rm.Highlighter || line.BrushType == rm.HighlighterV5
			opacity := 1.0
			if isHighlighter {
				color = "yellow"
				opacity = 0.5
				strokeWidth = 20.0
			}

			// Build path
			sb.WriteString(`  <path d="M `)
			for i, point := range line.Points {
				// Translate coordinates (center on page)
				x := float64(point.X) + float64(DeviceWidth)/2
				y := float64(point.Y) + float64(DeviceHeight)/2

				if i == 0 {
					sb.WriteString(fmt.Sprintf("%.2f %.2f", x, y))
				} else {
					sb.WriteString(fmt.Sprintf(" L %.2f %.2f", x, y))
				}
			}
			sb.WriteString(fmt.Sprintf(`" fill="none" stroke="%s" stroke-width="%.1f" stroke-linecap="round" stroke-linejoin="round"`, color, strokeWidth))
			if opacity < 1.0 {
				sb.WriteString(fmt.Sprintf(` opacity="%.1f"`, opacity))
			}
			sb.WriteString("/>\n")
		}
	}

	sb.WriteString("</svg>\n")
	return sb.String()
}
