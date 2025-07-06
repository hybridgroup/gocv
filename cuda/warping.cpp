//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_cuda_warping)

#include "warping.h"

OpenCVResult CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp, Stream s) {
    try {
        cv::Size sz(dsize.width, dsize.height);

        if (s == NULL) {
            cv::cuda::resize(*src, *dst, sz, fx, fy, interp);
        } else {
            cv::cuda::resize(*src, *dst, sz, fx, fy, interp, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaPyrDown(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::pyrDown(*src, *dst);
        } else {
            cv::cuda::pyrDown(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaPyrUp(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::pyrUp(*src, *dst);
        } else {
            cv::cuda::pyrUp(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaBuildWarpAffineMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        if (s == NULL) {
            cv::cuda::buildWarpAffineMaps(*M, inverse, sz, *xmap, *ymap);
        } else {
            cv::cuda::buildWarpAffineMaps(*M, inverse, sz, *xmap, *ymap, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaBuildWarpPerspectiveMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        if (s == NULL) {
            cv::cuda::buildWarpPerspectiveMaps(*M, inverse, sz, *xmap, *ymap);
        } else {
            cv::cuda::buildWarpPerspectiveMaps(*M, inverse, sz, *xmap, *ymap, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaRemap(GpuMat src, GpuMat dst, GpuMat xmap, GpuMat ymap, int interp, int borderMode, Scalar borderValue, Stream s) {
    try {
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        if (s == NULL) {
            cv::cuda::remap(*src, *dst, *xmap, *ymap, interp, borderMode, c);
        } else {
            cv::cuda::remap(*src, *dst, *xmap, *ymap, interp, borderMode, c, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaRotate(GpuMat src, GpuMat dst, Size dsize, double angle, double xShift, double yShift, int interp, Stream s) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        if (s == NULL) {
            cv::cuda::rotate(*src, *dst, sz, angle, xShift, yShift, interp);
        } else {
            cv::cuda::rotate(*src, *dst, sz, angle, xShift, yShift, interp, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaWarpAffine(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s) {
    try {
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::Size sz(dsize.width, dsize.height);
    
        if (s == NULL) {
            cv::cuda::warpAffine(*src, *dst, *M, sz, flags, borderMode, c);
        } else {
            cv::cuda::warpAffine(*src, *dst, *M, sz, flags, borderMode, c, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CudaWarpPerspective(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s) {
    try {
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::Size sz(dsize.width, dsize.height);
        if (s == NULL) {
            cv::cuda::warpPerspective(*src, *dst, *M, sz, flags, borderMode, c);
        } else {
            cv::cuda::warpPerspective(*src, *dst, *M, sz, flags, borderMode, c, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
