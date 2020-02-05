#ifndef _OPENCV3_CUDA_CORE_H_
#define _OPENCV3_CUDA_CORE_H_

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

void GpuRects_Close(struct Rects rs);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_CORE_H_
