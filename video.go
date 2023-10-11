package gocv

/*
#include <stdlib.h>
#include "video.h"
*/
import "C"
import (
	"image"
	"unsafe"
)

/*
*

	cv::OPTFLOW_USE_INITIAL_FLOW = 4,
	cv::OPTFLOW_LK_GET_MIN_EIGENVALS = 8,
	cv::OPTFLOW_FARNEBACK_GAUSSIAN = 256
	For further details, please see: https://docs.opencv.org/master/dc/d6b/group__video__track.html#gga2c6cc144c9eee043575d5b311ac8af08a9d4430ac75199af0cf6fcdefba30eafe
*/
const (
	OptflowUseInitialFlow    = 4
	OptflowLkGetMinEigenvals = 8
	OptflowFarnebackGaussian = 256
)

/*
*

	cv::MOTION_TRANSLATION = 0,
	cv::MOTION_EUCLIDEAN = 1,
	cv::MOTION_AFFINE = 2,
	cv::MOTION_HOMOGRAPHY = 3
	For further details, please see: https://docs.opencv.org/4.x/dc/d6b/group__video__track.html#ggaaedb1f94e6b143cef163622c531afd88a01106d6d20122b782ff25eaeffe9a5be
*/
const (
	MotionTranslation = 0
	MotionEuclidean   = 1
	MotionAffine      = 2
	MotionHomography  = 3
)

// BackgroundSubtractorMOG2 is a wrapper around the cv::BackgroundSubtractorMOG2.
type BackgroundSubtractorMOG2 struct {
	// C.BackgroundSubtractorMOG2
	p unsafe.Pointer
}

// NewBackgroundSubtractorMOG2 returns a new BackgroundSubtractor algorithm
// of type MOG2. MOG2 is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/de/de1/group__video__motion.html#ga2beb2dee7a073809ccec60f145b6b29c
// https://docs.opencv.org/master/d7/d7b/classcv_1_1BackgroundSubtractorMOG2.html
func NewBackgroundSubtractorMOG2() BackgroundSubtractorMOG2 {
	return BackgroundSubtractorMOG2{p: unsafe.Pointer(C.BackgroundSubtractorMOG2_Create())}
}

// NewBackgroundSubtractorMOG2WithParams returns a new BackgroundSubtractor algorithm
// of type MOG2 with customized parameters. MOG2 is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/de/de1/group__video__motion.html#ga2beb2dee7a073809ccec60f145b6b29c
// https://docs.opencv.org/master/d7/d7b/classcv_1_1BackgroundSubtractorMOG2.html
func NewBackgroundSubtractorMOG2WithParams(history int, varThreshold float64, detectShadows bool) BackgroundSubtractorMOG2 {
	return BackgroundSubtractorMOG2{p: unsafe.Pointer(C.BackgroundSubtractorMOG2_CreateWithParams(C.int(history), C.double(varThreshold), C.bool(detectShadows)))}
}

