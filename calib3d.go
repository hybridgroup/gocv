package gocv

/*
#include <stdlib.h>
#include "calib3d.h"
*/
import "C"
import (
	"image"
)

// Calib is a wrapper around OpenCV's "Camera Calibration and 3D Reconstruction" of
// Fisheye Camera model
//
// For more details, please see:
// https://docs.opencv.org/trunk/db/d58/group__calib3d__fisheye.html

// CalibFlag value for calibration
type CalibFlag int32

const (
	// CalibUseIntrinsicGuess indicates that cameraMatrix contains valid initial values
	// of fx, fy, cx, cy that are optimized further. Otherwise, (cx, cy) is initially
	// set to the image center ( imageSize is used), and focal distances are computed
	// in a least-squares fashion.
	CalibUseIntrinsicGuess CalibFlag = 1 << iota

	// CalibRecomputeExtrinsic indicates that extrinsic will be recomputed after each
	// iteration of intrinsic optimization.
	CalibRecomputeExtrinsic

	// CalibCheckCond indicates that the functions will check validity of condition number
	CalibCheckCond

	// CalibFixSkew indicates that skew coefficient (alpha) is set to zero and stay zero
	CalibFixSkew

	// CalibFixK1 indicates that selected distortion coefficients are set to zeros and stay zero
	CalibFixK1

	// CalibFixK2 indicates that selected distortion coefficients are set to zeros and stay zero
	CalibFixK2

	// CalibFixK3 indicates that selected distortion coefficients are set to zeros and stay zero
	CalibFixK3

	// CalibFixK4 indicates that selected distortion coefficients are set to zeros and stay zero
	CalibFixK4

	// CalibFixIntrinsic indicates that fix K1, K2? and D1, D2? so that only R, T matrices are estimated
	CalibFixIntrinsic

	// CalibFixPrincipalPoint indicates that the principal point is not changed during the global optimization.
	// It stays at the center or at a different location specified when CalibUseIntrinsicGuess is set too.
	CalibFixPrincipalPoint
)

// FisheyeUndistortImage transforms an image to compensate for fisheye lens distortion
func FisheyeUndistortImage(distorted Mat, undistorted *Mat, k, d Mat) {
	C.Fisheye_UndistortImage(distorted.Ptr(), undistorted.Ptr(), k.Ptr(), d.Ptr())
}

// FisheyeUndistortImageWithParams transforms an image to compensate for fisheye lens distortion with Knew matrix
func FisheyeUndistortImageWithParams(distorted Mat, undistorted *Mat, k, d, knew Mat, size image.Point) {
	sz := C.struct_Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}
	C.Fisheye_UndistortImageWithParams(distorted.Ptr(), undistorted.Ptr(), k.Ptr(), d.Ptr(), knew.Ptr(), sz)
}

// FisheyeUndistortPoints transforms points to compensate for fisheye lens distortion
//
// For further details, please see:
// https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#gab738cdf90ceee97b2b52b0d0e7511541
func FisheyeUndistortPoints(distorted Mat, undistorted *Mat, k, d, r, p Mat) {
	C.Fisheye_UndistortPoints(distorted.Ptr(), undistorted.Ptr(), k.Ptr(), d.Ptr(), r.Ptr(), p.Ptr())
}

// EstimateNewCameraMatrixForUndistortRectify estimates new camera matrix for undistortion or rectification.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#ga384940fdf04c03e362e94b6eb9b673c9
func EstimateNewCameraMatrixForUndistortRectify(k, d Mat, imgSize image.Point, r Mat, p *Mat, balance float64, newSize image.Point, fovScale float64) {
	imgSz := C.struct_Size{
		width:  C.int(imgSize.X),
		height: C.int(imgSize.Y),
	}
	newSz := C.struct_Size{
		width:  C.int(newSize.X),
		height: C.int(newSize.Y),
	}
	C.Fisheye_EstimateNewCameraMatrixForUndistortRectify(k.Ptr(), d.Ptr(), imgSz, r.Ptr(), p.Ptr(), C.double(balance), newSz, C.double(fovScale))
}

