#include "calib3d.h"

double Fisheye_Calibrate(Points3fVector objectPoints, Points2fVector imagePoints, Size size, Mat k, Mat d, Mat rvecs, Mat tvecs, int flags) {
    try {
        cv::Size sz(size.width, size.height);
        return cv::fisheye::calibrate(*objectPoints, *imagePoints, sz, *k, *d, *rvecs, *tvecs, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

OpenCVResult Fisheye_DistortPoints(Mat undistorted, Mat distorted, Mat k, Mat d) {
    try {
        cv::fisheye::distortPoints(*undistorted, *distorted, *k, *d);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d) {
    try {
        cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Fisheye_UndistortImageWithParams(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew, Size size) {
    try {
        cv::Size sz(size.width, size.height);
        cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d, *knew, sz);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Fisheye_UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p) {
    try {
        cv::fisheye::undistortPoints(*distorted, *undistorted, *k, *d, *r, *p);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Fisheye_EstimateNewCameraMatrixForUndistortRectify(Mat k, Mat d, Size imgSize, Mat r, Mat p, double balance, Size newSize, double fovScale) {
    try {
        cv::Size newSz(newSize.width, newSize.height);
        cv::Size imgSz(imgSize.width, imgSize.height);
        cv::fisheye::estimateNewCameraMatrixForUndistortRectify(*k, *d, imgSz, *r, *p, balance, newSz, fovScale);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult InitUndistortRectifyMap(Mat cameraMatrix,Mat distCoeffs,Mat r,Mat newCameraMatrix,Size size,int m1type,Mat map1,Mat map2) {
    try {
        cv::Size sz(size.width, size.height);
        cv::initUndistortRectifyMap(*cameraMatrix,*distCoeffs,*r,*newCameraMatrix,sz,m1type,*map1,*map2);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

Mat GetOptimalNewCameraMatrixWithParams(Mat cameraMatrix,Mat distCoeffs,Size size,double alpha,Size newImgSize,Rect* validPixROI,bool centerPrincipalPoint) {
    try {
        cv::Size sz(size.width, size.height);
        cv::Size newSize(newImgSize.width, newImgSize.height);
        cv::Rect rect(validPixROI->x,validPixROI->y,validPixROI->width,validPixROI->height);
        cv::Mat* mat = new cv::Mat(cv::getOptimalNewCameraMatrix(*cameraMatrix,*distCoeffs,sz,alpha,newSize,&rect,centerPrincipalPoint));
        validPixROI->x = rect.x;
        validPixROI->y = rect.y;
        validPixROI->width = rect.width;
        validPixROI->height = rect.height;
        return mat;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

double CalibrateCamera(Points3fVector objectPoints, Points2fVector imagePoints, Size imageSize, Mat cameraMatrix, Mat distCoeffs, Mat rvecs, Mat tvecs, int flag) {
    try {
        return cv::calibrateCamera(*objectPoints, *imagePoints, cv::Size(imageSize.width, imageSize.height), *cameraMatrix, *distCoeffs, *rvecs, *tvecs, flag);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

OpenCVResult Undistort(Mat src, Mat dst, Mat cameraMatrix, Mat distCoeffs, Mat newCameraMatrix) {
    try {
        cv::undistort(*src, *dst, *cameraMatrix, *distCoeffs, *newCameraMatrix);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p) {
    try {
        cv::undistortPoints(*distorted, *undistorted, *k, *d, *r, *p);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

bool CheckChessboard(Mat image, Size size) {
    try {
        cv::Size sz(size.width, size.height);
        return cv::checkChessboard(*image, sz);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool FindChessboardCorners(Mat image, Size patternSize, Mat corners, int flags) {
    try {
        cv::Size sz(patternSize.width, patternSize.height);
        return cv::findChessboardCorners(*image, sz, *corners, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool FindChessboardCornersSB(Mat image, Size patternSize, Mat corners, int flags) {
    try {
        cv::Size sz(patternSize.width, patternSize.height);
        return cv::findChessboardCornersSB(*image, sz, *corners, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool FindChessboardCornersSBWithMeta(Mat image, Size patternSize, Mat corners, int flags, Mat meta) {
    try {
        cv::Size sz(patternSize.width, patternSize.height);
        return cv::findChessboardCornersSB(*image, sz, *corners, flags, *meta);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

OpenCVResult DrawChessboardCorners(Mat image, Size patternSize, Mat corners, bool patternWasFound) {
    try {
        cv::Size sz(patternSize.width, patternSize.height);
        cv::drawChessboardCorners(*image, sz, *corners, patternWasFound);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

Mat EstimateAffinePartial2D(Point2fVector from, Point2fVector to) {
    try {
        return new cv::Mat(cv::estimateAffinePartial2D(*from, *to));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat EstimateAffinePartial2DWithParams(Point2fVector from, Point2fVector to, Mat inliers, int method, double ransacReprojThreshold, size_t maxIters, double confidence, size_t refineIters) {
    try {
        return new cv::Mat(cv::estimateAffinePartial2D(*from, *to, *inliers, method, ransacReprojThreshold, maxIters, confidence, refineIters));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat EstimateAffine2D(Point2fVector from, Point2fVector to) {
    try {
        return new cv::Mat(cv::estimateAffine2D(*from, *to));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat EstimateAffine2DWithParams(Point2fVector from, Point2fVector to, Mat inliers, int method, double ransacReprojThreshold, size_t maxIters, double confidence, size_t refineIters) {
    try {
        return new cv::Mat(cv::estimateAffine2D(*from, *to, *inliers, method, ransacReprojThreshold, maxIters, confidence, refineIters));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

OpenCVResult TriangulatePoints(Mat projMatr1, Mat projMatr2, Point2fVector projPoints1, Point2fVector projPoints2, Mat points4D) {
    try {
        cv::triangulatePoints(*projMatr1, *projMatr2, *projPoints1, *projPoints2, *points4D);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ConvertPointsFromHomogeneous(Mat src, Mat dst) {
    try {
        cv::convertPointsFromHomogeneous(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Rodrigues(Mat src, Mat dst) {
    try {
        cv::Rodrigues(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

bool SolvePnP(Point3fVector objectPoints, Point2fVector imagePoints, Mat cameraMatrix, Mat distCoeffs, Mat rvec, Mat tvec, bool useExtrinsicGuess, int flags) {
    try {
        return cv::solvePnP(*objectPoints, *imagePoints, *cameraMatrix, *distCoeffs, *rvec, *tvec, useExtrinsicGuess, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

OpenCVResult StereoRectify(Mat cameraMatrix1, Mat distCoeffs1, Mat cameraMatrix2, Mat distCoeffs2, Size imageSize, Mat r, Mat t, Mat R1, Mat R2, Mat P1, Mat P2, Mat Q, int flags) {
    try {
        cv::stereoRectify(*cameraMatrix1, *distCoeffs1, *cameraMatrix2, *distCoeffs2, cv::Size(imageSize.width, imageSize.height), *r, *t, *R1, *R2, *P1, *P2, *Q, flags);
        return successResult();
    } catch(const cv::Exception& e){
        return errorResult(e.code, e.what());
    }
}