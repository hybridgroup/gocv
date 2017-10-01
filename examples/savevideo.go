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
	
	webcam, err := opencv3.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}	
	defer webcam.Close()

	img := opencv3.NewMat()
	defer img.Close()

	if ok := webcam.Read(img); !ok {
		fmt.Printf("cannot read device %d\n", deviceID)
		return
	}

	writer, err := opencv3.VideoWriterFileMat(saveFile, 25, img)
	if err != nil {
		fmt.Printf("error opening video writer device: %v\n", saveFile)
		return
	}		
	defer writer.Close()

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
