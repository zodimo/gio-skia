package skia

// Canvas provides an interface for drawing, and how the drawing is clipped and transformed.
// Canvas contains a stack of Matrix and clip values.
//
// Canvas and Paint together provide the state to draw into Surface or Device.
// Each Canvas draw call transforms the geometry of the object by the concatenation of all
// Matrix values in the stack. The transformed geometry is clipped by the intersection
// of all of clip values in the stack. The Canvas draw calls use Paint to supply drawing
// state such as color, Typeface, text size, stroke width, Shader and so on.
type Canvas interface {
	// ── State Management ────────────────────────────────────────────────────

	// Save saves Matrix and clip.
	// Calling Restore discards changes to Matrix and clip,
	// restoring the Matrix and clip to their state when Save was called.
	Save() int

	// SaveLayer saves Matrix and clip, and allocates a Surface for subsequent drawing.
	// Calling Restore discards changes to Matrix and clip, and draws the Surface.
	SaveLayer(bounds *Rect, paint Paint) int

	// SaveLayerAlpha saves Matrix and clip, and allocates Surface for subsequent drawing.
	// Calling Restore discards changes to Matrix and clip,
	// and blends layer with alpha opacity onto prior layer.
	SaveLayerAlpha(bounds *Rect, alpha float32) int

	// Restore removes changes to Matrix and clip since Canvas state was last saved.
	Restore()

	// RestoreToCount restores state to Matrix and clip values when Save returned saveCount.
	RestoreToCount(saveCount int)

	// GetSaveCount returns the number of saved states.
	GetSaveCount() int

	// ── Matrix Operations ───────────────────────────────────────────────────

	// Translate translates Matrix by dx along the x-axis and dy along the y-axis.
	Translate(dx, dy Scalar)

	// Scale scales Matrix by sx on the x-axis and sy on the y-axis.
	Scale(sx, sy Scalar)

	// Rotate rotates Matrix by degrees. Positive degrees rotates clockwise.
	Rotate(degrees Scalar)

	// RotateAbout rotates Matrix by degrees about a point at (px, py).
	RotateAbout(degrees, px, py Scalar)

	// Skew skews Matrix by sx on the x-axis and sy on the y-axis.
	Skew(sx, sy Scalar)

	// Concat replaces Matrix with matrix premultiplied with existing Matrix.
	Concat(matrix Matrix)

	// SetMatrix replaces Matrix with matrix.
	SetMatrix(matrix Matrix)

	// ResetMatrix sets Matrix to the identity matrix.
	ResetMatrix()

	// GetTotalMatrix returns the current total matrix.
	GetTotalMatrix() Matrix

	// ── Clip Operations ─────────────────────────────────────────────────────

	// ClipRect replaces clip with the intersection or difference of clip and rect.
	ClipRect(rect Rect, op ClipOp, doAntiAlias bool)

	// ClipRRect replaces clip with the intersection or difference of clip and rrect.
	ClipRRect(rrect RRect, op ClipOp, doAntiAlias bool)

	// ClipPath replaces clip with the intersection or difference of clip and path.
	ClipPath(path Path, op ClipOp, doAntiAlias bool)

	// ClipRegion replaces clip with the intersection or difference of clip and region.
	ClipRegion(region Region, op ClipOp)

	// ClipShader replaces clip with the intersection or difference of clip and shader.
	ClipShader(shader Shader, op ClipOp)

	// QuickReject returns true if rect, transformed by Matrix, can be quickly determined to be outside of clip.
	QuickReject(rect Rect) bool

	// QuickRejectPath returns true if path, transformed by Matrix, can be quickly determined to be outside of clip.
	QuickRejectPath(path Path) bool

	// GetLocalClipBounds returns bounds of clip, transformed by inverse of Matrix.
	GetLocalClipBounds() Rect

	// GetDeviceClipBounds returns IRect bounds of clip, unaffected by Matrix.
	GetDeviceClipBounds() IRect

	// IsClipEmpty returns true if clip is empty; that is, nothing will draw.
	IsClipEmpty() bool

	// IsClipRect returns true if clip is Rect and not empty.
	IsClipRect() bool

	// ── Drawing Operations ──────────────────────────────────────────────────

	// DrawColor fills clip with color.
	DrawColor(color Color, mode BlendMode)

	// DrawColor4f fills clip with color.
	DrawColor4f(color Color4f, mode BlendMode)

	// Clear fills clip with color using BlendModeSrc.
	Clear(color Color)

	// Clear4f fills clip with color using BlendModeSrc.
	Clear4f(color Color4f)

	// Discard makes Canvas contents undefined.
	Discard()

	// DrawPaint fills clip with Paint.
	DrawPaint(paint Paint)

	// DrawPoint draws point at (x, y).
	DrawPoint(x, y Scalar, paint Paint)

	// DrawPointPoint draws point p.
	DrawPointPoint(p Point, paint Paint)

	// DrawPoints draws pts using clip, Matrix and Paint.
	DrawPoints(mode PointMode, points []Point, paint Paint)

	// DrawLine draws line segment from (x0, y0) to (x1, y1).
	DrawLine(x0, y0, x1, y1 Scalar, paint Paint)

	// DrawLinePoints draws line segment from p0 to p1.
	DrawLinePoints(p0, p1 Point, paint Paint)

	// DrawRect draws Rect using clip, Matrix, and Paint.
	DrawRect(rect Rect, paint Paint)

	// DrawIRect draws IRect using clip, Matrix, and Paint.
	DrawIRect(rect IRect, paint Paint)

	// DrawRegion draws Region using clip, Matrix, and Paint.
	DrawRegion(region Region, paint Paint)

	// DrawOval draws oval using clip, Matrix, and Paint.
	DrawOval(oval Rect, paint Paint)

	// DrawRRect draws RRect using clip, Matrix, and Paint.
	DrawRRect(rrect RRect, paint Paint)

	// DrawDRRect draws RRect outer and inner using clip, Matrix, and Paint.
	DrawDRRect(outer, inner RRect, paint Paint)

	// DrawCircle draws circle at (cx, cy) with radius.
	DrawCircle(cx, cy, radius Scalar, paint Paint)

	// DrawCirclePoint draws circle at center with radius.
	DrawCirclePoint(center Point, radius Scalar, paint Paint)

	// DrawArc draws arc using clip, Matrix, and Paint.
	DrawArc(oval Rect, startAngle, sweepAngle Scalar, useCenter bool, paint Paint)

	// DrawRoundRect draws RRect bounded by Rect, with corner radii (rx, ry).
	DrawRoundRect(rect Rect, rx, ry Scalar, paint Paint)

	// DrawPath draws Path using clip, Matrix, and Paint.
	DrawPath(path Path, paint Paint)

	// DrawImage draws Image at (x, y).
	DrawImage(image Image, x, y Scalar, sampling SamplingOptions, paint Paint)

	// DrawImageRect draws Image stretched proportionally to fit into Rect dst.
	DrawImageRect(image Image, src, dst Rect, sampling SamplingOptions, paint Paint, constraint SrcRectConstraint)

	// DrawImageRectDst draws Image stretched proportionally to fit into Rect dst.
	DrawImageRectDst(image Image, dst Rect, sampling SamplingOptions, paint Paint)

	// DrawImageNine draws Image stretched proportionally to fit into Rect dst.
	DrawImageNine(image Image, center IRect, dst Rect, filter FilterMode, paint Paint)

	// DrawImageLattice draws Image stretched proportionally to fit into Rect dst.
	DrawImageLattice(image Image, lattice Lattice, dst Rect, filter FilterMode, paint Paint)

	// DrawSimpleText draws text, with origin at (x, y).
	DrawSimpleText(text []byte, encoding TextEncoding, x, y Scalar, font Font, paint Paint)

	// DrawString draws null terminated string, with origin at (x, y).
	DrawString(str string, x, y Scalar, font Font, paint Paint)

	// DrawGlyphs draws glyphs, at positions relative to origin styled with font and paint.
	DrawGlyphs(glyphs []GlyphID, positions []Point, origin Point, font Font, paint Paint)

	// DrawGlyphsWithClusters draws glyphs with cluster information.
	DrawGlyphsWithClusters(glyphs []GlyphID, positions []Point, clusters []uint32, utf8text []byte, origin Point, font Font, paint Paint)

	// DrawGlyphsRSXform draws glyphs using RSXform transformations.
	DrawGlyphsRSXform(glyphs []GlyphID, xforms []RSXform, origin Point, font Font, paint Paint)

	// DrawTextBlob draws TextBlob at (x, y).
	DrawTextBlob(blob TextBlob, x, y Scalar, paint Paint)

	// DrawPicture draws Picture, using clip and Matrix.
	DrawPicture(picture Picture, matrix Matrix, paint Paint)

	// DrawVertices draws Vertices, a triangle mesh.
	DrawVertices(vertices Vertices, mode BlendMode, paint Paint)

	// DrawMesh draws a mesh using a user-defined specification.
	DrawMesh(mesh Mesh, blender Blender, paint Paint)

	// DrawPatch draws a Coons patch.
	DrawPatch(cubics [12]Point, colors [4]Color, texCoords [4]Point, mode BlendMode, paint Paint)

	// DrawAtlas draws a set of sprites from atlas.
	DrawAtlas(atlas Image, xform []RSXform, tex []Rect, colors []Color, mode BlendMode, sampling SamplingOptions, cullRect *Rect, paint Paint)

	// DrawDrawable draws Drawable using clip and Matrix.
	DrawDrawable(drawable Drawable, matrix Matrix)

	// DrawDrawableXY draws Drawable using clip and Matrix, offset by (x, y).
	DrawDrawableXY(drawable Drawable, x, y Scalar)

	// DrawAnnotation associates Rect on Canvas with an annotation.
	DrawAnnotation(rect Rect, key string, value Data)

	// ── Canvas Information ─────────────────────────────────────────────────

	// ImageInfo returns ImageInfo for Canvas.
	ImageInfo() ImageInfo

	// GetBaseLayerSize gets the size of the base or root layer in global canvas coordinates.
	GetBaseLayerSize() ISize

	// GetBaseProps returns the SurfaceProps associated with the canvas at the base of the layer stack.
	GetBaseProps() SurfaceProps

	// GetTopProps returns the SurfaceProps associated with the canvas that are currently active.
	GetTopProps() SurfaceProps

	// MakeSurface creates Surface matching info and props.
	MakeSurface(info ImageInfo, props SurfaceProps) Surface

	// PeekPixels returns true if Canvas has direct access to its pixels.
	PeekPixels(pixmap Pixmap) bool

	// ReadPixels copies Rect of pixels from Canvas into dstPixels.
	ReadPixels(dstInfo ImageInfo, dstPixels []byte, dstRowBytes int, srcX, srcY int) bool

	// ReadPixelsPixmap copies Rect of pixels from Canvas into pixmap.
	ReadPixelsPixmap(pixmap Pixmap, srcX, srcY int) bool

	// ReadPixelsBitmap copies Rect of pixels from Canvas into bitmap.
	ReadPixelsBitmap(bitmap Bitmap, srcX, srcY int) bool

	// WritePixels copies Rect from pixels to Canvas.
	WritePixels(info ImageInfo, pixels []byte, rowBytes int, x, y int) bool

	// WritePixelsBitmap copies Rect from bitmap to Canvas.
	WritePixelsBitmap(bitmap Bitmap, x, y int) bool
}

