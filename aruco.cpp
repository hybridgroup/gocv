#include "aruco.h"

ArucoDetector ArucoDetector_New() {
    return new cv::aruco::ArucoDetector();
}

ArucoDetector ArucoDetector_NewWithParams(ArucoDictionary dictionary, ArucoDetectorParameters params) {
    return new cv::aruco::ArucoDetector(*dictionary, *params);
}

void ArucoDetector_Close(ArucoDetector ad) {
    delete ad;
}

void ArucoDetector_DetectMarkers(ArucoDetector ad, Mat inputArr, Points2fVector markerCorners, IntVector *markerIds, Points2fVector rejectedCandidates) {
    std::vector<int> _markerIds;
    ad->detectMarkers(*inputArr, *markerCorners, _markerIds, *rejectedCandidates);

    int *ids = new int[_markerIds.size()];

    for (size_t i = 0; i < _markerIds.size(); ++i)
    {
        ids[i] = _markerIds[i];
    }

    markerIds->length = _markerIds.size();
    markerIds->val = ids;
}

ArucoDetectorParameters ArucoDetectorParameters_Create()
{
    return new cv::aruco::DetectorParameters();
}

void ArucoDetectorParameters_SetAdaptiveThreshWinSizeMin(ArucoDetectorParameters ap, int adaptiveThreshWinSizeMin) {
    ap->adaptiveThreshWinSizeMin = adaptiveThreshWinSizeMin;
}

int ArucoDetectorParameters_GetAdaptiveThreshWinSizeMin(ArucoDetectorParameters ap) {
    return ap->adaptiveThreshWinSizeMin;
}

void ArucoDetectorParameters_SetAdaptiveThreshWinSizeMax(ArucoDetectorParameters ap, int adaptiveThreshWinSizeMax) {
    ap->adaptiveThreshWinSizeMax = adaptiveThreshWinSizeMax;
}

int ArucoDetectorParameters_GetAdaptiveThreshWinSizeMax(ArucoDetectorParameters ap) {
    return ap->adaptiveThreshWinSizeMax;
}

void ArucoDetectorParameters_SetAdaptiveThreshWinSizeStep(ArucoDetectorParameters ap, int adaptiveThreshWinSizeStep) {
    ap->adaptiveThreshWinSizeStep = adaptiveThreshWinSizeStep;
}

int ArucoDetectorParameters_GetAdaptiveThreshWinSizeStep(ArucoDetectorParameters ap) {
    return ap->adaptiveThreshWinSizeStep;
}

void ArucoDetectorParameters_SetAdaptiveThreshConstant(ArucoDetectorParameters ap, double adaptiveThreshConstant) {
    ap->adaptiveThreshConstant = adaptiveThreshConstant;
}

double ArucoDetectorParameters_GetAdaptiveThreshConstant(ArucoDetectorParameters ap) {
    return ap->adaptiveThreshConstant;
}

void ArucoDetectorParameters_SetMinMarkerPerimeterRate(ArucoDetectorParameters ap, double minMarkerPerimeterRate) {
    ap->minMarkerPerimeterRate = minMarkerPerimeterRate;
}

double ArucoDetectorParameters_GetMinMarkerPerimeterRate(ArucoDetectorParameters ap){
    return ap->minMarkerPerimeterRate;
}

void ArucoDetectorParameters_SetMaxMarkerPerimeterRate(ArucoDetectorParameters ap, double maxMarkerPerimeterRate) {
    ap->maxMarkerPerimeterRate = maxMarkerPerimeterRate;
}

double ArucoDetectorParameters_GetMaxMarkerPerimeterRate(ArucoDetectorParameters ap){
    return ap->maxMarkerPerimeterRate;
}

void ArucoDetectorParameters_SetPolygonalApproxAccuracyRate(ArucoDetectorParameters ap, double polygonalApproxAccuracyRate) {
    ap->polygonalApproxAccuracyRate = polygonalApproxAccuracyRate;
}

double ArucoDetectorParameters_GetPolygonalApproxAccuracyRate(ArucoDetectorParameters ap){
    return ap->polygonalApproxAccuracyRate;
}

void ArucoDetectorParameters_SetMinCornerDistanceRate(ArucoDetectorParameters ap, double minCornerDistanceRate) {
    ap->minCornerDistanceRate = minCornerDistanceRate;
}

double ArucoDetectorParameters_GetMinCornerDistanceRate(ArucoDetectorParameters ap) {
    return ap->minCornerDistanceRate;
}

void ArucoDetectorParameters_SetMinDistanceToBorder(ArucoDetectorParameters ap, int minDistanceToBorder) {
    ap->minDistanceToBorder = minDistanceToBorder;
}

