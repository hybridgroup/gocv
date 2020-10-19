package gocv

/*
#include <stdlib.h>
#include "features2d.h"
*/
import "C"
import (
	"image/color"
	"reflect"
	"unsafe"
)

// AKAZE is a wrapper around the cv::AKAZE algorithm.
type AKAZE struct {
	// C.AKAZE
	p unsafe.Pointer
}

// NewAKAZE returns a new AKAZE algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d30/classcv_1_1AKAZE.html
//
func NewAKAZE() AKAZE {
	return AKAZE{p: unsafe.Pointer(C.AKAZE_Create())}
}

// Close AKAZE.
func (a *AKAZE) Close() error {
	C.AKAZE_Close((C.AKAZE)(a.p))
	a.p = nil
	return nil
}

// Detect keypoints in an image using AKAZE.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (a *AKAZE) Detect(src Mat) []KeyPoint {
	ret := C.AKAZE_Detect((C.AKAZE)(a.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// DetectAndCompute keypoints and compute in an image using AKAZE.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677
//
func (a *AKAZE) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
	desc := NewMat()
	ret := C.AKAZE_DetectAndCompute((C.AKAZE)(a.p), src.p, mask.p, desc.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret), desc
}

// AgastFeatureDetector is a wrapper around the cv::AgastFeatureDetector.
type AgastFeatureDetector struct {
	// C.AgastFeatureDetector
	p unsafe.Pointer
}

// NewAgastFeatureDetector returns a new AgastFeatureDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d19/classcv_1_1AgastFeatureDetector.html
//
func NewAgastFeatureDetector() AgastFeatureDetector {
	return AgastFeatureDetector{p: unsafe.Pointer(C.AgastFeatureDetector_Create())}
}

// Close AgastFeatureDetector.
func (a *AgastFeatureDetector) Close() error {
	C.AgastFeatureDetector_Close((C.AgastFeatureDetector)(a.p))
	a.p = nil
	return nil
}

// Detect keypoints in an image using AgastFeatureDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (a *AgastFeatureDetector) Detect(src Mat) []KeyPoint {
	ret := C.AgastFeatureDetector_Detect((C.AgastFeatureDetector)(a.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// BRISK is a wrapper around the cv::BRISK algorithm.
type BRISK struct {
	// C.BRISK
	p unsafe.Pointer
}

// NewBRISK returns a new BRISK algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d30/classcv_1_1AKAZE.html
//
func NewBRISK() BRISK {
	return BRISK{p: unsafe.Pointer(C.BRISK_Create())}
}

// Close BRISK.
func (b *BRISK) Close() error {
	C.BRISK_Close((C.BRISK)(b.p))
	b.p = nil
	return nil
}

// Detect keypoints in an image using BRISK.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (b *BRISK) Detect(src Mat) []KeyPoint {
	ret := C.BRISK_Detect((C.BRISK)(b.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// DetectAndCompute keypoints and compute in an image using BRISK.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677
//
func (b *BRISK) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
	desc := NewMat()
	ret := C.BRISK_DetectAndCompute((C.BRISK)(b.p), src.p, mask.p, desc.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret), desc
}

// FastFeatureDetectorType defines the detector type
//
// For further details, please see:
// https://docs.opencv.org/master/df/d74/classcv_1_1FastFeatureDetector.html#a4654f6fb0aa4b8e9123b223bfa0e2a08
type FastFeatureDetectorType int

const (
	//FastFeatureDetectorType58 is an alias of FastFeatureDetector::TYPE_5_8
	FastFeatureDetectorType58 FastFeatureDetectorType = 0
	//FastFeatureDetectorType712 is an alias of FastFeatureDetector::TYPE_7_12
	FastFeatureDetectorType712 FastFeatureDetectorType = 1
	//FastFeatureDetectorType916 is an alias of FastFeatureDetector::TYPE_9_16
	FastFeatureDetectorType916 FastFeatureDetectorType = 2
)

// FastFeatureDetector is a wrapper around the cv::FastFeatureDetector.
type FastFeatureDetector struct {
	// C.FastFeatureDetector
	p unsafe.Pointer
}

// NewFastFeatureDetector returns a new FastFeatureDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/df/d74/classcv_1_1FastFeatureDetector.html
//
func NewFastFeatureDetector() FastFeatureDetector {
	return FastFeatureDetector{p: unsafe.Pointer(C.FastFeatureDetector_Create())}
}

// NewFastFeatureDetectorWithParams returns a new FastFeatureDetector algorithm with parameters
//
// For further details, please see:
// https://docs.opencv.org/master/df/d74/classcv_1_1FastFeatureDetector.html#ab986f2ff8f8778aab1707e2642bc7f8e
//
func NewFastFeatureDetectorWithParams(threshold int, nonmaxSuppression bool, typ FastFeatureDetectorType) FastFeatureDetector {
	return FastFeatureDetector{p: unsafe.Pointer(C.FastFeatureDetector_CreateWithParams(C.int(threshold), C.bool(nonmaxSuppression), C.int(typ)))}
}

// Close FastFeatureDetector.
func (f *FastFeatureDetector) Close() error {
	C.FastFeatureDetector_Close((C.FastFeatureDetector)(f.p))
	f.p = nil
	return nil
}

// Detect keypoints in an image using FastFeatureDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (f *FastFeatureDetector) Detect(src Mat) []KeyPoint {
	ret := C.FastFeatureDetector_Detect((C.FastFeatureDetector)(f.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// GFTTDetector is a wrapper around the cv::GFTTDetector algorithm.
type GFTTDetector struct {
	// C.GFTTDetector
	p unsafe.Pointer
}

// NewGFTTDetector returns a new GFTTDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/df/d21/classcv_1_1GFTTDetector.html
//
func NewGFTTDetector() GFTTDetector {
	return GFTTDetector{p: unsafe.Pointer(C.GFTTDetector_Create())}
}

// Close GFTTDetector.
func (a *GFTTDetector) Close() error {
	C.GFTTDetector_Close((C.GFTTDetector)(a.p))
	a.p = nil
	return nil
}

// Detect keypoints in an image using GFTTDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (a *GFTTDetector) Detect(src Mat) []KeyPoint {
	ret := C.GFTTDetector_Detect((C.GFTTDetector)(a.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// KAZE is a wrapper around the cv::KAZE algorithm.
type KAZE struct {
	// C.KAZE
	p unsafe.Pointer
}

// NewKAZE returns a new KAZE algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d61/classcv_1_1KAZE.html
//
func NewKAZE() KAZE {
	return KAZE{p: unsafe.Pointer(C.KAZE_Create())}
}

// Close KAZE.
func (a *KAZE) Close() error {
	C.KAZE_Close((C.KAZE)(a.p))
	a.p = nil
	return nil
}

// Detect keypoints in an image using KAZE.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (a *KAZE) Detect(src Mat) []KeyPoint {
	ret := C.KAZE_Detect((C.KAZE)(a.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// DetectAndCompute keypoints and compute in an image using KAZE.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677
//
func (a *KAZE) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
	desc := NewMat()
	ret := C.KAZE_DetectAndCompute((C.KAZE)(a.p), src.p, mask.p, desc.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret), desc
}

// MSER is a wrapper around the cv::MSER algorithm.
type MSER struct {
	// C.MSER
	p unsafe.Pointer
}

// NewMSER returns a new MSER algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d28/classcv_1_1MSER.html
//
func NewMSER() MSER {
	return MSER{p: unsafe.Pointer(C.MSER_Create())}
}

// Close MSER.
func (a *MSER) Close() error {
	C.MSER_Close((C.MSER)(a.p))
	a.p = nil
	return nil
}

// Detect keypoints in an image using MSER.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (a *MSER) Detect(src Mat) []KeyPoint {
	ret := C.MSER_Detect((C.MSER)(a.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// ORB is a wrapper around the cv::ORB.
type ORB struct {
	// C.ORB
	p unsafe.Pointer
}

// NewORB returns a new ORB algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d19/classcv_1_1AgastFeatureDetector.html
//
func NewORB() ORB {
	return ORB{p: unsafe.Pointer(C.ORB_Create())}
}

// Close ORB.
func (o *ORB) Close() error {
	C.ORB_Close((C.ORB)(o.p))
	o.p = nil
	return nil
}

// Detect keypoints in an image using ORB.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (o *ORB) Detect(src Mat) []KeyPoint {
	ret := C.ORB_Detect((C.ORB)(o.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// DetectAndCompute detects keypoints and computes from an image using ORB.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677
//
func (o *ORB) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
	desc := NewMat()
	ret := C.ORB_DetectAndCompute((C.ORB)(o.p), src.p, mask.p, desc.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret), desc
}

// SimpleBlobDetector is a wrapper around the cv::SimpleBlobDetector.
type SimpleBlobDetector struct {
	// C.SimpleBlobDetector
	p unsafe.Pointer
}

// SimpleBlobDetector_Params is a wrapper around the cv::SimpleBlobdetector::Params
type SimpleBlobDetectorParams struct {
	p C.SimpleBlobDetectorParams
}

// NewSimpleBlobDetector returns a new SimpleBlobDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d7a/classcv_1_1SimpleBlobDetector.html
//
func NewSimpleBlobDetector() SimpleBlobDetector {
	return SimpleBlobDetector{p: unsafe.Pointer(C.SimpleBlobDetector_Create())}
}

// NewSimpleBlobDetectorWithParams returns a new SimpleBlobDetector with custom parameters
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d7a/classcv_1_1SimpleBlobDetector.html
//
func NewSimpleBlobDetectorWithParams(params SimpleBlobDetectorParams) SimpleBlobDetector {
	return SimpleBlobDetector{p: unsafe.Pointer(C.SimpleBlobDetector_Create_WithParams(params.p))}
}

// Close SimpleBlobDetector.
func (b *SimpleBlobDetector) Close() error {
	C.SimpleBlobDetector_Close((C.SimpleBlobDetector)(b.p))
	b.p = nil
	return nil
}

// NewSimpleBlobDetectorParams returns the default parameters for the SimpleBobDetector
func NewSimpleBlobDetectorParams() SimpleBlobDetectorParams {
	return SimpleBlobDetectorParams{p: C.SimpleBlobDetectorParams_Create()}
}

// SetBlobColor sets the blobColor field
func (p *SimpleBlobDetectorParams) SetBlobColor(blobColor int) {
	p.p.blobColor = C.uchar(blobColor)
}

// GetBlobColor gets the blobColor field
func (p *SimpleBlobDetectorParams) GetBlobColor() int {
	return int(p.p.blobColor)
}

// SetFilterByArea sets the filterByArea field
func (p *SimpleBlobDetectorParams) SetFilterByArea(filterByArea bool) {
	p.p.filterByArea = C.bool(filterByArea)
}

// GetFilterByArea gets the filterByArea field
func (p *SimpleBlobDetectorParams) GetFilterByArea() bool {
	return bool(p.p.filterByArea)
}

// SetFilterByCircularity sets the filterByCircularity field
func (p *SimpleBlobDetectorParams) SetFilterByCircularity(filterByCircularity bool) {
	p.p.filterByCircularity = C.bool(filterByCircularity)
}

// GetFilterByCircularity gets the filterByCircularity field
func (p *SimpleBlobDetectorParams) GetFilterByCircularity() bool {
	return bool(p.p.filterByCircularity)
}

// SetFilterByColor sets the filterByColor field
func (p *SimpleBlobDetectorParams) SetFilterByColor(filterByColor bool) {
	p.p.filterByColor = C.bool(filterByColor)
}

// GetFilterByColor gets the filterByColor field
func (p *SimpleBlobDetectorParams) GetFilterByColor() bool {
	return bool(p.p.filterByColor)
}

// SetFilterByConvexity sets the filterByConvexity field
func (p *SimpleBlobDetectorParams) SetFilterByConvexity(filterByConvexity bool) {
	p.p.filterByConvexity = C.bool(filterByConvexity)
}

// GetFilterByConvexity gets the filterByConvexity field
func (p *SimpleBlobDetectorParams) GetFilterByConvexity() bool {
	return bool(p.p.filterByConvexity)
}

// SetFilterByInertia sets the filterByInertia field
func (p *SimpleBlobDetectorParams) SetFilterByInertia(filterByInertia bool) {
	p.p.filterByInertia = C.bool(filterByInertia)
}

// GetFilterByInertia gets the filterByInertia field
func (p *SimpleBlobDetectorParams) GetFilterByInertia() bool {
	return bool(p.p.filterByInertia)
}

// SetMaxArea sets the maxArea parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMaxArea(maxArea float64) {
	p.p.maxArea = C.float(maxArea)
}

// GetMaxArea sets the maxArea parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMaxArea() float64 {
	return float64(p.p.maxArea)
}

// SetMaxCircularity sets the maxCircularity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMaxCircularity(maxCircularity float64) {
	p.p.maxCircularity = C.float(maxCircularity)
}

// GetMaxCircularity sets the maxCircularity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMaxCircularity() float64 {
	return float64(p.p.maxCircularity)
}

// SetMaxConvexity sets the maxConvexity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMaxConvexity(maxConvexity float64) {
	p.p.maxConvexity = C.float(maxConvexity)
}

// GetMaxConvexity sets the maxConvexity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMaxConvexity() float64 {
	return float64(p.p.maxConvexity)
}

// SetMaxInertiaRatio sets the maxInertiaRatio parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMaxInertiaRatio(maxInertiaRatio float64) {
	p.p.maxInertiaRatio = C.float(maxInertiaRatio)
}

// GetMaxInertiaRatio sets the maxCConvexity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMaxInertiaRatio() float64 {
	return float64(p.p.maxInertiaRatio)
}

// SetMaxThreshold sets the maxThreshold parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMaxThreshold(maxThreshold float64) {
	p.p.maxThreshold = C.float(maxThreshold)
}

// GetMaxThreshold sets the maxCConvexity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMaxThreshold() float64 {
	return float64(p.p.maxThreshold)
}

// SetMinArea sets the minArea parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinArea(minArea float64) {
	p.p.minArea = C.float(minArea)
}

// GetMinArea sets theinArea parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinArea() float64 {
	return float64(p.p.minArea)
}

