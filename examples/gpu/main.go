package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
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
			c := skia.NewCanvas(&ops)
			c.SetColor(color.NRGBA{R: 255, G: 200, B: 0, A: 255})
			c.Fill()

			c.Save()
			c.Translate(100, 100)
			c.Scale(2, 2)

			p := skia.NewPath()
			p.AddCircle(0, 0, 40)
			c.SetColor(color.NRGBA{R: 0, G: 0, B: 200, A: 255})
			c.DrawPath(p)

			p2 := skia.NewPath()
			p2.AddRect(-50, -50, 100, 100)
			c.SetStroke(skia.StrokeOpts{Width: 3, Join: skia.RoundJoin})
			c.SetColor(color.NRGBA{R: 200, G: 0, B: 0, A: 255})
			c.DrawPath(p2)

			c.Restore()
			frameEvent.Frame(&ops)
		}
	}
}
