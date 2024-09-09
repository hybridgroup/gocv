package gocv

import (
	"image"
	"os"
	"path/filepath"
	"testing"
)

func TestReadNetDiskFromTensorflow(t *testing.T) {
	path := os.Getenv("GOCV_TENSORFLOW_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Tensorflow model files for tests")
	}

	net := ReadNet(path+"/tensorflow_inception_graph.pb", "")
	if net.Empty() {
		t.Errorf("Unable to load Tensorflow model using ReadNet")
	}
	defer net.Close()

	checkTensorflowNet(t, net)
}

func TestReadNetMemoryFromTensorflow(t *testing.T) {
	path := os.Getenv("GOCV_TENSORFLOW_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Tensorflow model files for tests")
	}

	bModel, err := os.ReadFile(path + "/tensorflow_inception_graph.pb")
	if err != nil {
		t.Errorf("Failed to load model from file: %v", err)
	}

	_, err = ReadNetBytes("tensorflow", nil, nil)
	if err == nil {
		t.Errorf("Should have error for reading nil model bytes")
	}

	net, err := ReadNetBytes("tensorflow", bModel, nil)
	if err != nil {
		t.Errorf("Failed to read net bytes: %v", err)
	}
	if net.Empty() {
		t.Errorf("Unable to load Tensorflow model using ReadNetBytes")
	}
	defer net.Close()

	checkTensorflowNet(t, net)
}

func TestReadNetDiskFromONNX(t *testing.T) {
	path := os.Getenv("GOCV_ONNX_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate ONNX model files for tests")
	}

	net := ReadNet(filepath.Join(path, "googlenet-9.onnx"), "")
	if net.Empty() {
		t.Errorf("Unable to load ONNX model using ReadNet")
	}
	defer net.Close()

	checkONNXNet(t, net)
}

func TestReadNetMemoryFromONNX(t *testing.T) {
	path := os.Getenv("GOCV_ONNX_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate ONNX model files for tests")
	}

	bModel, err := os.ReadFile(filepath.Join(path, "googlenet-9.onnx"))
	if err != nil {
		t.Errorf("Failed to load model from file: %v", err)
	}

	_, err = ReadNetBytes("onnx", nil, nil)
	if err == nil {
		t.Errorf("Should have error for reading nil model bytes")
	}

	net, err := ReadNetBytes("onnx", bModel, nil)
	if err != nil {
		t.Errorf("Failed to read net bytes: %v", err)
	}
	if net.Empty() {
		t.Errorf("Unable to load Caffe model using ReadNetBytes")
	}
	defer net.Close()
	checkONNXNet(t, net)
}

func checkTensorflowNet(t *testing.T, net Net) {
	img := IMRead("images/space_shuttle.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in Tensorflow test")
	}
	defer img.Close()

	blob := BlobFromImage(img, 1.0, image.Pt(224, 224), NewScalar(0, 0, 0, 0), true, false)
	if blob.Empty() {
		t.Error("Invalid blob in Tensorflow test")
	}
	defer blob.Close()

	net.SetInput(blob, "input")
	prob := net.Forward("softmax2")
	defer prob.Close()
	if prob.Empty() {
		t.Error("Invalid softmax2 in Tensorflow test")
	}

	probMat := prob.Reshape(1, 1)
	defer probMat.Close()
	_, maxVal, minLoc, maxLoc := MinMaxLoc(probMat)

	if round(float64(maxVal), 0.00005) != 1.0 {
		t.Errorf("Tensorflow maxVal incorrect: %v\n", round(float64(maxVal), 0.00005))
	}

	if minLoc.X != 481 || minLoc.Y != 0 {
		t.Errorf("Tensorflow minLoc incorrect: %v\n", minLoc)
	}

	if maxLoc.X != 234 || maxLoc.Y != 0 {
		t.Errorf("Tensorflow maxLoc incorrect: %v\n", maxLoc)
	}
}

func TestTensorflowDisk(t *testing.T) {
	path := os.Getenv("GOCV_TENSORFLOW_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Tensorflow model file for tests")
	}

	net := ReadNetFromTensorflow(path + "/tensorflow_inception_graph.pb")
	if net.Empty() {
		t.Errorf("Unable to load Tensorflow model")
	}
	defer net.Close()

	checkTensorflowNet(t, net)
}