// SetMinCircularity sets the minCircularity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinCircularity(minCircularity float64) {
	p.p.minCircularity = C.float(minCircularity)
}

// GetMinCircularity sets the minCircularity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinCircularity() float64 {
	return float64(p.p.minCircularity)
}

// SetMinConvexity sets the minConvexity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinConvexity(minConvexity float64) {
	p.p.minConvexity = C.float(minConvexity)
}

// GetMinConvexity sets the minConvexity parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinConvexity() float64 {
	return float64(p.p.minConvexity)
}

// SetMinDistBetweenBlobs sets the minDistBetweenBlobs parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinDistBetweenBlobs(minDistBetweenBlobs float64) {
	p.p.minDistBetweenBlobs = C.float(minDistBetweenBlobs)
}

// GetMinDistBetweenBlobs sets the minDistBetweenBlobs parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinDistBetweenBlobs() float64 {
	return float64(p.p.minDistBetweenBlobs)
}

// SetMinInertiaRatio sets the minInertiaRatio parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinInertiaRatio(minInertiaRatio float64) {
	p.p.minInertiaRatio = C.float(minInertiaRatio)
}

// GetMinInertiaRatio sets the minInertiaRatio parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinInertiaRatio() float64 {
	return float64(p.p.minInertiaRatio)
}

