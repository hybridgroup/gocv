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

void GpuAbsDiff(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::absdiff(*src1, *src2, *dst);
        return;
    }
    cv::cuda::absdiff(*src1, *src2, *dst, *s);
}

void GpuAdd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::add(*src1, *src2, *dst);
        return;
    }
    cv::cuda::add(*src1, *src2, *dst, cv::noArray(), -1, *s);
}

void GpuBitwiseAnd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_and(*src1, *src2, *dst);
        return;
    }
    cv::cuda::bitwise_and(*src1, *src2, *dst, cv::noArray(), *s);
}

void GpuBitwiseNot(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_not(*src, *dst);
        return;
    }
    cv::cuda::bitwise_not(*src, *dst, cv::noArray(), *s);
}

void GpuBitwiseOr(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_or(*src1, *src2, *dst);
        return;
    }
    cv::cuda::bitwise_or(*src1, *src2, *dst, cv::noArray(), *s);
}

void GpuBitwiseXor(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_xor(*src1, *src2, *dst);
        return;
    }
    cv::cuda::bitwise_xor(*src1, *src2, *dst, cv::noArray(), *s);
}

void GpuDivide(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::divide(*src1, *src2, *dst);
        return;
    }
    cv::cuda::divide(*src1, *src2, *dst, 1, -1, *s);
}

void GpuExp(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::exp(*src, *dst);
        return;
    }
    cv::cuda::exp(*src, *dst, *s);
}

void GpuLog(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::log(*src, *dst);
        return;
    }
    cv::cuda::log(*src, *dst, *s);
}

void GpuMax(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::max(*src1, *src2, *dst);
        return;
    }
    cv::cuda::max(*src1, *src2, *dst, *s);
}

void GpuMin(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::min(*src1, *src2, *dst);
        return;
    }
    cv::cuda::min(*src1, *src2, *dst, *s);
}

void GpuMultiply(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::multiply(*src1, *src2, *dst);
        return;
    }
    cv::cuda::multiply(*src1, *src2, *dst, 1, -1, *s);
}

void GpuSqr(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::sqr(*src, *dst);
        return;
    }
    cv::cuda::sqr(*src, *dst, *s);
}

void GpuSqrt(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::sqrt(*src, *dst);
        return;
    }
    cv::cuda::sqrt(*src, *dst, *s);
}

void GpuSubtract(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::subtract(*src1, *src2, *dst);
        return;
    }
    cv::cuda::subtract(*src1, *src2, *dst, cv::noArray(), -1, *s);
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

void GpuMerge(struct GpuMats mats, GpuMat dst, Stream s) {
    std::vector<cv::cuda::GpuMat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }

    if (s == NULL) {
        cv::cuda::merge(images, *dst);
        return;
    }
    cv::cuda::merge(images, *dst, *s);
}

void GpuTranspose(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::transpose(*src, *dst);
        return;
    }
    cv::cuda::transpose(*src, *dst, *s);
}

void GpuAddWeighted(GpuMat src1, double alpha, GpuMat src2, double beta, double gamma, GpuMat dst, int dType, Stream s) {
    if (s == NULL) {
        cv::cuda::addWeighted(*src1, alpha, *src2, beta, gamma, *dst, dType);
        return;
    }

    cv::cuda::addWeighted(*src1, alpha, *src2, beta, gamma, *dst, dType, *s);
}

void GpuCopyMakeBorder(GpuMat src, GpuMat dst, int top, int bottom, int left, int right, int borderType, Scalar value, Stream s) {
    cv::Scalar cValue = cv::Scalar(value.val1, value.val2, value.val3, value.val4);

    if (s == NULL) {
        cv::cuda::copyMakeBorder(*src, *dst, top, bottom, left, right, borderType, cValue);
        return;
    }
    cv::cuda::copyMakeBorder(*src, *dst, top, bottom, left, right, borderType, cValue, *s);
}
