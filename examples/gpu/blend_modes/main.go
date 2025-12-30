package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/enums"
)

// This example demonstrates canonical Skia blend modes.
// Blend modes control how colors are composited when drawing overlapping shapes.
// This is a fundamental feature in Skia for creating complex visual effects.

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
			spacing := float32(150)
			startX, startY := spacing, spacing

			// Define common Porter-Duff blend modes to demonstrate
			blendModes := []struct {
				mode  enums.BlendMode
				name  string
				color color.NRGBA
			}{
				{enums.BlendModeSrcOver, "SrcOver", color.NRGBA{R: 255, G: 100, B: 100, A: 200}},
				{enums.BlendModeMultiply, "Multiply", color.NRGBA{R: 100, G: 200, B: 255, A: 200}},
				{enums.BlendModeScreen, "Screen", color.NRGBA{R: 100, G: 255, B: 100, A: 200}},
				{enums.BlendModeOverlay, "Overlay", color.NRGBA{R: 255, G: 200, B: 100, A: 200}},
				{enums.BlendModeDarken, "Darken", color.NRGBA{R: 200, G: 100, B: 255, A: 200}},
				{enums.BlendModeLighten, "Lighten", color.NRGBA{R: 100, G: 255, B: 255, A: 200}},
				{enums.BlendModeColorDodge, "ColorDodge", color.NRGBA{R: 255, G: 150, B: 150, A: 200}},
				{enums.BlendModeColorBurn, "ColorBurn", color.NRGBA{R: 150, G: 150, B: 255, A: 200}},
				{enums.BlendModeHardLight, "HardLight", color.NRGBA{R: 255, G: 100, B: 200, A: 200}},
				{enums.BlendModeSoftLight, "SoftLight", color.NRGBA{R: 200, G: 255, B: 100, A: 200}},
				{enums.BlendModeDifference, "Difference", color.NRGBA{R: 255, G: 255, B: 100, A: 200}},
				{enums.BlendModeExclusion, "Exclusion", color.NRGBA{R: 100, G: 200, B: 200, A: 200}},
			}

			// Draw base shapes (destination)
			baseColor := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
			basePaint := skia.NewPaintFill(baseColor)

			// Draw blend mode demonstrations in a grid
			cols := 4
			for i, bm := range blendModes {
				row := i / cols
				col := i % cols
				x := startX + float32(col)*spacing
				y := startY + float32(row)*spacing

				c.Save()
				c.Translate(x, y)

				// Draw base rectangle (destination)
				basePath := skia.NewPath()
				skia.PathAddRect(basePath, -40, -40, 80, 80)
				c.DrawPath(basePath, basePaint)

				// Draw overlay shape (source) with blend mode
				overlayPaint := skia.NewPaintFill(bm.color)
				overlayPaint.SetBlendMode(bm.mode)
				overlayPath := skia.NewPath()
				skia.PathAddCircle(overlayPath, 0, 0, 35)
				c.DrawPath(overlayPath, overlayPaint)

				c.Restore()
			}

			// Example: Multiple overlapping shapes with different blend modes
			c.Save()
			c.Translate(w/2, h*0.75)

			// Base circle
			baseCircle := skia.NewPath()
			skia.PathAddCircle(baseCircle, 0, 0, 60)
			basePaint = skia.NewPaintFill(color.NRGBA{R: 150, G: 150, B: 255, A: 255})
			c.DrawPath(baseCircle, basePaint)

			// Overlay circles with different blend modes
			overlays := []struct {
				x, y float32
				mode enums.BlendMode
				col  color.NRGBA
			}{
				{-30, -30, enums.BlendModeMultiply, color.NRGBA{R: 255, G: 100, B: 100, A: 200}},
				{30, -30, enums.BlendModeScreen, color.NRGBA{R: 100, G: 255, B: 100, A: 200}},
				{-30, 30, enums.BlendModeOverlay, color.NRGBA{R: 255, G: 255, B: 100, A: 200}},
				{30, 30, enums.BlendModeDifference, color.NRGBA{R: 255, G: 100, B: 255, A: 200}},
			}

			for _, overlay := range overlays {
				c.Save()
				c.Translate(overlay.x, overlay.y)
				overlayPath := skia.NewPath()
				skia.PathAddCircle(overlayPath, 0, 0, 35)
				overlayPaint := skia.NewPaintFill(overlay.col)
				overlayPaint.SetBlendMode(overlay.mode)
				c.DrawPath(overlayPath, overlayPaint)
				c.Restore()
			}

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}