// SetMinRepeatability sets the minRepeatability parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinRepeatability(minRepeatability int) {
	p.p.minRepeatability = C.size_t(minRepeatability)
}

// GetMinInertiaRatio sets the minRepeatability parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinRepeatability() int {
	return int(p.p.minRepeatability)
}

// SetMinThreshold sets the minThreshold parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetMinThreshold(minThreshold float64) {
	p.p.minThreshold = C.float(minThreshold)
}

// GetMinThreshold sets the minInertiaRatio parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetMinThreshold() float64 {
	return float64(p.p.minThreshold)
}

// SetMinThreshold sets the minThreshold parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) SetThresholdStep(thresholdStep float64) {
	p.p.thresholdStep = C.float(thresholdStep)
}

// GetMinThreshold sets the minInertiaRatio parameter for SimpleBlobDetector_Params
func (p *SimpleBlobDetectorParams) GetThresholdStep() float64 {
	return float64(p.p.thresholdStep)
}

// Detect keypoints in an image using SimpleBlobDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (b *SimpleBlobDetector) Detect(src Mat) []KeyPoint {
	ret := C.SimpleBlobDetector_Detect((C.SimpleBlobDetector)(b.p), src.p)
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// getKeyPoints returns a slice of KeyPoint given a pointer to a C.KeyPoints
func getKeyPoints(ret C.KeyPoints) []KeyPoint {
	cArray := ret.keypoints
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.KeyPoint)(unsafe.Pointer(&hdr))

	keys := make([]KeyPoint, length)
	for i, r := range s {
		keys[i] = KeyPoint{float64(r.x), float64(r.y), float64(r.size), float64(r.angle), float64(r.response),
			int(r.octave), int(r.classID)}
	}
	return keys
}

