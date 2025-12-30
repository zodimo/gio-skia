// Package main demonstrates the SkCanvas text drawing methods.
// It shows DrawTextBlob, DrawSimpleText, and DrawString using HarfBuzz shaper.
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
	"github.com/zodimo/go-skia-support/skia/shaper"
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

	// Create fonts with different sizes (using real font data for shaper)
	parsedFont, err := font.ParseTTF(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	typeface := impl.NewTypefaceWithTypefaceFace("regular", impl.FontStyle{}, parsedFont)

	smallFont := impl.NewFontWithTypefaceAndSize(typeface, 16)
	mediumFont := impl.NewFontWithTypefaceAndSize(typeface, 24)
	largeFont := impl.NewFontWithTypefaceAndSize(typeface, 36)

	// Pre-create text blobs for efficiency using a helper that shapes text
	helloBlob := makeTextBlob("Hello, Gio-Skia!", largeFont)
	numberBlob := makeTextBlob("12345", mediumFont)
	smallBlob := makeTextBlob("Small text rendering", smallFont)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()
			elapsed := time.Since(startTime).Seconds()

			// Dark background
			paint.ColorOp{Color: color.NRGBA{R: 15, G: 20, B: 35, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(e.Size.X), float32(e.Size.Y)

			// ─────────────────────────────────────────────────────────
			// Demo 1: DrawTextBlob - Pre-built text blob
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.5, h*0.12)

			// Draw section label
			drawSectionBox(c, "DrawTextBlob (Pre-built)", -120, -40, smallFont)

			// Draw text blob with color animation
			red := uint8(128 + 127*math.Sin(elapsed))
			green := uint8(128 + 127*math.Sin(elapsed+2))
			blue := uint8(128 + 127*math.Sin(elapsed+4))
			textPaint := skia.NewPaintFill(color.NRGBA{R: red, G: green, B: blue, A: 255})

			if helloBlob != nil {
				c.DrawTextBlob(helloBlob, -100, 20, textPaint)
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 2: DrawString - Simple string API
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.25, h*0.3)

			drawSectionBox(c, "DrawString", -60, -40, smallFont)

			// Draw strings at different positions
			textPaint = skia.NewPaintFill(color.NRGBA{R: 100, G: 200, B: 255, A: 255})
			c.DrawString("Gio + Skia", -50, 20, mediumFont, textPaint)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 3: DrawSimpleText with transforms
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.75, h*0.3)

			drawSectionBox(c, "DrawSimpleText + Scale", -100, -40, smallFont)

			// Scale animation
			scale := 0.7 + 0.3*float32(math.Sin(elapsed*2))
			c.Scale(scale, scale)

			textPaint = skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 100, A: 255})
			c.DrawSimpleText([]byte("Scaled!"), enums.TextEncodingUTF8, -40, 20, mediumFont, textPaint)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 4: Rotating text
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.25, h*0.55)

			drawSectionBox(c, "Rotating Text", -70, -80, smallFont)

			c.Rotate(float32(elapsed * 30))
			textPaint = skia.NewPaintFill(color.NRGBA{R: 150, G: 255, B: 150, A: 255})
			c.DrawString("Spin!", -30, 5, mediumFont, textPaint)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 5: Multiple text sizes
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.75, h*0.55)

			drawSectionBox(c, "Different Sizes", -80, -70, smallFont)

			textPaint = skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			c.DrawString("Large", -40, -20, largeFont, textPaint)

			textPaint = skia.NewPaintFill(color.NRGBA{R: 200, G: 200, B: 255, A: 255})
			c.DrawString("Medium text", -60, 10, mediumFont, textPaint)

			textPaint = skia.NewPaintFill(color.NRGBA{R: 200, G: 255, B: 200, A: 255})
			c.DrawString("Small text here", -60, 35, smallFont, textPaint)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 6: Text with clipping
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.5, h*0.8)

			drawSectionBox(c, "Clipped Text", -60, -60, smallFont)

			// Animated clip
			clipWidth := 50 + 50*float32(math.Sin(elapsed*1.5))
			clipRect := models.Rect{Left: -clipWidth, Top: -20, Right: clipWidth, Bottom: 20}
			c.ClipRect(clipRect, enums.ClipOpIntersect, true)

			textPaint = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 150, A: 255})
			c.DrawString("This text reveals!", -90, 5, mediumFont, textPaint)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 7: Text blob with HarfBuzz shaper
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.5, h*0.95)

			// Create blob dynamically with shaper
			timeStr := time.Now().Format("15:04:05")
			hbShaper := shaper.NewHarfbuzzShaper()
			handler := shaper.NewTextBlobBuilderRunHandler(timeStr, models.Point{X: 0, Y: 0})
			hbShaper.Shape(timeStr, mediumFont, true, 0, handler, nil)
			timeBlob := handler.MakeBlob()

			textPaint = skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 200, A: 255})
			if timeBlob != nil {
				c.DrawTextBlob(timeBlob, -50, 0, textPaint)
			}
			c.Restore()

			// Draw unused blobs info (for compilation)
			_ = numberBlob
			_ = smallBlob

			window.Invalidate()
			e.Frame(&ops)
		}
	}
}

func drawSectionBox(c skia.Canvas, label string, x, y float32, font interfaces.SkFont) {
	// Draw background box
	p := skia.NewPaintFill(color.NRGBA{R: 60, G: 70, B: 90, A: 150})
	box := skia.NewPath()
	skia.PathMoveTo(box, x, y)
	skia.PathLineTo(box, x+200, y)
	skia.PathLineTo(box, x+200, y+25)
	skia.PathLineTo(box, x, y+25)
	box.Close()
	c.DrawPath(box, p)

	// Draw label text inside the box
	textPaint := skia.NewPaintFill(color.NRGBA{R: 200, G: 200, B: 200, A: 255})
	c.DrawSimpleText([]byte(label), enums.TextEncodingUTF8, x+10, y+18, font, textPaint)
}
