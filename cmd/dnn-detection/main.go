// What it does:
//
// This example uses a deep neural network to perform face detection.
//
// Download the model file from:
// https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel
//
// You will also need the prototxt file:
// https://raw.githubusercontent.com/opencv/opencv/master/samples/dnn/face_detector/deploy.prototxt
//
// How to run:
//
// 		go run ./cmd/dnn-detection/main.go [videosource] [modelfile] [configfile] ([backend] [device])
//
// +build example

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("How to run:\ndnn-detection [videosource] [modelfile] [configfile] ([backend] [device])")
		return
	}

	// parse args
	deviceID := os.Args[1]
	proto := os.Args[2]
	model := os.Args[3]
	backend := gocv.NetBackendDefault
	if len(os.Args) > 4 {
		backend = gocv.ParseNetBackend(os.Args[4])
	}

	target := gocv.NetTargetCPU
	if len(os.Args) > 5 {
		target = gocv.ParseNetTarget(os.Args[5])
	}

	// open capture device
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("DNN Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	// open DNN face tracking model
	net := gocv.ReadNet(model, proto)
	if net.Empty() {
		fmt.Printf("Error reading network model from : %v %v\n", model, proto)
		return
	}
	defer net.Close()
	net.SetPreferableBackend(gocv.NetBackendType(backend))
	net.SetPreferableTarget(gocv.NetTargetType(target))

	// make sure we have valid output layer type
	outs := net.GetUnconnectedOutLayers()
	layer := net.GetLayer(outs[1])
	layerType := layer.GetType()
	layer.Close()
	if layerType != "DetectionOutput" {
		fmt.Printf("Error unknown output layer type %v\n", layerType)
		return
	}

	fmt.Printf("Start reading camera device: %v\n", deviceID)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image Mat to 300x300 blob that the object detector can analyze
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(300, 300), gocv.NewScalar(104, 177, 123, 0), false, false)

		// feed the blob into the detector
		net.SetInput(blob, "")

		// run a forward pass thru the network
		prob := net.Forward("detection_out")

		performDetection(&img, prob)

		prob.Close()
		blob.Close()

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

// performDetection analyzes the results from the detector network,
// which produces an output blob with a shape 1x1xNx7
// where N is the number of detections, and each detection
// is a vector of float values
// [batchId, classId, confidence, left, top, right, bottom]
func performDetection(frame *gocv.Mat, results gocv.Mat) {
	for i := 0; i < results.Total(); i += 7 {
		confidence := results.GetFloatAt(0, i+2)
		if confidence > 0.5 {
			left := int(results.GetFloatAt(0, i+3) * float32(frame.Cols()))
			top := int(results.GetFloatAt(0, i+4) * float32(frame.Rows()))
			right := int(results.GetFloatAt(0, i+5) * float32(frame.Cols()))
			bottom := int(results.GetFloatAt(0, i+6) * float32(frame.Rows()))
			gocv.Rectangle(frame, image.Rect(left, top, right, bottom), color.RGBA{0, 255, 0, 0}, 2)
		}
	}
}