// BFMatcher is a wrapper around the the cv::BFMatcher algorithm
type BFMatcher struct {
	// C.BFMatcher
	p unsafe.Pointer
}

// NewBFMatcher returns a new BFMatcher
//
// For further details, please see:
// https://docs.opencv.org/master/d3/da1/classcv_1_1BFMatcher.html#abe0bb11749b30d97f60d6ade665617bd
//
func NewBFMatcher() BFMatcher {
	return BFMatcher{p: unsafe.Pointer(C.BFMatcher_Create())}
}

// NewBFMatcherWithParams creates a new BFMatchers but allows setting parameters
// to values other than just the defaults.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/da1/classcv_1_1BFMatcher.html#abe0bb11749b30d97f60d6ade665617bd
//
func NewBFMatcherWithParams(normType NormType, crossCheck bool) BFMatcher {
	return BFMatcher{p: unsafe.Pointer(C.BFMatcher_CreateWithParams(C.int(normType), C.bool(crossCheck)))}
}

// Close BFMatcher
func (b *BFMatcher) Close() error {
	C.BFMatcher_Close((C.BFMatcher)(b.p))
	b.p = nil
	return nil
}

// KnnMatch Finds the k best matches for each descriptor from a query set.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d39/classcv_1_1DescriptorMatcher.html#aa880f9353cdf185ccf3013e08210483a
//
func (b *BFMatcher) KnnMatch(query, train Mat, k int) [][]DMatch {
	ret := C.BFMatcher_KnnMatch((C.BFMatcher)(b.p), query.p, train.p, C.int(k))
	defer C.MultiDMatches_Close(ret)

	return getMultiDMatches(ret)
}

