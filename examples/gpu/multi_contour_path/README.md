# Skia Example

https://skia.org/docs/user/api/skpath_overview/


```c
void draw(SkCanvas* canvas) {
    SkPaint paint;
    paint.setAntiAlias(true);
    SkPath path;
    path.moveTo(124, 108);
    path.lineTo(172, 24);
    path.addCircle(50, 50, 30);
    path.moveTo(36, 148);
    path.quadTo(66, 188, 120, 136);
    canvas->drawPath(path, paint);
    paint.setStyle(SkPaint::kStroke_Style);
    paint.setColor(SK_ColorBLUE);
    paint.setStrokeWidth(3);
    canvas->drawPath(path, paint);
}

```

