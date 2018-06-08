package gocv

/*
#include <stdlib.h>
#include "calib3d.h"
*/
import "C"
import "image"

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
