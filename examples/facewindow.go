// how to use
// 		go run ./examples/facewindow.go data/haarcascade_frontalface_default.xml
//
package main

import (
	"fmt"
	"os"

	opencv3 ".."
)

func main() {
	deviceID := 0
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(int(deviceID)); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	window := opencv3.NewWindow("Capture")

	img := opencv3.NewMat()
	defer img.Delete()

	classifier := opencv3.NewCascadeClassifier()
	classifier.Load(os.Args[1])

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d\n", len(rects))
		if len(rects) > 0 {
			opencv3.DrawRectsToImage(img, rects)
		}

		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
