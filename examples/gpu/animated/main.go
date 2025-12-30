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
	startTime := time.Now()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()

			// Calculate animation time
			elapsed := time.Since(startTime).Seconds()

			// Dark gradient background
			paint.ColorOp{Color: color.NRGBA{R: 15, G: 15, B: 25, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			w, h := float32(e.Size.X), float32(e.Size.Y)
			centerX, centerY := w/2, h/2

			// Animation 1: Rotating Gear
			c.Save()
			c.Translate(centerX*0.3, centerY*0.3)
			c.Rotate(float32(elapsed * 0.5 * 180.0 / math.Pi)) // Convert radians to degrees
			paint := skia.NewPaintFill(color.NRGBA{R: 100, G: 200, B: 255, A: 255})
			gear := skia.NewPath()
			gearRadius := float32(40)
			innerRadius := float32(25)
			teeth := 12
			for i := 0; i < teeth*2; i++ {
				angle := float32(i) * math.Pi / float32(teeth)
				var radius float32
				if i%2 == 0 {
					radius = gearRadius
				} else {
					radius = innerRadius
				}
				x := radius * float32(math.Cos(float64(angle)))
				y := radius * float32(math.Sin(float64(angle)))
				if i == 0 {
					skia.PathMoveTo(gear, x, y)
				} else {
					skia.PathLineTo(gear, x, y)
				}
			}
			gear.Close()
			c.DrawPath(gear, paint)
			c.Restore()

			// Animation 2: Pulsing Circles
			pulseScale := 0.5 + 0.5*float32(math.Sin(elapsed*2))
			c.Save()
			c.Translate(centerX*1.5, centerY*0.3)
			c.Scale(pulseScale, pulseScale)
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 150, A: 255})
			pulseCircle := skia.NewPath()
			skia.PathAddCircle(pulseCircle, 0, 0, 30)
			c.DrawPath(pulseCircle, paint)
			c.Restore()

			// Animation 3: Spinning Star
			c.Save()
			c.Translate(centerX*0.3, centerY*1.5)
			c.Rotate(float32(elapsed * 180.0 / math.Pi)) // Convert radians to degrees
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			star := skia.NewPath()
			starRadius := float32(35)
			starInnerRadius := float32(15)
			for i := 0; i < 5; i++ {
				angle := float32(i) * math.Pi * 2 / 5
				x := starRadius * float32(math.Cos(float64(angle-math.Pi/2)))
				y := starRadius * float32(math.Sin(float64(angle-math.Pi/2)))
				if i == 0 {
					skia.PathMoveTo(star, x, y)
				} else {
					skia.PathLineTo(star, x, y)
				}
				innerAngle := angle + math.Pi/5
				ix := starInnerRadius * float32(math.Cos(float64(innerAngle-math.Pi/2)))
				iy := starInnerRadius * float32(math.Sin(float64(innerAngle-math.Pi/2)))
				skia.PathLineTo(star, ix, iy)
			}
			star.Close()
			c.DrawPath(star, paint)
			c.Restore()

			// Animation 4: Waving Flag (using Bézier curves)
			c.Save()
			c.Translate(centerX*1.5, centerY*1.5)
			paint = skia.NewPaintFill(color.NRGBA{R: 200, G: 50, B: 50, A: 255})
			flag := skia.NewPath()
			flagWidth := float32(80)
			flagHeight := float32(50)
			waveAmplitude := float32(10)

			skia.PathMoveTo(flag, 0, 0)
			for i := 0; i <= 10; i++ {
				x := float32(i) * flagWidth / 10
				wave := waveAmplitude * float32(math.Sin(elapsed*3+float64(i)*0.5))
				y := wave

				if i == 0 {
					skia.PathMoveTo(flag, x, y)
				} else {
					// Use cubic curves for smooth wave
					prevX := float32(i-1) * flagWidth / 10
					prevWave := waveAmplitude * float32(math.Sin(elapsed*3+float64(i-1)*0.5))
					ctrl1X := prevX + flagWidth/20
					ctrl1Y := prevWave
					ctrl2X := x - flagWidth/20
					ctrl2Y := wave
					skia.PathCubeTo(flag, ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, x, y)
				}
			}
			// Complete the flag shape
			skia.PathLineTo(flag, flagWidth, flagHeight)
			skia.PathLineTo(flag, 0, flagHeight)
			flag.Close()
			c.DrawPath(flag, paint)
			c.Restore()

			// Animation 5: Orbiting Planets
			c.Save()
			c.Translate(centerX, centerY)

			// Central sun
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 50, A: 255})
			sun := skia.NewPath()
			skia.PathAddCircle(sun, 0, 0, 20)
			c.DrawPath(sun, paint)

			// Planet 1 - Fast orbit
			orbit1Radius := float32(60)
			orbit1Angle := elapsed * 2
			planet1X := orbit1Radius * float32(math.Cos(orbit1Angle))
			planet1Y := orbit1Radius * float32(math.Sin(orbit1Angle))

			c.Save()
			c.Translate(planet1X, planet1Y)
			c.Rotate(float32(elapsed * 3 * 180.0 / math.Pi)) // Convert radians to degrees
			paint = skia.NewPaintFill(color.NRGBA{R: 100, G: 150, B: 255, A: 255})
			planet1 := skia.NewPath()
			skia.PathAddCircle(planet1, 0, 0, 8)
			c.DrawPath(planet1, paint)
			c.Restore()

			// Planet 2 - Slow orbit
			orbit2Radius := float32(90)
			orbit2Angle := elapsed * 0.8
			planet2X := orbit2Radius * float32(math.Cos(orbit2Angle))
			planet2Y := orbit2Radius * float32(math.Sin(orbit2Angle))

			c.Save()
			c.Translate(planet2X, planet2Y)
			c.Rotate(float32(elapsed * 1.5 * 180.0 / math.Pi)) // Convert radians to degrees
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			planet2 := skia.NewPath()
			skia.PathAddCircle(planet2, 0, 0, 10)
			c.DrawPath(planet2, paint)
			c.Restore()

			// Orbit paths (dashed)
			strokeOpts := stroke.StrokeOpts{
				Width: 1,
				Miter: 4,
				Cap:   stroke.RoundCap,
				Join:  stroke.RoundJoin,
				Dash:  []float32{5, 5},
				Dash0: float32(elapsed * 10),
			}
			paint = skia.NewPaintStroke(color.NRGBA{R: 100, G: 100, B: 100, A: 100}, 1)
			paint = skia.ConfigureStrokePaint(paint, strokeOpts)
			orbit1 := skia.NewPath()
			skia.PathAddCircle(orbit1, 0, 0, orbit1Radius)
			c.DrawPath(orbit1, paint)
			orbit2 := skia.NewPath()
			skia.PathAddCircle(orbit2, 0, 0, orbit2Radius)
			c.DrawPath(orbit2, paint)

			c.Restore()

			// Animation 6: Morphing Shape
			morphPhase := float32(math.Sin(elapsed * 1.5))
			c.Save()
			c.Translate(centerX*0.2, centerY*0.8)
			paint = skia.NewPaintFill(color.NRGBA{R: 150, G: 255, B: 150, A: 255})
			morph := skia.NewPath()
			points := 8
			baseRadius := float32(30)
			for i := 0; i < points; i++ {
				angle := float32(i) * math.Pi * 2 / float32(points)
				radius := baseRadius + morphPhase*10
				x := radius * float32(math.Cos(float64(angle)))
				y := radius * float32(math.Sin(float64(angle)))
				if i == 0 {
					skia.PathMoveTo(morph, x, y)
				} else {
					skia.PathLineTo(morph, x, y)
				}
			}
			morph.Close()
			c.DrawPath(morph, paint)
			c.Restore()

			// Animation 7: Animated Dash Pattern
			c.Save()
			c.Translate(centerX*1.8, centerY*0.8)
			strokeOpts = stroke.StrokeOpts{
				Width: 6,
				Miter: 4,
				Cap:   stroke.RoundCap,
				Join:  stroke.RoundJoin,
				Dash:  []float32{15, 10},
				Dash0: float32(elapsed * 50),
			}
			paint = skia.NewPaintStroke(color.NRGBA{R: 255, G: 150, B: 255, A: 255}, 6)
			paint = skia.ConfigureStrokePaint(paint, strokeOpts)
			dashLine := skia.NewPath()
			skia.PathMoveTo(dashLine, -40, 0)
			skia.PathLineTo(dashLine, 40, 0)
			c.DrawPath(dashLine, paint)
			c.Restore()

			// Animation 8: Spiral Growth
			c.Save()
			c.Translate(centerX*0.2, centerY*0.2)
			spiral := skia.NewPath()
			maxTurns := float32(elapsed * 0.5)
			if maxTurns > 8 {
				maxTurns = 8
			}
			spiralRadius := float32(5)
			for t := float32(0); t < maxTurns*math.Pi*2; t += 0.1 {
				radius := spiralRadius + t*0.5
				x := radius * float32(math.Cos(float64(t)))
				y := radius * float32(math.Sin(float64(t)))
				if t == 0 {
					skia.PathMoveTo(spiral, x, y)
				} else {
					skia.PathLineTo(spiral, x, y)
				}
			}
			hue := int(elapsed*60) % 360
			r, g, b := hslToRGB(float32(hue), 0.7, 0.6)
			paint = skia.NewPaintFill(color.NRGBA{R: r, G: g, B: b, A: 255})
			c.DrawPath(spiral, paint)
			c.Restore()

			// Animation 9: Bouncing Balls
			c.Save()
			c.Translate(centerX*1.8, centerY*0.2)
			for i := 0; i < 3; i++ {
				bounceHeight := float32(math.Abs(math.Sin(elapsed*2+float64(i)*0.5))) * 30
				c.Save()
				c.Translate(float32(i*25-25), bounceHeight)
				paint = skia.NewPaintFill(color.NRGBA{
					R: uint8(100 + i*50),
					G: uint8(150 + i*30),
					B: uint8(200 + i*20),
					A: 255,
				})
				ball := skia.NewPath()
				skia.PathAddCircle(ball, 0, 0, 8)
				c.DrawPath(ball, paint)
				c.Restore()
			}
			c.Restore()

			// Animation 10: Heartbeat (pulsing heart)
			heartbeatScale := 1.0 + 0.3*float32(math.Abs(math.Sin(elapsed*4)))
			c.Save()
			c.Translate(centerX*0.5, centerY*1.8)
			c.Scale(heartbeatScale, heartbeatScale)
			paint = skia.NewPaintFill(color.NRGBA{R: 255, G: 50, B: 50, A: 255})
			heart := skia.NewPath()
			heartSize := float32(20)
			skia.PathMoveTo(heart, 0, heartSize*0.3)
			skia.PathCubeTo(heart, 0, 0, -heartSize*0.5, -heartSize*0.5, -heartSize*0.5, 0)
			skia.PathCubeTo(heart, -heartSize*0.5, heartSize*0.5, 0, heartSize*0.8, 0, heartSize)
			skia.PathCubeTo(heart, 0, heartSize*0.8, heartSize*0.5, heartSize*0.5, heartSize*0.5, 0)
			skia.PathCubeTo(heart, heartSize*0.5, -heartSize*0.5, 0, 0, 0, heartSize*0.3)
			heart.Close()
			c.DrawPath(heart, paint)
			c.Restore()

			window.Invalidate()

			e.Frame(&ops)
		}
	}
}

// hslToRGB converts HSL color to RGB
func hslToRGB(h, s, l float32) (r, g, b uint8) {
	h = h / 360.0
	var c, x, m float32
	c = (1 - abs(2*l-1)) * s
	x = c * (1 - abs(mod(h*6, 2)-1))
	m = l - c/2

	var r1, g1, b1 float32
	switch {
	case h < 1.0/6:
		r1, g1, b1 = c, x, 0
	case h < 2.0/6:
		r1, g1, b1 = x, c, 0
	case h < 3.0/6:
		r1, g1, b1 = 0, c, x
	case h < 4.0/6:
		r1, g1, b1 = 0, x, c
	case h < 5.0/6:
		r1, g1, b1 = x, 0, c
	default:
		r1, g1, b1 = c, 0, x
	}

	r = uint8((r1 + m) * 255)
	g = uint8((g1 + m) * 255)
	b = uint8((b1 + m) * 255)
	return
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func mod(a, b float32) float32 {
	return a - b*float32(math.Floor(float64(a/b)))
}
