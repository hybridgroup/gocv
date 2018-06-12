// What it does:
//
// This example uses a deep neural network to perform face detection
// of whatever is in front of the camera.
//
// Download the model file from:
// ...
//
// Also, you will need the prototxt file:
// ...
//
// How to run:
//
// 		go run ./cmd/dnn-detection/main.go 0 [modelfile] [configfile] [backend] [device]
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
		fmt.Println("How to run:\ndnn-detection [camera ID] [modelfile] [configfile] [descriptionsfile] ([backend] [device])")
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

	statusColor := color.RGBA{0, 255, 0, 0}
	fmt.Printf("Start reading camera device: %v\n", deviceID)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image Mat to 300x300 blob that the classifier can analyze
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(300, 300), gocv.NewScalar(104, 177, 123, 0), false, false)

		// feed the blob into the classifier
		net.SetInput(blob, "")

		// run a forward pass thru the network
		prob := net.Forward("detection_out")

		// make sure we have valid output layer type
		outs := net.GetUnconnectedOutLayers()
		layer := net.GetLayer(outs[1])
		layerType := layer.GetType()
		layer.Close()
		if layerType != "DetectionOutput" {
			fmt.Printf("Error unknown output layer type %v\n", layerType)
			continue
		}

		for i := 0; i < prob.Total(); i += 7 {
			confidence := prob.GetFloatAt(0, i+2)
			if confidence > 0.5 {
				left := int(prob.GetFloatAt(0, i+3) * float32(img.Cols()))
				top := int(prob.GetFloatAt(0, i+4) * float32(img.Rows()))
				right := int(prob.GetFloatAt(0, i+5) * float32(img.Cols()))
				bottom := int(prob.GetFloatAt(0, i+6) * float32(img.Rows()))
				gocv.Rectangle(&img, image.Rect(left, top, right, bottom), statusColor, 2)
			}
		}

		prob.Close()
		blob.Close()

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
