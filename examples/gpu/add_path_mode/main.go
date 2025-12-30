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
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/models"
)

// This example demonstrates the difference between AddPathModeAppend and AddPathModeExtend
// when combining paths. The example is based on the Skia fiddle code that shows how
// paths are combined differently depending on the mode:
// - Append: Adds the second path as a new contour
// - Extend: Connects the second path to the first if they share an endpoint

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

			// Create path1: triangle shape
			path1 := impl.NewSkPath(enums.PathFillTypeWinding)
			path1.MoveTo(models.Scalar(20), models.Scalar(20))
			path1.LineTo(models.Scalar(20), models.Scalar(40))
			path1.LineTo(models.Scalar(40), models.Scalar(20))

			// Create path2: L-shape
			path2 := impl.NewSkPath(enums.PathFillTypeWinding)
			path2.MoveTo(models.Scalar(60), models.Scalar(60))
			path2.LineTo(models.Scalar(80), models.Scalar(60))
			path2.LineTo(models.Scalar(80), models.Scalar(40))

			// Create stroke paint
			paint := skia.NewPaintStroke(color.NRGBA{R: 0, G: 0, B: 0, A: 255}, 2)

			// Loop twice: once with path1 open, once with path1 closed
			for i := 0; i < 2; i++ {
				// Test both AddPathMode options
				for _, addPathMode := range []enums.AddPathMode{enums.AddPathModeAppend, enums.AddPathModeExtend} {
					// Create a new path based on path1
					testPath := impl.NewSkPath(enums.PathFillTypeWinding)
					testPath.MoveTo(models.Scalar(20), models.Scalar(20))
					testPath.LineTo(models.Scalar(20), models.Scalar(40))
					testPath.LineTo(models.Scalar(40), models.Scalar(20))
					if i == 1 {
						// Close path1 for the second iteration
						testPath.Close()
					}

					// Add path2 using the specified mode
					testPath.AddPath(path2, models.Scalar(0), models.Scalar(0), addPathMode)

					// Draw the combined path
					c.DrawPath(testPath, paint)

					// Translate to the right for the next drawing
					c.Translate(100, 0)
				}

				// Move back left and down for the next row
				c.Translate(-200, 100)
			}

			frameEvent.Frame(&ops)
		}
	}
}
