#include "photo.h"

void ColorChange(Mat src, Mat mask, Mat dst, float red_mul, float green_mul, float blue_mul) {
    try {
        cv::colorChange(*src, *mask, *dst, red_mul, green_mul, blue_mul);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void IlluminationChange(Mat src, Mat mask, Mat dst, float alpha, float beta) {
    try {
        cv::illuminationChange(*src, *mask, *dst, alpha, beta);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags) {
    try {
        cv::Point pt(p.x, p.y);
        cv::seamlessClone(*src, *dst, *mask, pt, *blend, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void TextureFlattening(Mat src, Mat mask, Mat dst, float low_threshold, float high_threshold, int kernel_size) {
    try {
        cv::textureFlattening(*src, *mask, *dst, low_threshold, high_threshold, kernel_size);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}


void FastNlMeansDenoisingColoredMulti(	struct Mats src, Mat dst, int imgToDenoiseIndex, int 	temporalWindowSize){
    try {
        std::vector<cv::Mat> images;
        for (int i = 0; i < src.length; ++i) {
            images.push_back(*src.mats[i]);
        }
        cv::fastNlMeansDenoisingColoredMulti( images, *dst, imgToDenoiseIndex, 	temporalWindowSize );
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FastNlMeansDenoisingColoredMultiWithParams( struct Mats src, Mat dst, int imgToDenoiseIndex, int 	temporalWindowSize, float 	h, float 	hColor, int 	templateWindowSize, int 	searchWindowSize ){
    try {
        std::vector<cv::Mat> images;
        for (int i = 0; i < src.length; ++i) {
            images.push_back(*src.mats[i]);
        }
        cv::fastNlMeansDenoisingColoredMulti( images, *dst, imgToDenoiseIndex, 	temporalWindowSize, h, hColor, templateWindowSize, searchWindowSize );
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

MergeMertens MergeMertens_Create() {
    try {
        return new cv::Ptr<cv::MergeMertens>(cv::createMergeMertens());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

MergeMertens MergeMertens_CreateWithParams(float contrast_weight,
                                           float saturation_weight,
                                           float exposure_weight) {
    try {
        return new cv::Ptr<cv::MergeMertens>(cv::createMergeMertens(contrast_weight, saturation_weight, exposure_weight));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void MergeMertens_Close(MergeMertens b) {
    delete b;
}

void MergeMertens_Process(MergeMertens b, struct Mats src, Mat dst) {
    try {
        std::vector<cv::Mat> images;
        for (int i = 0; i < src.length; ++i) {
            images.push_back(*src.mats[i]);
        }
        (*b)->process(images, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

AlignMTB AlignMTB_Create() {
    try {
        return new cv::Ptr<cv::AlignMTB>(cv::createAlignMTB(6,4,false));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

AlignMTB AlignMTB_CreateWithParams(int max_bits, int exclude_range, bool cut) {
    try {
        return new cv::Ptr<cv::AlignMTB>(cv::createAlignMTB(max_bits, exclude_range, cut));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void AlignMTB_Close(AlignMTB b) { delete b; }

void AlignMTB_Process(AlignMTB b, struct Mats src, struct Mats *dst) {
    try {
        std::vector<cv::Mat> srcMats;
        for (int i = 0; i < src.length; ++i) {
            srcMats.push_back(*src.mats[i]);
        }
      
        std::vector<cv::Mat> dstMats;
        (*b)->process(srcMats, dstMats);
      
        dst->mats = new Mat[dstMats.size()];
        for (size_t i = 0; i < dstMats.size() ; ++i) {
            dst->mats[i] = new cv::Mat( dstMats[i] );
        }
        dst->length = (int)dstMats.size();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FastNlMeansDenoising(Mat src, Mat dst) {
    try {
        cv::fastNlMeansDenoising(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FastNlMeansDenoisingWithParams(Mat src, Mat dst, float h, int templateWindowSize, int searchWindowSize) {
    try {
        cv::fastNlMeansDenoising(*src, *dst, h, templateWindowSize, searchWindowSize);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FastNlMeansDenoisingColored(Mat src, Mat dst) {
    try {
        cv::fastNlMeansDenoisingColored(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FastNlMeansDenoisingColoredWithParams(Mat src, Mat dst, float h, float hColor, int templateWindowSize, int searchWindowSize) {
    try {
        cv::fastNlMeansDenoisingColored(*src, *dst, h, hColor, templateWindowSize, searchWindowSize);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void EdgePreservingFilter(Mat src, Mat dst, int filter, float sigma_s, float sigma_r) {
    try {
        cv::edgePreservingFilter(*src, *dst, filter, sigma_s, sigma_r);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void DetailEnhance(Mat src, Mat dst, float sigma_s, float sigma_r) {
    try {
        cv::detailEnhance(*src, *dst, sigma_s, sigma_r);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void PencilSketch(Mat src, Mat dst1, Mat dst2, float sigma_s, float sigma_r, float shade_factor) {
    try {
        cv::pencilSketch(*src, *dst1, *dst2, sigma_s, sigma_r, shade_factor);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Stylization(Mat src, Mat dst, float sigma_s, float sigma_r) {
    try {
        cv::stylization(*src, *dst, sigma_s, sigma_r);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void PhotoInpaint(Mat src, Mat mask, Mat dst, float inpaint_radius, int algorithm_type) {
    try {
        cv::inpaint(*src, *mask, *dst, inpaint_radius, algorithm_type);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
