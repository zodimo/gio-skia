package skia

// Paint controls options applied when drawing.
// Paint collects all options outside of the Canvas clip and Canvas matrix.
type Paint interface {
	// Reset sets all Paint contents to their initial values.
	Reset()

	// AntiAlias returns true if edge pixels may be drawn with partial transparency.
	AntiAlias() bool

	// SetAntiAlias requests that edge pixels draw opaque or with partial transparency.
	SetAntiAlias(aa bool)

	// Dither returns true if color error may be distributed to smooth color transition.
	Dither() bool

	// SetDither requests to distribute color error.
	SetDither(dither bool)

	// Style returns whether the geometry is filled, stroked, or filled and stroked.
	Style() PaintStyle

	// SetStyle sets whether the geometry is filled, stroked, or filled and stroked.
	SetStyle(style PaintStyle)

	// SetStroke sets paint's style to Stroke if true, or Fill if false.
	SetStroke(stroke bool)

	// Color returns alpha and RGB, unpremultiplied, packed into 32 bits.
	Color() Color

	// Color4f returns alpha and RGB, unpremultiplied, as four floating point values.
	Color4f() Color4f

	// SetColor sets alpha and RGB used when stroking and filling.
	SetColor(color Color)

	// SetColor4f sets alpha and RGB used when stroking and filling with floating point values.
	SetColor4f(color Color4f, colorSpace ColorSpace)

	// SetARGB sets color used when drawing solid fills.
	SetARGB(a, r, g, b uint8)

	// Alphaf returns alpha from the color used when stroking and filling.
	Alphaf() float32

	// Alpha returns alpha scaled by 255.
	Alpha() uint8

	// SetAlphaf replaces alpha, leaving RGB unchanged.
	SetAlphaf(a float32)

	// SetAlpha sets alpha using an int between 0 and 255.
	SetAlpha(a uint8)

	// StrokeWidth returns the thickness of the pen used to outline the shape.
	StrokeWidth() Scalar

	// SetStrokeWidth sets the thickness of the pen used to outline the shape.
	SetStrokeWidth(width Scalar)

	// StrokeMiter returns the limit at which a sharp corner is drawn beveled.
	StrokeMiter() Scalar

	// SetStrokeMiter sets the miter limit.
	SetStrokeMiter(miterLimit Scalar)

	// StrokeCap returns the geometry drawn at the beginning and end of strokes.
	StrokeCap() PaintCap

	// SetStrokeCap sets the geometry drawn at the beginning and end of strokes.
	SetStrokeCap(cap PaintCap)

	// StrokeJoin returns the geometry drawn at the corners of strokes.
	StrokeJoin() PaintJoin

	// SetStrokeJoin sets the geometry drawn at the corners of strokes.
	SetStrokeJoin(join PaintJoin)

	// Shader returns optional colors used when filling a path, such as a gradient.
	Shader() Shader

	// SetShader sets optional colors used when filling a path.
	SetShader(shader Shader)

	// ColorFilter returns the color filter if set, or nil.
	ColorFilter() ColorFilter

	// SetColorFilter sets the color filter.
	SetColorFilter(filter ColorFilter)

	// BlendMode returns the blend mode used to combine source color and destination.
	BlendMode() BlendMode

	// SetBlendMode sets the blend mode used to combine source color and destination.
	SetBlendMode(mode BlendMode)

	// MaskFilter returns the mask filter if set, or nil.
	MaskFilter() MaskFilter

	// SetMaskFilter sets the mask filter.
	SetMaskFilter(filter MaskFilter)

	// PathEffect returns the path effect if set, or nil.
	PathEffect() PathEffect

	// SetPathEffect sets the path effect.
	SetPathEffect(effect PathEffect)

	// ImageFilter returns the image filter if set, or nil.
	ImageFilter() ImageFilter

	// SetImageFilter sets the image filter.
	SetImageFilter(filter ImageFilter)
}

// Shader provides colors used when filling a path.
type Shader interface{}

// ColorFilter modifies the color of pixels drawn.
type ColorFilter interface{}

// MaskFilter modifies the alpha channel of pixels drawn.
type MaskFilter interface{}

// PathEffect modifies the geometry of paths before they are drawn.
type PathEffect interface{}

// ImageFilter modifies the pixels of images before they are drawn.
type ImageFilter interface{}

// ColorSpace describes the color space of colors.
type ColorSpace interface{}

