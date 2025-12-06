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
	Save() (saveCount int)

	// SaveLayer saves Matrix and clip, and allocates a Surface for subsequent drawing.
	// Calling Restore discards changes to Matrix and clip, and draws the Surface.
	SaveLayer(bounds *Rect, paint Paint) (saveCount int)

	// SaveLayerAlpha saves Matrix and clip, and allocates Surface for subsequent drawing.
	// Calling Restore discards changes to Matrix and clip,
	// and blends layer with alpha opacity onto prior layer.
	SaveLayerAlpha(bounds *Rect, alpha float32) (saveCount int)

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

	// DrawArcArc draws arc from Arc specification using clip, Matrix, and Paint.
	DrawArcArc(arc Arc, paint Paint)

	// DrawRoundRect draws RRect bounded by Rect, with corner radii (rx, ry).
	DrawRoundRect(rect Rect, rx, ry Scalar, paint Paint)

	// DrawPath draws Path using clip, Matrix, and Paint.
	DrawPath(path Path, paint Paint)

	// DrawImage draws Image at (x, y) with sampling options and paint.
	// paint may be nil to use default paint settings.
	DrawImage(image Image, x, y Scalar, sampling SamplingOptions, paint Paint)

	// DrawImageSimple draws Image at (x, y) with default sampling (nearest neighbor) and no paint.
	// This is a convenience method equivalent to DrawImage with zero-value SamplingOptions and nil paint.
	DrawImageSimple(image Image, x, y Scalar)

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
	// matrix may be nil to use identity matrix, paint may be nil to use default paint settings.
	DrawPicture(picture Picture, matrix Matrix, paint Paint)

	// DrawPictureSimple draws Picture using current clip and matrix, without additional transform or paint.
	// This is a convenience method equivalent to DrawPicture with nil matrix and nil paint.
	DrawPictureSimple(picture Picture)

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

	// ComputeRegionComplexity returns a value that increases with the number of elements.
	ComputeRegionComplexity() int

	// GetBoundaryPath appends outline of region to path builder.
	GetBoundaryPath(path PathBuilder) bool
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

	// MakeShader creates a shader from the image.
	MakeShader(tmx, tmy TileMode, sampling SamplingOptions, localMatrix Matrix) Shader

	// MakeRawShader creates a raw shader from the image.
	MakeRawShader(tmx, tmy TileMode, sampling SamplingOptions, localMatrix Matrix) Shader

	// MakeSubset creates a subset of the image.
	MakeSubset(subset IRect) Image

	// MakeWithFilter creates a filtered image.
	MakeWithFilter(filter ImageFilter, subset, clipBounds IRect, outSubset *IRect, offset *IPoint) Image

	// PeekPixels returns true if the image has direct access to its pixels.
	PeekPixels(pixmap Pixmap) bool

	// ReadPixels copies pixels from the image.
	ReadPixels(dstInfo ImageInfo, dstPixels []byte, dstRowBytes int, srcX, srcY int) bool

	// ReadPixelsPixmap copies pixels from the image into a pixmap.
	ReadPixelsPixmap(pixmap Pixmap, srcX, srcY int) bool

	// ScalePixels scales pixels from the image.
	ScalePixels(dst Pixmap, sampling SamplingOptions) bool
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
type Typeface interface {
	// FontStyle returns the typeface's intrinsic style attributes.
	FontStyle() FontStyle

	// IsBold returns true if style has the bold bit set.
	IsBold() bool

	// IsItalic returns true if style has the italic bit set.
	IsItalic() bool

	// IsFixedPitch returns true if the typeface claims to be fixed-pitch.
	IsFixedPitch() bool

	// UniqueID returns a 32bit value for this typeface, unique for the underlying font data.
	UniqueID() uint32

	// CountGlyphs returns the number of glyphs in the typeface.
	CountGlyphs() int

	// UnicharToGlyph returns the glyphID that corresponds to the specified unicode code-point.
	UnicharToGlyph(unichar uint32) GlyphID

	// UnicharsToGlyphs returns the corresponding glyph IDs for each character.
	UnicharsToGlyphs(unis []uint32, glyphs []GlyphID)

	// TextToGlyphs converts text to glyphs.
	TextToGlyphs(text []byte, encoding TextEncoding, glyphs []GlyphID) int

	// CountTables returns the number of tables in the font.
	CountTables() int

	// ReadTableTags copies the list of table tags in the font.
	ReadTableTags(tags []uint32) int

	// GetTableSize returns the size of a table's contents.
	GetTableSize(tag uint32) int

	// GetTableData copies the contents of a table.
	GetTableData(tag uint32, offset, length int, data []byte) int

	// GetBounds returns the bounds of the typeface.
	GetBounds() Rect

	// GetMetrics returns the font metrics.
	GetMetrics(metrics *FontMetrics)

	// GetKerningPairAdjustments returns the kerning pair adjustments.
	GetKerningPairAdjustments(firstGlyphs, secondGlyphs []GlyphID, adjustments []Scalar) int
}

