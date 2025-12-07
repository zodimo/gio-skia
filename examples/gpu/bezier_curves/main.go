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
			c.TranslateFloat32(startX, startY)

			// Draw control points and lines
			paint := skia.NewPaintFill(color.NRGBA{R: 100, G: 100, B: 100, A: 255})
			p0 := skia.NewPath()
			skia.PathMoveTo(p0, 0, 50)
			skia.PathLineTo(p0, 30, 0)
			skia.PathLineTo(p0, 60, 50)
			c.DrawPath(p0, paint)

			// Draw control points
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 255, B: 0, A: 255})
			cp1 := skia.NewPath()
			skia.PathAddCircle(cp1, 0, 50, 3)
			c.DrawPath(cp1, paint)
			cp2 := skia.NewPath()
			skia.PathAddCircle(cp2, 30, 0, 3)
			c.DrawPath(cp2, paint)
			cp3 := skia.NewPath()
			skia.PathAddCircle(cp3, 60, 50, 3)
			c.DrawPath(cp3, paint)

			// Draw the curve
			paint = skia.NewPaintFill(color.NRGBA{R: 100, G: 200, B: 255, A: 255})
			curve1 := skia.NewPath()
			skia.PathMoveTo(curve1, 0, 50)
			skia.PathQuadTo(curve1, 30, 0, 60, 50)
			c.DrawPath(curve1, paint)

			c.Restore()

			// Example 2: Cubic Bézier Curves
			c.Save()
			c.TranslateFloat32(startX+spacing, startY)

			// Draw control lines
			paint = skia.NewPaintFill(color.NRGBA{R: 100, G: 100, B: 100, A: 255})
			p1 := skia.NewPath()
			skia.PathMoveTo(p1, 0, 50)
			skia.PathLineTo(p1, 20, 0)
			skia.PathMoveTo(p1, 40, 0)
			skia.PathLineTo(p1, 60, 50)
			c.DrawPath(p1, paint)

			// Draw control points
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 0, A: 255})
			points := []struct{ x, y float32 }{
				{0, 50}, {20, 0}, {40, 0}, {60, 50},
			}
			for _, pt := range points {
				cp := skia.NewPath()
				skia.PathAddCircle(cp, pt.x, pt.y, 3)
				c.DrawPath(cp, paint)
			}

			// Draw the cubic curve
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 200, A: 255})
			curve2 := skia.NewPath()
			skia.PathMoveTo(curve2, 0, 50)
			skia.PathCubeTo(curve2, 20, 0, 40, 0, 60, 50)
			c.DrawPath(curve2, paint)

			c.Restore()

			// Example 3: Smooth Wave Pattern (Connected Quadratic Curves)
			c.Save()
			c.TranslateFloat32(startX, startY+spacing)
			paint = skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 150, A: 255})
			wave := skia.NewPath()
			skia.PathMoveTo(wave, 0, 25)
			for i := 0; i < 5; i++ {
				x1 := float32(i*30 + 15)
				y1 := float32(25 - 15*float32(math.Pow(-1, float64(i))))
				x2 := float32((i + 1) * 30)
				y2 := float32(25)
				skia.PathQuadTo(wave, x1, y1, x2, y2)
			}
			c.DrawPath(wave, paint)

			c.Restore()

			// Example 4: Complex Cubic Curve - Infinity Symbol
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing)
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 100, A: 255})
			infinity := skia.NewPath()
			skia.PathMoveTo(infinity, 0, 25)
			// Left loop
			skia.PathCubeTo(infinity, 15, 0, 15, 0, 30, 25)
			// Right loop
			skia.PathCubeTo(infinity, 45, 50, 45, 50, 60, 25)
			c.DrawPath(infinity, paint)

			c.Restore()

			// Example 5: Flower Petals using Bézier Curves
			c.Save()
			c.TranslateFloat32(startX, startY+spacing*2)

			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
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
					skia.PathMoveTo(flower, petalX, petalY)
				} else {
					skia.PathCubeTo(flower, ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, petalX, petalY)
				}
			}
			flower.Close()
			c.DrawPath(flower, paint)

			// Center circle
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 50, A: 255})
			center := skia.NewPath()
			skia.PathAddCircle(center, centerX, centerY, 8)
			c.DrawPath(center, paint)

			c.Restore()

			// Example 6: Spiral using Cubic Bézier Segments
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing*2)
			paint = skia.NewPaintFill(color.NRGBA{R: 150, G: 200, B: 255, A: 255})
			spiral := skia.NewPath()
			spiralCenterX, spiralCenterY := float32(40), float32(40)
			skia.PathMoveTo(spiral, spiralCenterX, spiralCenterY)

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

				skia.PathCubeTo(spiral, ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, x2, y2)
			}
			c.DrawPath(spiral, paint)

			c.Restore()

			// Example 7: Heart Shape with Bézier Curves
			c.Save()
			c.TranslateFloat32(startX, startY+spacing*3)

			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			heart := skia.NewPath()
			heartX, heartY := float32(40), float32(35)
			heartSize := float32(25)

			// Left curve
			skia.PathMoveTo(heart, heartX, heartY+heartSize*0.3)
			skia.PathCubeTo(heart, 
				heartX, heartY,
				heartX-heartSize*0.5, heartY-heartSize*0.5,
				heartX-heartSize*0.5, heartY,
			)
			skia.PathCubeTo(heart, 
				heartX-heartSize*0.5, heartY+heartSize*0.5,
				heartX, heartY+heartSize*0.8,
				heartX, heartY+heartSize,
			)
			// Right curve
			skia.PathCubeTo(heart, 
				heartX, heartY+heartSize*0.8,
				heartX+heartSize*0.5, heartY+heartSize*0.5,
				heartX+heartSize*0.5, heartY,
			)
			skia.PathCubeTo(heart, 
				heartX+heartSize*0.5, heartY-heartSize*0.5,
				heartX, heartY,
				heartX, heartY+heartSize*0.3,
			)
			heart.Close()
			c.DrawPath(heart, paint)

			c.Restore()

			// Example 8: Abstract Art - Flowing Curves
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing*3)
			colors := []color.NRGBA{
				{R: 255, G: 100, B: 150, A: 255},
				{R: 100, G: 200, B: 255, A: 255},
				{R: 150, G: 255, B: 100, A: 255},
				{R: 255, G: 200, B: 100, A: 255},
			}

			for i, col := range colors {
				paint = skia.NewPaintFill(col)
				flow := skia.NewPath()
				baseY := float32(20 + i*15)
				skia.PathMoveTo(flow, 0, baseY)
				for j := 0; j < 4; j++ {
					x := float32(j * 20)
					ctrl1X := x + 8
					ctrl1Y := baseY + float32(math.Sin(float64(j))*10)
					ctrl2X := x + 12
					ctrl2Y := baseY + float32(math.Cos(float64(j))*10)
					endX := x + 20
					endY := baseY + float32(math.Sin(float64(j+1))*8)
					skia.PathCubeTo(flow, ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, endX, endY)
				}
				c.DrawPath(flow, paint)
			}

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}
