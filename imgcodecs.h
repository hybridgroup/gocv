#ifndef _OPENCV3_IMGCODECS_H_
#define _OPENCV3_IMGCODECS_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

Mat Image_IMRead(const char* filename, int flags);
bool Image_IMWrite(const char* filename, Mat img);
bool Image_IMWrite_WithParams(const char* filename, Mat img, IntVector params);
void Image_IMEncode(const char* fileExt, Mat img, void* vector);

void Image_IMEncode_WithParams(const char* fileExt, Mat img, IntVector params, void* vector);
Mat Image_IMDecode(ByteArray buf, int flags);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGCODECS_H_
