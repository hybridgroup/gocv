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

	opencv3 ".."
)

func main() {
	deviceID := 0
	webcam := opencv3.NewVideoCapture()
	defer webcam.Close()

	if ok := webcam.OpenDevice(int(deviceID)); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	window := opencv3.NewWindow("Capture")
	defer window.Close()
	
	img := opencv3.NewMat()
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
		opencv3.WaitKey(1)
	}
}
