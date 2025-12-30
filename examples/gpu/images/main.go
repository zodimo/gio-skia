package main

import (
	"bytes"
	"image/color"
	"log"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/go-text/typesetting/font"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func Run(window *app.Window) error {
	var ops op.Ops
	startTime := time.Now()

	// Load font
	parsedFont, err := font.ParseTTF(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	typeface := impl.NewTypefaceWithTypefaceFace("regular", models.FontStyle{}, parsedFont)
	labelFont := impl.NewFontWithTypefaceAndSize(typeface, 14)

	// Create test images
	checkerboard := createCheckerboardImage(64, 64, 8)
	gradient := createGradientImage(100, 100)
	pattern := createPatternImage(80, 80)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()
			elapsed := time.Since(startTime).Seconds()

			// Dark gradient background
			paint.ColorOp{Color: color.NRGBA{R: 25, G: 30, B: 45, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(e.Size.X), float32(e.Size.Y)

			// ─────────────────────────────────────────────────────────
			// Demo 1: DrawImage - Simple image at position
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.2, h*0.25)

			// Draw label box
			drawLabelBox(c, "DrawImage", -40, -90, labelFont)

			// Draw image at position
			c.DrawImage(checkerboard, -32, -32, skia.NewPaint())
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 2: DrawImageRect - Scale image to fit
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.5, h*0.25)

			drawLabelBox(c, "DrawImageRect (Scale)", -80, -80, labelFont)

			// Draw scaled image - zoom in/out animation
			scaleFactor := 0.5 + 0.5*float32(math.Sin(elapsed))
			dstSize := 60.0 * scaleFactor
			dst := models.Rect{
				Left:   float32(-dstSize),
				Top:    float32(-dstSize),
				Right:  float32(dstSize),
				Bottom: float32(dstSize),
			}
			c.DrawImageRect(gradient, nil, dst, skia.NewPaint())
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 3: DrawImageRect - Crop source region
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.8, h*0.25)

			drawLabelBox(c, "DrawImageRect (Crop)", -70, -90, labelFont)

			// Animate crop region
			cropOffset := 20 + 15*float32(math.Sin(elapsed*2))
			src := models.Rect{
				Left:   cropOffset,
				Top:    cropOffset,
				Right:  cropOffset + 40,
				Bottom: cropOffset + 40,
			}
			dst = models.Rect{Left: -50, Top: -50, Right: 50, Bottom: 50}
			c.DrawImageRect(pattern, &src, dst, skia.NewPaint())
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 4: Multiple images with transforms
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.25, h*0.65)

			drawLabelBox(c, "Image + Transforms", -80, -100, labelFont)

			// Draw rotating images in a circle
			for i := 0; i < 6; i++ {
				angle := float32(i)*math.Pi/3 + float32(elapsed*0.5)
				radius := float32(70)

				c.Save()
				c.Translate(radius*float32(math.Cos(float64(angle))), radius*float32(math.Sin(float64(angle))))
				c.Rotate(float32(elapsed * 45)) // Rotate each image
				c.Scale(0.5, 0.5)
				c.DrawImage(checkerboard, -32, -32, skia.NewPaint())
				c.Restore()
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 5: Image grid with different scaling
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.7, h*0.65)

			drawLabelBox(c, "Image Grid", -50, -100, labelFont)

			// 3x3 grid of scaled images
			for row := -1; row <= 1; row++ {
				for col := -1; col <= 1; col++ {
					scale := 0.3 + 0.1*float32(row+col+2)
					dstW := 30 * scale
					dstH := 30 * scale

					dst = models.Rect{
						Left:   float32(col)*40 - dstW/2,
						Top:    float32(row)*40 - dstH/2,
						Right:  float32(col)*40 + dstW/2,
						Bottom: float32(row)*40 + dstH/2,
					}
					c.DrawImageRect(gradient, nil, dst, skia.NewPaint())
				}
			}
			c.Restore()

			window.Invalidate()
			e.Frame(&ops)
		}
	}
}

// createCheckerboardImage creates a checkerboard pattern image
func createCheckerboardImage(width, height, tileSize int) *impl.RasterImage {
	imgInfo := models.NewImageInfo(width, height, enums.ColorTypeRGBA8888, enums.AlphaTypePremul)
	pixels := make([]byte, width*height*4)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 4
			tileX := x / tileSize
			tileY := y / tileSize
			if (tileX+tileY)%2 == 0 {
				pixels[idx] = 255   // R
				pixels[idx+1] = 255 // G
				pixels[idx+2] = 255 // B
				pixels[idx+3] = 255 // A
			} else {
				pixels[idx] = 50    // R
				pixels[idx+1] = 50  // G
				pixels[idx+2] = 80  // B
				pixels[idx+3] = 255 // A
			}
		}
	}

	return impl.NewRasterImage(imgInfo, pixels, width*4)
}

// createGradientImage creates a radial gradient image
func createGradientImage(width, height int) *impl.RasterImage {
	imgInfo := models.NewImageInfo(width, height, enums.ColorTypeRGBA8888, enums.AlphaTypePremul)
	pixels := make([]byte, width*height*4)

	centerX := float64(width) / 2
	centerY := float64(height) / 2
	maxDist := math.Sqrt(centerX*centerX + centerY*centerY)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 4
			dx := float64(x) - centerX
			dy := float64(y) - centerY
			dist := math.Sqrt(dx*dx+dy*dy) / maxDist

			// Gradient from red center to blue edge
			pixels[idx] = uint8(255 * (1 - dist))   // R
			pixels[idx+1] = uint8(100 * (1 - dist)) // G
			pixels[idx+2] = uint8(100 + 155*dist)   // B
			pixels[idx+3] = 255                     // A
		}
	}

	return impl.NewRasterImage(imgInfo, pixels, width*4)
}

// createPatternImage creates a pattern image
func createPatternImage(width, height int) *impl.RasterImage {
	imgInfo := models.NewImageInfo(width, height, enums.ColorTypeRGBA8888, enums.AlphaTypePremul)
	pixels := make([]byte, width*height*4)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 4
			// Create a pattern based on position
			pattern := (x * y) % 256
			pixels[idx] = uint8((x * 3) % 256)   // R
			pixels[idx+1] = uint8((y * 3) % 256) // G
			pixels[idx+2] = uint8(pattern)       // B
			pixels[idx+3] = 255                  // A
		}
	}

	return impl.NewRasterImage(imgInfo, pixels, width*4)
}

func drawLabelBox(c skia.Canvas, label string, x, y float32, font interfaces.SkFont) {
	p := skia.NewPaintFill(color.NRGBA{R: 255, G: 255, B: 255, A: 80})
	box := skia.NewPath()
	skia.PathMoveTo(box, x, y)
	skia.PathLineTo(box, x+160, y)
	skia.PathLineTo(box, x+160, y+20)
	skia.PathLineTo(box, x, y+20)
	box.Close()
	c.DrawPath(box, p)

	// Draw text
	textPaint := skia.NewPaintFill(color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	c.DrawSimpleText([]byte(label), enums.TextEncodingUTF8, x+5, y+14, font, textPaint)
}
