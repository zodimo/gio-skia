package skia

// BlendMode defines how source and destination colors are combined.
type BlendMode uint8

const (
	BlendModeClear BlendMode = iota
	BlendModeSrc
	BlendModeDst
	BlendModeSrcOver
	BlendModeDstOver
	BlendModeSrcIn
	BlendModeDstIn
	BlendModeSrcOut
	BlendModeDstOut
	BlendModeSrcATop
	BlendModeDstATop
	BlendModeXor
	BlendModePlus
	BlendModeModulate
	BlendModeScreen
	BlendModeOverlay
	BlendModeDarken
	BlendModeLighten
	BlendModeColorDodge
	BlendModeColorBurn
	BlendModeHardLight
	BlendModeSoftLight
	BlendModeDifference
	BlendModeExclusion
	BlendModeMultiply
	BlendModeHue
	BlendModeSaturation
	BlendModeColor
	BlendModeLuminosity
)

// ClipOp defines the operation to apply when clipping.
type ClipOp uint8

const (
	ClipOpDifference ClipOp = iota
	ClipOpIntersect
)

// PathFillType determines how paths are filled.
type PathFillType uint8

const (
	PathFillTypeWinding PathFillType = iota
	PathFillTypeEvenOdd
	PathFillTypeInverseWinding
	PathFillTypeInverseEvenOdd
)

// PathDirection specifies the direction for adding closed contours.
type PathDirection uint8

const (
	PathDirectionCW PathDirection = iota
	PathDirectionCCW
)

// PathVerb represents a path command.
type PathVerb uint8

const (
	PathVerbMove PathVerb = iota
	PathVerbLine
	PathVerbQuad
	PathVerbConic
	PathVerbCubic
	PathVerbClose
)

// PaintStyle determines how geometry is drawn.
type PaintStyle uint8

const (
	PaintStyleFill PaintStyle = iota
	PaintStyleStroke
	PaintStyleStrokeAndFill
)

// PaintCap specifies the geometry drawn at the beginning and end of strokes.
type PaintCap uint8

const (
	PaintCapButt PaintCap = iota
	PaintCapRound
	PaintCapSquare
)

// PaintJoin specifies how corners are drawn when a shape is stroked.
type PaintJoin uint8

const (
	PaintJoinMiter PaintJoin = iota
	PaintJoinRound
	PaintJoinBevel
)

// PointMode selects how an array of points are drawn.
type PointMode uint8

const (
	PointModePoints PointMode = iota
	PointModeLines
	PointModePolygon
)

// SrcRectConstraint controls behavior at the edge of source rectangles.
type SrcRectConstraint uint8

const (
	SrcRectConstraintStrict SrcRectConstraint = iota
	SrcRectConstraintFast
)

// TextEncoding specifies how text is encoded.
type TextEncoding uint8

const (
	TextEncodingUTF8 TextEncoding = iota
	TextEncodingUTF16
	TextEncodingUTF32
	TextEncodingGlyphID
)

// AlphaType describes how alpha is stored.
type AlphaType uint8

const (
	AlphaTypeUnknown AlphaType = iota
	AlphaTypeOpaque
	AlphaTypePremul
	AlphaTypeUnpremul
)

// ColorType describes the bit depth and color components of a pixel.
type ColorType uint8

const (
	ColorTypeUnknown ColorType = iota
	ColorTypeAlpha8
	ColorTypeRGB565
	ColorTypeARGB4444
	ColorTypeRGBA8888
	ColorTypeRGB888x
	ColorTypeBGRA8888
	ColorTypeRGBA1010102
	ColorTypeBGRA1010102
	ColorTypeRGB101010x
	ColorTypeBGR101010x
	ColorTypeGray8
	ColorTypeRGBAF16
	ColorTypeRGBAF16Clamped
	ColorTypeRGBAF32
	ColorTypeR8G8UNorm
	ColorTypeA16Float
	ColorTypeR16G16Float
	ColorTypeA16UNorm
	ColorTypeR16G16UNorm
	ColorTypeR16G16B16A16UNorm
)

// TileMode specifies how shaders tile outside their bounds.
type TileMode uint8

const (
	TileModeClamp TileMode = iota
	TileModeRepeat
	TileModeMirror
	TileModeDecal
)
