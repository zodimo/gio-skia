// Package main demonstrates the SkCanvas clipping methods.
// It shows ClipRect, ClipRRect, and ClipPath with animated content.
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
	"github.com/zodimo/go-skia-support/skia/impl"
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
			paint.ColorOp{Color: color.NRGBA{R: 20, G: 25, B: 35, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(e.Size.X), float32(e.Size.Y)

			// ─────────────────────────────────────────────────────────
			// Demo 1: ClipRect - Rectangular clipping window
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.2, h*0.25)

			// Draw a label
			drawLabel(c, "ClipRect", -60, -80)

			// Apply rectangular clip
			clipRect := models.Rect{Left: -50, Top: -50, Right: 50, Bottom: 50}
			c.ClipRect(clipRect, enums.ClipOpIntersect, true)

			// Draw animated content inside clip
			// Rotating gradient of circles
			for i := 0; i < 8; i++ {
				angle := float32(i)*math.Pi/4 + float32(elapsed)
				radius := float32(60)
				x := radius * float32(math.Cos(float64(angle)))
				y := radius * float32(math.Sin(float64(angle)))

				circleColor := color.NRGBA{
					R: uint8(100 + i*20),
					G: uint8(150 + i*10),
					B: 255,
					A: 255,
				}
				p := skia.NewPaintFill(circleColor)
				circle := skia.NewPath()
				skia.PathAddCircle(circle, x, y, 25)
				c.DrawPath(circle, p)
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 2: ClipRRect - Rounded rectangle clipping
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.5, h*0.25)

			drawLabel(c, "ClipRRect", -60, -80)

			// Apply rounded rect clip
			var clipRRect models.RRect
			clipRRect.SetRectXY(models.Rect{Left: -60, Top: -40, Right: 60, Bottom: 40}, 15, 15)
			c.ClipRRect(clipRRect, enums.ClipOpIntersect, true)

			// Draw animated stripes
			for i := -10; i < 10; i++ {
				offset := float32(math.Mod(elapsed*50+float64(i*20), 200) - 100)
				stripeColor := color.NRGBA{R: 255, G: 100, B: 100, A: 200}
				if i%2 == 0 {
					stripeColor = color.NRGBA{R: 255, G: 200, B: 100, A: 200}
				}
				p := skia.NewPaintFill(stripeColor)
				stripe := skia.NewPath()
				skia.PathMoveTo(stripe, offset-10, -50)
				skia.PathLineTo(stripe, offset+10, -50)
				skia.PathLineTo(stripe, offset+10, 50)
				skia.PathLineTo(stripe, offset-10, 50)
				stripe.Close()
				c.DrawPath(stripe, p)
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 3: ClipPath - Star-shaped clipping
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.8, h*0.25)

			drawLabel(c, "ClipPath", -50, -80)

			// Create star-shaped clip path
			starPath := impl.NewSkPath(enums.PathFillTypeWinding)
			starRadius := float32(50)
			innerRadius := float32(25)
			for i := 0; i < 5; i++ {
				outerAngle := float32(i)*math.Pi*2/5 - math.Pi/2
				innerAngle := outerAngle + math.Pi/5

				ox := starRadius * float32(math.Cos(float64(outerAngle)))
				oy := starRadius * float32(math.Sin(float64(outerAngle)))
				ix := innerRadius * float32(math.Cos(float64(innerAngle)))
				iy := innerRadius * float32(math.Sin(float64(innerAngle)))

				if i == 0 {
					starPath.MoveTo(ox, oy)
				} else {
					starPath.LineTo(ox, oy)
				}
				starPath.LineTo(ix, iy)
			}
			starPath.Close()
			c.ClipPath(starPath, enums.ClipOpIntersect, true)

			// Draw pulsing gradient inside star
			scale := 0.5 + 0.5*float32(math.Sin(elapsed*3))
			for ring := 5; ring > 0; ring-- {
				ringPath := skia.NewPath()
				skia.PathAddCircle(ringPath, 0, 0, float32(ring)*15*scale)
				ringColor := color.NRGBA{
					R: uint8(50 + ring*40),
					G: uint8(100 + ring*30),
					B: uint8(200 - ring*20),
					A: 255,
				}
				c.DrawPath(ringPath, skia.NewPaintFill(ringColor))
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 4: Combined clipping and transforms
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.35, h*0.7)
			c.Rotate(float32(elapsed * 30)) // Rotate the whole demo

			drawLabel(c, "Clip + Rotate", -70, -100)

			// Clip to circle
			clipCircle := skia.NewPath()
			skia.PathAddCircle(clipCircle, 0, 0, 60)
			c.ClipPath(clipCircle, enums.ClipOpIntersect, true)

			// Draw checkerboard pattern
			tileSize := float32(20)
			for row := -4; row <= 4; row++ {
				for col := -4; col <= 4; col++ {
					tileColor := color.NRGBA{R: 50, G: 50, B: 80, A: 255}
					if (row+col)%2 == 0 {
						tileColor = color.NRGBA{R: 200, G: 200, B: 230, A: 255}
					}
					tile := skia.NewPath()
					x := float32(col) * tileSize
					y := float32(row) * tileSize
					skia.PathMoveTo(tile, x, y)
					skia.PathLineTo(tile, x+tileSize, y)
					skia.PathLineTo(tile, x+tileSize, y+tileSize)
					skia.PathLineTo(tile, x, y+tileSize)
					tile.Close()
					c.DrawPath(tile, skia.NewPaintFill(tileColor))
				}
			}
			c.Restore()

			// ─────────────────────────────────────────────────────────
			// Demo 5: Animated clip bounds
			// ─────────────────────────────────────────────────────────
			c.Save()
			c.Translate(w*0.7, h*0.7)

			drawLabel(c, "Animated Clip", -70, -100)

			// Animate clip size
			clipSize := 30 + 30*float32(math.Sin(elapsed*2))
			animRect := models.Rect{Left: -clipSize, Top: -clipSize, Right: clipSize, Bottom: clipSize}
			c.ClipRect(animRect, enums.ClipOpIntersect, true)

			// Draw static content that gets revealed/hidden
			p := skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 150, A: 255})
			bigCircle := skia.NewPath()
			skia.PathAddCircle(bigCircle, 0, 0, 80)
			c.DrawPath(bigCircle, p)

			p = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			innerCircle := skia.NewPath()
			skia.PathAddCircle(innerCircle, 0, 0, 40)
			c.DrawPath(innerCircle, p)
			c.Restore()

			window.Invalidate()
			e.Frame(&ops)
		}
	}
}

// drawLabel is a placeholder - in a real app you'd use text rendering
func drawLabel(c skia.Canvas, label string, x, y float32) {
	// Draw a simple colored box as label placeholder
	// Text rendering would go here with DrawString
	_ = label
	p := skia.NewPaintFill(color.NRGBA{R: 255, G: 255, B: 255, A: 100})
	box := skia.NewPath()
	skia.PathMoveTo(box, x, y)
	skia.PathLineTo(box, x+120, y)
	skia.PathLineTo(box, x+120, y+20)
	skia.PathLineTo(box, x, y+20)
	box.Close()
	c.DrawPath(box, p)
}
