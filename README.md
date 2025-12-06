# gio-skia

A Skia-style immediate-mode drawing API for [Gio](https://gioui.org), providing GPU-accelerated 2D graphics rendering.

## Features

- **Skia-style API**: Familiar drawing interface inspired by Skia graphics library
- **GPU-accelerated**: All rendering is hardware-accelerated via Gio's renderer
- **Path-based drawing**: Build complex shapes using paths with Bézier curves
- **Transformations**: Full support for translate, scale, rotate, and matrix concatenation
- **State management**: Save/Restore stack for managing drawing state
- **Advanced strokes**: Configurable stroke width, caps, joins, miter limits, and dash patterns
- **Fill and stroke modes**: Switch between filled and stroked rendering

## Installation

```bash
go get github.com/zodimo/gio-skia
```

## Quick Start

```go
package main

import (
    "image/color"
    "gioui.org/app"
    "gioui.org/op"
    "gioui.org/op/paint"
    "github.com/zodimo/gio-skia/skia"
)

func main() {
    go func() {
        w := app.NewWindow()
        var ops op.Ops
        for {
            switch e := w.Event().(type) {
            case app.FrameEvent:
                ops.Reset()
                
                // White background
                paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}.Add(&ops)
                paint.PaintOp{}.Add(&ops)
                
                // Create canvas and draw
                c := skia.NewCanvas(&ops)
                p := skia.NewPath()
                p.AddCircle(100, 100, 50)
                c.SetColor(color.NRGBA{R: 255, A: 255})
                c.DrawPath(p)
                
                e.Frame(&ops)
            }
        }
    }()
    app.Main()
}
```

## API Overview

### Canvas

The `Canvas` interface provides the main drawing context:

```go
type Canvas interface {
    // State management
    Save()                 // Push matrix & paint onto stack
    Restore()              // Pop state; restores matrix & paint
    Concat(m f32.Affine2D) // Multiply current transform matrix
    Translate(x, y float32)
    Scale(x, y float32)
    Rotate(angle float32) // radians, counter-clockwise
    
    // Paint state
    SetColor(col color.NRGBA)
    SetStroke(opt StrokeOpts) // Configure stroke style
    Fill()                    // Switch to fill mode (default)
    Stroke()                  // Switch to stroke mode
    
    // Drawing
    DrawPath(p Path) // Render path with current paint & transform
}
```

### Path Building

Build complex shapes using the `Path` interface:

```go
p := skia.NewPath()
p.MoveTo(x, y)                    // Move to point
p.LineTo(x, y)                    // Draw line
p.QuadTo(cx, cy, x, y)            // Quadratic Bézier curve
p.CubeTo(cx1, cy1, cx2, cy2, x, y) // Cubic Bézier curve
p.Close()                         // Close current contour
p.AddRect(x, y, w, h)             // Add rectangle
p.AddCircle(cx, cy, r)            // Add circle
```

### Stroke Options

Configure stroke appearance:

```go
strokeOpts := skia.StrokeOpts{
    Width: 5.0,                    // Stroke width in pixels
    Miter: 4.0,                    // Miter limit
    Cap:   skia.RoundCap,          // Cap style
    Join:  skia.RoundJoin,         // Join style
    Dash:  []float32{10, 5},       // Dash pattern
    Dash0: 0,                      // Dash phase
}
```

**Cap Styles:**
- `RoundCap` - Round caps
- `SquareCap` - Square caps
- `FlatCap` - Flat caps (default)
- `TriangularCap` - Triangular caps

**Join Styles:**
- `RoundJoin` - Round joins
- `MiterJoin` - Miter joins (default)
- `BevelJoin` - Bevel joins

## Examples

See the `examples/gpu/` directory for comprehensive examples:

- `basic_shapes.go` - Basic shapes and primitives
- `transformations.go` - Transformations and state management
- `strokes.go` - Stroke styles and dash patterns
- `bezier_curves.go` - Bézier curves and complex paths
- `animated.go` - Animated graphics example

## License

SPDX-License-Identifier: Unlicense OR MIT