// Additional types for Canvas operations

// Region represents a region (set of rectangles).
type Region interface {
	// SetEmpty sets the region to empty.
	SetEmpty()

	// SetRect sets the region to a rectangle.
	SetRect(rect IRect)

	// SetRects sets the region to a set of rectangles.
	SetRects(rects []IRect) bool

	// SetPath sets the region to the path.
	SetPath(path Path, clip IRect) bool

	// Intersect intersects the region with another region.
	Intersect(other Region) bool

	// Union unions the region with another region.
	Union(other Region) bool

	// Op performs a region operation.
	Op(other Region, op RegionOp) bool

	// Contains returns true if the region contains the point.
	Contains(x, y int32) bool

	// ContainsRect returns true if the region contains the rectangle.
	ContainsRect(rect IRect) bool

	// QuickContains returns true if the region quickly contains the rectangle.
	QuickContains(rect IRect) bool

	// QuickReject returns true if the region quickly rejects the rectangle.
	QuickReject(rect IRect) bool

	// Bounds returns the bounding rectangle of the region.
	Bounds() IRect

	// IsEmpty returns true if the region is empty.
	IsEmpty() bool

	// IsRect returns true if the region is a rectangle.
	IsRect() bool

	// IsComplex returns true if the region is complex.
	IsComplex() bool
}

