# Skia Example

https://skia.org/docs/user/api/skpath_overview/


```c
void draw(SkCanvas* canvas) {
    SkPaint paint;
    paint.setAntiAlias(true);
    SkPath path;
    path.moveTo(36, 48);
    path.quadTo(66, 88, 120, 36);
    canvas->drawPath(path, paint);
    paint.setStyle(SkPaint::kStroke_Style);
    paint.setColor(SK_ColorBLUE);
    paint.setStrokeWidth(8);
    canvas->translate(0, 50);
    canvas->drawPath(path, paint);
    paint.setStyle(SkPaint::kStrokeAndFill_Style);
    paint.setColor(SK_ColorRED);
    canvas->translate(0, 50);
    canvas->drawPath(path, paint);
}

```