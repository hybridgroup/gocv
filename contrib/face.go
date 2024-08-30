package contrib

/*
#include <stdlib.h>
#include "face.h"
*/
import "C"
import (
	"image"

	"gocv.io/x/gocv"
)

// PredictResponse represents a predicted label and associated confidence.
type PredictResponse struct {
	Label      int32   `json:"label"`
	Confidence float32 `json:"confidence"`
}

var _ FaceRecognizer = (*LBPHFaceRecognizer)(nil)

// LBPHFaceRecognizer is a wrapper for the OpenCV Local Binary Patterns
// Histograms face recognizer.
type LBPHFaceRecognizer struct {
	p C.LBPHFaceRecognizer
}

func (fr *LBPHFaceRecognizer) Empty() bool {
	return faceRecognizer_Empty(C.FaceRecognizer(fr.p))
}

// NewLBPHFaceRecognizer creates a new LBPH Recognizer model.
//
// For further information, see:
// https://docs.opencv.org/master/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html
func NewLBPHFaceRecognizer() *LBPHFaceRecognizer {
	return &LBPHFaceRecognizer{p: C.CreateLBPHFaceRecognizer()}
}

// Train loaded model with images and their labels
//
// see https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#ac8680c2aa9649ad3f55e27761165c0d6
func (fr *LBPHFaceRecognizer) Train(images []gocv.Mat, labels []int) {
	faceRecognizer_Train(C.FaceRecognizer(fr.p), images, labels)
}

// Update updates the existing trained model with new images and labels.
//
// For further information, see:
// https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#a8a4e73ea878dcd0c235d0487189d25f3
func (fr *LBPHFaceRecognizer) Update(newImages []gocv.Mat, newLabels []int) {
	faceRecognizer_Update(C.FaceRecognizer(fr.p), newImages, newLabels)
}

// Predict predicts a label for a given input image. It returns the label for
// correctly predicted image or -1 if not found.
//
// For further information, see:
// https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#aa2d2f02faffab1bf01317ae6502fb631
func (fr *LBPHFaceRecognizer) Predict(sample gocv.Mat) int {
	label := faceRecognizer_Predict(C.FaceRecognizer(fr.p), sample)

	return int(label)
}

// PredictExtendedResponse returns a label and associated confidence (e.g.
// distance) for a given input image. It is the extended version of
// `Predict()`.
//
// For further information, see:
// https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#ab0d593e53ebd9a0f350c989fcac7f251
func (fr *LBPHFaceRecognizer) PredictExtendedResponse(sample gocv.Mat) PredictResponse {
	resp := faceRecognizer_PredictExtendedResponse(C.FaceRecognizer(fr.p), sample)

	return resp
}

// GetThreshold gets the threshold value of the model, i.e. the threshold
// applied in the prediction.
//
// For further information, see:
// https://docs.opencv.org/4.x/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#acf2a6993eb4347b3f89009da693a3f70
func (fr *LBPHFaceRecognizer) GetThreshold() float32 {
	t := faceRecognizer_GetThreshold(C.FaceRecognizer(fr.p))
	return float32(t)
}

// SetThreshold sets the threshold value of the model, i.e. the threshold
// applied in the prediction.
//
// For further information, see:
// https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#a3182081e5f8023e658ad8ab96656dd63
func (fr *LBPHFaceRecognizer) SetThreshold(threshold float32) {
	faceRecognizer_SetThreshold(C.FaceRecognizer(fr.p), threshold)
}

// SetNeighbors sets the neighbors value of the model, i.e. the number of
// sample points to build a Circular Local Binary Pattern from. Note that wrong
// neighbors can raise OpenCV exception!
//
// For further information, see:
// https://docs.opencv.org/master/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#ab225f7bf353ce8697a506eda10124a92
func (fr *LBPHFaceRecognizer) SetNeighbors(neighbors int) {
	C.LBPHFaceRecognizer_SetNeighbors(fr.p, C.int(neighbors))
}

// GetNeighbors returns the neighbors value of the model.
//
// For further information, see:
// https://docs.opencv.org/master/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#a50a3e2ca6e8464166e153c9df84b0a77
func (fr *LBPHFaceRecognizer) GetNeighbors() int {

	n := C.LBPHFaceRecognizer_GetNeighbors(fr.p)

	return int(n)
}

// SetRadius sets the radius used for building the Circular Local Binary
// Pattern.
//
// For further information, see:
// https://docs.opencv.org/master/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#a62d94c75cade902fd3b487b1ef9883fc
func (fr *LBPHFaceRecognizer) SetRadius(radius int) {
	C.LBPHFaceRecognizer_SetRadius(fr.p, C.int(radius))
}

// SaveFile saves the trained model data to file.
//
// For further information, see:
// https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#a2adf2d555550194244b05c91fefcb4d6
func (fr *LBPHFaceRecognizer) SaveFile(fname string) {
	faceRecognizer_SaveFile(C.FaceRecognizer(fr.p), fname)
}

// LoadFile loads a trained model data from file.
//
// For further information, see:
// https://docs.opencv.org/master/dd/d65/classcv_1_1face_1_1FaceRecognizer.html#acc42e5b04595dba71f0777c7179af8c3
func (fr *LBPHFaceRecognizer) LoadFile(fname string) {
	faceRecognizer_LoadFile(C.FaceRecognizer(fr.p), fname)
}

// SetGridX sets grid's X value
//
// For further information, see:
// https://docs.opencv.org/4.x/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#ad65975baee31dbf3bd2a750feef74831
func (fr *LBPHFaceRecognizer) SetGridX(x int) {
	C.LBPHFaceRecognizer_SetGridX(fr.p, C.int(x))
}

// SetGridY sets grid's Y value
//
// For further information, see:
// https://docs.opencv.org/4.x/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#a9cebb0138dbb3553b27beb2df3924ae6
func (fr *LBPHFaceRecognizer) SetGridY(y int) {
	C.LBPHFaceRecognizer_SetGridY(fr.p, C.int(y))
}

// GetGridX gets grid's X value
//
// For further information, see:
// https://docs.opencv.org/4.x/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#ada6839bed931a8f68c5127e1af7a8b83
func (fr *LBPHFaceRecognizer) GetGridX() int {
	x := C.LBPHFaceRecognizer_GetGridX(fr.p)
	return int(x)
}

// GetGridY gets grid's Y value
//
// For further information, see:
// https://docs.opencv.org/4.x/df/d25/classcv_1_1face_1_1LBPHFaceRecognizer.html#a22c68c0baf3eb9e852f47ae9241dbf15
func (fr *LBPHFaceRecognizer) GetGridY() int {
	y := C.LBPHFaceRecognizer_GetGridY(fr.p)
	return int(y)
}

// SetGrid helper for SetGrid* functions
func (fr *LBPHFaceRecognizer) SetGrid(p image.Point) {
	fr.SetGridX(p.X)
	fr.SetGridY(p.Y)
}

// GetGrid helper for GetGrid* functions
func (fr *LBPHFaceRecognizer) GetGrid() image.Point {
	return image.Pt(fr.GetGridX(), fr.GetGridY())
}

func (fr *LBPHFaceRecognizer) Close() error {
	C.LBPHFaceRecognizer_Close(fr.p)
	fr.p = nil
	return nil
}
