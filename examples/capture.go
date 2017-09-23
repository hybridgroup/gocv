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
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(int(deviceID)); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
	}

	// streaming, capture from webcam
	buf := opencv3.NewMat()
	defer buf.Delete()
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
