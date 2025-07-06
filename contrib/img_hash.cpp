//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_contrib_img_hash)

#include "img_hash.h"

void pHashCompute(Mat inputArr, Mat outputArr) {
    try {
        cv::img_hash::pHash(*inputArr, *outputArr);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
double pHashCompare(Mat a, Mat b) {
    try {
        return cv::img_hash::PHash::create()->compare(*a, *b);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void averageHashCompute(Mat inputArr, Mat outputArr) {
    try {
        cv::img_hash::averageHash(*inputArr, *outputArr);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double averageHashCompare(Mat a, Mat b) {
    try {
        return cv::img_hash::AverageHash::create()->compare(*a, *b);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void blockMeanHashCompute(Mat inputArr, Mat outputArr, int mode) {
    try {
        cv::img_hash::blockMeanHash(*inputArr, *outputArr, mode);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double blockMeanHashCompare(Mat a, Mat b, int mode) {
    try {
        return cv::img_hash::BlockMeanHash::create(mode)->compare(*a, *b);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void colorMomentHashCompute(Mat inputArr, Mat outputArr) {
    try {
        cv::img_hash::colorMomentHash(*inputArr, *outputArr);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }    
}

double colorMomentHashCompare(Mat a, Mat b) {
    try {
        return cv::img_hash::ColorMomentHash::create()->compare(*a, *b);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void marrHildrethHashCompute(Mat inputArr, Mat outputArr, float alpha, float scale) {
    try {
        cv::img_hash::marrHildrethHash(*inputArr, *outputArr, alpha, scale);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double marrHildrethHashCompare(Mat a, Mat b, float alpha, float scale) {
    try {
        return cv::img_hash::MarrHildrethHash::create(alpha, scale)->compare(*a, *b);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void radialVarianceHashCompute(Mat inputArr, Mat outputArr, double sigma, int numOfAngleLine) {
    try {
        cv::img_hash::radialVarianceHash(*inputArr, *outputArr, sigma, numOfAngleLine);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double radialVarianceHashCompare(Mat a, Mat b, double sigma, int numOfAngleLine) {
    try {
        return cv::img_hash::RadialVarianceHash::create(sigma, numOfAngleLine)->compare(*a, *b);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}
