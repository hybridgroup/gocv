package gocv

/*
#include <stdlib.h>
#include "objdetect.h"
*/
import "C"
import (
	"image"
	"unsafe"
)

// CascadeClassifier is a cascade classifier class for object detection.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html
type CascadeClassifier struct {
	p C.CascadeClassifier
}

// NewCascadeClassifier returns a new CascadeClassifier.
func NewCascadeClassifier() CascadeClassifier {
	return CascadeClassifier{p: C.CascadeClassifier_New()}
}

// Close deletes the CascadeClassifier's pointer.
func (c *CascadeClassifier) Close() error {
	C.CascadeClassifier_Close(c.p)
	c.p = nil
	return nil
}

// Load cascade classifier from a file.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#a1a5884c8cc749422f9eb77c2471958bc
func (c *CascadeClassifier) Load(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.CascadeClassifier_Load(c.p, cName) != 0
}

// DetectMultiScale detects objects of different sizes in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
func (c *CascadeClassifier) DetectMultiScale(img Mat) []image.Rectangle {
	ret := C.CascadeClassifier_DetectMultiScale(c.p, img.p)
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// DetectMultiScaleWithParams calls DetectMultiScale but allows setting parameters
// to values other than just the defaults.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
func (c *CascadeClassifier) DetectMultiScaleWithParams(img Mat, scale float64,
	minNeighbors, flags int, minSize, maxSize image.Point) []image.Rectangle {

	minSz := C.struct_Size{
		width:  C.int(minSize.X),
		height: C.int(minSize.Y),
	}

	maxSz := C.struct_Size{
		width:  C.int(maxSize.X),
		height: C.int(maxSize.Y),
	}

	ret := C.CascadeClassifier_DetectMultiScaleWithParams(c.p, img.p, C.double(scale),
		C.int(minNeighbors), C.int(flags), minSz, maxSz)
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// HOGDescriptor is a Histogram Of Gradiants (HOG) for object detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a723b95b709cfd3f95cf9e616de988fc8
type HOGDescriptor struct {
	p C.HOGDescriptor
}

// NewHOGDescriptor returns a new HOGDescriptor.
func NewHOGDescriptor() HOGDescriptor {
	return HOGDescriptor{p: C.HOGDescriptor_New()}
}

// Close deletes the HOGDescriptor's pointer.
func (h *HOGDescriptor) Close() error {
	C.HOGDescriptor_Close(h.p)
	h.p = nil
	return nil
}

// DetectMultiScale detects objects in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
func (h *HOGDescriptor) DetectMultiScale(img Mat) []image.Rectangle {
	ret := C.HOGDescriptor_DetectMultiScale(h.p, img.p)
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// DetectMultiScaleWithParams calls DetectMultiScale but allows setting parameters
// to values other than just the defaults.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
func (h *HOGDescriptor) DetectMultiScaleWithParams(img Mat, hitThresh float64,
	winStride, padding image.Point, scale, finalThreshold float64, useMeanshiftGrouping bool) []image.Rectangle {
	wSz := C.struct_Size{
		width:  C.int(winStride.X),
		height: C.int(winStride.Y),
	}

	pSz := C.struct_Size{
		width:  C.int(padding.X),
		height: C.int(padding.Y),
	}

	ret := C.HOGDescriptor_DetectMultiScaleWithParams(h.p, img.p, C.double(hitThresh),
		wSz, pSz, C.double(scale), C.double(finalThreshold), C.bool(useMeanshiftGrouping))
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// HOGDefaultPeopleDetector returns a new Mat with the HOG DefaultPeopleDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
func HOGDefaultPeopleDetector() Mat {
	return newMat(C.HOG_GetDefaultPeopleDetector())
}

// SetSVMDetector sets the data for the HOGDescriptor.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a09e354ad701f56f9c550dc0385dc36f1
func (h *HOGDescriptor) SetSVMDetector(det Mat) error {
	C.HOGDescriptor_SetSVMDetector(h.p, det.p)
	return nil
}

// GroupRectangles groups the object candidate rectangles.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d54/group__objdetect.html#ga3dba897ade8aa8227edda66508e16ab9
func GroupRectangles(rects []image.Rectangle, groupThreshold int, eps float64) []image.Rectangle {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.Min.X),
			y:      C.int(r.Min.Y),
			width:  C.int(r.Size().X),
			height: C.int(r.Size().Y),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}

	ret := C.GroupRectangles(cRects, C.int(groupThreshold), C.double(eps))

	return toRectangles(ret)
}

// QRCodeDetector groups the object candidate rectangles.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html
type QRCodeDetector struct {
	p C.QRCodeDetector
}

// newQRCodeDetector returns a new QRCodeDetector from a C QRCodeDetector
func newQRCodeDetector(p C.QRCodeDetector) QRCodeDetector {
	return QRCodeDetector{p: p}
}

func NewQRCodeDetector() QRCodeDetector {
	return newQRCodeDetector(C.QRCodeDetector_New())
}

func (a *QRCodeDetector) Close() error {
	C.QRCodeDetector_Close(a.p)
	a.p = nil
	return nil
}

// DetectAndDecode Both detects and decodes QR code.
//
// Returns true as long as some QR code was detected even in case where the decoding failed
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a7290bd6a5d59b14a37979c3a14fbf394
func (a *QRCodeDetector) DetectAndDecode(input Mat, points *Mat, straight_qrcode *Mat) string {
	goResult := C.GoString(C.QRCodeDetector_DetectAndDecode(a.p, input.p, points.p, straight_qrcode.p))
	return string(goResult)
}

// Detect detects QR code in image and returns the quadrangle containing the code.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a64373f7d877d27473f64fe04bb57d22b
func (a *QRCodeDetector) Detect(input Mat, points *Mat) bool {
	result := C.QRCodeDetector_Detect(a.p, input.p, points.p)
	return bool(result)
}

// Decode decodes QR code in image once it's found by the detect() method. Returns UTF8-encoded output string or empty string if the code cannot be decoded.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a4172c2eb4825c844fb1b0ae67202d329
func (a *QRCodeDetector) Decode(input Mat, points Mat, straight_qrcode *Mat) string {
	goResult := C.GoString(C.QRCodeDetector_DetectAndDecode(a.p, input.p, points.p, straight_qrcode.p))
	return string(goResult)
}

// Detects QR codes in image and finds of the quadrangles containing the codes.
//
// Each quadrangle would be returned as a row in the `points` Mat and each point is a Vecf.
// Returns true if QR code was detected
// For usage please see TestQRCodeDetector
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#aaf2b6b2115b8e8fbc9acf3a8f68872b6
func (a *QRCodeDetector) DetectMulti(input Mat, points *Mat) bool {
	result := C.QRCodeDetector_DetectMulti(a.p, input.p, points.p)
	return bool(result)
}

// Detects QR codes in image, finds the quadrangles containing the codes, and decodes the QRCodes to strings.
//
// Each quadrangle would be returned as a row in the `points` Mat and each point is a Vecf.
// Returns true as long as some QR code was detected even in case where the decoding failed
// For usage please see TestQRCodeDetector
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a188b63ffa17922b2c65d8a0ab7b70775
func (a *QRCodeDetector) DetectAndDecodeMulti(input Mat, decoded []string, points *Mat, qrCodes []Mat) bool {
	cDecoded := C.CStrings{}
	defer C.CStrings_Close(cDecoded)
	cQrCodes := C.struct_Mats{}
	defer C.Mats_Close(cQrCodes)
	success := C.QRCodeDetector_DetectAndDecodeMulti(a.p, input.p, &cDecoded, points.p, &cQrCodes)
	if !success {
		return bool(success)
	}

	tmpCodes := make([]Mat, cQrCodes.length)
	for i := C.int(0); i < cQrCodes.length; i++ {
		tmpCodes[i].p = C.Mats_get(cQrCodes, i)
	}

	for _, qr := range tmpCodes {
		qrCodes = append(qrCodes, qr)
	}

	for _, s := range toGoStrings(cDecoded) {
		decoded = append(decoded, s)
	}
	return bool(success)
}

type FaceDetectorYN struct {
	p C.FaceDetectorYN
}

// NewFaceDetectorYN Creates an instance of face detector with given parameters.
//
// modelPath: the path to the requested model
//
// configPath: the path to the config file for compability, which is not requested for ONNX models
//
// size: the size of the input image
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516
func NewFaceDetectorYN(modelPath string, configPath string, size image.Point) FaceDetectorYN {

	c_model_path := C.CString(modelPath)
	defer C.free(unsafe.Pointer(c_model_path))

	c_config_path := C.CString(configPath)
	defer C.free(unsafe.Pointer(c_config_path))

	c_size := C.Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}

	return FaceDetectorYN{p: C.FaceDetectorYN_Create(c_model_path, c_config_path, c_size)}
}

// NewFaceDetectorYNWithParams Creates an instance of face detector with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516
func NewFaceDetectorYNWithParams(modelPath string, configPath string, size image.Point, scoreThreshold float32, nmsThreshold float32, topK int, backendId int, targetId int) FaceDetectorYN {

	c_model_path := C.CString(modelPath)
	defer C.free(unsafe.Pointer(c_model_path))

	c_config_path := C.CString(configPath)
	defer C.free(unsafe.Pointer(c_config_path))

	c_size := C.Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}

	return FaceDetectorYN{p: C.FaceDetectorYN_Create_WithParams(c_model_path, c_config_path, c_size, C.float(scoreThreshold), C.float(nmsThreshold), C.int(topK), C.int(backendId), C.int(targetId))}
}

