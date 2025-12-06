package skia

// ImageInfo describes pixel and encoding. ImageInfo can be created from ColorInfo by
// providing dimensions.
//
// It encodes how pixel bits describe alpha, transparency; color components red, blue,
// and green; and ColorSpace, the range and linearity of colors.
type ImageInfo struct {
	FWidth      int32
	FHeight     int32
	FColorType  ColorType
	FAlphaType  AlphaType
	FColorSpace ColorSpace
}

// ColorInfo describes pixel and encoding without dimensions.
type ColorInfo struct {
	FColorType  ColorType
	FAlphaType  AlphaType
	FColorSpace ColorSpace
}

// MakeImageInfo creates ImageInfo from dimensions and ColorInfo.
func MakeImageInfo(width, height int32, colorInfo ColorInfo) ImageInfo {
	return ImageInfo{
		FWidth:      width,
		FHeight:     height,
		FColorType:  colorInfo.FColorType,
		FAlphaType:  colorInfo.FAlphaType,
		FColorSpace: colorInfo.FColorSpace,
	}
}

// Width returns the width of the image.
func (info ImageInfo) Width() int32 {
	return info.FWidth
}

// Height returns the height of the image.
func (info ImageInfo) Height() int32 {
	return info.FHeight
}

// ColorType returns the color type.
func (info ImageInfo) ColorType() ColorType {
	return info.FColorType
}

// AlphaType returns the alpha type.
func (info ImageInfo) AlphaType() AlphaType {
	return info.FAlphaType
}

// ColorSpace returns the color space.
func (info ImageInfo) ColorSpace() ColorSpace {
	return info.FColorSpace
}

// IsEmpty returns true if width or height is zero or negative.
func (info ImageInfo) IsEmpty() bool {
	return info.FWidth <= 0 || info.FHeight <= 0
}

// IsOpaque returns true if the alpha type is opaque.
func (info ImageInfo) IsOpaque() bool {
	return info.FAlphaType == AlphaTypeOpaque || info.FAlphaType == AlphaTypeUnknown
}

// BytesPerPixel returns the number of bytes per pixel.
func (info ImageInfo) BytesPerPixel() int {
	return ColorTypeBytesPerPixel(info.FColorType)
}

// MinRowBytes returns the minimum row bytes for the given width.
func (info ImageInfo) MinRowBytes() int {
	return info.BytesPerPixel() * int(info.FWidth)
}

// ComputeByteSize returns the total byte size of the image.
func (info ImageInfo) ComputeByteSize(rowBytes int) int {
	if info.IsEmpty() {
		return 0
	}
	return rowBytes * int(info.FHeight)
}

// ColorTypeBytesPerPixel returns the number of bytes required to store a pixel.
func ColorTypeBytesPerPixel(ct ColorType) int {
	switch ct {
	case ColorTypeUnknown:
		return 0
	case ColorTypeAlpha8:
		return 1
	case ColorTypeRGB565:
		return 2
	case ColorTypeARGB4444:
		return 2
	case ColorTypeRGBA8888:
		return 4
	case ColorTypeRGB888x:
		return 4
	case ColorTypeBGRA8888:
		return 4
	case ColorTypeRGBA1010102:
		return 4
	case ColorTypeBGRA1010102:
		return 4
	case ColorTypeRGB101010x:
		return 4
	case ColorTypeBGR101010x:
		return 4
	case ColorTypeGray8:
		return 1
	case ColorTypeRGBAF16:
		return 8
	case ColorTypeRGBAF16Clamped:
		return 8
	case ColorTypeRGBAF32:
		return 16
	case ColorTypeR8G8UNorm:
		return 2
	case ColorTypeA16Float:
		return 2
	case ColorTypeR16G16Float:
		return 4
	case ColorTypeA16UNorm:
		return 2
	case ColorTypeR16G16UNorm:
		return 4
	case ColorTypeR16G16B16A16UNorm:
		return 8
	default:
		return 0
	}
}

// ColorTypeIsAlwaysOpaque returns true if ColorType always decodes alpha to 1.0.
func ColorTypeIsAlwaysOpaque(ct ColorType) bool {
	switch ct {
	case ColorTypeRGB565, ColorTypeRGB888x, ColorTypeRGB101010x, ColorTypeBGR101010x, ColorTypeGray8:
		return true
	default:
		return false
	}
}
