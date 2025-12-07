package main

import (
	"image/color"
	"log"
	"math"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
)

// This example demonstrates canonical Skia transparency and alpha compositing.
// Transparency is fundamental to Skia's rendering model, allowing for
// layered drawing with varying opacity levels.

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
	for {
		switch frameEvent := window.Event().(type) {
		case app.DestroyEvent:
			return frameEvent.Err
		case app.FrameEvent:
			ops.Reset()

			// White background
			paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(frameEvent.Size.X), float32(frameEvent.Size.Y)
			spacing := float32(180)
			startX, startY := spacing, spacing

			// Example 1: Varying alpha levels
			c.Save()
			c.TranslateFloat32(startX, startY)

			alphas := []uint8{255, 200, 150, 100, 50}
			for i, alpha := range alphas {
				p := skia.NewPath()
				skia.PathAddCircle(p, float32(i)*40-80, 0, 25)
				paint := skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 100, A: alpha})
				c.DrawPath(p, paint)
			}

			c.Restore()

			// Example 2: Overlapping transparent shapes
			c.Save()
			c.TranslateFloat32(startX+spacing, startY)

			// Base circle
			baseCircle := skia.NewPath()
			skia.PathAddCircle(baseCircle, 0, 0, 40)
			basePaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 150, B: 255, A: 200})
			c.DrawPath(baseCircle, basePaint)

			// Overlapping circles with transparency
			for i := 0; i < 3; i++ {
				angle := float32(i) * math.Pi * 2 / 3
				cx := 25 * float32(math.Cos(float64(angle)))
				cy := 25 * float32(math.Sin(float64(angle)))
				overlayCircle := skia.NewPath()
				skia.PathAddCircle(overlayCircle, cx, cy, 30)
				overlayPaint := skia.NewPaintFill(color.NRGBA{
					R: uint8(255 - i*50),
					G: uint8(100 + i*50),
					B: uint8(150 + i*30),
					A: 180,
				})
				c.DrawPath(overlayCircle, overlayPaint)
			}

			c.Restore()

			// Example 3: Layered transparency with different colors
			c.Save()
			c.TranslateFloat32(startX, startY+spacing)

			layers := []struct {
				x, y, r float32
				col     color.NRGBA
			}{
				{0, 0, 35, color.NRGBA{R: 255, G: 100, B: 100, A: 200}},
				{-20, -20, 30, color.NRGBA{R: 100, G: 255, B: 100, A: 200}},
				{20, -20, 30, color.NRGBA{R: 100, G: 100, B: 255, A: 200}},
				{0, 20, 30, color.NRGBA{R: 255, G: 255, B: 100, A: 200}},
			}

			for _, layer := range layers {
				p := skia.NewPath()
				skia.PathAddCircle(p, layer.x, layer.y, layer.r)
				paint := skia.NewPaintFill(layer.col)
				c.DrawPath(p, paint)
			}

			c.Restore()

			// Example 4: Transparency gradient effect
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing)

			// Draw overlapping circles with decreasing alpha
			for i := 0; i < 5; i++ {
				p := skia.NewPath()
				skia.PathAddCircle(p, 0, 0, float32(40-i*6))
				alpha := uint8(255 - i*40)
				paint := skia.NewPaintFill(color.NRGBA{R: 150, G: 200, B: 255, A: alpha})
				c.DrawPath(p, paint)
			}

			c.Restore()

			// Example 5: Semi-transparent overlay on background
			c.Save()
			c.TranslateFloat32(startX, startY+spacing*2)

			// Background pattern
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					p := skia.NewPath()
					skia.PathAddRect(p, float32(i-1)*30, float32(j-1)*30, 25, 25)
					paint := skia.NewPaintFill(color.NRGBA{
						R: uint8(100 + i*50),
						G: uint8(100 + j*50),
						B: 200,
						A: 255,
					})
					c.DrawPath(p, paint)
				}
			}

			// Semi-transparent overlay
			overlayPath := skia.NewPath()
			skia.PathAddCircle(overlayPath, 0, 0, 50)
			overlayPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 255, B: 255, A: 150})
			c.DrawPath(overlayPath, overlayPaint)

			c.Restore()

			// Example 6: Complex transparency composition
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing*2)

			// Base shape
			basePath := skia.NewPath()
			skia.PathAddRect(basePath, -40, -40, 80, 80)
			basePaint = skia.NewPaintFill(color.NRGBA{R: 200, G: 200, B: 200, A: 255})
			c.DrawPath(basePath, basePaint)

			// Transparent overlays
			for i := 0; i < 4; i++ {
				angle := float32(i) * math.Pi / 2
				cx := 20 * float32(math.Cos(float64(angle)))
				cy := 20 * float32(math.Sin(float64(angle)))
				p := skia.NewPath()
				skia.PathAddCircle(p, cx, cy, 25)
				paint := skia.NewPaintFill(color.NRGBA{
					R: uint8(255 - i*40),
					G: uint8(150 + i*30),
					B: uint8(100 + i*50),
					A: 180,
				})
				c.DrawPath(p, paint)
			}

			c.Restore()

			// Example 7: Alpha blending demonstration
			c.Save()
			c.TranslateFloat32(w/2, h*0.75)

			// Create a grid of overlapping shapes with different alpha values
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					c.Save()
					c.TranslateFloat32(float32(i-1)*60, float32(j-1)*60)

					p := skia.NewPath()
					skia.PathAddCircle(p, 0, 0, 30)
					alpha := uint8(100 + (i+j)*30)
					paint := skia.NewPaintFill(color.NRGBA{
						R: uint8(255 - i*80),
						G: uint8(255 - j*80),
						B: 200,
						A: alpha,
					})
					c.DrawPath(p, paint)
					c.Restore()
				}
			}

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}