// FontStyle represents font style attributes.
type FontStyle struct {
	Weight int32
	Width  int32
	Slant  FontSlant
}

// FontSlant specifies the font slant.
type FontSlant uint8

const (
	FontSlantUpright FontSlant = iota
	FontSlantItalic
	FontSlantOblique
)

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
type TextBlob interface {
	// Bounds returns conservative bounding box.
	Bounds() Rect

	// UniqueID returns a non-zero value unique among all text blobs.
	UniqueID() uint32

	// GetIntercepts returns the number of intervals that intersect bounds.
	GetIntercepts(bounds [2]Scalar, intervals []Scalar, paint Paint) int
}

// Picture represents a picture.
type Picture interface {
	// Playback replays the drawing commands on the specified canvas.
	Playback(canvas Canvas, callback AbortCallback)

	// CullRect returns cull Rect for this picture.
	CullRect() Rect

	// ApproximateBytesUsed returns approximate memory usage.
	ApproximateBytesUsed() int

	// ApproximateOpCount returns approximate number of operations.
	ApproximateOpCount(nested bool) int

	// UniqueID returns unique value identifying the content.
	UniqueID() uint32
}

// AbortCallback allows interruption of picture playback.
type AbortCallback interface {
	// Abort returns true to stop playback.
	Abort() bool
}

// Vertices represents vertices.
type Vertices interface {
	// UniqueID returns a unique value for this instance.
	UniqueID() uint32

	// Bounds returns the bounds of the vertices.
	Bounds() Rect

	// ApproximateSize returns approximate byte size of the vertices object.
	ApproximateSize() int

	// Mode returns the vertex mode.
	Mode() VertexMode

	// VertexCount returns the number of vertices.
	VertexCount() int

	// IndexCount returns the number of indices.
	IndexCount() int
}

// VertexMode specifies how vertices are interpreted.
type VertexMode uint8

const (
	VertexModeTriangles VertexMode = iota
	VertexModeTriangleStrip
	VertexModeTriangleFan
)

// Mesh represents a mesh.
type Mesh interface {
	// Bounds returns the bounds of the mesh.
	Bounds() Rect

	// UniqueID returns a unique value for this instance.
	UniqueID() uint32

	// Specification returns the mesh specification.
	Specification() MeshSpecification

	// VertexCount returns the number of vertices.
	VertexCount() int

	// IndexCount returns the number of indices.
	IndexCount() int
}

// MeshSpecification represents a mesh specification.
type MeshSpecification interface {
	// Stride returns the vertex stride.
	Stride() int

	// AttributeCount returns the number of attributes.
	AttributeCount() int

	// VaryingCount returns the number of varyings.
	VaryingCount() int
}

// Blender represents a blender.
type Blender interface {
	// Mode creates a blender that implements the specified BlendMode.
	Mode(mode BlendMode) Blender
}

// Drawable represents a drawable.
type Drawable interface {
	// Draw draws into the specified canvas.
	Draw(canvas Canvas, matrix Matrix)

	// DrawXY draws into the specified canvas at the given position.
	DrawXY(canvas Canvas, x, y Scalar)

	// GetGenerationID returns a unique value for this instance.
	GetGenerationID() uint32

	// GetBounds returns the (conservative) bounds of what the drawable will draw.
	GetBounds() Rect

	// ApproximateBytesUsed returns approximately how many bytes would be freed if this drawable is destroyed.
	ApproximateBytesUsed() int

	// NotifyDrawingChanged invalidates the previous generation ID.
	NotifyDrawingChanged()
}