// NewFaceDetectorYNFromBytes Creates an instance of face detector with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#aa0796a4bfe2d4709bef81abbae9a927a
func NewFaceDetectorYNFromBytes(framework string, bufferModel []byte, bufferConfig []byte, size image.Point) FaceDetectorYN {

	c_framework := C.CString(framework)
	defer C.free(unsafe.Pointer(c_framework))

	c_size := C.Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}

	return FaceDetectorYN{p: C.FaceDetectorYN_Create_FromBytes(c_framework,
		unsafe.Pointer(unsafe.SliceData(bufferModel)), C.int(len(bufferModel)),
		unsafe.Pointer(unsafe.SliceData(bufferConfig)), C.int(len(bufferConfig)), c_size)}
}

// NewFaceDetectorYNFromBuffers Creates an instance of face detector with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#aa0796a4bfe2d4709bef81abbae9a927a
func NewFaceDetectorYNFromBytesWithParams(framework string, bufferModel []byte, bufferConfig []byte, size image.Point, scoreThreshold float32, nmsThreshold float32, topK int, backendId int, targetId int) FaceDetectorYN {

	c_framework := C.CString(framework)
	defer C.free(unsafe.Pointer(c_framework))

	c_size := C.Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}

	return FaceDetectorYN{p: C.FaceDetectorYN_Create_FromBytes_WithParams(c_framework,
		unsafe.Pointer(unsafe.SliceData(bufferModel)), C.int(len(bufferModel)),
		unsafe.Pointer(unsafe.SliceData(bufferConfig)), C.int(len(bufferConfig)), c_size,
		C.float(scoreThreshold), C.float(nmsThreshold), C.int(topK), C.int(backendId), C.int(targetId))}
}

