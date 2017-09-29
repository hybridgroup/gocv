// What it does:
//
// This example uses the CascadeClassifier to detect faces, 
// then blurs them using a Gaussian blur before displaying in a window.
//
// How to run:
//
// faceblur [camera ID] [classifier XML file]
//
// 		go run ./examples/faceblur.go 0 data/haarcascade_frontalface_default.xml
//
// +build example

package main

import (
	"fmt"
	"os"
	"strconv"

	opencv3 ".."
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
	webcam := opencv3.NewVideoCapture()
	defer webcam.Close()

	if ok := webcam.OpenDevice(deviceID); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	// open display window
	window := opencv3.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := opencv3.NewMat()
	defer img.Close()

	// load classifier to recognize faces
	classifier := opencv3.NewCascadeClassifier()
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

		// draw a rectangle around each face on the original image,
		// along with text identifing as "Human"
		for _, r := range rects {
			imgFace := img.Region(r)
			defer imgFace.Close()
		
			// blur face
			opencv3.GaussianBlur(imgFace, imgFace, opencv3.Size{Width:23, Height:23}, 30, 50, 4)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
