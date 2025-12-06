# Skia Go Interfaces

This package provides Go interfaces that mirror the Skia C++ API from the core headers. These interfaces focus on Canvas and all its dependent types.

## Package Structure

### Core Types (`types.go`)
- `Scalar` - Floating-point value (float32)
- `Point`, `IPoint` - 2D points (float and int)
- `Size`, `ISize` - 2D sizes (float and int)
- `Rect`, `IRect` - Rectangles (float and int)
- `Color`, `Color4f` - Color representations
- `Point3` - 3D point

### Enums (`enums.go`)
- `BlendMode` - How source and destination colors are combined
- `ClipOp` - Clip operations (Difference, Intersect)
- `PathFillType` - How paths are filled (Winding, EvenOdd, etc.)
- `PathDirection` - Direction for closed contours (CW, CCW)
- `PathVerb` - Path commands (Move, Line, Quad, Conic, Cubic, Close)
- `PaintStyle` - How geometry is drawn (Fill, Stroke, StrokeAndFill)
- `PaintCap` - Stroke cap styles (Butt, Round, Square)
- `PaintJoin` - Stroke join styles (Miter, Round, Bevel)
- `PointMode` - How points are drawn (Points, Lines, Polygon)
- `SrcRectConstraint` - Source rectangle constraint behavior
- `TextEncoding` - Text encoding formats
- `AlphaType` - Alpha storage types
- `ColorType` - Pixel color component types
- `TileMode` - Shader tiling modes

### Paint (`paint.go`)
The `Paint` interface controls options applied when drawing:
- Color and alpha management
- Stroke width, cap, join, and miter
- Shader, ColorFilter, MaskFilter, PathEffect, ImageFilter
- Blend mode
- Anti-aliasing and dithering

### Path (`path.go`)
The `Path` interface contains geometry:
- Fill type management
- Path construction (MoveTo, LineTo, QuadTo, CubicTo, etc.)
- Path operations (AddRect, AddOval, AddCircle, AddRRect)
- Transformations and offsets
- Bounds computation

Also includes `RRect` interface for rounded rectangles.

### Matrix (`matrix.go`)
The `Matrix` interface provides 3x3 matrix transformations:
- Translation, scale, rotation, skew
- Concatenation operations
- Point and rectangle mapping
- Matrix inversion
- Type queries

### Canvas (`canvas.go`)
The main `Canvas` interface provides:
- **State Management**: Save/Restore operations
- **Matrix Operations**: Translate, Scale, Rotate, Skew, Concat
- **Clip Operations**: ClipRect, ClipRRect, ClipPath, ClipRegion
- **Drawing Operations**:
  - Colors: DrawColor, Clear
  - Points and Lines: DrawPoint, DrawPoints, DrawLine
  - Shapes: DrawRect, DrawOval, DrawCircle, DrawArc, DrawRoundRect
  - Paths: DrawPath
  - Images: DrawImage, DrawImageRect, DrawImageNine, DrawImageLattice
  - Text: DrawSimpleText, DrawString, DrawGlyphs, DrawTextBlob
  - Advanced: DrawPicture, DrawVertices, DrawMesh, DrawPatch, DrawAtlas

### Image Info (`image.go`)
The `ImageInfo` struct describes pixel dimensions and encoding:
- Width, height, color type, alpha type, color space
- Helper functions for byte calculations
- Opaque detection

## Usage Example

```go
package main

import (
    "github.com/zodimo/gio-skia/pkg/skia"
)

func drawExample(canvas skia.Canvas, paint skia.Paint) {
    // Set up paint
    paint.SetColor(skia.Color(0xFF0000FF)) // Blue
    paint.SetStyle(skia.PaintStyleFill)
    
    // Draw a rectangle
    rect := skia.Rect{
        Left:   10,
        Top:    10,
        Right:  100,
        Bottom: 100,
    }
    canvas.DrawRect(rect, paint)
    
    // Transform and draw again
    canvas.Save()
    canvas.Translate(50, 50)
    canvas.Rotate(45)
    canvas.DrawRect(rect, paint)
    canvas.Restore()
}
```

## Design Notes

1. **Interfaces over Structs**: Following Go best practices, the API uses interfaces to allow for multiple implementations.

2. **C++ to Go Mapping**:
   - C++ `SkScalar` → Go `Scalar` (float32)
   - C++ `SkPoint` → Go `Point` struct
   - C++ `SkRect` → Go `Rect` struct
   - C++ `SkPaint` → Go `Paint` interface
   - C++ `SkPath` → Go `Path` interface
   - C++ `SkCanvas` → Go `Canvas` interface

3. **Method Naming**: Methods follow Go conventions (PascalCase, no underscores).

4. **Optional Parameters**: Go doesn't support optional parameters, so methods with optional Paint parameters use `Paint` (can be nil).

5. **Error Handling**: Some methods return `bool` to indicate success/failure (e.g., `ReadPixels`).

## Dependencies

The Canvas interface depends on:
- Paint (for drawing state)
- Path (for path drawing)
- Matrix (for transformations)
- Image (for image drawing)
- Font (for text drawing)
- Various enums and types

All these are defined within this package to provide a complete, self-contained API surface.

