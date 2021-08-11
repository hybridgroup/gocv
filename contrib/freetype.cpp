#include "freetype.h"

FreeType2 FreeType2_CreateFreeType2() {
    return new cv::Ptr<cv::freetype::FreeType2>(cv::freetype::createFreeType2());
}

void FreeType2_Close(FreeType2 f) {
    delete f;
}

void FreeType2_LoadFontData(FreeType2 f, const char *fontFileName, int id) {
    (*f)->loadFontData(fontFileName, id);
}

void FreeType2_SetSplitNumber(FreeType2 f, int num) {
    (*f)->setSplitNumber(num);
}

void FreeType2_PutText(FreeType2 f, Mat img, const char *text, Point org,
                       int fontHeight, Scalar color,
                       int thickness, int line_type, bool bottomLeftOrigin) {
    cv::Point pt(org.x, org.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    (*f)->putText(*img, text, pt, fontHeight, c, thickness, line_type, bottomLeftOrigin);
}

Size FreeType2_GetTextSize(FreeType2 f, const char *text,
                           int fontHeight, int thickness, int *baseLine) {
    cv::Size sz = (*f)->getTextSize(text, fontHeight, thickness, baseLine);
    return Size{sz.width, sz.height};
}