//
// Created by rainy on 2024/3/2.
//

#ifndef OPENCV_DART_LIBRARY_EXCEPTION_H
#define OPENCV_DART_LIBRARY_EXCEPTION_H

#include "core.h"

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

typedef void (*ErrorCallback)(int status, const char *func_name,
                             const char *err_msg, const char *file_name,
                             int line, void *userdata);

void registerErrorCallback(ErrorCallback callback);

int CvException_GetCode(CvException exception);

const char *CvException_GetErr(CvException exception);

const char *CvException_GetFunc(CvException exception);

const char *CvException_GetFile(CvException exception);

int CvException_GetLine(CvException exception);

void CvException_Close(CvException exception);

#ifdef __cplusplus
}
#endif

#endif //OPENCV_DART_LIBRARY_EXCEPTION_H