func (fd *FaceDetectorYN) Close() {
	C.FaceDetectorYN_Close(fd.p)
}

// Detect Detects faces in the input image.
//
// image: an image to detect
//
// faces: detection results stored in a 2D cv::Mat of shape [num_faces, 15]
//
// 0-1: x, y of bbox top left corner
//
// 2-3: width, height of bbox
//
// 4-5: x, y of right eye (blue point in the example image)
//
// 6-7: x, y of left eye (red point in the example image)
//
// 8-9: x, y of nose tip (green point in the example image)
//
// 10-11: x, y of right corner of mouth (pink point in the example image)
//
// 12-13: x, y of left corner of mouth (yellow point in the example image)
//
// 14: face score
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#ac05bd075ca3e6edc0e328927aae6f45b
func (fd *FaceDetectorYN) Detect(image Mat, faces *Mat) int {
	c_rv := C.FaceDetectorYN_Detect(fd.p, image.p, faces.p)
	return int(c_rv)
}

func (fd *FaceDetectorYN) GetInputSize() image.Point {
	sz := C.FaceDetectorYN_GetInputSize(fd.p)

	return image.Pt(int(sz.width), int(sz.height))
}

func (fd *FaceDetectorYN) GetNMSThreshold() float32 {
	t := C.FaceDetectorYN_GetNMSThreshold(fd.p)
	return float32(t)
}

func (fd *FaceDetectorYN) GetScoreThreshold() float32 {
	t := C.FaceDetectorYN_GetScoreThreshold(fd.p)
	return float32(t)
}

func (fd *FaceDetectorYN) GetTopK() int {
	i := C.FaceDetectorYN_GetTopK(fd.p)
	return int(i)
}

// SetInputSize Set the size for the network input, which overwrites the input size of creating model.
// Call this method when the size of input image does not match the input size when creating model.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a072418e5ce7beeb69c41edda75c41d2e
func (fd *FaceDetectorYN) SetInputSize(sz image.Point) {
	c_sz := C.Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}
	C.FaceDetectorYN_SetInputSize(fd.p, c_sz)
}

// SetNMSThreshold Set the Non-maximum-suppression threshold to suppress
// bounding boxes that have IoU greater than the given value.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#ab6011efee7e12dca3857d82de5269ac5
func (fd *FaceDetectorYN) SetNMSThreshold(nmsThreshold float32) {
	C.FaceDetectorYN_SetNMSThreshold(fd.p, C.float(nmsThreshold))
}

