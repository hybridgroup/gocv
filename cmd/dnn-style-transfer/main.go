// What it does:
//
// This example uses a deep neural network to perform style transfer.
//
// Download the model file from:
// http://cs.stanford.edu/people/jcjohns/fast-neural-style/models/eccv16/starry_night.t7
//
// How to run:
//
// 		go run ./cmd/dnn-style-transfer/main.go [videosource] [modelfile] ([backend] [device])
//
// +build example

package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\ndnn-style-transfer [videosource] [modelfile] ([backend] [device])")
		return
	}

	// parse args
	deviceID := os.Args[1]
	model := os.Args[2]
	backend := gocv.NetBackendDefault
	if len(os.Args) > 3 {
		backend = gocv.ParseNetBackend(os.Args[3])
	}

	target := gocv.NetTargetCPU
	if len(os.Args) > 4 {
		target = gocv.ParseNetTarget(os.Args[4])
	}

	// open capture device
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("DNN Style Transfer")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	// open DNN style transfer model
	net := gocv.ReadNet(model, "")
	if net.Empty() {
		fmt.Printf("Error reading network model from : %v\n", model)
		return
	}
	defer net.Close()
	net.SetPreferableBackend(gocv.NetBackendType(backend))
	net.SetPreferableTarget(gocv.NetTargetType(target))

	fmt.Printf("Start reading camera device: %v\n", deviceID)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image Mat to 640x480 blob that the style transfer can analyze
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(640, 480), gocv.NewScalar(103.939, 116.779, 123.68, 0), false, false)

		// feed the blob into the detector
		net.SetInput(blob, "")

		// run a forward pass thru the network
		prob := net.Forward("")

		out := gocv.NewMatWithSize(480, 640, gocv.MatTypeCV8UC3)
		probMat := prob

		for i := 0; i < probMat.Total(); i += 3 {
			r := probMat.GetFloatAt(0, i)
			r += 103.939

			g := probMat.GetFloatAt(0, i+1)
			g += 116.779

			b := probMat.GetFloatAt(0, i+2)
			b += 123.68

			out.SetUCharAt(0, i, uint8(b))
			out.SetUCharAt(0, i+1, uint8(g))
			out.SetUCharAt(0, i+2, uint8(r))
		}

		window.IMShow(out)
		if window.WaitKey(1) >= 0 {
			break
		}

		prob.Close()
		blob.Close()
		out.Close()
	}
}