// InitUndistortRectifyMap computes the joint undistortion and rectification transformation and represents the result in the form of maps for remap
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga7dfb72c9cf9780a347fbe3d1c47e5d5a
//
func InitUndistortRectifyMap(cameraMatrix Mat, distCoeffs Mat, r Mat, newCameraMatrix Mat, size image.Point, m1type int, map1 Mat, map2 Mat) {
	sz := C.struct_Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}
	C.InitUndistortRectifyMap(cameraMatrix.Ptr(), distCoeffs.Ptr(), r.Ptr(), newCameraMatrix.Ptr(), sz, C.int(m1type), map1.Ptr(), map2.Ptr())
}

// GetOptimalNewCameraMatrixWithParams computes and returns the optimal new camera matrix based on the free scaling parameter.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga7a6c4e032c97f03ba747966e6ad862b1
//
func GetOptimalNewCameraMatrixWithParams(cameraMatrix Mat, distCoeffs Mat, imageSize image.Point, alpha float64, newImgSize image.Point, centerPrincipalPoint bool) (Mat, image.Rectangle) {
	sz := C.struct_Size{
		width:  C.int(imageSize.X),
		height: C.int(imageSize.Y),
	}
	newSize := C.struct_Size{
		width:  C.int(newImgSize.X),
		height: C.int(newImgSize.Y),
	}
	rt := C.struct_Rect{}
	return newMat(C.GetOptimalNewCameraMatrixWithParams(cameraMatrix.Ptr(), distCoeffs.Ptr(), sz, C.double(alpha), newSize, &rt, C.bool(centerPrincipalPoint))), toRect(rt)
}

// CalibrateCamera finds the camera intrinsic and extrinsic parameters from several views of a calibration pattern.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga3207604e4b1a1758aa66acb6ed5aa65d
//
func CalibrateCamera(objectPoints Points3fVector, imagePoints Points2fVector, imageSize image.Point,
	cameraMatrix *Mat, distCoeffs *Mat, rvecs *Mat, tvecs *Mat, calibFlag CalibFlag) float64 {
	sz := C.struct_Size{
		width:  C.int(imageSize.X),
		height: C.int(imageSize.Y),
	}

	res := C.CalibrateCamera(objectPoints.p, imagePoints.p, sz, cameraMatrix.p, distCoeffs.p, rvecs.p, tvecs.p, C.int(calibFlag))
	return float64(res)
}

func Undistort(src Mat, dst *Mat, cameraMatrix Mat, distCoeffs Mat, newCameraMatrix Mat) {
	C.Undistort(src.Ptr(), dst.Ptr(), cameraMatrix.Ptr(), distCoeffs.Ptr(), newCameraMatrix.Ptr())
}

// UndistortPoints transforms points to compensate for lens distortion
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga55c716492470bfe86b0ee9bf3a1f0f7e
func UndistortPoints(src Mat, dst *Mat, cameraMatrix, distCoeffs, rectificationTransform, newCameraMatrix Mat) {
	C.UndistortPoints(src.Ptr(), dst.Ptr(), cameraMatrix.Ptr(), distCoeffs.Ptr(), rectificationTransform.Ptr(), newCameraMatrix.Ptr())
}

// CalibCBFlag value for chessboard calibration
// For more details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga93efa9b0aa890de240ca32b11253dd4a
type CalibCBFlag int

const (
	// Various operation flags that can be zero or a combination of the following values:
	//  Use adaptive thresholding to convert the image to black and white, rather than a fixed threshold level (computed from the average image brightness).
	CalibCBAdaptiveThresh CalibCBFlag = 1 << iota
	//  Normalize the image gamma with equalizeHist before applying fixed or adaptive thresholding.
	CalibCBNormalizeImage
	//  Use additional criteria (like contour area, perimeter, square-like shape) to filter out false quads extracted at the contour retrieval stage.
	CalibCBFilterQuads
	//  Run a fast check on the image that looks for chessboard corners, and shortcut the call if none is found. This can drastically speed up the call in the degenerate condition when no chessboard is observed.
	CalibCBFastCheck
	//  Run an exhaustive search to improve detection rate.
	CalibCBExhaustive
	//  Up sample input image to improve sub-pixel accuracy due to aliasing effects.
	CalibCBAccuracy
	//  The detected pattern is allowed to be larger than patternSize (see description).
	CalibCBLarger
	//  The detected pattern must have a marker (see description). This should be used if an accurate camera calibration is required.
	CalibCBMarker
)