// Close BackgroundSubtractorMOG2.
func (b *BackgroundSubtractorMOG2) Close() error {
	C.BackgroundSubtractorMOG2_Close((C.BackgroundSubtractorMOG2)(b.p))
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorMOG2.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df6/classcv_1_1BackgroundSubtractor.html#aa735e76f7069b3fa9c3f32395f9ccd21
func (b *BackgroundSubtractorMOG2) Apply(src Mat, dst *Mat) {
	C.BackgroundSubtractorMOG2_Apply((C.BackgroundSubtractorMOG2)(b.p), src.p, dst.p)
	return
}

// BackgroundSubtractorKNN is a wrapper around the cv::BackgroundSubtractorKNN.
type BackgroundSubtractorKNN struct {
	// C.BackgroundSubtractorKNN
	p unsafe.Pointer
}

// NewBackgroundSubtractorKNN returns a new BackgroundSubtractor algorithm
// of type KNN. K-Nearest Neighbors (KNN) uses a Background/Foreground
// Segmentation Algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/de/de1/group__video__motion.html#gac9be925771f805b6fdb614ec2292006d
// https://docs.opencv.org/master/db/d88/classcv_1_1BackgroundSubtractorKNN.html
func NewBackgroundSubtractorKNN() BackgroundSubtractorKNN {
	return BackgroundSubtractorKNN{p: unsafe.Pointer(C.BackgroundSubtractorKNN_Create())}
}

// NewBackgroundSubtractorKNNWithParams returns a new BackgroundSubtractor algorithm
// of type KNN with customized parameters. K-Nearest Neighbors (KNN) uses a Background/Foreground
// Segmentation Algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/de/de1/group__video__motion.html#gac9be925771f805b6fdb614ec2292006d
// https://docs.opencv.org/master/db/d88/classcv_1_1BackgroundSubtractorKNN.html
func NewBackgroundSubtractorKNNWithParams(history int, dist2Threshold float64, detectShadows bool) BackgroundSubtractorKNN {
	return BackgroundSubtractorKNN{p: unsafe.Pointer(C.BackgroundSubtractorKNN_CreateWithParams(C.int(history), C.double(dist2Threshold), C.bool(detectShadows)))}
}

// Close BackgroundSubtractorKNN.
func (k *BackgroundSubtractorKNN) Close() error {
	C.BackgroundSubtractorKNN_Close((C.BackgroundSubtractorKNN)(k.p))
	k.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorKNN.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df6/classcv_1_1BackgroundSubtractor.html#aa735e76f7069b3fa9c3f32395f9ccd21
func (k *BackgroundSubtractorKNN) Apply(src Mat, dst *Mat) {
	C.BackgroundSubtractorKNN_Apply((C.BackgroundSubtractorKNN)(k.p), src.p, dst.p)
	return
}

// CalcOpticalFlowFarneback computes a dense optical flow using
// Gunnar Farneback's algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga5d10ebbd59fe09c5f650289ec0ece5af
func CalcOpticalFlowFarneback(prevImg Mat, nextImg Mat, flow *Mat, pyrScale float64, levels int, winsize int,
	iterations int, polyN int, polySigma float64, flags int) {
	C.CalcOpticalFlowFarneback(prevImg.p, nextImg.p, flow.p, C.double(pyrScale), C.int(levels), C.int(winsize),
		C.int(iterations), C.int(polyN), C.double(polySigma), C.int(flags))
	return
}

// CalcOpticalFlowPyrLK calculates an optical flow for a sparse feature set using
// the iterative Lucas-Kanade method with pyramids.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga473e4b886d0bcc6b65831eb88ed93323
func CalcOpticalFlowPyrLK(prevImg Mat, nextImg Mat, prevPts Mat, nextPts Mat, status *Mat, err *Mat) {
	C.CalcOpticalFlowPyrLK(prevImg.p, nextImg.p, prevPts.p, nextPts.p, status.p, err.p)
	return
}

// CalcOpticalFlowPyrLKWithParams calculates an optical flow for a sparse feature set using
// the iterative Lucas-Kanade method with pyramids.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga473e4b886d0bcc6b65831eb88ed93323
func CalcOpticalFlowPyrLKWithParams(prevImg Mat, nextImg Mat, prevPts Mat, nextPts Mat, status *Mat, err *Mat,
	winSize image.Point, maxLevel int, criteria TermCriteria, flags int, minEigThreshold float64) {
	winSz := C.struct_Size{
		width:  C.int(winSize.X),
		height: C.int(winSize.Y),
	}
	C.CalcOpticalFlowPyrLKWithParams(prevImg.p, nextImg.p, prevPts.p, nextPts.p, status.p, err.p, winSz, C.int(maxLevel), criteria.p, C.int(flags), C.double(minEigThreshold))
	return
}

// FindTransformECC finds the geometric transform (warp) between two images in terms of the ECC criterion.
//
// For futther details, please see:
// https://docs.opencv.org/4.x/dc/d6b/group__video__track.html#ga1aa357007eaec11e9ed03500ecbcbe47
func FindTransformECC(templateImage Mat, inputImage Mat, warpMatrix *Mat, motionType int, criteria TermCriteria, inputMask Mat, gaussFiltSize int) float64 {
	return float64(C.FindTransformECC(templateImage.p, inputImage.p, warpMatrix.p, C.int(motionType), criteria.p, inputMask.p, C.int(gaussFiltSize)))
}

// Tracker is the base interface for object tracking.
//
// see: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html
type Tracker interface {
	// Close closes, as Trackers need to be Closed manually.
	//
	Close() error

	// Init initializes the tracker with a known bounding box that surrounded the target.
	// Note: this can only be called once. If you lose the object, you have to Close() the instance,
	// create a new one, and call Init() on it again.
	//
	// see: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html#a4d285747589b1bdd16d2e4f00c3255dc
	//
	Init(image Mat, boundingBox image.Rectangle) bool

	// Update updates the tracker, returns a new bounding box and a boolean determining whether the tracker lost the target.
	//
	// see: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html#a549159bd0553e6a8de356f3866df1f18
	//
	Update(image Mat) (image.Rectangle, bool)
}

func trackerInit(trk C.Tracker, img Mat, boundingBox image.Rectangle) bool {
	cBox := C.struct_Rect{
		x:      C.int(boundingBox.Min.X),
		y:      C.int(boundingBox.Min.Y),
		width:  C.int(boundingBox.Size().X),
		height: C.int(boundingBox.Size().Y),
	}

	ret := C.Tracker_Init(trk, C.Mat(img.Ptr()), cBox)
	return bool(ret)
}

func trackerUpdate(trk C.Tracker, img Mat) (image.Rectangle, bool) {
	cBox := C.struct_Rect{}

	ret := C.Tracker_Update(trk, C.Mat(img.Ptr()), &cBox)

	rect := image.Rect(int(cBox.x), int(cBox.y), int(cBox.x+cBox.width), int(cBox.y+cBox.height))
	return rect, bool(ret)
}

// TrackerMIL is a Tracker that uses the MIL algorithm. MIL trains a classifier in an online manner
// to separate the object from the background.
// Multiple Instance Learning avoids the drift problem for a robust tracking.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d26/classcv_1_1TrackerMIL.html
type TrackerMIL struct {
	p C.TrackerMIL
}

// NewTrackerMIL returns a new TrackerMIL.
func NewTrackerMIL() Tracker {
	return TrackerMIL{p: C.TrackerMIL_Create()}
}

// Close closes the TrackerMIL.
func (trk TrackerMIL) Close() error {
	C.TrackerMIL_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes the TrackerMIL.
func (trk TrackerMIL) Init(img Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates the TrackerMIL.
func (trk TrackerMIL) Update(img Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// KalmanFilter implements a standard Kalman filter http://en.wikipedia.org/wiki/Kalman_filter.
// However, you can modify transitionMatrix, controlMatrix, and measurementMatrix
// to get an extended Kalman filter functionality.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html
type KalmanFilter struct {
	p C.KalmanFilter
}

// NewKalmanFilter returns a new KalmanFilter.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#ac0799f0611baee9e7e558f016e4a7b40
func NewKalmanFilter(dynamParams int, measureParams int) KalmanFilter {
	return KalmanFilter{p: C.KalmanFilter_New(C.int(dynamParams), C.int(measureParams))}
}

// NewKalmanFilterWithParams returns a new KalmanFilter.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#abac82ecfa530611a163255bc7d91c088
func NewKalmanFilterWithParams(dynamParams int, measureParams int, controlParams int, matType MatType) KalmanFilter {
	return KalmanFilter{p: C.KalmanFilter_NewWithParams(C.int(dynamParams), C.int(measureParams), C.int(controlParams), C.int(matType))}
}

// Init re-initializes the Kalman filter. The previous content is destroyed.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a4f136c39c016d3530c7c5801dd1ddb3b
func (kf *KalmanFilter) Init(dynamParams int, measureParams int) {
	C.KalmanFilter_Init(kf.p, C.int(dynamParams), C.int(measureParams))
}

// Predict computes a predicted state.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#aa710d2255566bec8d6ce608d103d4fa7
func (kf *KalmanFilter) Predict() Mat {
	return newMat(C.KalmanFilter_Predict(kf.p))
}

// PredictWithParams computes a predicted state.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#aa710d2255566bec8d6ce608d103d4fa7
func (kf *KalmanFilter) PredictWithParams(control Mat) Mat {
	return newMat(C.KalmanFilter_PredictWithParams(kf.p, control.p))
}

// Correct the predicted state from the measurement.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a60eb7feb569222ad0657ef1875884b5e
func (kf *KalmanFilter) Correct(measurement Mat) Mat {
	return newMat(C.KalmanFilter_Correct(kf.p, measurement.p))
}

// Close closes the kalman filter.
func (kf *KalmanFilter) Close() {
	C.KalmanFilter_Close(kf.p)
	kf.p = nil
}

// GetStatePre returns the Kalman filter's statePre Mat.
//
// predicted state (x'(k)): x(k)=A*x(k-1)+B*u(k)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a60eb7feb569222ad0657ef1875884b5e
func (kf *KalmanFilter) GetStatePre() Mat {
	return newMat(C.KalmanFilter_GetStatePre(kf.p))
}

// GetStatePost returns the Kalman filter's statePost Mat.
//
// corrected state (x(k)): x(k)=x'(k)+K(k)*(z(k)-H*x'(k))
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#add8fb5ac9c04b4600b679698dcb0447d
func (kf *KalmanFilter) GetStatePost() Mat {
	return newMat(C.KalmanFilter_GetStatePost(kf.p))
}

// GetTransitionMatrix returns the Kalman filter's transitionMatrix Mat.
//
// state transition matrix (A)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a0657173e411acbf40d2d3c6b46e03b19
func (kf *KalmanFilter) GetTransitionMatrix() Mat {
	return newMat(C.KalmanFilter_GetTransitionMatrix(kf.p))
}

// GetControlMatrix returns the Kalman filter's controlMatrix Mat.
//
// control matrix (B) (not used if there is no control)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a6486e7287114810636fb33953280ed52
func (kf *KalmanFilter) GetControlMatrix() Mat {
	return newMat(C.KalmanFilter_GetControlMatrix(kf.p))
}

// GetMeasurementMatrix returns the Kalman filter's measurementMatrix Mat.
//
// measurement matrix (H)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a0f60b78726d8eccf74a1f2479c2d1f97
func (kf *KalmanFilter) GetMeasurementMatrix() Mat {
	return newMat(C.KalmanFilter_GetMeasurementMatrix(kf.p))
}

// GetProcessNoiseCov returns the Kalman filter's processNoiseCov Mat.
//
// process noise covariance matrix (Q)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#af19be9c0630d0f658bdbaea409a35cda
func (kf *KalmanFilter) GetProcessNoiseCov() Mat {
	return newMat(C.KalmanFilter_GetProcessNoiseCov(kf.p))
}

// GetMeasurementNoiseCov returns the Kalman filter's measurementNoiseCov Mat.
//
// measurement noise covariance matrix (R)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a828d051035ba807966ad65edf288a08e
func (kf *KalmanFilter) GetMeasurementNoiseCov() Mat {
	return newMat(C.KalmanFilter_GetMeasurementNoiseCov(kf.p))
}

// GetErrorCovPre returns the Kalman filter's errorCovPre Mat.
//
// priori error estimate covariance matrix (P'(k)): P'(k)=A*P(k-1)*At + Q)*/
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#ae1bd3a86f10753d723e7174d570d9ac1
func (kf *KalmanFilter) GetErrorCovPre() Mat {
	return newMat(C.KalmanFilter_GetErrorCovPre(kf.p))
}

// GetGain returns the Kalman filter's gain Mat.
//
// Kalman gain matrix (K(k)): K(k)=P'(k)*Ht*inv(H*P'(k)*Ht+R)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a077d73eb075b00779dc009a9057c27c3
func (kf *KalmanFilter) GetGain() Mat {
	return newMat(C.KalmanFilter_GetGain(kf.p))
}

// GetErrorCovPost returns the Kalman filter's errorCovPost Mat.
//
// posteriori error estimate covariance matrix (P(k)): P(k)=(I-K(k)*H)*P'(k)
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a446d8e9a0105b0aa35cd66119c529803
func (kf *KalmanFilter) GetErrorCovPost() Mat {
	return newMat(C.KalmanFilter_GetErrorCovPost(kf.p))
}

// GetTemp1 returns the Kalman filter's temp1 Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#aa3d064a9194c2815dbe19c056b6dc763
func (kf *KalmanFilter) GetTemp1() Mat {
	return newMat(C.KalmanFilter_GetTemp1(kf.p))
}

// GetTemp2 returns the Kalman filter's temp2 Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a14866bd506668eb0ed57b3974b3a1ee7
func (kf *KalmanFilter) GetTemp2() Mat {
	return newMat(C.KalmanFilter_GetTemp2(kf.p))
}

// GetTemp3 returns the Kalman filter's temp3 Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#afdbe36066a7d7f560aa02abe6be114d8
func (kf *KalmanFilter) GetTemp3() Mat {
	return newMat(C.KalmanFilter_GetTemp3(kf.p))
}

// GetTemp4 returns the Kalman filter's temp4 Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a84342f2d9dec1e6389025ad229401809
func (kf *KalmanFilter) GetTemp4() Mat {
	return newMat(C.KalmanFilter_GetTemp4(kf.p))
}

// GetTemp5 returns the Kalman filter's temp5 Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/dd/d6a/classcv_1_1KalmanFilter.html#a846c2a6222c6e5d8b1385dfbccc83ae0
func (kf *KalmanFilter) GetTemp5() Mat {
	return newMat(C.KalmanFilter_GetTemp5(kf.p))
}

// SetStatePre sets the Kalman filter's statePre Mat.
func (kf *KalmanFilter) SetStatePre(statePre Mat) {
	C.KalmanFilter_SetStatePre(kf.p, statePre.p)
}

// SetStatePost sets the Kalman filter's statePost Mat.
func (kf *KalmanFilter) SetStatePost(statePost Mat) {
	C.KalmanFilter_SetStatePost(kf.p, statePost.p)
}

// SetTransitionMatrix sets the Kalman filter's transitionMatrix Mat.
func (kf *KalmanFilter) SetTransitionMatrix(transitionMatrix Mat) {
	C.KalmanFilter_SetTransitionMatrix(kf.p, transitionMatrix.p)
}

// SetControlMatrix sets the Kalman filter's controlMatrix Mat.
func (kf *KalmanFilter) SetControlMatrix(controlMatrix Mat) {
	C.KalmanFilter_SetControlMatrix(kf.p, controlMatrix.p)
}

// SetMeasurementMatrix sets the Kalman filter's measurementMatrix Mat.
func (kf *KalmanFilter) SetMeasurementMatrix(measurementMatrix Mat) {
	C.KalmanFilter_SetMeasurementMatrix(kf.p, measurementMatrix.p)
}

// SetProcessNoiseCov sets the Kalman filter's processNoiseCov Mat.
func (kf *KalmanFilter) SetProcessNoiseCov(processNoiseCov Mat) {
	C.KalmanFilter_SetProcessNoiseCov(kf.p, processNoiseCov.p)
}

// SetMeasurementNoiseCov sets the Kalman filter's measurementNoiseCov Mat.
func (kf *KalmanFilter) SetMeasurementNoiseCov(measurementNoiseCov Mat) {
	C.KalmanFilter_SetMeasurementNoiseCov(kf.p, measurementNoiseCov.p)
}

// SetErrorCovPre sets the Kalman filter's errorCovPre Mat.
func (kf *KalmanFilter) SetErrorCovPre(errorCovPre Mat) {
	C.KalmanFilter_SetErrorCovPre(kf.p, errorCovPre.p)
}

// SetGain sets the Kalman filter's gain Mat.
func (kf *KalmanFilter) SetGain(gain Mat) {
	C.KalmanFilter_SetGain(kf.p, gain.p)
}

// SetErrorCovPost sets the Kalman filter's errorCovPost Mat.
func (kf *KalmanFilter) SetErrorCovPost(errorCovPost Mat) {
	C.KalmanFilter_SetErrorCovPost(kf.p, errorCovPost.p)
}
