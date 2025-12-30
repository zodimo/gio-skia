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
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/models"
)

// This example demonstrates canonical Skia fill rules.
// Fill rules determine how overlapping path contours are filled:
// - Winding: Non-zero winding rule (default)
// - EvenOdd: Even-odd rule (alternating)
// This is essential for complex shapes with holes or overlapping regions.

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

			// Light gray background
			paint.ColorOp{Color: color.NRGBA{R: 240, G: 240, B: 240, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			spacing := float32(200)
			startX, startY := spacing, spacing

			// Example 1: Concentric circles - Winding vs EvenOdd
			// Winding rule: All circles filled (same direction)
			c.Save()
			c.Translate(startX, startY)

			windingPath := impl.NewSkPath(enums.PathFillTypeWinding)
			// Outer circle (clockwise)
			windingPath.AddCircle(0, 0, 50, enums.PathDirectionCW)
			// Inner circle (clockwise) - creates a hole with winding rule
			windingPath.AddCircle(0, 0, 30, enums.PathDirectionCW)
			windingPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 150, B: 255, A: 255})
			c.DrawPath(windingPath, windingPaint)

			c.Restore()

			// EvenOdd rule: Alternating fill
			c.Save()
			c.Translate(startX+spacing, startY)

			evenOddPath := impl.NewSkPath(enums.PathFillTypeEvenOdd)
			// Outer circle (clockwise)
			evenOddPath.AddCircle(0, 0, 50, enums.PathDirectionCW)
			// Inner circle (clockwise) - creates a hole with even-odd rule
			evenOddPath.AddCircle(0, 0, 30, enums.PathDirectionCW)
			evenOddPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 100, A: 255})
			c.DrawPath(evenOddPath, evenOddPaint)

			c.Restore()

			// Example 2: Overlapping rectangles
			c.Save()
			c.Translate(startX, startY+spacing)

			// Winding rule - overlapping region filled
			windingRectPath := impl.NewSkPath(enums.PathFillTypeWinding)
			windingRectPath.AddRect(
				models.Rect{Left: -40, Top: -30, Right: 40, Bottom: 30},
				enums.PathDirectionCW, 0)
			windingRectPath.AddRect(
				models.Rect{Left: -30, Top: -40, Right: 30, Bottom: 40},
				enums.PathDirectionCW, 0)
			windingRectPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 150, A: 255})
			c.DrawPath(windingRectPath, windingRectPaint)

			c.Restore()

			c.Save()
			c.Translate(startX+spacing, startY+spacing)

			// EvenOdd rule - overlapping region not filled
			evenOddRectPath := impl.NewSkPath(enums.PathFillTypeEvenOdd)
			evenOddRectPath.AddRect(
				models.Rect{Left: -40, Top: -30, Right: 40, Bottom: 30},
				enums.PathDirectionCW, 0)
			evenOddRectPath.AddRect(
				models.Rect{Left: -30, Top: -40, Right: 30, Bottom: 40},
				enums.PathDirectionCW, 0)
			evenOddRectPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 200, A: 255})
			c.DrawPath(evenOddRectPath, evenOddRectPaint)

			c.Restore()

			// Example 3: Star with hole (donut shape)
			c.Save()
			c.Translate(startX, startY+spacing*2)

			// Winding rule - requires opposite direction for hole
			windingStarPath := impl.NewSkPath(enums.PathFillTypeWinding)
			starRadius := float32(40)
			innerRadius := float32(20)
			// Outer star (clockwise)
			for i := 0; i < 10; i++ {
				angle := float32(i) * math.Pi / 5
				var radius float32
				if i%2 == 0 {
					radius = starRadius
				} else {
					radius = innerRadius
				}
				x := radius * float32(math.Cos(float64(angle-math.Pi/2)))
				y := radius * float32(math.Sin(float64(angle-math.Pi/2)))
				if i == 0 {
					windingStarPath.MoveTo(models.Scalar(x), models.Scalar(y))
				} else {
					windingStarPath.LineTo(models.Scalar(x), models.Scalar(y))
				}
			}
			windingStarPath.Close()
			// Inner circle (counter-clockwise) creates hole
			windingStarPath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(15), enums.PathDirectionCCW)
			windingStarPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			c.DrawPath(windingStarPath, windingStarPaint)

			c.Restore()

			c.Save()
			c.Translate(startX+spacing, startY+spacing*2)

			// EvenOdd rule - same direction creates hole
			evenOddStarPath := impl.NewSkPath(enums.PathFillTypeEvenOdd)
			// Outer star (clockwise)
			for i := 0; i < 10; i++ {
				angle := float32(i) * math.Pi / 5
				var radius float32
				if i%2 == 0 {
					radius = starRadius
				} else {
					radius = innerRadius
				}
				x := radius * float32(math.Cos(float64(angle-math.Pi/2)))
				y := radius * float32(math.Sin(float64(angle-math.Pi/2)))
				if i == 0 {
					evenOddStarPath.MoveTo(models.Scalar(x), models.Scalar(y))
				} else {
					evenOddStarPath.LineTo(models.Scalar(x), models.Scalar(y))
				}
			}
			evenOddStarPath.Close()
			// Inner circle (clockwise) creates hole with even-odd
			evenOddStarPath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(15), enums.PathDirectionCW)
			evenOddStarPaint := skia.NewPaintFill(color.NRGBA{R: 200, G: 100, B: 255, A: 255})
			c.DrawPath(evenOddStarPath, evenOddStarPaint)

			c.Restore()

			// Example 4: Complex overlapping shapes
			c.Save()
			c.Translate(startX, startY+spacing*3)

			// Winding rule - complex overlapping
			windingComplexPath := impl.NewSkPath(enums.PathFillTypeWinding)
			// Multiple overlapping circles
			for i := 0; i < 3; i++ {
				angle := float32(i) * math.Pi * 2 / 3
				cx := 20 * float32(math.Cos(float64(angle)))
				cy := 20 * float32(math.Sin(float64(angle)))
				windingComplexPath.AddCircle(models.Scalar(cx), models.Scalar(cy), models.Scalar(30), enums.PathDirectionCW)
			}
			windingComplexPaint := skia.NewPaintFill(color.NRGBA{R: 150, G: 200, B: 255, A: 255})
			c.DrawPath(windingComplexPath, windingComplexPaint)

			c.Restore()

			c.Save()
			c.Translate(startX+spacing, startY+spacing*3)

			// EvenOdd rule - complex overlapping
			evenOddComplexPath := impl.NewSkPath(enums.PathFillTypeEvenOdd)
			// Multiple overlapping circles
			for i := 0; i < 3; i++ {
				angle := float32(i) * math.Pi * 2 / 3
				cx := 20 * float32(math.Cos(float64(angle)))
				cy := 20 * float32(math.Sin(float64(angle)))
				evenOddComplexPath.AddCircle(models.Scalar(cx), models.Scalar(cy), models.Scalar(30), enums.PathDirectionCW)
			}
			evenOddComplexPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 150, A: 255})
			c.DrawPath(evenOddComplexPath, evenOddComplexPaint)

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}
