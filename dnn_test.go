package gocv

import (
	"image"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadNet(t *testing.T) {
	checkNet := func(t *testing.T, net Net) {
		net.SetPreferableBackend(NetBackendDefault)
		net.SetPreferableTarget(NetTargetCPU)

		img := IMRead("images/space_shuttle.jpg", IMReadColor)
		if img.Empty() {
			t.Error("Invalid Mat in ReadNet test")
		}
		defer img.Close()

		blob := BlobFromImage(img, 1.0, image.Pt(224, 224), NewScalar(0, 0, 0, 0), false, false)
		if blob.Empty() {
			t.Error("Invalid blob in ReadNet test")
		}
		defer blob.Close()

		net.SetInput(blob, "data")

		layer := net.GetLayer(0)
		defer layer.Close()

		if layer.InputNameToIndex("notthere") != -1 {
			t.Error("Invalid layer in ReadNet test")
		}
		if layer.OutputNameToIndex("notthere") != -1 {
			t.Error("Invalid layer in ReadNet test")
		}
		if layer.GetName() != "_input" {
			t.Errorf("Invalid layer name in ReadNet test: %s\n", layer.GetName())
		}
		if layer.GetType() != "" {
			t.Errorf("Invalid layer type in ReadNet test: %s\n", layer.GetType())
		}

		ids := net.GetUnconnectedOutLayers()
		if len(ids) != 1 {
			t.Errorf("Invalid len output layers in ReadNet test: %d\n", len(ids))
		}

		lnames := net.GetLayerNames()
		if len(lnames) != 142 {
			t.Errorf("Invalid len layer names in ReadNet test: %d\n", len(lnames))
		}

		prob := net.ForwardLayers([]string{"prob"})
		if len(prob) == 0 {
			t.Error("Invalid len prob in ReadNet test")
		}

		if prob[0].Empty() {
			t.Error("Invalid prob[0] in ReadNet test")
		}

		probMat := prob[0].Reshape(1, 1)
		defer probMat.Close()
		_, maxVal, minLoc, maxLoc := MinMaxLoc(probMat)

		if round(float64(maxVal), 0.00005) != 0.9998 {
			t.Errorf("ReadNet maxVal incorrect: %v\n", round(float64(maxVal), 0.00005))
		}

		if minLoc.X != 955 || minLoc.Y != 0 {
			t.Errorf("ReadNet minLoc incorrect: %v\n", minLoc)
		}

		if maxLoc.X != 812 || maxLoc.Y != 0 {
			t.Errorf("ReadNet maxLoc incorrect: %v\n", maxLoc)
		}

		perf := net.GetPerfProfile()
		if perf == 0 {
			t.Error("ReadNet GetPerfProfile error")
		}
	}

	path := os.Getenv("GOCV_CAFFE_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Caffe model files for tests")
	}

	t.Run("net from disk", func(t *testing.T) {
		net := ReadNet(path+"/bvlc_googlenet.caffemodel", path+"/bvlc_googlenet.prototxt")
		if net.Empty() {
			t.Errorf("Unable to load Caffe model using ReadNet")
		}
		defer net.Close()

		checkNet(t, net)
	})

	t.Run("net from memory", func(t *testing.T) {
		bModel, err := ioutil.ReadFile(path + "/bvlc_googlenet.caffemodel")
		if err != nil {
			t.Errorf("Failed to load model from file: %v", err)
		}
		bConfig, err := ioutil.ReadFile(path + "/bvlc_googlenet.prototxt")
		if err != nil {
			t.Errorf("Failed to load config from file: %v", err)
		}
		net, err := ReadNetBytes("caffe", bModel, bConfig)
		if err != nil {
			t.Errorf("Failed to read net bytes: %v", err)
		}
		if net.Empty() {
			t.Errorf("Unable to load Caffe model using ReadNetBytes")
		}
		defer net.Close()

		checkNet(t, net)
	})
}

func TestCaffe(t *testing.T) {
	checkNet := func(t *testing.T, net Net) {
		img := IMRead("images/space_shuttle.jpg", IMReadColor)
		if img.Empty() {
			t.Error("Invalid Mat in Caffe test")
		}
		defer img.Close()

		blob := BlobFromImage(img, 1.0, image.Pt(224, 224), NewScalar(0, 0, 0, 0), false, false)
		if blob.Empty() {
			t.Error("Invalid blob in Caffe test")
		}
		defer blob.Close()

		net.SetInput(blob, "data")
		prob := net.Forward("prob")
		defer prob.Close()
		if prob.Empty() {
			t.Error("Invalid prob in Caffe test")
		}

		probMat := prob.Reshape(1, 1)
		defer probMat.Close()
		_, maxVal, minLoc, maxLoc := MinMaxLoc(probMat)

		if round(float64(maxVal), 0.00005) != 0.9998 {
			t.Errorf("Caffe maxVal incorrect: %v\n", round(float64(maxVal), 0.00005))
		}

		if minLoc.X != 955 || minLoc.Y != 0 {
			t.Errorf("Caffe minLoc incorrect: %v\n", minLoc)
		}

		if maxLoc.X != 812 || maxLoc.Y != 0 {
			t.Errorf("Caffe maxLoc incorrect: %v\n", maxLoc)
		}
	}

	path := os.Getenv("GOCV_CAFFE_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Caffe model files for tests")
	}

	t.Run("net from disk", func(t *testing.T) {
		net := ReadNetFromCaffe(path+"/bvlc_googlenet.prototxt", path+"/bvlc_googlenet.caffemodel")
		if net.Empty() {
			t.Errorf("Unable to load Caffe model")
		}
		defer net.Close()

		checkNet(t, net)
	})

	t.Run("net from memory", func(t *testing.T) {
		bPrototxt, err := ioutil.ReadFile(path + "/bvlc_googlenet.prototxt")
		if err != nil {
			t.Errorf("Failed to load Caffe prototxt from file: %v", err)
		}
		bCaffeModel, err := ioutil.ReadFile(path + "/bvlc_googlenet.caffemodel")
		if err != nil {
			t.Errorf("Failed to load Caffe caffemodel from file: %v", err)
		}
		net, err := ReadNetFromCaffeBytes(bPrototxt, bCaffeModel)
		if err != nil {
			t.Errorf("Error reading caffe from bytes: %v", err)
		}
		if net.Empty() {
			t.Errorf("Unable to load Caffe model")
		}
		defer net.Close()

		checkNet(t, net)
	})
}

func TestTensorflow(t *testing.T) {
	checkNet := func(t *testing.T, net Net) {
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

	path := os.Getenv("GOCV_TENSORFLOW_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Tensorflow model file for tests")
	}

	t.Run("net from disk", func(t *testing.T) {
		net := ReadNetFromTensorflow(path + "/tensorflow_inception_graph.pb")
		if net.Empty() {
			t.Errorf("Unable to load Tensorflow model")
		}
		defer net.Close()

		checkNet(t, net)
	})

	t.Run("net from memory", func(t *testing.T) {
		b, err := ioutil.ReadFile(path + "/tensorflow_inception_graph.pb")
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

		checkNet(t, net)
	})
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
