#include "photo.h"

void FastNlMeansDenoising(Mat src, Mat dst) {
    cv::fastNlMeansDenoising(*src, *dst);
}

void FastNlMeansDenoisingWithParams(Mat src, Mat dst, float h, int templateWindowSize, int searchWindowSize) {
    cv::fastNlMeansDenoising(*src, *dst, h, templateWindowSize, searchWindowSize);
}

void FastNlMeansDenoisingColored(Mat src, Mat dst) {
    cv::fastNlMeansDenoisingColored(*src, *dst);
}

void FastNlMeansDenoisingColoredWithParams(Mat src, Mat dst, float h, float hColor, int templateWindowSize, int searchWindowSize) {
    cv::fastNlMeansDenoisingColored(*src, *dst, h, hColor, templateWindowSize, searchWindowSize);
}