int ArucoDetectorParameters_GetMinDistanceToBorder(ArucoDetectorParameters ap) {
    return ap->minDistanceToBorder;
}

void ArucoDetectorParameters_SetMinMarkerDistanceRate(ArucoDetectorParameters ap, double minMarkerDistanceRate) {
    ap->minMarkerDistanceRate = minMarkerDistanceRate;
}

double ArucoDetectorParameters_GetMinMarkerDistanceRate(ArucoDetectorParameters ap) {
    return ap->minMarkerDistanceRate;
}

void ArucoDetectorParameters_SetCornerRefinementMethod(ArucoDetectorParameters ap, int cornerRefinementMethod) {
    ap->cornerRefinementMethod = cv::aruco::CornerRefineMethod(cornerRefinementMethod);
}

int ArucoDetectorParameters_GetCornerRefinementMethod(ArucoDetectorParameters ap) {
    return ap->cornerRefinementMethod;
}

void ArucoDetectorParameters_SetCornerRefinementWinSize(ArucoDetectorParameters ap, int cornerRefinementWinSize) {
    ap->cornerRefinementWinSize = cornerRefinementWinSize;   
}

int ArucoDetectorParameters_GetCornerRefinementWinSize(ArucoDetectorParameters ap) {
    return ap->cornerRefinementWinSize;
}

void ArucoDetectorParameters_SetCornerRefinementMaxIterations(ArucoDetectorParameters ap, int cornerRefinementMaxIterations) {
    ap->cornerRefinementMaxIterations = cornerRefinementMaxIterations;
}

int ArucoDetectorParameters_GetCornerRefinementMaxIterations(ArucoDetectorParameters ap) {
    return ap->cornerRefinementMaxIterations;
}

void ArucoDetectorParameters_SetCornerRefinementMinAccuracy(ArucoDetectorParameters ap, double cornerRefinementMinAccuracy) {
    ap->cornerRefinementMinAccuracy = cornerRefinementMinAccuracy;
}

double ArucoDetectorParameters_GetCornerRefinementMinAccuracy(ArucoDetectorParameters ap) {
    return ap->cornerRefinementMinAccuracy;
}

void ArucoDetectorParameters_SetMarkerBorderBits(ArucoDetectorParameters ap, int markerBorderBits) {
    ap->markerBorderBits = markerBorderBits;
}

int ArucoDetectorParameters_GetMarkerBorderBits(ArucoDetectorParameters ap) {
    return ap->markerBorderBits;
}

void ArucoDetectorParameters_SetPerspectiveRemovePixelPerCell(ArucoDetectorParameters ap, int perspectiveRemovePixelPerCell) {
    ap->perspectiveRemovePixelPerCell = perspectiveRemovePixelPerCell;
}

int ArucoDetectorParameters_GetPerspectiveRemovePixelPerCell(ArucoDetectorParameters ap) {
    return ap->perspectiveRemovePixelPerCell;
}

void ArucoDetectorParameters_SetPerspectiveRemoveIgnoredMarginPerCell(ArucoDetectorParameters ap, double perspectiveRemoveIgnoredMarginPerCell) {
    ap->perspectiveRemoveIgnoredMarginPerCell = perspectiveRemoveIgnoredMarginPerCell;
}

double ArucoDetectorParameters_GetPerspectiveRemoveIgnoredMarginPerCell(ArucoDetectorParameters ap) {
    return ap->perspectiveRemoveIgnoredMarginPerCell;
}

void ArucoDetectorParameters_SetMaxErroneousBitsInBorderRate(ArucoDetectorParameters ap, double maxErroneousBitsInBorderRate) {
    ap->maxErroneousBitsInBorderRate = maxErroneousBitsInBorderRate;
}

double ArucoDetectorParameters_GetMaxErroneousBitsInBorderRate(ArucoDetectorParameters ap) {
    return ap->maxErroneousBitsInBorderRate;
}

void ArucoDetectorParameters_SetMinOtsuStdDev(ArucoDetectorParameters ap, double minOtsuStdDev) {
    ap->minOtsuStdDev = minOtsuStdDev;
}

double ArucoDetectorParameters_GetMinOtsuStdDev(ArucoDetectorParameters ap) {
    return ap->minOtsuStdDev;
}

void ArucoDetectorParameters_SetErrorCorrectionRate(ArucoDetectorParameters ap, double errorCorrectionRate) {
    ap->errorCorrectionRate = errorCorrectionRate;
}

double ArucoDetectorParameters_GetErrorCorrectionRate(ArucoDetectorParameters ap) {
    return ap->errorCorrectionRate;
}

