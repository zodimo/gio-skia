package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/models"
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

			// Create a path with multiple contours:
			// 1. Line: moveTo(124, 108), lineTo(172, 24)
			// 2. Circle: addCircle(50, 50, 30)
			// 3. Quadratic curve: moveTo(36, 148), quadTo(66, 188, 120, 136)
			path := skia.NewPath()
			skia.PathMoveTo(path, 124, 108)
			skia.PathLineTo(path, 172, 24)
			skia.PathAddCircle(path, 50, 50, 30)
			skia.PathMoveTo(path, 36, 148)
			skia.PathQuadTo(path, 66, 188, 120, 136)

			// Create paint with anti-aliasing enabled
			paint := skia.NewPaint()
			paint.SetAntiAlias(true)

			// Draw filled path (default style is Fill)
			c.DrawPath(path, paint)

			// Set style to Stroke, color to blue, stroke width to 3
			paint.SetStyle(skia.PaintStyleStroke)
			blueColor := models.Color4f{
				R: skia.Scalar(0),
				G: skia.Scalar(0),
				B: skia.Scalar(1),
				A: skia.Scalar(1),
			}
			paint.SetColor(blueColor)
			paint.SetStrokeWidth(3)

			// Draw stroked path in blue
			c.DrawPath(path, paint)

			frameEvent.Frame(&ops)
		}
	}
}

