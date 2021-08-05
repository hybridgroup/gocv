#include "aruco.h"

cv::Ptr<cv::aruco::DetectorParameters> ArucoConvertCParamsToCPPParams(ArucoDetectorParameters params)
{
    cv::aruco::DetectorParameters converted;

    converted.adaptiveThreshWinSizeMin = params.adaptiveThreshWinSizeMin;
    converted.adaptiveThreshWinSizeMax = params.adaptiveThreshWinSizeMax;
    converted.adaptiveThreshWinSizeStep = params.adaptiveThreshWinSizeStep;
    converted.adaptiveThreshConstant = params.adaptiveThreshConstant;
    converted.minMarkerPerimeterRate = params.minMarkerPerimeterRate;
    converted.maxMarkerPerimeterRate = params.maxMarkerPerimeterRate;
    converted.polygonalApproxAccuracyRate = params.polygonalApproxAccuracyRate;
    converted.minCornerDistanceRate = params.minCornerDistanceRate;
    converted.minDistanceToBorder = params.minDistanceToBorder;
    converted.minMarkerDistanceRate = params.minMarkerDistanceRate;
    converted.cornerRefinementMethod = params.cornerRefinementMethod;
    converted.cornerRefinementWinSize = params.cornerRefinementWinSize;
    converted.cornerRefinementMaxIterations = params.cornerRefinementMaxIterations;
    converted.cornerRefinementMinAccuracy = params.cornerRefinementMinAccuracy;
    converted.markerBorderBits = params.markerBorderBits;
    converted.perspectiveRemovePixelPerCell = params.perspectiveRemovePixelPerCell;
    converted.perspectiveRemoveIgnoredMarginPerCell = params.perspectiveRemoveIgnoredMarginPerCell;
    converted.maxErroneousBitsInBorderRate = params.maxErroneousBitsInBorderRate;
    converted.minOtsuStdDev = params.minOtsuStdDev;
    converted.errorCorrectionRate = params.errorCorrectionRate;

    converted.aprilTagQuadDecimate = params.aprilTagQuadDecimate;
    converted.aprilTagQuadSigma = params.aprilTagQuadSigma;

    converted.aprilTagMinClusterPixels = params.aprilTagMinClusterPixels;
    converted.aprilTagMaxNmaxima = params.aprilTagMaxNmaxima;
    converted.aprilTagCriticalRad = params.aprilTagCriticalRad;
    converted.aprilTagMaxLineFitMse = params.aprilTagMaxLineFitMse;
    converted.aprilTagMinWhiteBlackDiff = params.aprilTagMinWhiteBlackDiff;
    converted.aprilTagDeglitch = params.aprilTagDeglitch;

    converted.detectInvertedMarker = params.detectInvertedMarker;

    cv::Ptr<cv::aruco::DetectorParameters> ptr = cv::makePtr<cv::aruco::DetectorParameters>(converted);

    return ptr;
}

ArucoDetectorParameters ArucoConvertCPPParamsToCParams(cv::aruco::DetectorParameters params)
{
    ArucoDetectorParameters converted;

    converted.adaptiveThreshWinSizeMin = params.adaptiveThreshWinSizeMin;
    converted.adaptiveThreshWinSizeMax = params.adaptiveThreshWinSizeMax;
    converted.adaptiveThreshWinSizeStep = params.adaptiveThreshWinSizeStep;
    converted.adaptiveThreshConstant = params.adaptiveThreshConstant;
    converted.minMarkerPerimeterRate = params.minMarkerPerimeterRate;
    converted.maxMarkerPerimeterRate = params.maxMarkerPerimeterRate;
    converted.polygonalApproxAccuracyRate = params.polygonalApproxAccuracyRate;
    converted.minCornerDistanceRate = params.minCornerDistanceRate;
    converted.minDistanceToBorder = params.minDistanceToBorder;
    converted.minMarkerDistanceRate = params.minMarkerDistanceRate;
    converted.cornerRefinementMethod = params.cornerRefinementMethod;
    converted.cornerRefinementWinSize = params.cornerRefinementWinSize;
    converted.cornerRefinementMaxIterations = params.cornerRefinementMaxIterations;
    converted.cornerRefinementMinAccuracy = params.cornerRefinementMinAccuracy;
    converted.markerBorderBits = params.markerBorderBits;
    converted.perspectiveRemovePixelPerCell = params.perspectiveRemovePixelPerCell;
    converted.perspectiveRemoveIgnoredMarginPerCell = params.perspectiveRemoveIgnoredMarginPerCell;
    converted.maxErroneousBitsInBorderRate = params.maxErroneousBitsInBorderRate;
    converted.minOtsuStdDev = params.minOtsuStdDev;
    converted.errorCorrectionRate = params.errorCorrectionRate;

    converted.aprilTagQuadDecimate = params.aprilTagQuadDecimate;
    converted.aprilTagQuadSigma = params.aprilTagQuadSigma;

    converted.aprilTagMinClusterPixels = params.aprilTagMinClusterPixels;
    converted.aprilTagMaxNmaxima = params.aprilTagMaxNmaxima;
    converted.aprilTagCriticalRad = params.aprilTagCriticalRad;
    converted.aprilTagMaxLineFitMse = params.aprilTagMaxLineFitMse;
    converted.aprilTagMinWhiteBlackDiff = params.aprilTagMinWhiteBlackDiff;
    converted.aprilTagDeglitch = params.aprilTagDeglitch;

    converted.detectInvertedMarker = params.detectInvertedMarker;

    return converted;
}

