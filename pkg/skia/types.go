// Package skia provides Go interfaces for Skia graphics library.
// These interfaces mirror the C++ Skia API from the core headers.
package skia

// Scalar represents a floating-point value used throughout Skia.
// In C++, this is typically a float32.
type Scalar = float32

// Point represents a 2D point with floating-point coordinates.
type Point struct {
	X, Y Scalar
}

// IPoint represents a 2D point with integer coordinates.
type IPoint struct {
	X, Y int32
}

// Size represents a 2D size with floating-point dimensions.
type Size struct {
	Width, Height Scalar
}

// ISize represents a 2D size with integer dimensions.
type ISize struct {
	Width, Height int32
}

// Rect represents a rectangle with floating-point coordinates.
type Rect struct {
	Left, Top, Right, Bottom Scalar
}

// IRect represents a rectangle with integer coordinates.
type IRect struct {
	Left, Top, Right, Bottom int32
}

// Color represents a 32-bit ARGB color value (unpremultiplied).
type Color uint32

// Color4f represents a color with floating-point RGBA components (unpremultiplied).
type Color4f struct {
	R, G, B, A float32
}

// Point3 represents a 3D point with floating-point coordinates.
type Point3 struct {
	X, Y, Z Scalar
}

// CubicResampler specifies B and C coefficients for cubic reconstruction filter.
// Used in SamplingOptions for high-quality image resampling.
//
// Example values:
//   - B = 1/3, C = 1/3: "Mitchell" filter
//   - B = 0, C = 1/2: "Catmull-Rom" filter
type CubicResampler struct {
	B, C float32
}

// Mitchell returns a CubicResampler with Mitchell filter coefficients (B=1/3, C=1/3).
func Mitchell() CubicResampler {
	return CubicResampler{B: 1.0 / 3.0, C: 1.0 / 3.0}
}

// CatmullRom returns a CubicResampler with Catmull-Rom filter coefficients (B=0, C=1/2).
func CatmullRom() CubicResampler {
	return CubicResampler{B: 0.0, C: 0.5}
}

// Arc represents an arc specification for drawing.
type Arc struct {
	Oval       Rect
	StartAngle Scalar
	SweepAngle Scalar
	IsWedge    bool // If true, includes lines from center to arc endpoints
}

