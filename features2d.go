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
	FastFeatureDetectorType712 = 1
	//FastFeatureDetectorType916 is an alias of FastFeatureDetector::TYPE_9_16
	FastFeatureDetectorType916 = 2
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
	DrawOverOutImg = 1
	// NotDrawSinglePoints will not draw single points
	NotDrawSinglePoints = 2
	// DrawRichKeyPoints draws the circle around each keypoint with keypoint size and orientation
	DrawRichKeyPoints = 3
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
		val1: C.double(color.R),
		val2: C.double(color.G),
		val3: C.double(color.B),
		val4: C.double(color.A),
	}

	C.DrawKeyPoints(src.p, cKeyPoints, dst.p, scalar, C.int(flag))
}
