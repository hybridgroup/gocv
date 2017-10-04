// Package saveimage uses the VideoCapture class to capture a frame from a connected webcam,
// then save it to an image file on disk.
package saveimage

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
	if img.Empty() {
		fmt.Printf("no image on device %d\n", deviceID)
		return
	}

	gocv.IMWrite(saveFile, img)
}