// FlannBasedMatcher is a wrapper around the the cv::FlannBasedMatcher algorithm
type FlannBasedMatcher struct {
	// C.FlannBasedMatcher
	p unsafe.Pointer
}

// NewFlannBasedMatcher returns a new FlannBasedMatcher
//
// For further details, please see:
// https://docs.opencv.org/master/dc/de2/classcv_1_1FlannBasedMatcher.html#ab9114a6471e364ad221f89068ca21382
//
func NewFlannBasedMatcher() FlannBasedMatcher {
	return FlannBasedMatcher{p: unsafe.Pointer(C.FlannBasedMatcher_Create())}
}

// Close FlannBasedMatcher
func (f *FlannBasedMatcher) Close() error {
	C.FlannBasedMatcher_Close((C.FlannBasedMatcher)(f.p))
	f.p = nil
	return nil
}

// KnnMatch Finds the k best matches for each descriptor from a query set.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d39/classcv_1_1DescriptorMatcher.html#aa880f9353cdf185ccf3013e08210483a
//
func (f *FlannBasedMatcher) KnnMatch(query, train Mat, k int) [][]DMatch {
	ret := C.FlannBasedMatcher_KnnMatch((C.FlannBasedMatcher)(f.p), query.p, train.p, C.int(k))
	defer C.MultiDMatches_Close(ret)

	return getMultiDMatches(ret)
}

