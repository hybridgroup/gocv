// What it does:
//
// This example uses the VideoCapture class to capture a frame from a connected webcam,
// then save it to an image file on disk
//
// How to run:
//
// saveimage [camera ID] [image file]
//
// 		go run ./examples/saveimage.go filename.jpg
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
		fmt.Println("How to run:\n\tsaveimage [camera ID] [image file]")
		return
	}

	deviceID, _ := strconv.Atoi(os.Args[1])
	saveFile := os.Args[2]

	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(deviceID); !ok {
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

	opencv3.IMWrite(saveFile, img)
}
