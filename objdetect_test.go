package gocv

import (
	"image"
	"image/color"
	"os"
	"testing"
)

func TestCascadeClassifier(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CascadeClassifier test")
	}
	defer img.Close()

	// load classifier to recognize faces
	classifier := NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("data/haarcascade_frontalface_default.xml")

	rects := classifier.DetectMultiScale(img)
	if len(rects) != 1 {
		t.Error("Error in TestCascadeClassifier test")
	}
}

func TestCascadeClassifierWithParams(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CascadeClassifierWithParams test")
	}
	defer img.Close()

	// load classifier to recognize faces
	classifier := NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("data/haarcascade_frontalface_default.xml")

	rects := classifier.DetectMultiScaleWithParams(img, 1.1, 3, 0, image.Pt(0, 0), image.Pt(0, 0))
	if len(rects) != 1 {
		t.Errorf("Error in CascadeClassifierWithParams test: %v", len(rects))
	}
}

func TestHOGDescriptor(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in HOGDescriptor test")
	}
	defer img.Close()

	// load HOGDescriptor to recognize people
	hog := NewHOGDescriptor()
	defer hog.Close()

	d := HOGDefaultPeopleDetector()
	defer d.Close()
	hog.SetSVMDetector(d)

	rects := hog.DetectMultiScale(img)
	if len(rects) != 1 {
		t.Errorf("Error in TestHOGDescriptor test: %d", len(rects))
	}
}

func TestHOGDescriptorWithParams(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in HOGDescriptorWithParams test")
	}
	defer img.Close()

	// load HOGDescriptor to recognize people
	hog := NewHOGDescriptor()
	defer hog.Close()

	d := HOGDefaultPeopleDetector()
	defer d.Close()
	hog.SetSVMDetector(d)

	rects := hog.DetectMultiScaleWithParams(img, 0, image.Pt(0, 0), image.Pt(0, 0),
		1.05, 2.0, false)
	if len(rects) != 1 {
		t.Errorf("Error in TestHOGDescriptorWithParams test: %d", len(rects))
	}
}

func TestGroupRectangles(t *testing.T) {
	rects := []image.Rectangle{
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 30, 30),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
		image.Rect(10, 10, 35, 35),
	}

	results := GroupRectangles(rects, 1, 0.2)
	if len(results) != 2 {
		t.Errorf("Error in TestGroupRectangles test: %d", len(results))
	}
}

func TestQRCodeDetector(t *testing.T) {
	img := IMRead("images/qrcode.png", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in QRCodeDetector test")
	}
	defer img.Close()

	// load QRCodeDetector to QR codes

	detector := NewQRCodeDetector()
	defer detector.Close()

	bbox := NewMat()
	qr := NewMat()
	defer bbox.Close()
	defer qr.Close()

	res := detector.Detect(img, &bbox)
	if !res {
		t.Errorf("Error in TestQRCodeDetector test: res == false")
	}
	res2 := detector.Decode(img, bbox, &qr)

	res3 := detector.DetectAndDecode(img, &bbox, &qr)

	if res2 != res3 {
		t.Errorf("Error in TestQRCodeDetector res2: %s != res3: %s", res2, res3)
	}

	// multi
	img2 := IMRead("images/multi_qrcodes.png", IMReadColor)
	defer img2.Close()
	if img2.Empty() {
		t.Error("Invalid Mat in QRCodeDetector test")
	}

	multiBox := NewMat()
	defer multiBox.Close()
	res4 := detector.DetectMulti(img2, &multiBox)
	if !res4 {
		t.Errorf("Error in TestQRCodeDetector Multi test: res == false")
	}

	if multiBox.Rows() != 2 {
		t.Errorf("Error in TestQRCodeDetector Multi test: number of Rows = %d", multiBox.Rows())
	}

	multiBox2 := NewMat()
	defer multiBox2.Close()
	decoded := []string{}
	qrCodes := make([]Mat, 0)
	defer func() {
		for _, q := range qrCodes {
			q.Close()
		}
	}()
	success := detector.DetectAndDecodeMulti(img2, decoded, &multiBox2, qrCodes)
	if !success {
		t.Errorf("Error in TestQRCodeDetector Multi test: returned false")
	}

	tmpPoints := NewMat()
	defer tmpPoints.Close()
	tmpQr := NewMat()
	defer tmpQr.Close()
	var tmpDecoded string
	for i, s := range decoded {
		tmpInput := padQr(&(qrCodes[i]))
		defer tmpInput.Close()
		tmpDecoded = detector.Decode(tmpInput, tmpPoints, &tmpQr)
		if tmpDecoded != s {
			t.Errorf("Error in TestQRCodeDetector Multi test: decoded straight QR code=%s, decoded[%d] = %s", tmpDecoded, i, s)
		}
	}

	emptyMat := NewMatWithSize(100, 200, MatTypeCV8UC3)
	success = detector.DetectAndDecodeMulti(emptyMat, decoded, &multiBox2, qrCodes)
	if success {
		t.Errorf("Error in TestQRCodeDetector Multi test: empty Mat returned success=true")
	}
	emptyMat.Close()
}

