#ifndef _OPENCV3_ARUCO_H_
#define _OPENCV3_ARUCO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/aruco.hpp>

extern "C"
{
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::aruco::Dictionary> *ArucoDictionary;
#else
typedef void *ArucoDictionary;
#endif

// Wrapper for aruco::DetectorParameters
typedef struct ArucoDetectorParameters
{
    int adaptiveThreshWinSizeMin;
    int adaptiveThreshWinSizeMax;
    int adaptiveThreshWinSizeStep;
    double adaptiveThreshConstant;
    double minMarkerPerimeterRate;
    double maxMarkerPerimeterRate;
    double polygonalApproxAccuracyRate;
    double minCornerDistanceRate;
    int minDistanceToBorder;
    double minMarkerDistanceRate;
    int cornerRefinementMethod;
    int cornerRefinementWinSize;
    int cornerRefinementMaxIterations;
    double cornerRefinementMinAccuracy;
    int markerBorderBits;
    int perspectiveRemovePixelPerCell;
    double perspectiveRemoveIgnoredMarginPerCell;
    double maxErroneousBitsInBorderRate;
    double minOtsuStdDev;
    double errorCorrectionRate;

    float aprilTagQuadDecimate;
    float aprilTagQuadSigma;

    int aprilTagMinClusterPixels;
    int aprilTagMaxNmaxima;
    float aprilTagCriticalRad;
    float aprilTagMaxLineFitMse;
    int aprilTagMinWhiteBlackDiff;
    int aprilTagDeglitch;

    bool detectInvertedMarker;
} ArucoDetectorParameters;

ArucoDetectorParameters ArucoDetectorParameters_Create();

ArucoDictionary getPredefinedDictionary(int dictionaryId);

void detectMarkers(Mat inputArr, ArucoDictionary dictionary, Points2fVector markerCorners, IntVector *markerIds, ArucoDetectorParameters params, Points2fVector rejectedCandidates);

void detectMarkersWithDictId(Mat inputArr, int dictionaryId, Points2fVector markerCorners, IntVector *markerIds, ArucoDetectorParameters params, Points2fVector rejectedCandidates);

void drawDetectedMarkers(Mat image, Points2fVector markerCorners, IntVector markerIds, Scalar borderColor);

void drawMarker(int dictionaryId, int id, int sidePixels, Mat img, int borderBits);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_ARUCO_H_
