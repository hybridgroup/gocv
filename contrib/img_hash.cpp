#include "img_hash.h"

void pHashCompute(Mat inputArr, Mat outputArr) {
    cv::img_hash::pHash(*inputArr, *outputArr);
}
double pHashCompare(Mat a, Mat b) {
    return cv::img_hash::PHash::create()->compare(*a, *b);
}

void averageHashCompute(Mat inputArr, Mat outputArr) {
    cv::img_hash::averageHash(*inputArr, *outputArr);
}
double averageHashCompare(Mat a, Mat b) {
    return cv::img_hash::AverageHash::create()->compare(*a, *b);
}

void blockMeanHashCompute(Mat inputArr, Mat outputArr, int mode) {
    cv::img_hash::blockMeanHash(*inputArr, *outputArr, mode);
}
double blockMeanHashCompare(Mat a, Mat b, int mode) {
    return cv::img_hash::BlockMeanHash::create(mode)->compare(*a, *b);
}

void colorMomentHashCompute(Mat inputArr, Mat outputArr) {
    cv::img_hash::colorMomentHash(*inputArr, *outputArr);
}
double colorMomentHashCompare(Mat a, Mat b) {
    return cv::img_hash::ColorMomentHash::create()->compare(*a, *b);
}

void marrHildrethHashCompute(Mat inputArr, Mat outputArr, float alpha, float scale) {
    cv::img_hash::marrHildrethHash(*inputArr, *outputArr, alpha, scale);
}
double marrHildrethHashCompare(Mat a, Mat b, float alpha, float scale) {
    return cv::img_hash::MarrHildrethHash::create(alpha, scale)->compare(*a, *b);
}

void radialVarianceHashCompute(Mat inputArr, Mat outputArr, double sigma, int numOfAngleLine) {
    cv::img_hash::radialVarianceHash(*inputArr, *outputArr, sigma, numOfAngleLine);
}
double radialVarianceHashCompare(Mat a, Mat b, double sigma, int numOfAngleLine) {
    return cv::img_hash::RadialVarianceHash::create(sigma, numOfAngleLine)->compare(*a, *b);
}