func padQr(qr *Mat) Mat {
	l := 101
	d := 10
	L := l + 2*d

	out := NewMatWithSizeFromScalar(NewScalar(255, 255, 255, 255), L, L, MatTypeCV8UC3)
	qrCodes0 := NewMat()
	defer qrCodes0.Close()
	qr.ConvertTo(&qrCodes0, MatTypeCV8UC3)

	Resize(qrCodes0, &qrCodes0, image.Point{L, L}, 0, 0, InterpolationArea)
	CopyMakeBorder(qrCodes0, &out, d, d, d, d, BorderConstant, color.RGBA{255, 255, 255, 255})
	return out
}

func TestFaceDetectorYN(t *testing.T) {

	img := IMRead("images/face.jpg", IMReadAnyColor)
	defer img.Close()

	s := image.Pt(img.Size()[1], img.Size()[0])

	faces := NewMat()
	defer faces.Close()

	// https://github.com/opencv/opencv_zoo/raw/refs/heads/main/models/face_detection_yunet/face_detection_yunet_2023mar.onnx
	fd := NewFaceDetectorYN("testdata/face_detection_yunet_2023mar.onnx", "", s)
	defer fd.Close()

	sz := fd.GetInputSize()
	if sz.X != 640 && sz.Y != 480 {
		t.Error("error on FaceDetectorYN.GetInputSize()")
	}
	fd.SetInputSize(sz)

	t1 := fd.GetMNSThreshold()
	fd.SetNMSThreshold(t1)

	t2 := fd.GetScoreThreshold()
	fd.SetScoreThreshold(t2)

	topK := fd.GetTopK()
	fd.SetTopK(topK)

	fd.Detect(img, &faces)

	facesCount := faces.Rows()

	if facesCount < 1 {
		t.Error("no face detected")
	}
}

func TestFaceDetectorYNWithParams(t *testing.T) {

	img := IMRead("images/face.jpg", IMReadAnyColor)
	defer img.Close()

	s := image.Pt(img.Size()[1], img.Size()[0])

	faces := NewMat()
	defer faces.Close()

	fd := NewFaceDetectorYNWithParams("testdata/face_detection_yunet_2023mar.onnx", "", s, 0.9, 0.3, 5000, 0, 0)
	defer fd.Close()

	sz := fd.GetInputSize()
	if sz.X != 640 && sz.Y != 480 {
		t.Error("error on FaceDetectorYN.GetInputSize()")
	}
	fd.SetInputSize(sz)

	t1 := fd.GetMNSThreshold()
	fd.SetNMSThreshold(t1)

	t2 := fd.GetScoreThreshold()
	fd.SetScoreThreshold(t2)

	topK := fd.GetTopK()
	fd.SetTopK(topK)

	fd.Detect(img, &faces)

	facesCount := faces.Rows()

	if facesCount < 1 {
		t.Error("no face detected")
	}

}

