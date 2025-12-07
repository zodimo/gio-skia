# Skia Examples Comparison Links

This document provides links to official Skia examples and documentation that correspond to the canonical examples created in this repository. Use these links to compare developer experience and functionality between gio-skia and native Skia implementations.

## General Skia Resources

### Official Documentation & Examples
- **Skia Official Documentation**: https://skia.org/docs/
- **Skia Samples and Tutorials**: https://skia.org/docs/user/sample/
- **Skia API Reference**: https://api.skia.org/
- **Skia Source Code Examples**: https://skia.googlesource.com/skia/+/refs/heads/main/docs/examples/

### Interactive Tools
- **Skia Fiddle** (Interactive Skia Playground): https://fiddle.skia.org/
  - Great for testing and comparing Skia code interactively
  - Search for examples by feature: blend modes, paths, matrices, etc.

## Example-Specific Comparison Links

### 1. Blend Modes (`blend_modes/main.go`)

**Skia Documentation:**
- **SkBlendMode Enum**: https://api.skia.org/classSkBlendMode.html
- **SkPaint::setBlendMode()**: https://api.skia.org/classSkPaint.html#a8c48c5a62987c997f664bcae8c42af2a

**Example Code References:**
- **Skia Fiddle - Blend Modes**: Search "blend mode" at https://fiddle.skia.org/
- **SkiaSharp Blend Modes Example**: https://learn.microsoft.com/en-us/dotnet/api/skiasharp.skblendmode
- **React Native Skia Blend Modes**: https://shopify.github.io/react-native-skia/docs/effects/blend-modes

**Key Concepts:**
- Porter-Duff blend modes (SrcOver, Multiply, Screen, etc.)
- Advanced blend modes (Overlay, Darken, Lighten, etc.)
- Compositing operations

### 2. Fill Rules (`fill_rules/main.go`)

**Skia Documentation:**
- **SkPath::FillType**: https://api.skia.org/classSkPath.html#a0e0a4c0c8c8c8c8c8c8c8c8c8c8c8c8
- **Path Fill Types**: https://api.skia.org/namespaceSkPathFillType.html

**Example Code References:**
- **Skia Fiddle - Fill Rules**: Search "fill type" or "winding" at https://fiddle.skia.org/
- **Skia Path Fill Types Documentation**: https://skia.org/docs/user/api/skpath_overview/#fill-types

**Key Concepts:**
- Winding rule (non-zero winding)
- EvenOdd rule (alternating)
- Path direction (CW vs CCW)
- Creating holes in paths

### 3. Path Operations (`path_operations/main.go`)

**Skia Documentation:**
- **SkPath Class Reference**: https://api.skia.org/classSkPath.html
- **Path Operations**: https://skia.org/docs/user/api/skpath_overview/
- **SkPath::addPath()**: https://api.skia.org/classSkPath.html#a8c48c5a62987c997f664bcae8c42af2a
- **SkPath::transform()**: https://api.skia.org/classSkPath.html#a8c48c5a62987c997f664bcae8c42af2a

**Example Code References:**
- **Skia Fiddle - Path Operations**: Search "path addPath" or "path transform" at https://fiddle.skia.org/
- **Skia Path Examples**: https://skia.googlesource.com/skia/+/refs/heads/main/docs/examples/Path_Overview.cpp

**Key Concepts:**
- Path addition and combination
- Path transformation with matrices
- Path offset operations
- Multiple contours
- Bounds computation

### 4. Transparency (`transparency/main.go`)

**Skia Documentation:**
- **SkPaint::setAlpha()**: https://api.skia.org/classSkPaint.html#a8c48c5a62987c997f664bcae8c42af2a
- **SkPaint::setAlphaf()**: https://api.skia.org/classSkPaint.html#a8c48c5a62987c997f664bcae8c42af2a
- **Alpha Compositing**: https://skia.org/docs/user/api/skpaint_overview/#alpha

**Example Code References:**
- **Skia Fiddle - Transparency**: Search "alpha" or "transparency" at https://fiddle.skia.org/
- **SkiaSharp Alpha Examples**: https://learn.microsoft.com/en-us/dotnet/api/skiasharp.skpaint.alpha

**Key Concepts:**
- Alpha channel compositing
- Semi-transparent shapes
- Layered transparency
- Alpha blending

### 5. Matrix Operations (`matrix_operations/main.go`)

**Skia Documentation:**
- **SkMatrix Class Reference**: https://api.skia.org/classSkMatrix.html
- **Matrix Transformations**: https://skia.org/docs/user/api/skmatrix_overview/
- **SkCanvas::concat()**: https://api.skia.org/classSkCanvas.html#a8c48c5a62987c997f664bcae8c42af2a

**Example Code References:**
- **Skia Fiddle - Matrix Operations**: Search "matrix" or "transform" at https://fiddle.skia.org/
- **Skia Matrix Examples**: https://skia.googlesource.com/skia/+/refs/heads/main/docs/examples/Matrix_Overview.cpp

**Key Concepts:**
- Matrix concatenation (pre vs post)
- Matrix multiplication
- Matrix inversion
- Point mapping
- Pivot point transformations

## Additional Comparison Resources

### SkiaSharp Examples (.NET)
- **SkiaSharp Samples**: https://github.com/mono/SkiaSharp/tree/main/samples
- **SkiaSharp MAUI Demos**: https://learn.microsoft.com/en-us/samples/dotnet/maui-samples/skiasharpmaui-demos/

### React Native Skia Examples
- **React Native Skia Documentation**: https://shopify.github.io/react-native-skia/
- **React Native Skia Examples**: https://github.com/Shopify/react-native-skia/tree/main/example

### C++ Skia Examples
- **Skia C++ Examples**: https://skia.googlesource.com/skia/+/refs/heads/main/docs/examples/
- **Skia Viewer Source**: https://skia.googlesource.com/skia/+/refs/heads/main/tools/viewer/

## How to Use These Links for Comparison

1. **Functionality Comparison**: Compare the visual output of gio-skia examples with Skia Fiddle examples
2. **API Comparison**: Compare the API calls in gio-skia with the official Skia C++ API documentation
3. **Developer Experience**: Compare code structure, ease of use, and verbosity between implementations
4. **Performance**: Note any differences in rendering performance (if measurable)

## Running Comparisons

### Using Skia Fiddle:
1. Go to https://fiddle.skia.org/
2. Search for examples matching each feature (blend modes, fill rules, etc.)
3. Compare the code structure and visual output
4. Try porting gio-skia examples to Skia Fiddle format

### Using Official Examples:
1. Clone the Skia repository: `git clone https://skia.googlesource.com/skia`
2. Build the examples: See https://skia.org/docs/user/build/
3. Run the viewer: `out/Release/viewer`
4. Compare visual output with gio-skia examples

## Notes

- **Skia Fiddle** is the easiest way to quickly test and compare Skia code
- **API Documentation** provides the most detailed reference for method signatures
- **Source Code Examples** show real-world usage patterns
- Some features may have different names or implementations across language bindings

