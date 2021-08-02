#include "../core.h"
#include "arithm.h"
#include <string.h>

void GpuAbs(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::abs(*src, *dst);
        return;
    }
    cv::cuda::abs(*src, *dst, *s);
}

void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ, Stream s) {
    if (s == NULL) {
        cv::cuda::threshold(*src, *dst, thresh, maxval, typ);
        return;
    }

    cv::cuda::threshold(*src, *dst, thresh, maxval, typ, *s);
}

void GpuFlip(GpuMat src, GpuMat dst, int flipCode, Stream s) {
    if (s == NULL) {
        cv::cuda::flip(*src, *dst, flipCode);
        return;
    }
    cv::cuda::flip(*src, *dst, flipCode, *s);
}
