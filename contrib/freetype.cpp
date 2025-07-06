//go:build linux && (!gocv_specific_modules || (gocv_specific_modules && gocv_contrib_freetype))

#ifndef _WIN32  // Exclude compiling on Windows platforms

#include "freetype.h"

FreeType2 FreeType2_CreateFreeType2() {
    try {
        return new cv::Ptr<cv::freetype::FreeType2>(cv::freetype::createFreeType2());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void FreeType2_Close(FreeType2 f) {
    delete f;
}

OpenCVResult FreeType2_LoadFontData(FreeType2 f, const char *fontFileName, int id) {
    try {
        (*f)->loadFontData(fontFileName, id);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

void FreeType2_SetSplitNumber(FreeType2 f, int num) {
    try {
        (*f)->setSplitNumber(num);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

OpenCVResult FreeType2_PutText(FreeType2 f, Mat img, const char *text, Point org,
                       int fontHeight, Scalar color,
                       int thickness, int line_type, bool bottomLeftOrigin) {
    try {
        cv::Point pt(org.x, org.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        (*f)->putText(*img, text, pt, fontHeight, c, thickness, line_type, bottomLeftOrigin);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }                    
}

Size FreeType2_GetTextSize(FreeType2 f, const char *text, int fontHeight, int thickness, int *baseLine) {
    try {
        cv::Size sz = (*f)->getTextSize(text, fontHeight, thickness, baseLine);
        return Size{sz.width, sz.height};
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Size sz = {0, 0};
        return sz;
    }                        
}

#endif // _WIN32