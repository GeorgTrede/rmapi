package annotations

import (
	"bytes"
	"errors"
	"fmt"

	"os"

	"github.com/juruen/rmapi/archive"
	"github.com/juruen/rmapi/encoding/rm"
	"github.com/juruen/rmapi/log"
	"github.com/unidoc/unipdf/v3/annotator"
	"github.com/unidoc/unipdf/v3/contentstream"
	"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/core"
	"github.com/unidoc/unipdf/v3/creator"
	pdf "github.com/unidoc/unipdf/v3/model"
)

const (
	DeviceWidth  = 1404
	DeviceHeight = 1872
)

var rmPageSize = creator.PageSize{445, 594}

type PdfGenerator struct {
	zipName        string
	outputFilePath string
	options        PdfGeneratorOptions
	pdfReader      *pdf.PdfReader
	template       bool
}

type PdfGeneratorOptions struct {
	AddPageNumbers  bool
	AllPages        bool
	AnnotationsOnly bool //export the annotations without the background/pdf
}

func CreatePdfGenerator(zipName, outputFilePath string, options PdfGeneratorOptions) *PdfGenerator {
	return &PdfGenerator{zipName: zipName, outputFilePath: outputFilePath, options: options}
}

// getBoundingBox calculates the bounding box of all points in the rm data
func getBoundingBox(rmData *rm.Rm) (xMin, xMax, yMin, yMax float64) {
	// Default bounding box matching Python rmscene
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

func normalized(p1 rm.Point, ratioX float64, xOffset, yOffset float64) (float64, float64) {
	// Shift coordinates by offsets to make them positive for PDF
	x := (float64(p1.X) - xOffset) * ratioX
	y := (float64(p1.Y) - yOffset) * ratioX
	return x, y
}

// brushColorToRGB converts a BrushColor to RGB values (0.0-1.0)
func brushColorToRGB(color rm.BrushColor) (float64, float64, float64) {
	switch color {
	case rm.Black:
		return 0.0, 0.0, 0.0
	case rm.Grey, rm.GreyOverlap:
		return 0.56, 0.56, 0.56 // 144/255
	case rm.White:
		return 1.0, 1.0, 1.0
	case rm.Yellow, rm.Yellow2, rm.Highlight:
		return 0.98, 0.97, 0.10 // 251, 247, 25
	case rm.Green:
		return 0.0, 1.0, 0.0
	case rm.Green2:
		return 0.63, 0.85, 0.49 // 161, 216, 125
	case rm.Pink:
		return 1.0, 0.75, 0.80 // 255, 192, 203
	case rm.Blue:
		return 0.31, 0.41, 0.79 // 78, 105, 201
	case rm.Red:
		return 0.70, 0.24, 0.22 // 179, 62, 57
	case rm.Cyan:
		return 0.55, 0.82, 0.90 // 139, 208, 229
	case rm.Magenta:
		return 0.72, 0.51, 0.80 // 183, 130, 205
	default:
		return 0.0, 0.0, 0.0 // Default to black
	}
}

func (p *PdfGenerator) Generate() error {
	file, err := os.Open(p.zipName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	zip := archive.NewZip()

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	err = zip.Read(file, fi.Size())
	if err != nil {
		return err
	}

	if zip.Content.FileType == "epub" {
		return errors.New("only pdf and notebooks supported")
	}

	if err = p.initBackgroundPages(zip.Payload); err != nil {
		return err
	}

	if len(zip.Pages) == 0 {
		return errors.New("the document has no pages")
	}

	c := creator.New()
	if p.template {
		// use the standard page size
		c.SetPageSize(rmPageSize)
	}

	if p.pdfReader != nil && p.options.AllPages {
		outlines := p.pdfReader.GetOutlineTree()
		c.SetOutlineTree(outlines)
	}

	for _, pageAnnotations := range zip.Pages {
		hasContent := pageAnnotations.Data != nil

		// do not add a page when there are no annotations
		if !p.options.AllPages && !hasContent {
			continue
		}
		//1 based, redirected page
		pageNum := pageAnnotations.DocPage + 1

		page, err := p.addBackgroundPage(c, pageNum)
		if err != nil {
			return err
		}

		ratio := c.Height() / c.Width()

		var scale float64
		if ratio < 1.33 {
			scale = c.Width() / DeviceWidth
		} else {
			scale = c.Height() / DeviceHeight
		}
		if page == nil {
			log.Error.Fatal("page is null")
		}

		if err != nil {
			return err
		}
		if !hasContent {
			continue
		}

		// Calculate bounding box for V6 format coordinate translation
		xMin, _, yMin, _ := getBoundingBox(pageAnnotations.Data)

		contentCreator := contentstream.NewContentCreator()
		contentCreator.Add_q()

		for _, layer := range pageAnnotations.Data.Layers {
			for _, line := range layer.Lines {
				if len(line.Points) < 1 {
					continue
				}
				if line.BrushType == rm.Eraser || line.BrushType == rm.EraseArea {
					continue
				}

				if line.BrushType == rm.HighlighterV5 || line.BrushType == rm.Highlighter {
					last := len(line.Points) - 1
					x1, y1 := normalized(line.Points[0], scale, xMin, yMin)
					x2, _ := normalized(line.Points[last], scale, xMin, yMin)
					// make horizontal lines only, use y1
					width := scale * 30
					y1 += width / 2

					lineDef := annotator.LineAnnotationDef{X1: x1 - 1, Y1: c.Height() - y1, X2: x2, Y2: c.Height() - y1}
					// Use actual line color instead of hardcoded yellow
					r, g, b := brushColorToRGB(line.BrushColor)
					lineDef.LineColor = pdf.NewPdfColorDeviceRGB(r, g, b)
					// Opacity 0.3 matches Python rmc library's Highlighter.base_opacity
					lineDef.Opacity = 0.3
					lineDef.LineWidth = width
					ann, err := annotator.CreateLineAnnotation(lineDef)
					if err != nil {
						return err
					}
					page.AddAnnotation(ann)
				} else {
					path := draw.NewPath()
					for i := 0; i < len(line.Points); i++ {
						x1, y1 := normalized(line.Points[i], scale, xMin, yMin)
						path = path.AppendPoint(draw.NewPoint(x1, c.Height()-y1))
					}

					contentCreator.Add_w(float64(line.BrushSize*6.0 - 10.8))

					// Use actual color for strokes
					r, g, b := brushColorToRGB(line.BrushColor)
					contentCreator.Add_RG(r, g, b) // Add_RG sets stroke color

					//TODO: use bezier
					draw.DrawPathWithCreator(path, contentCreator)

					contentCreator.Add_S()
				}
			}
		}
		contentCreator.Add_Q()
		drawingOperations := contentCreator.Operations().String()
		pageContentStreams, err := page.GetAllContentStreams()
		//hack: wrap the page content in a context to prevent transformation matrix misalignment
		wrapper := []string{"q", pageContentStreams, "Q", drawingOperations}
		page.SetContentStreams(wrapper, core.NewFlateEncoder())
	}

	return c.WriteToFile(p.outputFilePath)
}

func (p *PdfGenerator) initBackgroundPages(pdfArr []byte) error {
	if len(pdfArr) > 0 {
		pdfReader, err := pdf.NewPdfReader(bytes.NewReader(pdfArr))
		if err != nil {
			return err
		}

		encrypted, err := pdfReader.IsEncrypted()
		if err != nil {
			return nil
		}
		if encrypted {
			valid, err := pdfReader.Decrypt([]byte(""))
			if err != nil {
				return err
			}
			if !valid {
				return fmt.Errorf("cannot decrypt")
			}

		}

		p.pdfReader = pdfReader
		p.template = false
		return nil
	}

	p.template = true
	return nil
}

func (p *PdfGenerator) addBackgroundPage(c *creator.Creator, pageNum int) (*pdf.PdfPage, error) {
	var page *pdf.PdfPage

	// if page == 0 then empty page
	if !p.template && !p.options.AnnotationsOnly && pageNum > 0 {
		tmpPage, err := p.pdfReader.GetPage(pageNum)
		if err != nil {
			return nil, err
		}
		mbox, err := tmpPage.GetMediaBox()
		if err != nil {
			return nil, err
		}

		// TODO: adjust the page if cropped
		pageHeight := mbox.Ury - mbox.Lly
		pageWidth := mbox.Urx - mbox.Llx
		// use the pdf's page size
		c.SetPageSize(creator.PageSize{pageWidth, pageHeight})
		c.AddPage(tmpPage)
		page = tmpPage
	} else {
		page = c.NewPage()
	}

	if p.options.AddPageNumbers {
		c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
			p := c.NewParagraph(fmt.Sprintf("%d", args.PageNum))
			p.SetFontSize(8)
			w := block.Width() - 20
			h := block.Height() - 10
			p.SetPos(w, h)
			block.Draw(p)
		})
	}
	return page, nil
}
