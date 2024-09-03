#ifndef _OPENCV3_FACE_H_
#define _OPENCV3_FACE_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/face.hpp>

extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::face::FaceRecognizer>* FaceRecognizer;
typedef cv::Ptr<cv::face::BasicFaceRecognizer>* BasicFaceRecognizer;
typedef cv::Ptr<cv::face::LBPHFaceRecognizer>* LBPHFaceRecognizer;
typedef cv::Ptr<cv::face::EigenFaceRecognizer>* EigenFaceRecognizer;
typedef cv::Ptr<cv::face::FisherFaceRecognizer>* FisherFaceRecognizer;
#else
typedef void* FaceRecognizer;
typedef void* BasicFaceRecognizer;
typedef void* LBPHFaceRecognizer;
typedef void* EigenFaceRecognizer;
typedef void* FisherFaceRecognizer;
#endif

struct PredictResponse {
    int label;
    double confidence;
};

bool FaceRecognizer_Empty(FaceRecognizer fr);
void FaceRecognizer_Train(FaceRecognizer fr, Mats images, IntVector labels);
void FaceRecognizer_Update(FaceRecognizer fr, Mats images, IntVector labels);
int FaceRecognizer_Predict(FaceRecognizer fr, Mat sample);
struct PredictResponse FaceRecognizer_PredictExtended(FaceRecognizer fr, Mat sample);
double FaceRecognizer_GetThreshold(FaceRecognizer fr);
void FaceRecognizer_SetThreshold(FaceRecognizer fr, double threshold);
void FaceRecognizer_SaveFile(FaceRecognizer fr, const char*  filename);
void FaceRecognizer_LoadFile(FaceRecognizer fr, const char*  filename);


void BasicFaceRecognizer_Train(BasicFaceRecognizer fr, Mats images, IntVector labels);
void BasicFaceRecognizer_Update(BasicFaceRecognizer fr, Mats images, IntVector labels);
Mat BasicFaceRecognizer_getEigenValues(BasicFaceRecognizer fr);
Mat BasicFaceRecognizer_getEigenVectors(BasicFaceRecognizer fr);
Mat BasicFaceRecognizer_getLabels(BasicFaceRecognizer fr);
Mat BasicFaceRecognizer_getMean(BasicFaceRecognizer fr);
int BasicFaceRecognizer_getNumComponents(BasicFaceRecognizer fr);
Mats BasicFaceRecognizer_getProjections(BasicFaceRecognizer fr);
void BasicFaceRecognizer_setNumComponents(BasicFaceRecognizer fr, int val);	
void BasicFaceRecognizer_SaveFile(BasicFaceRecognizer fr, const char*  filename);
void BasicFaceRecognizer_LoadFile(BasicFaceRecognizer fr, const char*  filename);

LBPHFaceRecognizer CreateLBPHFaceRecognizer(void);
void LBPHFaceRecognizer_SetRadius(LBPHFaceRecognizer fr, int radius);
void LBPHFaceRecognizer_SetNeighbors(LBPHFaceRecognizer fr, int neighbors);
int LBPHFaceRecognizer_GetNeighbors(LBPHFaceRecognizer fr);
void LBPHFaceRecognizer_SetGridX(LBPHFaceRecognizer fr, int x);
void LBPHFaceRecognizer_SetGridY(LBPHFaceRecognizer fr, int y);
int LBPHFaceRecognizer_GetGridX(LBPHFaceRecognizer fr);
int LBPHFaceRecognizer_GetGridY(LBPHFaceRecognizer fr);
void LBPHFaceRecognizer_Close(LBPHFaceRecognizer fr);


FisherFaceRecognizer FisherFaceRecognizer_Create(void);
FisherFaceRecognizer FisherFaceRecognizer_CreateWithParams(int num_components, float threshold);
void FisherFaceRecognizer_Close(FisherFaceRecognizer fr);

EigenFaceRecognizer EigenFaceRecognizer_Create(void);
EigenFaceRecognizer EigenFaceRecognizer_CreateWithParams(int num_components, float threshold);
void EigenFaceRecognizer_Close(EigenFaceRecognizer fr);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_FACE_H_
