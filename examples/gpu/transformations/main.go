package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/pkg/stroke"
	"github.com/zodimo/gio-skia/skia"
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

			// Dark background
			paint.ColorOp{Color: color.NRGBA{R: 30, G: 30, B: 40, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(frameEvent.Size.X), float32(frameEvent.Size.Y)
			centerX, centerY := w/2, h/2

			// Example 1: Nested transformations with Save/Restore
			// Draw a rotating square pattern
			c.Save()
			c.TranslateFloat32(centerX, centerY)
			for i := 0; i < 8; i++ {
				c.Save()
				c.RotateFloat32(float32(i) * 180.0 / 4) // Convert radians to degrees
				c.TranslateFloat32(80, 0)
				skPaint := skia.NewPaintFill(color.NRGBA{R: 100 + uint8(i*20), G: 150, B: 200, A: 255})
				p := skia.NewPath()
				skia.PathAddRect(p, -20, -20, 40, 40)
				c.DrawPath(p, skPaint)
				c.Restore()
			}
			c.Restore()

			// Example 2: Scaling demonstration
			c.Save()
			c.TranslateFloat32(centerX*0.3, centerY*0.3)
			scale := float32(0.8)
			for i := 0; i < 5; i++ {
				c.Save()
				c.ScaleFloat32(scale, scale)
				skPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 200 - uint8(i*30), B: 100, A: 255})
				p := skia.NewPath()
				skia.PathAddCircle(p, 0, 0, 30)
				c.DrawPath(p, skPaint)
				c.Restore()
				scale *= 0.7
			}
			c.Restore()

			// Example 3: Translation grid
			c.Save()
			c.TranslateFloat32(centerX*1.5, centerY*0.3)
			for x := 0; x < 4; x++ {
				for y := 0; y < 4; y++ {
					c.Save()
					c.TranslateFloat32(float32(x*40), float32(y*40))
					skPaint := skia.NewPaintFill(color.NRGBA{
						R: uint8(100 + x*40),
						G: uint8(100 + y*40),
						B: 150,
						A: 255,
					})
					p := skia.NewPath()
					skia.PathAddRect(p, -15, -15, 30, 30)
					c.DrawPath(p, skPaint)
					c.Restore()
				}
			}
			c.Restore()

			// Example 4: Rotation around different points
			c.Save()
			c.TranslateFloat32(centerX*0.3, centerY*1.5)
			for i := 0; i < 12; i++ {
				c.Save()
				c.RotateFloat32(float32(i) * 180.0 / 6) // Convert radians to degrees
				c.TranslateFloat32(50, 0)
				skPaint := skia.NewPaintFill(color.NRGBA{
					R: uint8(150 + i*8),
					G: uint8(200 - i*10),
					B: 255,
					A: 255,
				})
				p := skia.NewPath()
				skia.PathAddRect(p, -10, -10, 20, 20)
				c.DrawPath(p, skPaint)
				c.Restore()
			}
			c.Restore()

			// Example 5: Combined transformations - Spiral pattern
			c.Save()
			c.TranslateFloat32(centerX*1.5, centerY*1.5)
			for i := 0; i < 20; i++ {
				c.Save()
				angle := float32(i) * 180.0 / 10 // Already in degrees
				radius := float32(i * 8)
				c.RotateFloat32(angle)
				c.TranslateFloat32(radius, 0)
				c.ScaleFloat32(1.0-float32(i)*0.03, 1.0-float32(i)*0.03)
				skPaint := skia.NewPaintFill(color.NRGBA{
					R: uint8(255 - i*10),
					G: uint8(100 + i*5),
					B: uint8(150 + i*5),
					A: 255,
				})
				p := skia.NewPath()
				skia.PathAddCircle(p, 0, 0, 15)
				c.DrawPath(p, skPaint)
				c.Restore()
			}
			c.Restore()

			// Example 6: Nested Save/Restore with different stroke styles
			c.Save()
			c.TranslateFloat32(centerX, centerY*1.8)
			strokeOpts := stroke.StrokeOpts{Width: 2, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin}
			skPaint := skia.NewPaintStroke(color.NRGBA{R: 255, G: 255, B: 255, A: 255}, 2)
			skPaint = skia.ConfigureStrokePaint(skPaint, strokeOpts)
			
			// Outer ring
			p1 := skia.NewPath()
			skia.PathAddCircle(p1, 0, 0, 60)
			c.DrawPath(p1, skPaint)
			
			// Middle ring
			c.Save()
			c.ScaleFloat32(0.6, 0.6)
			p2 := skia.NewPath()
			skia.PathAddCircle(p2, 0, 0, 60)
			c.DrawPath(p2, skPaint)
			
			// Inner ring
			c.Save()
			c.ScaleFloat32(0.5, 0.5)
			p3 := skia.NewPath()
			skia.PathAddCircle(p3, 0, 0, 60)
			c.DrawPath(p3, skPaint)
			c.Restore()
			
			c.Restore()
			c.Restore()

			// Example 7: Transform chain - rotating squares
			c.Save()
			c.TranslateFloat32(centerX*0.2, centerY*1.8)
			for i := 0; i < 6; i++ {
				c.Save()
				c.RotateFloat32(float32(i) * 180.0 / 3) // Convert radians to degrees
				c.TranslateFloat32(40, 0)
				c.RotateFloat32(float32(i) * 180.0 / 6) // Convert radians to degrees
				skPaint := skia.NewPaintFill(color.NRGBA{
					R: uint8(200 + i*8),
					G: uint8(150 - i*10),
					B: uint8(100 + i*15),
					A: 255,
				})
				p := skia.NewPath()
				skia.PathAddRect(p, -12, -12, 24, 24)
				c.DrawPath(p, skPaint)
				c.Restore()
			}
			c.Restore()

			// Example 8: Skew-like effect using rotation and scale
			c.Save()
			c.TranslateFloat32(centerX*1.8, centerY*1.8)
			for i := 0; i < 5; i++ {
				c.Save()
				c.RotateFloat32(float32(i) * 180.0 / 8) // Convert radians to degrees
				c.ScaleFloat32(1.0+float32(i)*0.1, 1.0-float32(i)*0.05)
				skPaint := skia.NewPaintFill(color.NRGBA{
					R: uint8(150 + i*20),
					G: uint8(200 - i*15),
					B: uint8(255 - i*10),
					A: 255,
				})
				p := skia.NewPath()
				skia.PathAddRect(p, -20, -20, 40, 40)
				c.DrawPath(p, skPaint)
				c.Restore()
			}
			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}

