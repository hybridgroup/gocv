#ifndef _OPENCV3_XPHOTO_H_
#define _OPENCV3_XPHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xphoto/bm3d_image_denoising.hpp>
#include <opencv2/xphoto/white_balance.hpp>
#include <opencv2/xphoto/tonemap.hpp>
#include <opencv2/xphoto/inpainting.hpp>
#include <opencv2/xphoto/oilpainting.hpp>
extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::xphoto::GrayworldWB>* GrayworldWB;
typedef cv::Ptr<cv::xphoto::LearningBasedWB>* LearningBasedWB;
typedef cv::Ptr<cv::xphoto::SimpleWB>* SimpleWB;
typedef cv::Ptr<cv::xphoto::TonemapDurand>* TonemapDurand;
#else
typedef void* GrayworldWB;
typedef void* LearningBasedWB;
typedef void* SimpleWB;
typedef void* TonemapDurand;
#endif


// ----------------------- bm3d_image_denoising -----------------------

void Xphoto_ApplyChannelGains(Mat src, Mat dst, float gainB, float gainG, float gainR) ;

void Xphoto_Bm3dDenoising_Step(Mat src, Mat dststep1, Mat dststep2) ;
void Xphoto_Bm3dDenoising_Step_WithParams(
    Mat src, Mat dststep1, Mat dststep2,
    float h, int templateWindowSize,
    int searchWindowSize, int blockMatchingStep1,
    int blockMatchingStep2, int groupSize,
    int slidingStep, float beta,
    int normType, int step,
    int transformType
) ;

void Xphoto_Bm3dDenoising(Mat src, Mat dst) ;
void Xphoto_Bm3dDenoising_WithParams(
    Mat src, Mat dst,
    float h, int templateWindowSize,
    int searchWindowSize, int blockMatchingStep1,
    int blockMatchingStep2, int groupSize,
    int slidingStep, float beta,
    int normType, int step,
    int transformType
) ;


// ----------------------- GrayworldWB -----------------------

GrayworldWB GrayworldWB_Create();
void GrayworldWB_Close(GrayworldWB b);
void GrayworldWB_SetSaturationThreshold(GrayworldWB b, float saturationThreshold);
float GrayworldWB_GetSaturationThreshold(GrayworldWB b);
void GrayworldWB_BalanceWhite(GrayworldWB b, Mat src, Mat dst);

// ----------------------- LearningBasedWB -----------------------

LearningBasedWB LearningBasedWB_Create();
LearningBasedWB LearningBasedWB_CreateWithParams(const char* pathmodel);
void LearningBasedWB_Close(LearningBasedWB b);
void LearningBasedWB_ExtractSimpleFeatures(LearningBasedWB b, Mat src, Mat dst);
int LearningBasedWB_GetHistBinNum(LearningBasedWB b) ;
int LearningBasedWB_GetRangeMaxVal(LearningBasedWB b) ;
float LearningBasedWB_GetSaturationThreshold(LearningBasedWB b) ;
void LearningBasedWB_SetHistBinNum(LearningBasedWB b, int val);
void LearningBasedWB_SetRangeMaxVal(LearningBasedWB b, int val);
void LearningBasedWB_SetSaturationThreshold(LearningBasedWB b, float val);
void LearningBasedWB_BalanceWhite(LearningBasedWB b, Mat src, Mat dst);

// ----------------------- SimpleWB -----------------------

SimpleWB SimpleWB_Create();
void SimpleWB_Close(SimpleWB b);
float SimpleWB_GetInputMax(SimpleWB b);
float SimpleWB_GetInputMin(SimpleWB b);
float SimpleWB_GetOutputMax(SimpleWB b);
float SimpleWB_GetOutputMin(SimpleWB b);
float SimpleWB_GetP(SimpleWB b);
void SimpleWB_SetInputMax(SimpleWB b, float val);
void SimpleWB_SetInputMin(SimpleWB b, float val);
void SimpleWB_SetOutputMax(SimpleWB b, float val);
void SimpleWB_SetOutputMin(SimpleWB b, float val);
void SimpleWB_SetP(SimpleWB b, float val);
void SimpleWB_BalanceWhite(SimpleWB b, Mat src, Mat dst);


// -------------------- TonemapDurand --------------------

TonemapDurand TonemapDurand_Create();
TonemapDurand TonemapDurand_CreateWithParams(float gamma, float contrast, float saturation,
        float sigma_color, float sigma_space) ;
void TonemapDurand_Close(TonemapDurand b);

float TonemapDurand_GetContrast(TonemapDurand b);
float TonemapDurand_GetSaturation(TonemapDurand b);
float TonemapDurand_GetSigmaColor(TonemapDurand b);
float TonemapDurand_GetSigmaSpace(TonemapDurand b);
void TonemapDurand_SetContrast(TonemapDurand b, float contrast);
void TonemapDurand_SetSaturation(TonemapDurand b, float saturation);
void TonemapDurand_SetSigmaColor(TonemapDurand b, float sigma_color);
void TonemapDurand_SetSigmaSpace(TonemapDurand b, float sigma_space);

float TonemapDurand_GetGamma(TonemapDurand b);
void TonemapDurand_Process(TonemapDurand b, Mat src, Mat dst);
void TonemapDurand_SetGamma(TonemapDurand b, float gamma);


// ------------------------ Inpaint -----------------------

void Inpaint(Mat src, Mat mask, Mat dst, int algorithmType);
void OilPaintingWithParams(Mat src, Mat dst, int size, int dynRatio, int code);
void OilPainting(Mat src, Mat dst, int size, int dynRatio);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XPHOTO_H
