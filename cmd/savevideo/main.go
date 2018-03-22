// What it does:
//
// This example uses the VideoCapture class to capture video from a connected webcam,
// then saves 100 frames to a video file on disk.
//
// How to run:
//
// savevideo [camera ID] [video file]
//
// 		go run ./cmd/savevideo/main.go 0 testvideo.mp4
//
// +build example

package main

import (
	"fmt"
	"os"
	"strconv"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tsavevideo [camera ID] [video file]")
		return
	}

	deviceID, _ := strconv.Atoi(os.Args[1])
	saveFile := os.Args[2]

	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if err := webcam.Read(&img); err != nil {
		fmt.Println(err)
		return
	}

	writer, err := gocv.VideoWriterFile(saveFile, "MJPG", 25, img.Cols(), img.Rows())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer writer.Close()

	for i := 0; i < 100; i++ {
		if err := webcam.Read(&img); err != nil {
			fmt.Println(err)
			return
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
}