void ArucoDetectorParameters_SetAprilTagQuadDecimate(ArucoDetectorParameters ap, float aprilTagQuadDecimate) {
    ap->aprilTagQuadDecimate = aprilTagQuadDecimate;
}

float ArucoDetectorParameters_GetAprilTagQuadDecimate(ArucoDetectorParameters ap) {
    return ap->aprilTagQuadDecimate;
}

void ArucoDetectorParameters_SetAprilTagQuadSigma(ArucoDetectorParameters ap, float aprilTagQuadSigma) {
    ap->aprilTagQuadSigma = aprilTagQuadSigma;
}

float ArucoDetectorParameters_GetAprilTagQuadSigma(ArucoDetectorParameters ap) {
    return ap->aprilTagQuadSigma;
}

void ArucoDetectorParameters_SetAprilTagMinClusterPixels(ArucoDetectorParameters ap, int aprilTagMinClusterPixels) {
    ap->aprilTagMinClusterPixels = aprilTagMinClusterPixels;
}

int ArucoDetectorParameters_GetAprilTagMinClusterPixels(ArucoDetectorParameters ap) {
    return ap->aprilTagMinClusterPixels;
}

void ArucoDetectorParameters_SetAprilTagMaxNmaxima(ArucoDetectorParameters ap, int aprilTagMaxNmaxima) {
    ap->aprilTagMaxNmaxima = aprilTagMaxNmaxima;
}

int ArucoDetectorParameters_GetAprilTagMaxNmaxima(ArucoDetectorParameters ap) {
    return ap->aprilTagMaxNmaxima;
}

void ArucoDetectorParameters_SetAprilTagCriticalRad(ArucoDetectorParameters ap, float aprilTagCriticalRad) {
    ap->aprilTagCriticalRad = aprilTagCriticalRad;
}

float ArucoDetectorParameters_GetAprilTagCriticalRad(ArucoDetectorParameters ap) {
    return ap->aprilTagCriticalRad;
}

void ArucoDetectorParameters_SetAprilTagMaxLineFitMse(ArucoDetectorParameters ap, float aprilTagMaxLineFitMse) {
    ap->aprilTagMaxLineFitMse = aprilTagMaxLineFitMse;
}

float ArucoDetectorParameters_GetAprilTagMaxLineFitMse(ArucoDetectorParameters ap) {
    return ap->aprilTagMaxLineFitMse;
}

void ArucoDetectorParameters_SetAprilTagMinWhiteBlackDiff(ArucoDetectorParameters ap, int aprilTagMinWhiteBlackDiff) {
    ap->aprilTagMinWhiteBlackDiff = aprilTagMinWhiteBlackDiff;
}

int ArucoDetectorParameters_GetAprilTagMinWhiteBlackDiff(ArucoDetectorParameters ap) {
    return ap->aprilTagMinWhiteBlackDiff;
}

void ArucoDetectorParameters_SetAprilTagDeglitch(ArucoDetectorParameters ap, int aprilTagDeglitch) {
    ap->aprilTagDeglitch = aprilTagDeglitch;
}

int ArucoDetectorParameters_GetAprilTagDeglitch(ArucoDetectorParameters ap) {
    return ap->aprilTagDeglitch;
}

void ArucoDetectorParameters_SetDetectInvertedMarker(ArucoDetectorParameters ap, bool detectInvertedMarker) {
    ap->detectInvertedMarker = detectInvertedMarker;
}

bool ArucoDetectorParameters_GetDetectInvertedMarker(ArucoDetectorParameters ap) {
    return ap->detectInvertedMarker;
}

void ArucoDrawDetectedMarkers(Mat image, Points2fVector markerCorners, IntVector markerIds, Scalar borderColor)
{
    std::vector<int> _markerIds;
    for (int i = 0, *v = markerIds.val; i < markerIds.length; ++v, ++i)
    {
        _markerIds.push_back(*v);
    }
    cv::Scalar _borderColor = cv::Scalar(borderColor.val1, borderColor.val2, borderColor.val3);
    cv::aruco::drawDetectedMarkers(*image, *markerCorners, _markerIds, _borderColor);
}

void ArucoGenerateImageMarker(int dictionaryId, int id, int sidePixels, Mat img, int borderBits)
{
    cv::aruco::Dictionary dict = cv::aruco::getPredefinedDictionary(dictionaryId);
    cv::aruco::generateImageMarker(dict, id, sidePixels, *img, borderBits);
}

ArucoDictionary getPredefinedDictionary(int dictionaryId)
{
    return new cv::aruco::Dictionary(cv::aruco::getPredefinedDictionary(dictionaryId));
}