// SetScoreThreshold Set the score threshold to filter out bounding boxes of score less than the given value.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a37f3c23b82158fac7fdad967d315f85a
func (fd *FaceDetectorYN) SetScoreThreshold(scoreThreshold float32) {
	C.FaceDetectorYN_SetScoreThreshold(fd.p, C.float(scoreThreshold))
}

// SetTopK Set the number of bounding boxes preserved before NMS.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#aa88d20e1e2df75ea36b851534089856a
func (fd *FaceDetectorYN) SetTopK(topK int) {
	C.FaceDetectorYN_SetTopK(fd.p, C.int(topK))
}

type FaceRecognizerSFDisType int

const (
	FaceRecognizerSFDisTypeCosine FaceRecognizerSFDisType = 0
	FaceRecognizerSFDisTypeNormL2 FaceRecognizerSFDisType = 1
)

type FaceRecognizerSF struct {
	p C.FaceRecognizerSF
}

// NewFaceRecognizerSF Creates an instance with given parameters.
//
// model: the path of the onnx model used for face recognition
//
// config: the path to the config file for compability, which is not requested for ONNX models
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a04df90b0cd7d26d350acd92621a35743
func NewFaceRecognizerSF(modelPath string, configPath string) FaceRecognizerSF {
	c_model := C.CString(modelPath)
	defer C.free(unsafe.Pointer(c_model))

	c_config := C.CString(configPath)
	defer C.free(unsafe.Pointer(c_config))

	return FaceRecognizerSF{p: C.FaceRecognizerSF_Create(c_model, c_config)}
}

// NewFaceRecognizerSFWithParams Creates an instance with given parameters.
//
// model: the path of the onnx model used for face recognition
//
// config: the path to the config file for compability, which is not requested for ONNX models
//
// backend_id: the id of backend
//
// target_id: the id of target device
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a04df90b0cd7d26d350acd92621a35743
func NewFaceRecognizerSFWithParams(modelPath string, configPath string, backendId int, targetId int) FaceRecognizerSF {
	c_model := C.CString(modelPath)
	defer C.free(unsafe.Pointer(c_model))

	c_config := C.CString(configPath)
	defer C.free(unsafe.Pointer(c_config))

	return FaceRecognizerSF{p: C.FaceRecognizerSF_Create_WithParams(c_model, c_config, C.int(backendId), C.int(targetId))}
}

// Close Releases FaceRecognizerSF resources.
func (fr *FaceRecognizerSF) Close() {
	C.FaceRecognizerSF_Close(fr.p)
}

// AlignCrop Aligns detected face with the source input image and crops it.
//
// srcImg: input image
//
// faceBox: the detected face result from the input image
//
// alignedImg: output aligned image
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a84492908abecbc9362b4ddc8d46b8345
func (fr *FaceRecognizerSF) AlignCrop(srcImg Mat, faceBox Mat, alignedImg *Mat) {
	C.FaceRecognizerSF_AlignCrop(fr.p, srcImg.p, faceBox.p, alignedImg.p)
}

// Feature Extracts face feature from aligned image.
//
// alignedImg: input aligned image
//
// faceFeature: output face feature
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#ab1b4a3c12213e89091a490c573dc5aba
func (fr *FaceRecognizerSF) Feature(alignedImg Mat, faceFeature *Mat) {
	C.FaceRecognizerSF_Feature(fr.p, alignedImg.p, faceFeature.p)
}

// Match Calculates the distance between two face features.
//
// faceFeature1: the first input feature
//
// faceFeature2: the second input feature of the same size and the same type as face_feature1
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a2f0362ca1e64320a1f3ba7e1386d0219
func (fr *FaceRecognizerSF) Match(faceFeature1 Mat, faceFeature2 Mat) float32 {
	rv := C.FaceRecognizerSF_Match(fr.p, faceFeature1.p, faceFeature2.p)
	return float32(rv)
}

// MatchWithParams Calculates the distance between two face features.
//
// faceFeature1: the first input feature
//
// faceFeature2: the second input feature of the same size and the same type as face_feature1
//
// disType: defines how to calculate the distance between two face features
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a2f0362ca1e64320a1f3ba7e1386d0219
func (fr *FaceRecognizerSF) MatchWithParams(faceFeature1 Mat, faceFeature2 Mat, disType FaceRecognizerSFDisType) float32 {
	rv := C.FaceRecognizerSF_Match_WithParams(fr.p, faceFeature1.p, faceFeature2.p, C.int(disType))
	return float32(rv)
}
