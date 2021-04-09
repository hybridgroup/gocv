#include "photo.h"

void ColorChange(Mat src, Mat mask, Mat dst, float red_mul, float green_mul, float blue_mul) {
    cv::colorChange(*src, *mask, *dst, red_mul, green_mul, blue_mul);
}

void IlluminationChange(Mat src, Mat mask, Mat dst, float alpha, float beta) {
    cv::illuminationChange(*src, *mask, *dst, alpha, beta);
}

void SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags) {
    cv::Point pt(p.x, p.y);
    cv::seamlessClone(*src, *dst, *mask, pt, *blend, flags);
}

void TextureFlattening(Mat src, Mat mask, Mat dst, float low_threshold, float high_threshold, int kernel_size) {
    cv::textureFlattening(*src, *mask, *dst, low_threshold, high_threshold, kernel_size);
}

void MergeMertensProcessCSE(struct Mats src, Mat dst,float contrast_weight, float saturation_weight , float exposure_weight ) {
    std::vector<cv::Mat> images;
    for (int i = 0; i < src.length; ++i) {
        images.push_back(*src.mats[i]);
    }   
    cv::createMergeMertens(contrast_weight , saturation_weight, exposure_weight)->process(images, *dst);
}

void MergeMertensProcess(struct Mats src, Mat dst) {
    std::vector<cv::Mat> images;
    for (int i = 0; i < src.length; ++i) {
        images.push_back(*src.mats[i]);
    }  
    cv::createMergeMertens()->process(images, *dst);
}