// RegionOp specifies a region operation.
type RegionOp uint8

const (
	RegionOpDifference RegionOp = iota
	RegionOpIntersect
	RegionOpUnion
	RegionOpXOR
	RegionOpReverseDifference
	RegionOpReplace
)

// Image represents an image.
type Image interface {
	// Width returns the width of the image.
	Width() int32

	// Height returns the height of the image.
	Height() int32

	// ImageInfo returns the ImageInfo of the image.
	ImageInfo() ImageInfo

	// UniqueID returns the unique ID of the image.
	UniqueID() uint32

	// AlphaType returns the alpha type of the image.
	AlphaType() AlphaType

	// ColorType returns the color type of the image.
	ColorType() ColorType

	// ColorSpace returns the color space of the image.
	ColorSpace() ColorSpace

	// RefColorSpace returns a reference to the color space.
	RefColorSpace() ColorSpace

	// Bounds returns the bounds of the image.
	Bounds() IRect

	// Dimension returns the dimensions of the image.
	Dimension() ISize
}

// Font represents a font.
type Font interface {
	// Typeface returns the typeface.
	Typeface() Typeface

	// SetTypeface sets the typeface.
	SetTypeface(typeface Typeface)

	// Size returns the font size.
	Size() Scalar

	// SetSize sets the font size.
	SetSize(size Scalar)

	// ScaleX returns the x-axis scale factor.
	ScaleX() Scalar

	// SetScaleX sets the x-axis scale factor.
	SetScaleX(scaleX Scalar)

	// SkewX returns the x-axis skew factor.
	SkewX() Scalar

	// SetSkewX sets the x-axis skew factor.
	SetSkewX(skewX Scalar)

	// Edging returns the edging mode.
	Edging() FontEdging

	// SetEdging sets the edging mode.
	SetEdging(edging FontEdging)

	// Hinting returns the hinting mode.
	Hinting() FontHinting

	// SetHinting sets the hinting mode.
	SetHinting(hinting FontHinting)

	// Subpixel returns true if subpixel positioning is enabled.
	Subpixel() bool

	// SetSubpixel sets subpixel positioning.
	SetSubpixel(subpixel bool)

	// Embolden returns true if emboldening is enabled.
	Embolden() bool

	// SetEmbolden sets emboldening.
	SetEmbolden(embolden bool)

	// LinearMetrics returns true if linear metrics are enabled.
	LinearMetrics() bool

	// SetLinearMetrics sets linear metrics.
	SetLinearMetrics(linearMetrics bool)

	// BaselineSnap returns true if baseline snapping is enabled.
	BaselineSnap() bool

	// SetBaselineSnap sets baseline snapping.
	SetBaselineSnap(baselineSnap bool)

	// Metrics returns the font metrics.
	Metrics() FontMetrics

	// MeasureText measures the text.
	MeasureText(text []byte, encoding TextEncoding) Scalar

	// GetBounds returns the bounds of the glyphs.
	GetBounds(glyphs []GlyphID, paint Paint) Rect

	// GetWidths returns the widths of the glyphs.
	GetWidths(glyphs []GlyphID, widths []Scalar) bool
}

