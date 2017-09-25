// What it does:
//
// This example uses the CascadeClassifier class to detect faces,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// How to run:
//
// facedetect [camera ID] [classifier XML file]
//
// 		go run ./examples/facedetect.go 0 data/haarcascade_frontalface_default.xml
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
		fmt.Println("How to run:\n\tfacedetect [camera ID] [classifier XML file]")
		return
	}

	deviceID, _ := strconv.Atoi(os.Args[1])
	xmlFile := os.Args[2]

	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(deviceID); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	window := opencv3.NewWindow("Capture")

	img := opencv3.NewMat()
	defer img.Delete()

	classifier := opencv3.NewCascadeClassifier()
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

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d\n", len(rects))
		if len(rects) > 0 {
			opencv3.DrawRectsToImage(img, rects)
		}

		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
