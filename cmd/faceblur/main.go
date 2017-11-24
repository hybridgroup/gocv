// What it does:
//
// This example captures video from a connected camera,
// then uses the CascadeClassifier to detect faces, blurs them
// using a Gaussian blur, then displays the blurred video in a window.
//
// How to run:
//
// faceblur [camera ID] [classifier XML file]
//
// 		go run ./cmd/faceblur/main.go 0 data/haarcascade_frontalface_default.xml
//
// +build example

package main

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfaceblur [camera ID] [classifier XML file]")
		return
	}

	// parse args
	deviceID, _ := strconv.Atoi(os.Args[1])
	xmlFile := os.Args[2]

	// open webcam
	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load(xmlFile)

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// blur each face on the original image
		for _, r := range rects {
			imgFace := img.Region(r)
			defer imgFace.Close()

			// blur face
			gocv.GaussianBlur(imgFace, imgFace, image.Pt(23, 23), 30, 50, 4)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)
	}
}
