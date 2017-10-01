// What it does:
//
// This example uses the VideoCapture class to capture frames from a connected webcam.
//
// How to run:
//
// 		go run ./examples/capture.go
//
// +build example

package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	deviceID := 0
	webcam, err := opencv3.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// streaming, capture from webcam
	buf := opencv3.NewMat()
	defer buf.Close()
	
	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(buf); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if buf.Empty() {
			continue
		}

		fmt.Println("frame")
	}
}
