// Package captest uses the VideoCapture class to test if you can capture video
// from a connected webcam, by trying to read 100 frames.
package captest

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

	// streaming, capture from webcam
	buf := gocv.NewMat()
	defer buf.Close()

	fmt.Printf("Start reading camera device: %v\n", deviceID)
	for i := 0; i < 100; i++ {
		if ok := webcam.Read(buf); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if buf.Empty() {
			continue
		}

		fmt.Printf("Read frame %d\n", i+1)
	}
	fmt.Println("Done.")
}