func TestTensorflowMemory(t *testing.T) {
	path := os.Getenv("GOCV_TENSORFLOW_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Tensorflow model file for tests")
	}

	b, err := os.ReadFile(path + "/tensorflow_inception_graph.pb")
	if err != nil {
		t.Errorf("Failed to load tensorflow model from file: %v", err)
	}
	net, err := ReadNetFromTensorflowBytes(b)
	if err != nil {
		t.Errorf("Failed to load Tensorflow model from bytes: %v", err)
	}
	if net.Empty() {
		t.Errorf("Unable to load Tensorflow model")
	}
	defer net.Close()

	checkTensorflowNet(t, net)
}

func TestOnnxMemory(t *testing.T) {
	path := os.Getenv("GOCV_ONNX_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate ONNX model file for tests")
	}

	b, err := os.ReadFile(filepath.Join(path, "googlenet-9.onnx"))
	if err != nil {
		t.Errorf("Failed to load ONNX from file: %v", err)
	}

	net, err := ReadNetFromONNXBytes(b)
	if err != nil {
		t.Errorf("Failed to load Tensorflow model from bytes: %v", err)
	}
	if net.Empty() {
		t.Errorf("Unable to load Tensorflow model")
	}
	defer net.Close()

	checkONNXNet(t, net)
}

func TestOnnxDisk(t *testing.T) {
	path := os.Getenv("GOCV_ONNX_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate ONNX model file for tests")
	}

	net := ReadNetFromONNX(filepath.Join(path, "googlenet-9.onnx"))
	if net.Empty() {
		t.Errorf("Unable to load ONNX model")
	}
	defer net.Close()

	checkONNXNet(t, net)
}

func checkONNXNet(t *testing.T, net Net) {
	img := IMRead("images/space_shuttle.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in ONNX test")
	}
	defer img.Close()

	blob := BlobFromImage(img, 1.0, image.Pt(224, 224), NewScalar(0, 0, 0, 0), true, false)
	if blob.Empty() {
		t.Error("Invalid blob in ONNX test")
	}
	defer blob.Close()

	net.SetInput(blob, "data_0")
	prob := net.Forward("prob_1")
	defer prob.Close()
	if prob.Empty() {
		t.Error("Invalid output in ONNX test")
	}

	probMat := prob.Reshape(1, 1)
	defer probMat.Close()
	_, maxVal, minLoc, maxLoc := MinMaxLoc(probMat)

	if round(float64(maxVal), 0.0005) != 0.9965 {
		t.Errorf("ONNX maxVal incorrect: %v\n", round(float64(maxVal), 0.0005))
	}

	if minLoc.X != 955 || minLoc.Y != 0 {
		t.Errorf("ONNX minLoc incorrect: %v\n", minLoc)
	}

	if maxLoc.X != 812 || maxLoc.Y != 0 {
		t.Errorf("ONNX maxLoc incorrect: %v\n", maxLoc)
	}
}

func TestBlobFromImages(t *testing.T) {
	imgs := make([]Mat, 0)

	img := IMRead("images/space_shuttle.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in BlobFromImages test")
	}
	defer img.Close()

	imgs = append(imgs, img)
	imgs = append(imgs, img)

	blob := NewMat()
	BlobFromImages(imgs, &blob, 1.0, image.Pt(25, 25), NewScalar(0, 0, 0, 0), false, false, MatTypeCV32F)
	defer blob.Close()

	sz := GetBlobSize(blob)
	if sz.Val1 != 2 || sz.Val2 != 3 || sz.Val3 != 25 || sz.Val4 != 25 {
		t.Errorf("GetBlobSize in BlobFromImages retrieved wrong values")
	}
}

func TestBlobFromImageGreyscale(t *testing.T) {
	img := IMRead("images/space_shuttle.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in TestBlobFromImageGreyscale test")
	}
	defer img.Close()

	blob := BlobFromImage(img, 1.0, image.Pt(100, 100), NewScalar(0, 0, 0, 0), false, false)
	defer blob.Close()

	if blob.Empty() {
		t.Errorf("BlobFromImageGreyscale failed to create blob")
	}
}