ArucoDetectorParameters ArucoDetectorParameters_Create()
{
    cv::Ptr<cv::aruco::DetectorParameters> p = cv::aruco::DetectorParameters::create();

    ArucoDetectorParameters cparam = ArucoConvertCPPParamsToCParams(*p);
    p.release();
    return cparam;
}

void detectMarkers(Mat inputArr, ArucoDictionary dictionary, Points2fVector markerCorners, IntVector *markerIds, ArucoDetectorParameters params, Points2fVector rejectedCandidates)
{
    std::vector<int> _markerIds;
    cv::Ptr<cv::aruco::DetectorParameters> parameters = ArucoConvertCParamsToCPPParams(params);
    cv::aruco::detectMarkers(*inputArr, *dictionary, *markerCorners, _markerIds, parameters, *rejectedCandidates);

    int *ids = new int[_markerIds.size()];

    for (size_t i = 0; i < _markerIds.size(); ++i)
    {
        ids[i] = _markerIds[i];
    }

    markerIds->length = _markerIds.size();
    markerIds->val = ids;
}

void detectMarkersWithDictId(Mat inputArr, int dictionaryId, Points2fVector markerCorners, IntVector *markerIds, ArucoDetectorParameters params, Points2fVector rejectedCandidates)
{
    std::vector<int> _markerIds;
    cv::Ptr<cv::aruco::DetectorParameters> parameters = ArucoConvertCParamsToCPPParams(params);
    cv::Ptr<cv::aruco::Dictionary> dictionary = cv::aruco::getPredefinedDictionary(dictionaryId);
    cv::aruco::detectMarkers(*inputArr, dictionary, *markerCorners, _markerIds, parameters, *rejectedCandidates);

    int *ids = new int[_markerIds.size()];

    for (size_t i = 0; i < _markerIds.size(); ++i)
    {
        ids[i] = _markerIds[i];
    }

    markerIds->length = _markerIds.size();
    markerIds->val = ids;
}

void drawDetectedMarkers(Mat image, Points2fVector markerCorners, IntVector markerIds, Scalar borderColor)
{
    std::vector<int> _markerIds;
    for (int i = 0, *v = markerIds.val; i < markerIds.length; ++v, ++i)
    {
        _markerIds.push_back(*v);
    }
    cv::Scalar _borderColor = cv::Scalar(borderColor.val1, borderColor.val2, borderColor.val3);
    // cv::Scalar borderColor = cv::Scalar(0, 255, 0);
    cv::aruco::drawDetectedMarkers(*image, *markerCorners, _markerIds, _borderColor);
}

void drawMarker(int dictionaryId, int id, int sidePixels, Mat img, int borderBits)
{
    cv::Ptr<cv::aruco::Dictionary> dict = cv::aruco::getPredefinedDictionary(dictionaryId);
    cv::aruco::drawMarker(dict, id, sidePixels, *img, borderBits);
}

ArucoDictionary getPredefinedDictionary(int dictionaryId)
{
    return new cv::Ptr<cv::aruco::Dictionary>(cv::aruco::getPredefinedDictionary(dictionaryId));
}