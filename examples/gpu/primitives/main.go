// Package main demonstrates the SkCanvas drawing primitive methods.
// It shows DrawRect, DrawRRect, DrawDRRect, DrawOval, DrawCircle, DrawArc,
// DrawPoints, and DrawLine.
package main

import (
	"image/color"
	"log"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/enums"
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
	startTime := time.Now()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()
			elapsed := time.Since(startTime).Seconds()

			// Dark background
			paint.ColorOp{Color: color.NRGBA{R: 20, G: 25, B: 40, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(e.Size.X), float32(e.Size.Y)

			// ─────────────────────────────────────────────────────────
			// Row 1: Rectangles
			// ─────────────────────────────────────────────────────────

			// DrawRect
			c.Save()
			c.Translate(w*0.15, h*0.15)
			rect := models.Rect{Left: -40, Top: -30, Right: 40, Bottom: 30}
			p := skia.NewPaintFill(color.NRGBA{R: 100, G: 200, B: 255, A: 255})
			c.DrawRect(rect, p)
			c.Restore()

			// DrawRRect (rounded rectangle)
			c.Save()
			c.Translate(w*0.4, h*0.15)
			var rrect models.RRect
			rrect.SetRectXY(models.Rect{Left: -45, Top: -30, Right: 45, Bottom: 30}, 12, 12)
			p = skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 100, A: 255})
			c.DrawRRect(rrect, p)
			c.Restore()

			// DrawDRRect (donut shape)
			c.Save()
			c.Translate(w*0.65, h*0.15)
			c.Rotate(float32(elapsed * 30))
			var outer, inner models.RRect
			outer.SetRectXY(models.Rect{Left: -50, Top: -50, Right: 50, Bottom: 50}, 15, 15)
			inner.SetRectXY(models.Rect{Left: -25, Top: -25, Right: 25, Bottom: 25}, 8, 8)
			p = skia.NewPaintFill(color.NRGBA{R: 150, G: 100, B: 255, A: 255})
			c.DrawDRRect(outer, inner, p)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Row 2: Circles and Ovals
			// ─────────────────────────────────────────────────────────

			// DrawCircle
			c.Save()
			c.Translate(w*0.15, h*0.4)
			scale := 0.7 + 0.3*float32(math.Sin(elapsed*2))
			p = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 150, A: 255})
			c.DrawCircle(models.Point{X: 0, Y: 0}, 40*scale, p)
			c.Restore()

			// DrawOval
			c.Save()
			c.Translate(w*0.4, h*0.4)
			c.Rotate(float32(elapsed * 45))
			oval := models.Rect{Left: -50, Top: -30, Right: 50, Bottom: 30}
			p = skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 150, A: 255})
			c.DrawOval(oval, p)
			c.Restore()

			// DrawArc (without center)
			c.Save()
			c.Translate(w*0.65, h*0.4)
			arcOval := models.Rect{Left: -40, Top: -40, Right: 40, Bottom: 40}
			sweepAngle := 180 + 90*float32(math.Sin(elapsed))
			p = skia.NewPaintStroke(color.NRGBA{R: 255, G: 200, B: 50, A: 255}, 6)
			c.DrawArc(arcOval, float32(elapsed*60), sweepAngle, false, p)
			c.Restore()

			// DrawArc (with center - pie wedge)
			c.Save()
			c.Translate(w*0.85, h*0.4)
			wedgeSweep := 45 + 45*float32(math.Sin(elapsed*1.5))
			p = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			c.DrawArc(arcOval, float32(-elapsed*30), wedgeSweep, true, p)
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Row 3: Points and Lines
			// ─────────────────────────────────────────────────────────

			// DrawPoints - Points mode
			c.Save()
			c.Translate(w*0.15, h*0.65)
			points := make([]models.Point, 12)
			for i := range points {
				angle := float32(i) * math.Pi * 2 / 12
				points[i] = models.Point{
					X: 35 * float32(math.Cos(float64(angle)+elapsed)),
					Y: 35 * float32(math.Sin(float64(angle)+elapsed)),
				}
			}
			p = skia.NewPaintStroke(color.NRGBA{R: 255, G: 255, B: 100, A: 255}, 8)
			c.DrawPoints(enums.PointModePoints, points, p)
			c.Restore()

			// DrawPoints - Lines mode
			c.Save()
			c.Translate(w*0.4, h*0.65)
			linePoints := make([]models.Point, 6)
			for i := range linePoints {
				offset := float32(math.Sin(elapsed+float64(i)*0.5)) * 15
				linePoints[i] = models.Point{
					X: float32(i-3) * 15,
					Y: offset,
				}
			}
			p = skia.NewPaintStroke(color.NRGBA{R: 100, G: 200, B: 255, A: 255}, 3)
			c.DrawPoints(enums.PointModeLines, linePoints, p)
			c.Restore()

			// DrawPoints - Polygon mode
			c.Save()
			c.Translate(w*0.65, h*0.65)
			polyPoints := make([]models.Point, 5)
			for i := range polyPoints {
				angle := float32(i)*math.Pi*2/5 - math.Pi/2
				polyPoints[i] = models.Point{
					X: 35 * float32(math.Cos(float64(angle))),
					Y: 35 * float32(math.Sin(float64(angle))),
				}
			}
			p = skia.NewPaintStroke(color.NRGBA{R: 255, G: 150, B: 255, A: 255}, 4)
			c.DrawPoints(enums.PointModePolygon, polyPoints, p)
			c.Restore()

			// DrawLine
			c.Save()
			c.Translate(w*0.85, h*0.65)
			for i := 0; i < 8; i++ {
				angle := float32(i)*math.Pi/4 + float32(elapsed)
				x := 40 * float32(math.Cos(float64(angle)))
				y := 40 * float32(math.Sin(float64(angle)))
				lineColor := color.NRGBA{
					R: uint8(100 + i*20),
					G: uint8(200 - i*20),
					B: 255,
					A: 255,
				}
				p = skia.NewPaintStroke(lineColor, 3)
				c.DrawLine(models.Point{X: 0, Y: 0}, models.Point{X: x, Y: y}, p)
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Row 4: DrawPaint (fills entire clip)
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.5, h*0.85)

			// Create a clip, then use DrawPaint to fill it
			clipPath := skia.NewPath()
			skia.PathAddCircle(clipPath, 0, 0, 50)
			c.ClipPath(clipPath, enums.ClipOpIntersect, true)

			// Animate color
			red := uint8(128 + 127*math.Sin(elapsed))
			green := uint8(128 + 127*math.Sin(elapsed+2))
			blue := uint8(128 + 127*math.Sin(elapsed+4))
			p = skia.NewPaintFill(color.NRGBA{R: red, G: green, B: blue, A: 255})
			c.DrawPaint(p)
			c.Restore()

			window.Invalidate()
			e.Frame(&ops)
		}
	}
}
