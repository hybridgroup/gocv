// What it does:
//
// This example uses the VideoCapture class to capture video from a connected webcam,
// and save it to a video file on disk
//
// How to run:
//
// savevideo [camera ID] [video file]
//
// 		go run ./examples/savevideo.go 0 testvideo.mp4
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
		fmt.Println("How to run:\n\tsavevideo [camera ID] [video file]")
		return
	}

	deviceID, _ := strconv.Atoi(os.Args[1])
	saveFile := os.Args[2]
	
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	writer := opencv3.NewVideoWriter()
	defer writer.Delete()

	if ok := webcam.OpenDevice(deviceID); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
	}

	img := opencv3.NewMat()
	defer img.Delete()

	if ok := webcam.Read(img); !ok {
		fmt.Printf("cannot read device %d\n", deviceID)
		return
	}

	writer.OpenWithMat(saveFile, 25, img)

	for i := 0; i < 100; i++ {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
}
