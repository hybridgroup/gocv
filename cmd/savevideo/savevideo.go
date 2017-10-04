// Package savevideo uses the VideoCapture class to capture video from a connected webcam,
// then saves 100 frames to a video file on disk.

package savevideo

import (
	"fmt"

	"github.com/hybridgroup/gocv"
)

func Run(deviceID int, saveFile string) {
	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(img); !ok {
		fmt.Printf("cannot read device %d\n", deviceID)
		return
	}

	writer, err := gocv.VideoWriterFileMat(saveFile, 25, img)
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
