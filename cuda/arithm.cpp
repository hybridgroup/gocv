//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_cuda_arithm)

#include "../core.h"
#include "arithm.h"
#include <string.h>

OpenCVResult GpuAbs(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::abs(*src, *dst);
        } else {
            cv::cuda::abs(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuAbsDiff(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::absdiff(*src1, *src2, *dst);
        } else {
            cv::cuda::absdiff(*src1, *src2, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuAdd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::add(*src1, *src2, *dst);
        } else {
            cv::cuda::add(*src1, *src2, *dst, cv::noArray(), -1, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuBitwiseAnd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::bitwise_and(*src1, *src2, *dst);
        } else {
            cv::cuda::bitwise_and(*src1, *src2, *dst, cv::noArray(), *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuBitwiseNot(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::bitwise_not(*src, *dst);
        } else {
            cv::cuda::bitwise_not(*src, *dst, cv::noArray(), *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuBitwiseOr(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::bitwise_or(*src1, *src2, *dst);
        } else {
            cv::cuda::bitwise_or(*src1, *src2, *dst, cv::noArray(), *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuBitwiseXor(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::bitwise_xor(*src1, *src2, *dst);
        } else {
            cv::cuda::bitwise_xor(*src1, *src2, *dst, cv::noArray(), *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuDivide(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::divide(*src1, *src2, *dst);
        } else {
            cv::cuda::divide(*src1, *src2, *dst, 1, -1, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuExp(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::exp(*src, *dst);
        } else {
            cv::cuda::exp(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuLog(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::log(*src, *dst);
        } else {
            cv::cuda::log(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuMax(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::max(*src1, *src2, *dst);
        } else {
            cv::cuda::max(*src1, *src2, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuMin(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::min(*src1, *src2, *dst);
        } else {
            cv::cuda::min(*src1, *src2, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuMultiply(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::multiply(*src1, *src2, *dst);
        } else {
            cv::cuda::multiply(*src1, *src2, *dst, 1, -1, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuSqr(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::sqr(*src, *dst);
        } else {
            cv::cuda::sqr(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuSqrt(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::sqrt(*src, *dst);
        } else {
            cv::cuda::sqrt(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuSubtract(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::subtract(*src1, *src2, *dst);
        } else {
            cv::cuda::subtract(*src1, *src2, *dst, cv::noArray(), -1, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::threshold(*src, *dst, thresh, maxval, typ);
        } else {
            cv::cuda::threshold(*src, *dst, thresh, maxval, typ, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuFlip(GpuMat src, GpuMat dst, int flipCode, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::flip(*src, *dst, flipCode);
        } else {
            cv::cuda::flip(*src, *dst, flipCode, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuMerge(struct GpuMats mats, GpuMat dst, Stream s) {
    try {
        std::vector<cv::cuda::GpuMat> images;

        for (int i = 0; i < mats.length; ++i) {
            images.push_back(*mats.mats[i]);
        }
    
        if (s == NULL) {
            cv::cuda::merge(images, *dst);
        } else {
            cv::cuda::merge(images, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuTranspose(GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::transpose(*src, *dst);
        } else {
            cv::cuda::transpose(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuAddWeighted(GpuMat src1, double alpha, GpuMat src2, double beta, double gamma, GpuMat dst, int dType, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::addWeighted(*src1, alpha, *src2, beta, gamma, *dst, dType);
        } else {    
            cv::cuda::addWeighted(*src1, alpha, *src2, beta, gamma, *dst, dType, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuCopyMakeBorder(GpuMat src, GpuMat dst, int top, int bottom, int left, int right, int borderType, Scalar value, Stream s) {
    try {
        cv::Scalar cValue = cv::Scalar(value.val1, value.val2, value.val3, value.val4);

        if (s == NULL) {
            cv::cuda::copyMakeBorder(*src, *dst, top, bottom, left, right, borderType, cValue);
        } else {
            cv::cuda::copyMakeBorder(*src, *dst, top, bottom, left, right, borderType, cValue, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

LookUpTable Cuda_Create_LookUpTable(GpuMat lut){
    try {
        return new cv::Ptr<cv::cuda::LookUpTable>(cv::cuda::createLookUpTable(*lut));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void Cuda_LookUpTable_Close(LookUpTable lt) {
    delete lt;
}

bool Cuda_LookUpTable_Empty(LookUpTable lut) {
    try {
        return lut->empty();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return true;
    }
}

OpenCVResult Cuda_LookUpTable_Transform(LookUpTable lt, GpuMat src, GpuMat dst, Stream s) {
    try {
        cv::Ptr< cv::cuda::LookUpTable> p = cv::Ptr< cv::cuda::LookUpTable>(*lt);

        if(s == NULL) {
            p->transform(*src, *dst);
        } else {
            p->transform(*src, *dst, *s);
        }    
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_Split(GpuMat src, GpuMats dst, Stream s) {
    try {
        std::vector< cv::cuda::GpuMat > dstv;

        for(int i = 0; i < dst.length; i++) {
            dstv.push_back(*(dst.mats[i]));
        }
    
        if(s == NULL){
            cv::cuda::split(*src, dstv);
        } else {
            cv::cuda::split(*src, dstv, *s);
        }    
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuCalcNorm(GpuMat src, GpuMat dst, int typ, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::calcNorm(*src, *dst, typ);
        } else {
            cv::cuda::calcNorm(*src, *dst, typ, cv::noArray(), *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuCalcNormDiff(GpuMat src1, GpuMat src2, GpuMat dst, int typ, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::calcNormDiff(*src1, *src2, *dst, typ);
        } else {
            cv::cuda::calcNormDiff(*src1, *src2, *dst, typ, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

double GpuNorm(GpuMat src1, GpuMat src2, int typ) {
    try {
        return cv::cuda::norm(*src1, *src2, typ);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

OpenCVResult GpuCompare(GpuMat src1, GpuMat src2, GpuMat dst, int typ, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::compare(*src1, *src2, *dst, typ);
        } else {
            cv::cuda::compare(*src1, *src2, *dst, typ, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuLShift(GpuMat src, Scalar shift, GpuMat dst, Stream s) {
    try {
        cv::Scalar cValue = cv::Scalar(shift.val1, shift.val2, shift.val3, shift.val4);

        if (s == NULL) {
            cv::cuda::lshift(*src, cValue, *dst);
        } else {
            cv::cuda::lshift(*src, cValue, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuRShift(GpuMat src, Scalar shift, GpuMat dst, Stream s) {
    try {
        cv::Scalar cValue = cv::Scalar(shift.val1, shift.val2, shift.val3, shift.val4);

        if (s == NULL) {
            cv::cuda::rshift(*src, cValue, *dst);
        } else {
            cv::cuda::rshift(*src, cValue, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
