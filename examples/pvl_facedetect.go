// What it does:
//
// This example uses the Intel CV SDK PVL FaceDetect class to detect faces,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// How to run:
//
// pvl_facedetect [camera ID]
//
// 		go run ./examples/pvl_facedetect.go 0
//
// +build example

package main

import (
	"fmt"
	"os"
	"strconv"

	opencv3 ".."
	pvl "../pvl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tpvl_facedetect [camera ID]")
		return
	}

	deviceID, _ := strconv.Atoi(os.Args[1])

	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(deviceID); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	window := opencv3.NewWindow("Capture")

	img := opencv3.NewMat()
	defer img.Delete()

	fd := pvl.NewFaceDetector()
	defer fd.Close()

	fd.SetTrackingModeEnabled(true)

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		faces := fd.DetectFaceRect(img)

		fmt.Printf("found %d\n", len(faces))
		if len(faces) > 0 {
			rects := []opencv3.Rect{faces[0].Rect()}
			opencv3.DrawRectsToImage(img, rects)
		}

		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