func getMultiDMatches(ret C.MultiDMatches) [][]DMatch {
	cArray := ret.dmatches
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.DMatches)(unsafe.Pointer(&hdr))

	keys := make([][]DMatch, length)
	for i := range s {
		keys[i] = getDMatches(C.MultiDMatches_get(ret, C.int(i)))
	}
	return keys
}

func getDMatches(ret C.DMatches) []DMatch {
	cArray := ret.dmatches
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.DMatch)(unsafe.Pointer(&hdr))

	keys := make([]DMatch, length)
	for i, r := range s {
		keys[i] = DMatch{int(r.queryIdx), int(r.trainIdx), int(r.imgIdx),
			float64(r.distance)}
	}
	return keys
}

// DrawMatchesFlag are the flags setting drawing feature
//
// For further details please see:
// https://docs.opencv.org/master/de/d30/structcv_1_1DrawMatchesFlags.html
type DrawMatchesFlag int

const (
	// DrawDefault creates new image and for each keypoint only the center point will be drawn
	DrawDefault DrawMatchesFlag = 0
	// DrawOverOutImg draws matches on existing content of image
	DrawOverOutImg DrawMatchesFlag = 1
	// NotDrawSinglePoints will not draw single points
	NotDrawSinglePoints DrawMatchesFlag = 2
	// DrawRichKeyPoints draws the circle around each keypoint with keypoint size and orientation
	DrawRichKeyPoints DrawMatchesFlag = 3
)

// DrawKeyPoints draws keypoints
//
// For further details please see:
// https://docs.opencv.org/master/d4/d5d/group__features2d__draw.html#gab958f8900dd10f14316521c149a60433
func DrawKeyPoints(src Mat, keyPoints []KeyPoint, dst *Mat, color color.RGBA, flag DrawMatchesFlag) {
	cKeyPointArray := make([]C.struct_KeyPoint, len(keyPoints))

	for i, kp := range keyPoints {
		cKeyPointArray[i].x = C.double(kp.X)
		cKeyPointArray[i].y = C.double(kp.Y)
		cKeyPointArray[i].size = C.double(kp.Size)
		cKeyPointArray[i].angle = C.double(kp.Angle)
		cKeyPointArray[i].response = C.double(kp.Response)
		cKeyPointArray[i].octave = C.int(kp.Octave)
		cKeyPointArray[i].classID = C.int(kp.ClassID)
	}

	cKeyPoints := C.struct_KeyPoints{
		keypoints: (*C.struct_KeyPoint)(&cKeyPointArray[0]),
		length:    (C.int)(len(keyPoints)),
	}

	scalar := C.struct_Scalar{
		val1: C.double(color.B),
		val2: C.double(color.G),
		val3: C.double(color.R),
		val4: C.double(color.A),
	}

	C.DrawKeyPoints(src.p, cKeyPoints, dst.p, scalar, C.int(flag))
}

// SIFT is a wrapper around the cv::SIFT algorithm.
// Due to the patent having expired, this is now in the main OpenCV code modules.
type SIFT struct {
	// C.SIFT
	p unsafe.Pointer
}

// NewSIFT returns a new SIFT algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d3c/classcv_1_1xfeatures2d_1_1SIFT.html
//
func NewSIFT() SIFT {
	return SIFT{p: unsafe.Pointer(C.SIFT_Create())}
}

// Close SIFT.
func (d *SIFT) Close() error {
	C.SIFT_Close((C.SIFT)(d.p))
	d.p = nil
	return nil
}

