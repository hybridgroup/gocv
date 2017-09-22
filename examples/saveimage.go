// +build example
//
// Do not build by default.
//
// This example uses the VideoCapture class to capture a frame from a connected webcam,
// then save it to a file on disk
//
// how to run:
// 		go run ./examples/saveimage.go filename.jpg
//
package main

import (
	"fmt"
	"os"

	opencv3 ".."
)

func main() {
	deviceID := 0
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(int(deviceID)); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
	}

	img := opencv3.NewMat()
	defer img.Delete()

	if ok := webcam.Read(img); !ok {
		fmt.Printf("cannot read device %d\n", deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %d\n", deviceID)
		return
	}

	opencv3.IMWrite(os.Args[1], img)
}
