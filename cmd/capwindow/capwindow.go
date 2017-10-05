// Package capwindow uses the VideoCapture class to capture frames from a connected webcam,
// and displays the video in a Window class.
package capwindow

import (
	"fmt"

	"github.com/hybridgroup/gocv"
)

func Run(deviceID int) {
	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Capture Window")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("Start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("Error cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		gocv.WaitKey(1)
	}
}