// Data represents data.
type Data interface {
	// Size returns the number of bytes stored.
	Size() int

	// Data returns the ptr to the data.
	Data() []byte

	// Empty returns true if the data is empty.
	Empty() bool

	// Bytes returns the data as bytes.
	Bytes() []byte

	// WritableData returns writable data pointer.
	WritableData() []byte

	// CopySubset attempts to create a deep copy of a subset of the original data.
	CopySubset(offset, length int) Data

	// ShareSubset attempts to return a data that is a reference to a subset of the original data.
	ShareSubset(offset, length int) Data

	// CopyRange copies a range of the data into a caller-provided buffer.
	CopyRange(offset, length int, buffer []byte) int
}

// RSXform represents a RSXform transformation.
type RSXform struct {
	FSx, FKy, FTx, FTy Scalar
}

// SamplingOptions represents sampling options for image drawing operations.
// It controls how pixels are sampled when images are scaled or transformed.
type SamplingOptions struct {
	// MaxAniso specifies maximum anisotropy for anisotropic filtering.
	// Zero means anisotropic filtering is disabled.
	MaxAniso int

	// UseCubic indicates whether cubic resampling should be used.
	UseCubic bool

	// Cubic specifies the cubic resampling coefficients (B, C).
	// Only used when UseCubic is true.
	Cubic CubicResampler

	// Filter specifies the filter mode (nearest neighbor or linear interpolation).
	Filter FilterMode

	// Mipmap specifies how mipmap levels are sampled.
	Mipmap MipmapMode
}

// NewSamplingOptions creates SamplingOptions with the specified filter and mipmap modes.
func NewSamplingOptions(filter FilterMode, mipmap MipmapMode) SamplingOptions {
	return SamplingOptions{
		Filter: filter,
		Mipmap: mipmap,
	}
}

// NewSamplingOptionsFilter creates SamplingOptions with the specified filter mode
// and no mipmap sampling.
func NewSamplingOptionsFilter(filter FilterMode) SamplingOptions {
	return SamplingOptions{
		Filter: filter,
		Mipmap: MipmapModeNone,
	}
}

// NewSamplingOptionsCubic creates SamplingOptions with cubic resampling.
func NewSamplingOptionsCubic(cubic CubicResampler) SamplingOptions {
	return SamplingOptions{
		UseCubic: true,
		Cubic:    cubic,
		Filter:   FilterModeLinear, // Cubic implies linear filtering
		Mipmap:   MipmapModeNone,
	}
}

// NewSamplingOptionsAniso creates SamplingOptions with anisotropic filtering.
func NewSamplingOptionsAniso(maxAniso int) SamplingOptions {
	if maxAniso < 1 {
		maxAniso = 1
	}
	return SamplingOptions{
		MaxAniso: maxAniso,
		Filter:   FilterModeLinear,
		Mipmap:   MipmapModeNone,
	}
}

