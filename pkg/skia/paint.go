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
type Shader interface {
	// IsOpaque returns true if the shader is guaranteed to produce only opaque colors.
	IsOpaque() bool

	// IsAImage returns the image if this shader is backed by a single image.
	IsAImage(localMatrix Matrix, xy []TileMode) Image

	// MakeWithLocalMatrix returns a shader that will apply the specified localMatrix to this shader.
	MakeWithLocalMatrix(matrix Matrix) Shader

	// MakeWithColorFilter creates a new shader that produces the same colors as invoking this shader and then applying the colorfilter.
	MakeWithColorFilter(filter ColorFilter) Shader

	// MakeWithWorkingColorSpace returns a shader that will compute this shader in a context such that any child shaders return RGBA values converted to the inputCS colorspace.
	MakeWithWorkingColorSpace(inputCS, outputCS ColorSpace) Shader
}

// ColorFilter modifies the color of pixels drawn.
type ColorFilter interface {
	// AsAColorMode returns true if the filter can be represented by a source color plus Mode.
	AsAColorMode(color *Color, mode *BlendMode) bool

	// AsAColorMatrix returns true if the filter can be represented by a 5x4 matrix.
	AsAColorMatrix(matrix [20]float32) bool

	// IsAlphaUnchanged returns true if the filter is guaranteed to never change the alpha of a color it filters.
	IsAlphaUnchanged() bool

	// FilterColor4f converts the src color (in src colorspace), into the dst colorspace, then applies this filter to it.
	FilterColor4f(srcColor Color4f, srcCS, dstCS ColorSpace) Color4f

	// MakeComposed constructs a colorfilter whose effect is to first apply the inner filter and then apply this filter.
	MakeComposed(inner ColorFilter) ColorFilter

	// MakeWithWorkingColorSpace returns a colorfilter that will compute this filter in a specific color space.
	MakeWithWorkingColorSpace(cs ColorSpace) ColorFilter
}

// MaskFilter modifies the alpha channel of pixels drawn.
type MaskFilter interface {
	// MakeBlur creates a blur maskfilter.
	MakeBlur(style BlurStyle, sigma Scalar, respectCTM bool) MaskFilter
}

// BlurStyle specifies the blur style.
type BlurStyle uint8

const (
	BlurStyleNormal BlurStyle = iota
	BlurStyleSolid
	BlurStyleOuter
	BlurStyleInner
)

// PathEffect modifies the geometry of paths before they are drawn.
type PathEffect interface {
	// MakeSum returns a patheffect that applies each effect (first and second) to the original path.
	MakeSum(first, second PathEffect) PathEffect

	// MakeCompose returns a patheffect that applies the inner effect to the path, and then applies the outer effect to the result.
	MakeCompose(outer, inner PathEffect) PathEffect

	// FilterPath applies this effect to the src path, returning the new path in dst.
	FilterPath(dst PathBuilder, src Path, strokeRec StrokeRec, cullRect *Rect, ctm Matrix) bool

	// NeedsCTM returns true if this path effect requires a valid CTM.
	NeedsCTM() bool
}

// StrokeRec represents stroke recording information.
type StrokeRec interface {
	// Style returns the stroke style.
	Style() PaintStyle

	// Width returns the stroke width.
	Width() Scalar

	// Miter returns the miter limit.
	Miter() Scalar

	// Cap returns the cap style.
	Cap() PaintCap

	// Join returns the join style.
	Join() PaintJoin
}

// PathBuilder is used to build paths.
type PathBuilder interface {
	// MoveTo starts a new contour at the specified point.
	MoveTo(x, y Scalar)

	// LineTo adds a line from the last point to the specified point.
	LineTo(x, y Scalar)

	// QuadTo adds a quadratic bezier.
	QuadTo(cx, cy, x, y Scalar)

	// CubicTo adds a cubic bezier.
	CubicTo(cx1, cy1, cx2, cy2, x, y Scalar)

	// Close closes the current contour.
	Close()
}

// ImageFilter modifies the pixels of images before they are drawn.
type ImageFilter interface {
	// FilterBounds maps a device-space rect recursively forward or backward through the filter DAG.
	FilterBounds(src IRect, ctm Matrix, direction MapDirection, inputRect *IRect) IRect

	// IsColorFilterNode returns whether this image filter is a color filter.
	IsColorFilterNode(filterPtr **ColorFilter) bool

	// AsAColorFilter returns true if this imagefilter can be completely replaced by the returned colorfilter.
	AsAColorFilter(filterPtr **ColorFilter) bool

	// CountInputs returns the number of inputs this filter will accept.
	CountInputs() int

	// GetInput returns the input filter at a given index.
	GetInput(i int) ImageFilter

	// ComputeFastBounds returns the fast bounds of the filter.
	ComputeFastBounds(bounds Rect) Rect

	// CanComputeFastBounds returns true if this filter DAG can compute the resulting bounds.
	CanComputeFastBounds() bool

	// MakeWithLocalMatrix returns a filter with a local matrix applied.
	MakeWithLocalMatrix(matrix Matrix) ImageFilter
}

// MapDirection specifies the direction for mapping bounds.
type MapDirection uint8

const (
	MapDirectionForward MapDirection = iota
	MapDirectionReverse
)

// ColorSpace describes the color space of colors.
type ColorSpace interface {
	// IsSRGB returns true if the color space is sRGB.
	IsSRGB() bool

	// IsLinearGamma returns true if the color space has a linear gamma.
	IsLinearGamma() bool

	// GammaIsCloseToSRGB returns true if the gamma is close to sRGB.
	GammaIsCloseToSRGB() bool

	// IsNumericalTransferFunction returns true if the transfer function is numerical.
	IsNumericalTransferFunction() bool

	// ToXYZD50 converts the color space to XYZ D50.
	ToXYZD50() [9]float32

	// TransferFn returns the transfer function.
	TransferFn() TransferFn

	// Hash returns a hash of the color space.
	Hash() uint32
}

// TransferFn represents a transfer function.
type TransferFn struct {
	FA, FB, FC, FD, FE, FF, FG float32
}