func TestBlobFromImageWithParams(t *testing.T) {
	img := IMRead("images/space_shuttle.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in BlobFromImages test")
	}
	defer img.Close()

	params := NewImageToBlobParams(1.0, image.Pt(25, 25), NewScalar(0, 0, 0, 0), false, MatTypeCV32F, DataLayoutNCHW, PaddingModeCropCenter, NewScalar(0, 0, 0, 0))
	blob := BlobFromImageWithParams(img, params)
	defer blob.Close()

	sz := GetBlobSize(blob)
	if sz.Val1 != 1 || sz.Val2 != 3 || sz.Val3 != 25 || sz.Val4 != 25 {
		t.Errorf("GetBlobSize in BlobFromImagesWithParams retrieved wrong values: %v\n", sz)
	}
}

func TestBlobFromImagesWithParams(t *testing.T) {
	imgs := make([]Mat, 0)

	img := IMRead("images/space_shuttle.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in BlobFromImagesWithParams test")
	}
	defer img.Close()

	imgs = append(imgs, img)
	imgs = append(imgs, img)

	params := NewImageToBlobParams(1.0, image.Pt(25, 25), NewScalar(0, 0, 0, 0), false, MatTypeCV32F, DataLayoutNCHW, PaddingModeCropCenter, NewScalar(0, 0, 0, 0))
	blob := NewMat()
	BlobFromImagesWithParams(imgs, &blob, params)
	defer blob.Close()

	sz := GetBlobSize(blob)
	if sz.Val1 != 2 || sz.Val2 != 3 || sz.Val3 != 25 || sz.Val4 != 25 {
		t.Errorf("GetBlobSize in BlobFromImagesWithParams retrieved wrong values: %v\n", sz)
	}
}

func TestImagesFromBlob(t *testing.T) {
	imgs := make([]Mat, 0)

	img := IMRead("images/space_shuttle.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in BlobFromImages test")
	}
	defer img.Close()

	imgs = append(imgs, img)
	imgs = append(imgs, img)

	blob := NewMat()
	defer blob.Close()
	BlobFromImages(imgs, &blob, 1.0, image.Pt(img.Size()[0], img.Size()[1]), NewScalar(0, 0, 0, 0), false, false, MatTypeCV32F)

	imgsFromBlob := make([]Mat, len(imgs))
	ImagesFromBlob(blob, imgsFromBlob)

	for i := 0; i < len(imgs); i++ {
		func() {
			imgFromBlob := NewMat()
			defer imgFromBlob.Close()
			imgsFromBlob[i].ConvertTo(&imgFromBlob, imgs[i].Type())
			diff := NewMat()
			defer diff.Close()
			Compare(imgs[i], imgFromBlob, &diff, CompareNE)
			nz := CountNonZero(diff)
			if nz != 0 {
				t.Error("imgFromBlob is different from img!")
			}
		}()
	}
}

func TestGetBlobChannel(t *testing.T) {
	img := NewMatWithSize(100, 100, 5+16)
	defer img.Close()

	blob := BlobFromImage(img, 1.0, image.Pt(0, 0), NewScalar(0, 0, 0, 0), true, false)
	defer blob.Close()

	ch2 := GetBlobChannel(blob, 0, 1)
	defer ch2.Close()

	if ch2.Empty() {
		t.Errorf("GetBlobChannel failed to retrieve 2nd chan of a 3channel blob")
	}
	if ch2.Rows() != img.Rows() || ch2.Cols() != img.Cols() {
		t.Errorf("GetBlobChannel: retrieved image size does not match original")
	}
}

func TestGetBlobSize(t *testing.T) {
	img := NewMatWithSize(100, 100, 5+16)
	defer img.Close()

	blob := BlobFromImage(img, 1.0, image.Pt(0, 0), NewScalar(0, 0, 0, 0), true, false)
	defer blob.Close()

	sz := GetBlobSize(blob)
	if sz.Val1 != 1 || sz.Val2 != 3 || sz.Val3 != 100 || sz.Val4 != 100 {
		t.Errorf("GetBlobSize retrieved wrong values")
	}
}