// IsAniso returns true if anisotropic filtering is enabled.
func (so SamplingOptions) IsAniso() bool {
	return so.MaxAniso != 0
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
type SurfaceProps interface {
	// Flags returns the flags.
	Flags() uint32

	// PixelGeometry returns the pixel geometry.
	PixelGeometry() PixelGeometry

	// TextContrast returns the text contrast.
	TextContrast() Scalar

	// TextGamma returns the text gamma.
	TextGamma() Scalar

	// IsUseDeviceIndependentFonts returns true if device independent fonts are enabled.
	IsUseDeviceIndependentFonts() bool

	// IsAlwaysDither returns true if dithering is always enabled.
	IsAlwaysDither() bool
}

// PixelGeometry specifies how LCD strips are arranged for each pixel.
type PixelGeometry uint8

const (
	PixelGeometryUnknown PixelGeometry = iota
	PixelGeometryRGBH
	PixelGeometryBGRH
	PixelGeometryRGBV
	PixelGeometryBGRV
)

// Surface represents a surface.
type Surface interface {
	// Width returns pixel count in each row.
	Width() int

	// Height returns pixel row count.
	Height() int

	// ImageInfo returns an ImageInfo describing the surface.
	ImageInfo() ImageInfo

	// GenerationID returns unique value identifying the content of Surface.
	GenerationID() uint32

	// GetCanvas returns the canvas associated with this surface.
	GetCanvas() Canvas

	// MakeImageSnapshot creates an image from the current surface contents.
	MakeImageSnapshot(subset IRect) Image

	// MakeSurface creates a compatible surface.
	MakeSurface(info ImageInfo, props SurfaceProps) Surface

	// PeekPixels returns true if Surface has direct access to its pixels.
	PeekPixels(pixmap Pixmap) bool

	// ReadPixels copies Rect of pixels from Surface into dstPixels.
	ReadPixels(dstInfo ImageInfo, dstPixels []byte, dstRowBytes int, srcX, srcY int) bool

	// ReadPixelsPixmap copies Rect of pixels from Surface into pixmap.
	ReadPixelsPixmap(pixmap Pixmap, srcX, srcY int) bool

	// WritePixels copies Rect from pixels to Surface.
	WritePixels(info ImageInfo, pixels []byte, rowBytes int, x, y int) bool

	// Flush flushes any pending operations.
	Flush()

	// FlushAndSubmit flushes and submits pending operations.
	FlushAndSubmit(syncCpu bool)
}

// Pixmap represents a pixmap.
type Pixmap interface {
	// Reset sets width, height, row bytes to zero; pixel address to nullptr.
	Reset()

	// ResetInfo sets width, height, SkAlphaType, and SkColorType from info.
	ResetInfo(info ImageInfo, addr []byte, rowBytes int)

	// SetColorSpace changes SkColorSpace in SkImageInfo.
	SetColorSpace(colorSpace ColorSpace)

	// ExtractSubset sets subset width, height, pixel address to intersection of Pixmap with area.
	ExtractSubset(subset Pixmap, area IRect) bool

	// Info returns width, height, SkAlphaType, SkColorType, and SkColorSpace.
	Info() ImageInfo

	// RowBytes returns row bytes, the interval from one pixel row to the next.
	RowBytes() int

	// Addr returns pixel address, the base address corresponding to the pixel origin.
	Addr() []byte

	// Width returns pixel count in each pixel row.
	Width() int

	// Height returns pixel row count.
	Height() int

	// Dimensions returns the dimensions of the pixmap.
	Dimensions() ISize

	// ColorType returns the color type.
	ColorType() ColorType

	// AlphaType returns the alpha type.
	AlphaType() AlphaType

	// ColorSpace returns the color space.
	ColorSpace() ColorSpace

	// RefColorSpace returns smart pointer to color space.
	RefColorSpace() ColorSpace

	// IsOpaque returns true if SkAlphaType is opaque.
	IsOpaque() bool

	// GetPixel returns the pixel at the specified coordinates.
	GetPixel(x, y int) Color

	// ReadPixels copies pixels from the pixmap.
	ReadPixels(dstInfo ImageInfo, dstPixels []byte, dstRowBytes int, srcX, srcY int) bool

	// ScalePixels scales pixels from the pixmap.
	ScalePixels(dst Pixmap, sampling SamplingOptions) bool
}

// Bitmap represents a bitmap.
type Bitmap interface {
	// Info returns width, height, SkAlphaType, SkColorType, and SkColorSpace.
	Info() ImageInfo

	// Width returns pixel count in each row.
	Width() int

	// Height returns pixel row count.
	Height() int

	// ColorType returns the color type.
	ColorType() ColorType

	// AlphaType returns the alpha type.
	AlphaType() AlphaType

	// ColorSpace returns the color space.
	ColorSpace() ColorSpace

	// RefColorSpace returns smart pointer to color space.
	RefColorSpace() ColorSpace

	// BytesPerPixel returns number of bytes per pixel required by SkColorType.
	BytesPerPixel() int

	// RowBytes returns number of bytes per row.
	RowBytes() int

	// RowBytesAsPixels returns number of pixels that fit on row.
	RowBytesAsPixels() int

	// IsEmpty returns true if width or height is zero.
	IsEmpty() bool

	// IsNull returns true if pixels address is nullptr.
	IsNull() bool

	// DrawsNothing returns true if width or height is zero, or if pixels is nullptr.
	DrawsNothing() bool

	// GetAddr returns the address of the pixels.
	GetAddr(x, y int) []byte

	// ReadPixels copies pixels from the bitmap.
	ReadPixels(dstInfo ImageInfo, dstPixels []byte, dstRowBytes int, srcX, srcY int) bool

	// ReadPixelsPixmap copies pixels from the bitmap into a pixmap.
	ReadPixelsPixmap(pixmap Pixmap, srcX, srcY int) bool

	// WritePixels copies pixels to the bitmap.
	WritePixels(info ImageInfo, pixels []byte, rowBytes int, x, y int) bool

	// PeekPixels returns true if bitmap has direct access to its pixels.
	PeekPixels(pixmap Pixmap) bool

	// ExtractSubset sets subset to intersection of bitmap and area.
	ExtractSubset(subset Bitmap, area IRect) bool
}