func TestFaceDetectorYNFromBytes(t *testing.T) {

	modelBuffer, err := os.ReadFile("testdata/face_detection_yunet_2023mar.onnx")
	if err != nil {
		t.Errorf("%s reading testdata/face_detection_yunet_2023mar.onnx", err.Error())
	}

	img := IMRead("images/face.jpg", IMReadAnyColor)
	defer img.Close()

	s := image.Pt(img.Size()[1], img.Size()[0])

	faces := NewMat()
	defer faces.Close()

	fd := NewFaceDetectorYNFromBytes("onnx", modelBuffer, []byte(""), s)
	defer fd.Close()

	sz := fd.GetInputSize()
	if sz.X != 640 && sz.Y != 480 {
		t.Error("error on FaceDetectorYN.GetInputSize()")
	}
	fd.SetInputSize(sz)

	t1 := fd.GetMNSThreshold()
	fd.SetNMSThreshold(t1)

	t2 := fd.GetScoreThreshold()
	fd.SetScoreThreshold(t2)

	topK := fd.GetTopK()
	fd.SetTopK(topK)

	fd.Detect(img, &faces)

	facesCount := faces.Rows()

	if facesCount < 1 {
		t.Error("no face detected")
	}
}

func TestFaceDetectorYNFromBytesWithParams(t *testing.T) {

	modelBuffer, err := os.ReadFile("testdata/face_detection_yunet_2023mar.onnx")
	if err != nil {
		t.Errorf("%s reading testdata/face_detection_yunet_2023mar.onnx", err.Error())
	}

	img := IMRead("images/face.jpg", IMReadAnyColor)
	defer img.Close()

	s := image.Pt(img.Size()[1], img.Size()[0])

	faces := NewMat()
	defer faces.Close()

	fd := NewFaceDetectorYNFromBytesWithParams("onnx", modelBuffer, []byte(""), s, 0.9, 0.3, 5000, 0, 0)
	defer fd.Close()

	sz := fd.GetInputSize()
	if sz.X != 640 && sz.Y != 480 {
		t.Error("error on FaceDetectorYN.GetInputSize()")
	}
	fd.SetInputSize(sz)

	t1 := fd.GetMNSThreshold()
	fd.SetNMSThreshold(t1)

	t2 := fd.GetScoreThreshold()
	fd.SetScoreThreshold(t2)

	topK := fd.GetTopK()
	fd.SetTopK(topK)

	fd.Detect(img, &faces)

	facesCount := faces.Rows()

	if facesCount < 1 {
		t.Error("no face detected")
	}
}

func TestFaceRecognizerSF(t *testing.T) {

	img := IMRead("images/face.jpg", IMReadUnchanged)
	defer img.Close()

	IMWrite("testdata/original.png", img)

	imgSz := img.Size()
	s := image.Pt(imgSz[1], imgSz[0])

	faces := NewMat()
	defer faces.Close()

	// https://github.com/opencv/opencv_zoo/raw/refs/heads/main/models/face_detection_yunet/face_detection_yunet_2023mar.onnx
	fd := NewFaceDetectorYN("testdata/face_detection_yunet_2023mar.onnx", "", s)
	defer fd.Close()

	detectRv := fd.Detect(img, &faces)
	t.Log("detect rv is", detectRv)

	facesCount := faces.Rows()
	if facesCount < 1 {
		t.Error("no face detected")
	}

	ronsFaceX := faces.GetFloatAt(0, 0)
	ronsFaceY := faces.GetFloatAt(0, 1)
	ronsFaceW := faces.GetFloatAt(0, 2)
	ronsFaceH := faces.GetFloatAt(0, 3)

	ronsFace := img.Region(image.Rect(int(ronsFaceX), int(ronsFaceY), int(ronsFaceW), int(ronsFaceH)))
	defer ronsFace.Close()
	IMWrite("testdata/ronsface.png", ronsFace)

	//https://github.com/opencv/opencv_zoo/raw/refs/heads/main/models/face_recognition_sface/face_recognition_sface_2021dec.onnx
	fr := NewFaceRecognizerSF("testdata/face_recognition_sface_2021dec.onnx", "")
	defer fr.Close()

	aligned := NewMat()
	defer aligned.Close()

	fr.AlignCrop(img, ronsFace, &aligned)

	if aligned.Empty() {
		t.Error("aligned is empty")
	}

	IMWrite("testdata/ronsfaceSF.png", ronsFace)
	IMWrite("testdata/aligned.png", aligned)
}
