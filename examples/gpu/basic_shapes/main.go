package main

import (
	"image/color"
	"log"
	"math"
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

			// Light gray background
			paint.ColorOp{Color: color.NRGBA{R: 240, G: 240, B: 240, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)

			// Draw a grid of shapes showcasing different primitives
			spacing := float32(120)
			startX, startY := spacing, spacing

			// Row 1: Filled shapes
			// Filled rectangle
			c.SetColor(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			p1 := skia.NewPath()
			p1.AddRect(startX, startY, 80, 60)
			c.DrawPath(p1)

			// Filled circle
			c.SetColor(color.NRGBA{R: 100, G: 255, B: 100, A: 255})
			p2 := skia.NewPath()
			p2.AddCircle(startX+spacing, startY, 40)
			c.DrawPath(p2)

			// Filled triangle (using path)
			c.SetColor(color.NRGBA{R: 100, G: 100, B: 255, A: 255})
			p3 := skia.NewPath()
			p3.MoveTo(startX+spacing*2, startY+40)
			p3.LineTo(startX+spacing*2+40, startY+40)
			p3.LineTo(startX+spacing*2+20, startY)
			p3.Close()
			c.DrawPath(p3)

			// Filled star (5-pointed)
			c.SetColor(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			p4 := skia.NewPath()
			starCenterX := startX + spacing*3
			starCenterY := startY + 30
			outerRadius := float32(35)
			innerRadius := float32(15)
			for i := 0; i < 5; i++ {
				angle := float32(i) * math.Pi * 2 / 5
				x := starCenterX + outerRadius*float32(math.Cos(float64(angle-math.Pi/2)))
				y := starCenterY + outerRadius*float32(math.Sin(float64(angle-math.Pi/2)))
				if i == 0 {
					p4.MoveTo(x, y)
				} else {
					p4.LineTo(x, y)
				}
				// Inner point
				innerAngle := angle + math.Pi/5
				ix := starCenterX + innerRadius*float32(math.Cos(float64(innerAngle-math.Pi/2)))
				iy := starCenterY + innerRadius*float32(math.Sin(float64(innerAngle-math.Pi/2)))
				p4.LineTo(ix, iy)
			}
			p4.Close()
			c.DrawPath(p4)

			// Row 2: Stroked shapes
			startY2 := startY + spacing

			// Stroked rectangle
			c.SetStroke(skia.StrokeOpts{Width: 4, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 200, G: 50, B: 50, A: 255})
			p5 := skia.NewPath()
			p5.AddRect(startX, startY2, 80, 60)
			c.DrawPath(p5)

			// Stroked circle
			c.SetStroke(skia.StrokeOpts{Width: 5, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 50, G: 200, B: 50, A: 255})
			p6 := skia.NewPath()
			p6.AddCircle(startX+spacing, startY2+30, 40)
			c.DrawPath(p6)

			// Stroked hexagon
			c.SetStroke(skia.StrokeOpts{Width: 3, Miter: 4, Cap: stroke.SquareCap, Join: stroke.MiterJoin})
			c.SetColor(color.NRGBA{R: 50, G: 50, B: 200, A: 255})
			p7 := skia.NewPath()
			hexCenterX := startX + spacing*2 + 40
			hexCenterY := startY2 + 30
			hexRadius := float32(35)
			for i := 0; i < 6; i++ {
				angle := float32(i) * math.Pi / 3
				x := hexCenterX + hexRadius*float32(math.Cos(float64(angle)))
				y := hexCenterY + hexRadius*float32(math.Sin(float64(angle)))
				if i == 0 {
					p7.MoveTo(x, y)
				} else {
					p7.LineTo(x, y)
				}
			}
			p7.Close()
			c.DrawPath(p7)

			// Stroked arrow
			c.SetStroke(skia.StrokeOpts{Width: 6, Miter: 4, Cap: stroke.TriangularCap, Join: stroke.MiterJoin})
			c.SetColor(color.NRGBA{R: 150, G: 100, B: 200, A: 255})
			p8 := skia.NewPath()
			arrowX := startX + spacing*3
			arrowY := startY2 + 30
			p8.MoveTo(arrowX-30, arrowY)
			p8.LineTo(arrowX+30, arrowY)
			p8.MoveTo(arrowX+15, arrowY-15)
			p8.LineTo(arrowX+30, arrowY)
			p8.LineTo(arrowX+15, arrowY+15)
			c.DrawPath(p8)

			// Row 3: Complex filled shapes
			startY3 := startY + spacing*2

			// Reset to fill mode for Row 3 shapes
			c.Fill()

			// Rounded rectangle (approximated with arcs)
			c.SetColor(color.NRGBA{R: 255, G: 150, B: 200, A: 255})
			p9 := skia.NewPath()
			rectX, rectY := startX, startY3
			rectW, rectH := float32(80), float32(60)
			radius := float32(15)
			// Top-left corner
			p9.MoveTo(rectX+radius, rectY)
			p9.LineTo(rectX+rectW-radius, rectY)
			// Top-right corner (arc)
			p9.CubeTo(rectX+rectW-radius*0.552, rectY, rectX+rectW, rectY+radius*0.552, rectX+rectW, rectY+radius)
			p9.LineTo(rectX+rectW, rectY+rectH-radius)
			// Bottom-right corner (arc)
			p9.CubeTo(rectX+rectW, rectY+rectH-radius*0.552, rectX+rectW-radius*0.552, rectY+rectH, rectX+rectW-radius, rectY+rectH)
			p9.LineTo(rectX+radius, rectY+rectH)
			// Bottom-left corner (arc)
			p9.CubeTo(rectX+radius*0.552, rectY+rectH, rectX, rectY+rectH-radius*0.552, rectX, rectY+rectH-radius)
			p9.LineTo(rectX, rectY+radius)
			// Top-left corner (arc)
			p9.CubeTo(rectX, rectY+radius*0.552, rectX+radius*0.552, rectY, rectX+radius, rectY)
			p9.Close()
			c.DrawPath(p9)

			// Heart shape
			c.SetColor(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			p10 := skia.NewPath()
			heartX := startX + spacing
			heartY := startY3 + 30
			heartSize := float32(30)
			// Left curve
			p10.MoveTo(heartX, heartY+heartSize*0.3)
			p10.CubeTo(heartX, heartY, heartX-heartSize*0.5, heartY-heartSize*0.5, heartX-heartSize*0.5, heartY)
			p10.CubeTo(heartX-heartSize*0.5, heartY+heartSize*0.5, heartX, heartY+heartSize*0.8, heartX, heartY+heartSize)
			// Right curve
			p10.CubeTo(heartX, heartY+heartSize*0.8, heartX+heartSize*0.5, heartY+heartSize*0.5, heartX+heartSize*0.5, heartY)
			p10.CubeTo(heartX+heartSize*0.5, heartY-heartSize*0.5, heartX, heartY, heartX, heartY+heartSize*0.3)
			p10.Close()
			c.DrawPath(p10)

			// Spiral
			c.SetStroke(skia.StrokeOpts{Width: 2, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 100, G: 150, B: 255, A: 255})
			p11 := skia.NewPath()
			spiralX := startX + spacing*2 + 40
			spiralY := startY3 + 30
			spiralRadius := float32(5)
			for i := 0; i < 20; i++ {
				angle := float32(i) * math.Pi / 4
				radius := spiralRadius + float32(i)*2
				x := spiralX + radius*float32(math.Cos(float64(angle)))
				y := spiralY + radius*float32(math.Sin(float64(angle)))
				if i == 0 {
					p11.MoveTo(x, y)
				} else {
					p11.LineTo(x, y)
				}
			}
			c.DrawPath(p11)

			// Reset to fill mode for gear shape
			c.Fill()

			// Gear shape
			c.SetColor(color.NRGBA{R: 200, G: 200, B: 100, A: 255})
			p12 := skia.NewPath()
			gearX := startX + spacing*3
			gearY := startY3 + 30
			gearOuterRadius := float32(30)
			gearInnerRadius := float32(20)
			teeth := 12
			for i := 0; i < teeth*2; i++ {
				angle := float32(i) * math.Pi / float32(teeth)
				var radius float32
				if i%2 == 0 {
					radius = gearOuterRadius
				} else {
					radius = gearInnerRadius
				}
				x := gearX + radius*float32(math.Cos(float64(angle)))
				y := gearY + radius*float32(math.Sin(float64(angle)))
				if i == 0 {
					p12.MoveTo(x, y)
				} else {
					p12.LineTo(x, y)
				}
			}
			p12.Close()
			c.DrawPath(p12)

			frameEvent.Frame(&ops)
		}
	}
}
