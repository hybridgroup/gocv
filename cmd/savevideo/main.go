// What it does:
//
// This example uses the VideoCapture class to capture AVI video from a connected webcam,
// then saves 100 frames to a video file on disk.
//
// How to run:
//
// savevideo [camera ID] [video file]
//
// 		go run ./cmd/savevideo/main.go 0 testvideo.avi
//

package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tsavevideo [camera ID] [video file]")
		return
	}

	deviceID := os.Args[1]
	saveFile := os.Args[2]

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Cannot read device %v\n", deviceID)
		return
	}

	writer, err := gocv.VideoWriterFile(saveFile, "MJPG", 25, img.Cols(), img.Rows(), true)
	if err != nil {
		fmt.Printf("error opening video writer device: %v\n", saveFile)
		return
	}
	defer writer.Close()

	for i := 0; i < 100; i++ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
}
