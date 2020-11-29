#include "../core.h"
#include "arithm.h"
#include <string.h>

void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ) {
    cv::cuda::threshold(*src, *dst, thresh, maxval, typ);
}
