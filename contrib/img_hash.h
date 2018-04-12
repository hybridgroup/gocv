#ifndef _OPENCV3_IMG_HASH_H_
#define _OPENCV3_IMG_HASH_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/img_hash.hpp>
extern "C" {
#endif

#include "../core.h"

void pHashCompute(Mat inputArr, Mat outputArr);
double pHashCompare(Mat a, Mat b);
void averageHashCompute(Mat inputArr, Mat outputArr);
double averageHashCompare(Mat a, Mat b);
void blockMeanHashCompute(Mat inputArr, Mat outputArr, int mode);
double blockMeanHashCompare(Mat a, Mat b, int mode);
void colorMomentHashCompute(Mat inputArr, Mat outputArr);
double colorMomentHashCompare(Mat a, Mat b);
void marrHildrethHashCompute(Mat inputArr, Mat outputArr, float alpha, float scale);
double marrHildrethHashCompare(Mat a, Mat b, float alpha, float scale);
void radialVarianceHashCompute(Mat inputArr, Mat outputArr, double sigma, int numOfAngleLine);
double radialVarianceHashCompare(Mat a, Mat b, double sigma, int numOfAngleLine);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMG_HASH_H_
