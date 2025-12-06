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

			// Dark background
			paint.ColorOp{Color: color.NRGBA{R: 20, G: 20, B: 30, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			spacing := float32(200)
			startX, startY := spacing, spacing

			// Example 1: Simple Quadratic Bézier Curves
			c.Save()
			c.Translate(startX, startY)

			// Draw control points and lines
			c.SetStroke(skia.StrokeOpts{Width: 1, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255})
			p0 := skia.NewPath()
			p0.MoveTo(0, 50)
			p0.LineTo(30, 0)
			p0.LineTo(60, 50)
			c.DrawPath(p0)

			// Draw control points
			c.Fill()
			c.SetColor(color.NRGBA{R: 255, G: 255, B: 0, A: 255})
			cp1 := skia.NewPath()
			cp1.AddCircle(0, 50, 3)
			c.DrawPath(cp1)
			cp2 := skia.NewPath()
			cp2.AddCircle(30, 0, 3)
			c.DrawPath(cp2)
			cp3 := skia.NewPath()
			cp3.AddCircle(60, 50, 3)
			c.DrawPath(cp3)

			// Draw the curve
			c.SetStroke(skia.StrokeOpts{Width: 3, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 100, G: 200, B: 255, A: 255})
			curve1 := skia.NewPath()
			curve1.MoveTo(0, 50)
			curve1.QuadTo(30, 0, 60, 50)
			c.DrawPath(curve1)

			c.Restore()

			// Example 2: Cubic Bézier Curves
			c.Save()
			c.Translate(startX+spacing, startY)

			// Draw control lines
			c.SetStroke(skia.StrokeOpts{Width: 1, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255})
			p1 := skia.NewPath()
			p1.MoveTo(0, 50)
			p1.LineTo(20, 0)
			p1.MoveTo(40, 0)
			p1.LineTo(60, 50)
			c.DrawPath(p1)

			// Draw control points
			c.Fill()
			c.SetColor(color.NRGBA{R: 255, G: 200, B: 0, A: 255})
			points := []struct{ x, y float32 }{
				{0, 50}, {20, 0}, {40, 0}, {60, 50},
			}
			for _, pt := range points {
				cp := skia.NewPath()
				cp.AddCircle(pt.x, pt.y, 3)
				c.DrawPath(cp)
			}

			// Draw the cubic curve
			c.SetStroke(skia.StrokeOpts{Width: 3, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 255, G: 100, B: 200, A: 255})
			curve2 := skia.NewPath()
			curve2.MoveTo(0, 50)
			curve2.CubeTo(20, 0, 40, 0, 60, 50)
			c.DrawPath(curve2)

			c.Restore()

			// Example 3: Smooth Wave Pattern (Connected Quadratic Curves)
			c.Save()
			c.Translate(startX, startY+spacing)

			c.SetStroke(skia.StrokeOpts{Width: 4, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 100, G: 255, B: 150, A: 255})
			wave := skia.NewPath()
			wave.MoveTo(0, 25)
			for i := 0; i < 5; i++ {
				x1 := float32(i*30 + 15)
				y1 := float32(25 - 15*float32(math.Pow(-1, float64(i))))
				x2 := float32((i + 1) * 30)
				y2 := float32(25)
				wave.QuadTo(x1, y1, x2, y2)
			}
			c.DrawPath(wave)

			c.Restore()

			// Example 4: Complex Cubic Curve - Infinity Symbol
			c.Save()
			c.Translate(startX+spacing, startY+spacing)

			c.SetStroke(skia.StrokeOpts{Width: 5, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 255, G: 150, B: 100, A: 255})
			infinity := skia.NewPath()
			infinity.MoveTo(0, 25)
			// Left loop
			infinity.CubeTo(15, 0, 15, 0, 30, 25)
			// Right loop
			infinity.CubeTo(45, 50, 45, 50, 60, 25)
			c.DrawPath(infinity)

			c.Restore()

			// Example 5: Flower Petals using Bézier Curves
			c.Save()
			c.Translate(startX, startY+spacing*2)

			c.SetColor(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			flower := skia.NewPath()
			centerX, centerY := float32(40), float32(40)
			petalRadius := float32(30)
			petals := 8

			for i := 0; i < petals; i++ {
				angle := float32(i) * math.Pi * 2 / float32(petals)
				petalX := centerX + petalRadius*float32(math.Cos(float64(angle)))
				petalY := centerY + petalRadius*float32(math.Sin(float64(angle)))

				// Control points for smooth petal shape
				ctrl1X := centerX + petalRadius*0.7*float32(math.Cos(float64(angle-math.Pi/float32(petals))))
				ctrl1Y := centerY + petalRadius*0.7*float32(math.Sin(float64(angle-math.Pi/float32(petals))))
				ctrl2X := centerX + petalRadius*0.7*float32(math.Cos(float64(angle+math.Pi/float32(petals))))
				ctrl2Y := centerY + petalRadius*0.7*float32(math.Sin(float64(angle+math.Pi/float32(petals))))

				if i == 0 {
					flower.MoveTo(petalX, petalY)
				} else {
					flower.CubeTo(ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, petalX, petalY)
				}
			}
			flower.Close()
			c.DrawPath(flower)

			// Center circle
			c.SetColor(color.NRGBA{R: 255, G: 150, B: 50, A: 255})
			center := skia.NewPath()
			center.AddCircle(centerX, centerY, 8)
			c.DrawPath(center)

			c.Restore()

			// Example 6: Spiral using Cubic Bézier Segments
			c.Save()
			c.Translate(startX+spacing, startY+spacing*2)

			c.SetStroke(skia.StrokeOpts{Width: 3, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			c.SetColor(color.NRGBA{R: 150, G: 200, B: 255, A: 255})
			spiral := skia.NewPath()
			spiralCenterX, spiralCenterY := float32(40), float32(40)
			spiral.MoveTo(spiralCenterX, spiralCenterY)

			for i := 0; i < 12; i++ {
				angle1 := float32(i) * math.Pi / 6
				angle2 := float32(i+1) * math.Pi / 6
				radius1 := float32(5 + i*3)
				radius2 := float32(5 + (i+1)*3)

				x2 := spiralCenterX + radius2*float32(math.Cos(float64(angle2)))
				y2 := spiralCenterY + radius2*float32(math.Sin(float64(angle2)))

				// Control points for smooth curve
				ctrl1X := spiralCenterX + radius1*1.2*float32(math.Cos(float64(angle1+math.Pi/12)))
				ctrl1Y := spiralCenterY + radius1*1.2*float32(math.Sin(float64(angle1+math.Pi/12)))
				ctrl2X := spiralCenterX + radius2*1.2*float32(math.Cos(float64(angle2-math.Pi/12)))
				ctrl2Y := spiralCenterY + radius2*1.2*float32(math.Sin(float64(angle2-math.Pi/12)))

				spiral.CubeTo(ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, x2, y2)
			}
			c.DrawPath(spiral)

			c.Restore()

			// Example 7: Heart Shape with Bézier Curves
			c.Save()
			c.Translate(startX, startY+spacing*3)

			c.SetColor(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			heart := skia.NewPath()
			heartX, heartY := float32(40), float32(35)
			heartSize := float32(25)

			// Left curve
			heart.MoveTo(heartX, heartY+heartSize*0.3)
			heart.CubeTo(
				heartX, heartY,
				heartX-heartSize*0.5, heartY-heartSize*0.5,
				heartX-heartSize*0.5, heartY,
			)
			heart.CubeTo(
				heartX-heartSize*0.5, heartY+heartSize*0.5,
				heartX, heartY+heartSize*0.8,
				heartX, heartY+heartSize,
			)
			// Right curve
			heart.CubeTo(
				heartX, heartY+heartSize*0.8,
				heartX+heartSize*0.5, heartY+heartSize*0.5,
				heartX+heartSize*0.5, heartY,
			)
			heart.CubeTo(
				heartX+heartSize*0.5, heartY-heartSize*0.5,
				heartX, heartY,
				heartX, heartY+heartSize*0.3,
			)
			heart.Close()
			c.DrawPath(heart)

			c.Restore()

			// Example 8: Abstract Art - Flowing Curves
			c.Save()
			c.Translate(startX+spacing, startY+spacing*3)

			c.SetStroke(skia.StrokeOpts{Width: 2, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin})
			colors := []color.NRGBA{
				{R: 255, G: 100, B: 150, A: 255},
				{R: 100, G: 200, B: 255, A: 255},
				{R: 150, G: 255, B: 100, A: 255},
				{R: 255, G: 200, B: 100, A: 255},
			}

			for i, col := range colors {
				c.SetColor(col)
				flow := skia.NewPath()
				baseY := float32(20 + i*15)
				flow.MoveTo(0, baseY)
				for j := 0; j < 4; j++ {
					x := float32(j * 20)
					ctrl1X := x + 8
					ctrl1Y := baseY + float32(math.Sin(float64(j))*10)
					ctrl2X := x + 12
					ctrl2Y := baseY + float32(math.Cos(float64(j))*10)
					endX := x + 20
					endY := baseY + float32(math.Sin(float64(j+1))*8)
					flow.CubeTo(ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, endX, endY)
				}
				c.DrawPath(flow)
			}

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}