// Detect keypoints in an image using SIFT.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (d *SIFT) Detect(src Mat) []KeyPoint {
	ret := C.SIFT_Detect((C.SIFT)(d.p), C.Mat(src.Ptr()))
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret)
}

// DetectAndCompute detects and computes keypoints in an image using SIFT.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677
//
func (d *SIFT) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
	desc := NewMat()
	ret := C.SIFT_DetectAndCompute((C.SIFT)(d.p), C.Mat(src.Ptr()), C.Mat(mask.Ptr()),
		C.Mat(desc.Ptr()))
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret), desc
}

// DrawMatches draws matches on combined train and querry images.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d5d/group__features2d__draw.html#gad8f463ccaf0dc6f61083abd8717c261a
func DrawMatches(img1 Mat, kp1 []KeyPoint, img2 Mat, kp2 []KeyPoint, matches1to2 []DMatch, outImg *Mat, matchColor color.RGBA, singlePointColor color.RGBA, matchesMask []byte, flags DrawMatchesFlag) {
	kp1arr := make([]C.struct_KeyPoint, len(kp1))
	kp2arr := make([]C.struct_KeyPoint, len(kp2))

	for i, kp := range kp1 {
		kp1arr[i].x = C.double(kp.X)
		kp1arr[i].y = C.double(kp.Y)
		kp1arr[i].size = C.double(kp.Size)
		kp1arr[i].angle = C.double(kp.Angle)
		kp1arr[i].response = C.double(kp.Response)
		kp1arr[i].octave = C.int(kp.Octave)
		kp1arr[i].classID = C.int(kp.ClassID)
	}

	for i, kp := range kp2 {
		kp2arr[i].x = C.double(kp.X)
		kp2arr[i].y = C.double(kp.Y)
		kp2arr[i].size = C.double(kp.Size)
		kp2arr[i].angle = C.double(kp.Angle)
		kp2arr[i].response = C.double(kp.Response)
		kp2arr[i].octave = C.int(kp.Octave)
		kp2arr[i].classID = C.int(kp.ClassID)
	}

	cKeyPoints1 := C.struct_KeyPoints{
		keypoints: (*C.struct_KeyPoint)(&kp1arr[0]),
		length:    (C.int)(len(kp1)),
	}

	cKeyPoints2 := C.struct_KeyPoints{
		keypoints: (*C.struct_KeyPoint)(&kp2arr[0]),
		length:    (C.int)(len(kp2)),
	}

	dMatchArr := make([]C.struct_DMatch, len(matches1to2))

	for i, dm := range matches1to2 {
		dMatchArr[i].queryIdx = C.int(dm.QueryIdx)
		dMatchArr[i].trainIdx = C.int(dm.TrainIdx)
		dMatchArr[i].imgIdx = C.int(dm.ImgIdx)
		dMatchArr[i].distance = C.float(dm.Distance)
	}

	cDMatches := C.struct_DMatches{
		dmatches: (*C.struct_DMatch)(&dMatchArr[0]),
		length:   (C.int)(len(matches1to2)),
	}

	scalarMatchColor := C.struct_Scalar{
		val1: C.double(matchColor.R),
		val2: C.double(matchColor.G),
		val3: C.double(matchColor.B),
		val4: C.double(matchColor.A),
	}

	scalarPointColor := C.struct_Scalar{
		val1: C.double(singlePointColor.B),
		val2: C.double(singlePointColor.G),
		val3: C.double(singlePointColor.R),
		val4: C.double(singlePointColor.A),
	}

	mask := make([]C.char, len(matchesMask))

	cByteArray := C.struct_ByteArray{
		length: (C.int)(len(matchesMask)),
	}

	if len(matchesMask) > 0 {
		cByteArray = C.struct_ByteArray{
			data:   (*C.char)(&mask[0]),
			length: (C.int)(len(matchesMask)),
		}
	}

	C.DrawMatches(img1.p, cKeyPoints1, img2.p, cKeyPoints2, cDMatches, outImg.p, scalarMatchColor, scalarPointColor, cByteArray, C.int(flags))
}