// FindChessboardCorners finds the positions of internal corners of the chessboard.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga93efa9b0aa890de240ca32b11253dd4a
//
func FindChessboardCorners(image Mat, patternSize image.Point, corners *Mat, flags CalibCBFlag) bool {
	sz := C.struct_Size{
		width:  C.int(patternSize.X),
		height: C.int(patternSize.Y),
	}
	return bool(C.FindChessboardCorners(image.Ptr(), sz, corners.Ptr(), C.int(flags)))
}

// FindChessboardCorners finds the positions of internal corners of the chessboard using a sector based approach.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#gadc5bcb05cb21cf1e50963df26986d7c9
//
func FindChessboardCornersSB(image Mat, patternSize image.Point, corners *Mat, flags CalibCBFlag) bool {
	sz := C.struct_Size{
		width:  C.int(patternSize.X),
		height: C.int(patternSize.Y),
	}
	return bool(C.FindChessboardCornersSB(image.Ptr(), sz, corners.Ptr(), C.int(flags)))
}

// FindChessboardCornersSBWithMeta finds the positions of internal corners of the chessboard using a sector based approach.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga93efa9b0aa890de240ca32b11253dd4a
//
func FindChessboardCornersSBWithMeta(image Mat, patternSize image.Point, corners *Mat, flags CalibCBFlag, meta *Mat) bool {
	sz := C.struct_Size{
		width:  C.int(patternSize.X),
		height: C.int(patternSize.Y),
	}
	return bool(C.FindChessboardCornersSBWithMeta(image.Ptr(), sz, corners.Ptr(), C.int(flags), meta.Ptr()))
}

// DrawChessboardCorners renders the detected chessboard corners.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga6a10b0bb120c4907e5eabbcd22319022
//
func DrawChessboardCorners(image *Mat, patternSize image.Point, corners Mat, patternWasFound bool) {
	sz := C.struct_Size{
		width:  C.int(patternSize.X),
		height: C.int(patternSize.Y),
	}
	C.DrawChessboardCorners(image.Ptr(), sz, corners.Ptr(), C.bool(patternWasFound))
}

// EstimateAffinePartial2D computes an optimal limited affine transformation
// with 4 degrees of freedom between two 2D point sets.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#gad767faff73e9cbd8b9d92b955b50062d
func EstimateAffinePartial2D(from, to Point2fVector) Mat {
	return newMat(C.EstimateAffinePartial2D(from.p, to.p))
}

// EstimateAffinePartial2DWithParams computes an optimal limited affine transformation
// with 4 degrees of freedom between two 2D point sets
// with additional optional parameters.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#gad767faff73e9cbd8b9d92b955b50062d
func EstimateAffinePartial2DWithParams(from Point2fVector, to Point2fVector, inliers Mat, method int, ransacReprojThreshold float64, maxIters uint, confidence float64, refineIters uint) Mat {
	return newMat(C.EstimateAffinePartial2DWithParams(from.p, to.p, inliers.p, C.int(method), C.double(ransacReprojThreshold), C.size_t(maxIters), C.double(confidence), C.size_t(refineIters)))
}

// EstimateAffine2D Computes an optimal affine transformation between two 2D point sets.
//
// For further details, please see:
// https://docs.opencv.org/4.0.0/d9/d0c/group__calib3d.html#ga27865b1d26bac9ce91efaee83e94d4dd
func EstimateAffine2D(from, to Point2fVector) Mat {
	return newMat(C.EstimateAffine2D(from.p, to.p))
}

// EstimateAffine2DWithParams Computes an optimal affine transformation between two 2D point sets
// with additional optional parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.0.0/d9/d0c/group__calib3d.html#ga27865b1d26bac9ce91efaee83e94d4dd
func EstimateAffine2DWithParams(from Point2fVector, to Point2fVector, inliers Mat, method int, ransacReprojThreshold float64, maxIters uint, confidence float64, refineIters uint) Mat {
	return newMat(C.EstimateAffine2DWithParams(from.p, to.p, inliers.p, C.int(method), C.double(ransacReprojThreshold), C.size_t(maxIters), C.double(confidence), C.size_t(refineIters)))
}
