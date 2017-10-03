// What it does:
//
// This example uses the VideoCapture class to capture frames from a connected webcam,
// then displays the image in a Window class.
//
// How to run:
//
// 		go run ./examples/capwindow.go
//
// +build example

package main

import (
	"fmt"

	gocv ".."
)

func main() {
	deviceID := 0
	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Capture")
	defer window.Close()
	
	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		gocv.WaitKey(1)
	}
}
