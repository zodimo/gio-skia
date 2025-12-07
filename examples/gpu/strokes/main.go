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

			// Light background
			paint.ColorOp{Color: color.NRGBA{R: 250, G: 250, B: 255, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			spacing := float32(150)
			startX, startY := spacing, spacing

			// Row 1: Cap Styles
			capStyles := []stroke.CapStyle{stroke.FlatCap, stroke.RoundCap, stroke.SquareCap, stroke.TriangularCap}
			
			for i, capStyle := range capStyles {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing, startY)
				
				// Draw a line showing the cap style
				strokeOpts := stroke.StrokeOpts{Width: 8, Miter: 4, Cap: capStyle, Join: stroke.RoundJoin}
				paint := skia.NewPaintStroke(color.NRGBA{R: 100, G: 150, B: 255, A: 255}, 8)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				skia.PathMoveTo(p, -40, 0)
				skia.PathLineTo(p, 40, 0)
				c.DrawPath(p, paint)
				
				// Draw endpoints as circles for reference
				paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
				p1 := skia.NewPath()
				skia.PathAddCircle(p1, -40, 0, 3)
				c.DrawPath(p1, paint)
				p2 := skia.NewPath()
				skia.PathAddCircle(p2, 40, 0, 3)
				c.DrawPath(p2, paint)
				
				c.Restore()
			}

			// Row 2: Join Styles
			joinStyles := []stroke.JoinStyle{stroke.MiterJoin, stroke.RoundJoin, stroke.BevelJoin}
			startY2 := startY + spacing

			for i, joinStyle := range joinStyles {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing, startY2)
				strokeOpts := stroke.StrokeOpts{Width: 6, Miter: 4, Cap: stroke.RoundCap, Join: joinStyle}
				paint := skia.NewPaintStroke(color.NRGBA{R: 255, G: 100, B: 150, A: 255}, 6)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				skia.PathMoveTo(p, -30, -20)
				skia.PathLineTo(p, 0, 20)
				skia.PathLineTo(p, 30, -20)
				c.DrawPath(p, paint)
				
				c.Restore()
			}

			// Row 3: Different Stroke Widths
			startY3 := startY + spacing*2
			widths := []float32{2, 5, 10, 20, 30}

			for i, width := range widths {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing*0.6, startY3)
				strokeOpts := stroke.StrokeOpts{Width: width, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin}
				paint := skia.NewPaintStroke(color.NRGBA{R: 50, G: 200, B: 100, A: 255}, width)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				skia.PathAddCircle(p, 0, 0, 30)
				c.DrawPath(p, paint)
				
				c.Restore()
			}

			// Row 4: Dash Patterns
			startY4 := startY + spacing*3
			dashPatterns := [][]float32{
				{10, 5},           // Simple dash
				{20, 10, 5, 10},   // Dash-dot
				{5, 5},            // Dotted
				{15, 5, 5, 5},     // Dash-dot-dot
				{30, 10},          // Long dash
			}

			for i, pattern := range dashPatterns {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing*0.6, startY4)
				strokeOpts := stroke.StrokeOpts{Width: 4, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin, Dash: pattern, Dash0: 0}
				paint := skia.NewPaintStroke(color.NRGBA{R: 200, G: 150, B: 50, A: 255}, 4)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				skia.PathMoveTo(p, -50, 0)
				skia.PathLineTo(p, 50, 0)
				c.DrawPath(p, paint)
				
				c.Restore()
			}

			// Row 5: Dash Phase Animation (static demonstration)
			startY5 := startY + spacing*4
			phases := []float32{0, 5, 10, 15, 20}

			for i, phase := range phases {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing*0.6, startY5)
				strokeOpts := stroke.StrokeOpts{Width: 4, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin, Dash: []float32{10, 5}, Dash0: phase}
				paint := skia.NewPaintStroke(color.NRGBA{R: 150, G: 100, B: 255, A: 255}, 4)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				skia.PathMoveTo(p, -40, 0)
				skia.PathLineTo(p, 40, 0)
				c.DrawPath(p, paint)
				
				c.Restore()
			}

			// Row 6: Miter Limit Demonstration
			startY6 := startY + spacing*5
			miterLimits := []float32{1, 2, 4, 8, 16}

			for i, miter := range miterLimits {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing*0.6, startY6)
				strokeOpts := stroke.StrokeOpts{Width: 8, Miter: miter, Cap: stroke.RoundCap, Join: stroke.MiterJoin}
				paint := skia.NewPaintStroke(color.NRGBA{R: 255, G: 120, B: 80, A: 255}, 8)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				// Sharp angle to show miter limit effect
				skia.PathMoveTo(p, -20, -20)
				skia.PathLineTo(p, 0, 20)
				skia.PathLineTo(p, 20, -20)
				c.DrawPath(p, paint)
				
				c.Restore()
			}

			// Row 7: Complex Path with Various Stroke Styles
			startY7 := startY + spacing*6
			
			// Wavy path with different styles
			c.Save()
			c.TranslateFloat32(startX, startY7)
			
			// Style 1: Thick round cap
			strokeOpts1 := stroke.StrokeOpts{Width: 8, Miter: 4, Cap: stroke.RoundCap, Join: stroke.RoundJoin}
			paint := skia.NewPaintStroke(color.NRGBA{R: 100, G: 200, B: 255, A: 255}, 8)
			paint = skia.ConfigureStrokePaint(paint, strokeOpts1)
			p1 := skia.NewPath()
			for i := 0; i < 20; i++ {
				x := float32(i-10) * 8
				y := 20 * float32(math.Sin(float64(i)*0.5))
				if i == 0 {
					skia.PathMoveTo(p1, x, y)
				} else {
					skia.PathLineTo(p1, x, y)
				}
			}
			c.DrawPath(p1, paint)
			
			// Style 2: Dashed square cap
			strokeOpts2 := stroke.StrokeOpts{Width: 6, Miter: 4, Cap: stroke.SquareCap, Join: stroke.RoundJoin, Dash: []float32{10, 5}, Dash0: 0}
			paint = skia.NewPaintStroke(color.NRGBA{R: 255, G: 150, B: 100, A: 255}, 6)
			paint = skia.ConfigureStrokePaint(paint, strokeOpts2)
			p2 := skia.NewPath()
			for i := 0; i < 20; i++ {
				x := float32(i-10) * 8
				y := 20*float32(math.Sin(float64(i)*0.5)) + 60
				if i == 0 {
					skia.PathMoveTo(p2, x, y)
				} else {
					skia.PathLineTo(p2, x, y)
				}
			}
			c.DrawPath(p2, paint)
			
			// Style 3: Thin triangular cap
			strokeOpts3 := stroke.StrokeOpts{Width: 3, Miter: 4, Cap: stroke.TriangularCap, Join: stroke.RoundJoin}
			paint = skia.NewPaintStroke(color.NRGBA{R: 150, G: 255, B: 150, A: 255}, 3)
			paint = skia.ConfigureStrokePaint(paint, strokeOpts3)
			p3 := skia.NewPath()
			for i := 0; i < 20; i++ {
				x := float32(i-10) * 8
				y := 20*float32(math.Sin(float64(i)*0.5)) + 120
				if i == 0 {
					skia.PathMoveTo(p3, x, y)
				} else {
					skia.PathLineTo(p3, x, y)
				}
			}
			c.DrawPath(p3, paint)
			
			c.Restore()

			// Row 8: Star with different stroke styles
			startY8 := startY + spacing*7
			
			starStyles := []struct {
				width float32
				cap   stroke.CapStyle
				join  stroke.JoinStyle
				color color.NRGBA
			}{
				{4, stroke.RoundCap, stroke.RoundJoin, color.NRGBA{R: 255, G: 100, B: 100, A: 255}},
				{8, stroke.SquareCap, stroke.MiterJoin, color.NRGBA{R: 100, G: 255, B: 100, A: 255}},
				{12, stroke.TriangularCap, stroke.BevelJoin, color.NRGBA{R: 100, G: 100, B: 255, A: 255}},
			}

			for i, style := range starStyles {
				c.Save()
				c.TranslateFloat32(startX+float32(i)*spacing, startY8)
				strokeOpts := stroke.StrokeOpts{Width: style.width, Miter: 4, Cap: style.cap, Join: style.join}
				paint := skia.NewPaintStroke(style.color, style.width)
				paint = skia.ConfigureStrokePaint(paint, strokeOpts)
				p := skia.NewPath()
				starRadius := float32(40)
				innerRadius := float32(20)
				for j := 0; j < 10; j++ {
					angle := float32(j) * math.Pi / 5
					var radius float32
					if j%2 == 0 {
						radius = starRadius
					} else {
						radius = innerRadius
					}
					x := radius * float32(math.Cos(float64(angle-math.Pi/2)))
					y := radius * float32(math.Sin(float64(angle-math.Pi/2)))
					if j == 0 {
						skia.PathMoveTo(p, x, y)
					} else {
						skia.PathLineTo(p, x, y)
					}
				}
				p.Close()
				c.DrawPath(p, paint)
				
				c.Restore()
			}

			frameEvent.Frame(&ops)
		}
	}
}

