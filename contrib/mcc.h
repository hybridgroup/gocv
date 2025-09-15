#ifndef _OPENCV3_MCC_H_
#define _OPENCV3_MCC_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/mcc.hpp>
#include <opencv2/dnn.hpp>
extern "C" {
    #endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::mcc::CChecker* MccCChecker;
typedef cv::Ptr<cv::mcc::CCheckerDetector>* MccCCheckerDetector;
typedef cv::Ptr<cv::mcc::CCheckerDraw>* MccCCheckerDraw;
typedef cv::mcc::DetectorParameters* MccDetectorParameters;
typedef cv::dnn::Net* MccDnnNet;
typedef std::vector<cv::Rect>* MccRectVector;
typedef std::vector<cv::Ptr<cv::mcc::CChecker>>* MccCCheckerVector;
#else
typedef void *MccCChecker;
typedef void *MccCCheckerDetector;
typedef void *MccCCheckerDraw;
typedef void *MccDetectorParameters;
typedef void *MccDnnNet;
typedef void *MccRectVector;
typedef void *MccCCheckerVector;
#endif

MccDetectorParameters MccDetectorParameters_Create();
void MccDetectorParameters_Close(MccDetectorParameters mp);
void MccDetectorParameters_SetAdaptiveThreshWinSizeMin(MccDetectorParameters mp, int adaptiveThreshWinSizeMin);
int MccDetectorParameters_GetAdaptiveThreshWinSizeMin(MccDetectorParameters mp);
void MccDetectorParameters_SetAdaptiveThreshWinSizeMax(MccDetectorParameters mp, int adaptiveThreshWinSizeMax);
int MccDetectorParameters_GetAdaptiveThreshWinSizeMax(MccDetectorParameters mp);
void MccDetectorParameters_SetAdaptiveThreshWinSizeStep(MccDetectorParameters mp, int adaptiveThreshWinSizeStep);
int MccDetectorParameters_GetAdaptiveThreshWinSizeStep(MccDetectorParameters mp);
void MccDetectorParameters_SetBorderWidth(MccDetectorParameters mp, int borderWidth);
int MccDetectorParameters_GetBorderWidth(MccDetectorParameters mp);
void MccDetectorParameters_SetMinContourLengthAllowed(MccDetectorParameters mp, int minContourLengthAllowed);
int MccDetectorParameters_GetMinContourLengthAllowed(MccDetectorParameters mp);
void MccDetectorParameters_SetMinContourPointsAllowed(MccDetectorParameters mp, int minContourPointsAllowed);
int MccDetectorParameters_GetMinContourPointsAllowed(MccDetectorParameters mp);
void MccDetectorParameters_SetMinImageSize(MccDetectorParameters mp, int minImageSize);
int MccDetectorParameters_GetMinImageSize(MccDetectorParameters mp);
void MccDetectorParameters_SetMinInterCheckerDistance(MccDetectorParameters mp, int minInterCheckerDistance);
int MccDetectorParameters_GetMinInterCheckerDistance(MccDetectorParameters mp);
void MccDetectorParameters_SetMinInterContourDistance(MccDetectorParameters mp, int minInterContourDistance);
int MccDetectorParameters_GetMinInterContourDistance(MccDetectorParameters mp);

void MccDetectorParameters_SetAdaptiveThreshConstant(MccDetectorParameters mp, double adaptiveThreshConstant);
double MccDetectorParameters_GetAdaptiveThreshConstant(MccDetectorParameters mp);
void MccDetectorParameters_SetConfidenceThreshold(MccDetectorParameters mp, double confidenceThreshold);
double MccDetectorParameters_GetConfidenceThreshold(MccDetectorParameters mp);
void MccDetectorParameters_SetFindCandidatesApproxPolyDPEpsMultiplier(MccDetectorParameters mp, double findCandidatesApproxPolyDPEpsMultiplier);
double MccDetectorParameters_GetFindCandidatesApproxPolyDPEpsMultiplier(MccDetectorParameters mp);
void MccDetectorParameters_SetMinContourSolidity(MccDetectorParameters mp, double minContourSolidity);
double MccDetectorParameters_GetMinContourSolidity(MccDetectorParameters mp);
void MccDetectorParameters_SetMinContoursArea(MccDetectorParameters mp, double minContoursArea);
double MccDetectorParameters_GetMinContoursArea(MccDetectorParameters mp);
void MccDetectorParameters_SetMinContoursAreaRate(MccDetectorParameters mp, double minContoursAreaRate);
double MccDetectorParameters_GetMinContoursAreaRate(MccDetectorParameters mp);

void MccDetectorParameters_SetB0factor(MccDetectorParameters mp, float B0factor);
float MccDetectorParameters_GetB0factor(MccDetectorParameters mp);
void MccDetectorParameters_SetMaxError(MccDetectorParameters mp, float maxError);
float MccDetectorParameters_GetMaxError(MccDetectorParameters mp);
void MccDetectorParameters_SetMinGroupSize(MccDetectorParameters mp, unsigned minGroupSize);
unsigned MccDetectorParameters_GetMinGroupSize(MccDetectorParameters mp);

void MccCChecker_SetTarget(MccCChecker mc, int target);
int MccCChecker_GetTarget(MccCChecker mc);
void MccCChecker_SetBox(MccCChecker mc, Point2f* pts, int length);
Points2f MccCChecker_GetBox(MccCChecker mc);
void MccCChecker_SetChartsRGB(MccCChecker mc, Mat mat);
Mat MccCChecker_GetChartsRGB(MccCChecker mc);
void MccCChecker_SetChartsYCbCr(MccCChecker mc, Mat mat);
Mat MccCChecker_GetChartsYCbCr(MccCChecker mc);
void MccCChecker_SetCost(MccCChecker mc, float cost);
float MccCChecker_GetCost(MccCChecker mc);
void MccCChecker_SetCenter(MccCChecker mc, Point2f pt);
Point2f MccCChecker_GetCenter(MccCChecker mc);
Points2f MccCChecker_GetColorCharts(MccCChecker mc);


MccCCheckerDraw MccCCheckerDraw_Create(MccCChecker mc, double b, double g, double r, double a, int thickness);
void MccCCheckerDraw_Draw(MccCCheckerDraw md, Mat img);
void MccCCheckerDraw_Close(MccCCheckerDraw md);

MccCCheckerDetector MccCCheckerDetector_New();
void MccCCheckerDetector_Close(MccCCheckerDetector md);
MccCCheckerVector MccCCheckerDetector_GetListColorChecker(MccCCheckerDetector md);
MccCChecker MccCCheckerDetector_GetBestColorChecker(MccCCheckerDetector md);
bool MccCCheckerDetector_Process(MccCCheckerDetector md, Mat inputArr, int chartType, int nc, bool useNet);
bool MccCCheckerDetector_ProcessWithRegionsOfInterest(MccCCheckerDetector md, Mat inputArr, int chartType, const MccRectVector regionsOfInterest, int nc, bool useNet);
bool MccCCheckerDetector_SetNet(MccCCheckerDetector md, MccDnnNet net);

int MccCCheckerVector_Size(MccCCheckerVector mv);
MccCChecker MccCCheckerVector_At(MccCCheckerVector mv, int idx);
void MccCCheckerVector_Close(MccCCheckerVector mv);

MccRectVector MccRectVector_New();
void MccRectVector_PushBack(MccRectVector mv, int x, int y, int width, int height);
void MccRectVector_Close(MccRectVector mv);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_MCC_H_
