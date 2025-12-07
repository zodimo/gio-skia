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
                skia.PathAddCircle(p, 100, 100, 50)
                skPaint := skia.NewPaintFill(color.NRGBA{R: 255, A: 255})
                c.DrawPath(p, skPaint)
                
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
    Save() int                    // Push matrix onto stack, returns save count
    Restore()                     // Pop state; restores matrix
    Concat(matrix SkMatrix)       // Multiply current transform matrix
    Translate(dx, dy Scalar)     // Translate by dx, dy
    Scale(sx, sy Scalar)          // Scale by sx, sy
    Rotate(degrees Scalar)       // Rotate by degrees (clockwise)
    
    // Convenience methods (accept float32)
    TranslateFloat32(x, y float32)
    ScaleFloat32(x, y float32)
    RotateFloat32(degrees float32)
    
    // Drawing
    DrawPath(path SkPath, paint SkPaint) // Render path with paint & transform
}
```

### Path Building

Build complex shapes using `SkPath` and helper functions:

```go
p := skia.NewPath()                    // Create new path
skia.PathMoveTo(p, x, y)               // Move to point
skia.PathLineTo(p, x, y)               // Draw line
skia.PathQuadTo(p, cx, cy, x, y)       // Quadratic Bézier curve
skia.PathCubeTo(p, cx1, cy1, cx2, cy2, x, y) // Cubic Bézier curve
p.Close()                              // Close current contour
skia.PathAddRect(p, x, y, w, h)       // Add rectangle
skia.PathAddCircle(p, cx, cy, r)      // Add circle
```

### Paint

Create and configure paint for drawing operations:

```go
// Create paint with default settings
paint := skia.NewPaint()

// Create paint with color (defaults to fill style)
paint := skia.NewPaintWithColor(color.NRGBA{R: 255, G: 0, B: 0, A: 255})

// Create paint configured for filling
paint := skia.NewPaintFill(color.NRGBA{R: 255, G: 0, B: 0, A: 255})

// Create paint configured for stroking
paint := skia.NewPaintStroke(color.NRGBA{R: 255, G: 0, B: 0, A: 255}, 5.0)

// Configure stroke options
import "github.com/zodimo/gio-skia/pkg/stroke"

strokeOpts := stroke.StrokeOpts{
    Width: 5.0,                    // Stroke width in pixels
    Miter: 4.0,                    // Miter limit
    Cap:   stroke.RoundCap,        // Cap style
    Join:  stroke.RoundJoin,       // Join style
}
paint = skia.ConfigureStrokePaint(paint, strokeOpts)
```

**Paint Styles:**
- `skia.PaintStyleFill` - Fill mode (default)
- `skia.PaintStyleStroke` - Stroke mode
- `skia.PaintStyleStrokeAndFill` - Both fill and stroke

**Cap Styles:**
- `stroke.RoundCap` - Round caps
- `stroke.SquareCap` - Square caps
- `stroke.FlatCap` - Flat caps (default)

**Join Styles:**
- `stroke.RoundJoin` - Round joins
- `stroke.MiterJoin` - Miter joins (default)
- `stroke.BevelJoin` - Bevel joins

## Examples

See the `examples/gpu/` directory for comprehensive examples:

- `basic_shapes/` - Basic shapes and primitives
- `transformations/` - Transformations and state management
- `strokes/` - Stroke styles and dash patterns
- `bezier_curves/` - Bézier curves and complex paths
- `animated/` - Animated graphics example

## Type Aliases

The library provides convenient type aliases for better developer experience:

- `skia.Scalar` - Floating-point value for coordinates, dimensions, and angles
- `skia.SkPath` - Path interface for building shapes
- `skia.SkPaint` - Paint interface for configuring drawing properties
- `skia.SkMatrix` - Transformation matrix interface

These allow you to use Skia-style types without importing the underlying packages.

## License

SPDX-License-Identifier: Unlicense OR MIT
