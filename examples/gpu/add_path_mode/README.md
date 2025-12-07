# Skia Fiddle example

https://fiddle.skia.org/c/@Path_AddPathMode

```c
void draw(SkCanvas* canvas) {
    SkPathBuilder path;
    path.moveTo(20, 20)
        .lineTo(20, 40)
        .lineTo(40, 20);
    SkPath path2 = SkPathBuilder()
                   .moveTo(60, 60)
                   .lineTo(80, 60)
                   .lineTo(80, 40)
                   .detach();
    SkPaint paint;
    paint.setStyle(SkPaint::kStroke_Style);
    for (int i = 0; i < 2; i++) {
        for (auto addPathMode : { SkPath::kAppend_AddPathMode, SkPath::kExtend_AddPathMode } ) {
            SkPath test = SkPathBuilder(path)
                          .addPath(path2, addPathMode)
                          .detach();
            canvas->drawPath(test, paint);
            canvas->translate(100, 0);
        }
        canvas->translate(-200, 100);
        path.close();
    }
}

```