// Typeface represents a typeface.
type Typeface interface{}

// FontEdging specifies the edging mode.
type FontEdging uint8

const (
	FontEdgingAlias FontEdging = iota
	FontEdgingAntiAlias
	FontEdgingSubpixelAntiAlias
)

// FontHinting specifies the hinting mode.
type FontHinting uint8

const (
	FontHintingNone FontHinting = iota
	FontHintingSlight
	FontHintingNormal
	FontHintingFull
)

// FontMetrics contains font metrics.
type FontMetrics struct {
	FFlags              uint32
	FTop                Scalar
	FAscent             Scalar
	FDescent            Scalar
	FBottom             Scalar
	FLeading            Scalar
	FAvgCharWidth       Scalar
	FMaxCharWidth       Scalar
	FXMin               Scalar
	FYMin               Scalar
	FXMax               Scalar
	FYMax               Scalar
	FUnderlineThickness Scalar
	FUnderlinePosition  Scalar
	FStrikeoutThickness Scalar
	FStrikeoutPosition  Scalar
}

// GlyphID represents a glyph ID.
type GlyphID uint16

// TextBlob represents a text blob.
type TextBlob interface{}

// Picture represents a picture.
type Picture interface{}

// Vertices represents vertices.
type Vertices interface{}

// Mesh represents a mesh.
type Mesh interface{}

// Blender represents a blender.
type Blender interface{}

// Drawable represents a drawable.
type Drawable interface{}

// Data represents data.
type Data interface{}

// RSXform represents a RSXform transformation.
type RSXform struct {
	FSx, FKy, FTx, FTy Scalar
}

// SamplingOptions represents sampling options.
type SamplingOptions struct {
	// TODO: Add fields based on SkSamplingOptions
}

// FilterMode specifies the filter mode.
type FilterMode uint8

const (
	FilterModeNearest FilterMode = iota
	FilterModeLinear
)

// Lattice represents a lattice.
type Lattice struct {
	FXDivs     []int32
	FYDivs     []int32
	FRectTypes []LatticeRectType
	FXCount    int32
	FYCount    int32
	FBounds    *IRect
	FColors    []Color
}

// LatticeRectType specifies the type of a lattice rectangle.
type LatticeRectType uint8

const (
	LatticeRectTypeDefault LatticeRectType = iota
	LatticeRectTypeTransparent
	LatticeRectTypeFixedColor
)

// SurfaceProps represents surface properties.
type SurfaceProps interface{}

// Surface represents a surface.
type Surface interface{}

// Pixmap represents a pixmap.
type Pixmap interface{}

// Bitmap represents a bitmap.
type Bitmap interface{}