func TestParseNetBackend(t *testing.T) {
	val := ParseNetBackend("halide")
	if val != NetBackendHalide {
		t.Errorf("ParseNetBackend invalid")
	}
	val = ParseNetBackend("openvino")
	if val != NetBackendOpenVINO {
		t.Errorf("ParseNetBackend invalid")
	}
	val = ParseNetBackend("opencv")
	if val != NetBackendOpenCV {
		t.Errorf("ParseNetBackend invalid")
	}
	val = ParseNetBackend("vulkan")
	if val != NetBackendVKCOM {
		t.Errorf("ParseNetBackend invalid")
	}
	val = ParseNetBackend("cuda")
	if val != NetBackendCUDA {
		t.Errorf("ParseNetBackend invalid")
	}
	val = ParseNetBackend("crazytrain")
	if val != NetBackendDefault {
		t.Errorf("ParseNetBackend invalid")
	}
}

func TestParseNetTarget(t *testing.T) {
	val := ParseNetTarget("cpu")
	if val != NetTargetCPU {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("fp32")
	if val != NetTargetFP32 {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("fp16")
	if val != NetTargetFP16 {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("vpu")
	if val != NetTargetVPU {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("cuda")
	if val != NetTargetCUDA {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("vulkan")
	if val != NetTargetVulkan {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("fpga")
	if val != NetTargetFPGA {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("cudafp16")
	if val != NetTargetCUDAFP16 {
		t.Errorf("ParseNetTarget invalid")
	}
	val = ParseNetTarget("idk")
	if val != NetTargetCPU {
		t.Errorf("ParseNetTarget invalid")
	}
}

func TestFP16BlobFromImage(t *testing.T) {
	img := NewMatWithSize(100, 100, 5+16)
	defer img.Close()

	data := FP16BlobFromImage(img, 1.0, image.Pt(100, 100), 0, false, false)

	if len(data) != 60000 {
		t.Errorf("FP16BlobFromImage incorrect length: %v\n", len(data))
	}

	img2 := NewMatWithSize(100, 50, 5+16)
	defer img2.Close()

	data = FP16BlobFromImage(img2, 2.0, image.Pt(50, 100), -0.1, true, false)

	if len(data) != 30000 {
		t.Errorf("FP16BlobFromImage incorrect length: %v\n", len(data))
	}
}

func TestNMSBoxes(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in NMSBoxes test")
	}
	defer img.Close()

	img.ConvertTo(&img, MatTypeCV32F)

	bboxes := []image.Rectangle{
		image.Rect(53, 47, 589, 451),
		image.Rect(118, 54, 618, 450),
		image.Rect(53, 66, 605, 480),
		image.Rect(111, 65, 630, 480),
		image.Rect(156, 51, 640, 480),
	}
	scores := []float32{0.82094115, 0.7998236, 0.9809663, 0.99717456, 0.89628726}
	scoreThreshold := float32(0.5)
	nmsThreshold := float32(0.4)

	indices := NMSBoxes(bboxes, scores, scoreThreshold, nmsThreshold)

	if indices[0] != 3 {
		t.Errorf("Invalid NMSBoxes test indices: %v", indices)
	}
}

func TestNMSBoxesWithParams(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in NMSBoxesWithParams test")
	}
	defer img.Close()

	img.ConvertTo(&img, MatTypeCV32F)

	bboxes := []image.Rectangle{
		image.Rect(53, 47, 589, 451),
		image.Rect(118, 54, 618, 450),
		image.Rect(53, 66, 605, 480),
		image.Rect(111, 65, 630, 480),
		image.Rect(156, 51, 640, 480),
	}
	scores := []float32{0.82094115, 0.7998236, 0.9809663, 0.99717456, 0.89628726}
	scoreThreshold := float32(0.5)
	nmsThreshold := float32(0.4)

	indices := NMSBoxesWithParams(bboxes, scores, scoreThreshold, nmsThreshold, float32(1.0), 0)

	if indices[0] != 3 {
		t.Errorf("Invalid NMSBoxesWithParams test indices: %v", indices)
	}
}
