// What it does:
//
// This example uses the FaceDetectorYN class to detect faces,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// model files download link:
// https://github.com/opencv/opencv_zoo/tree/main/models/face_detection_yunet
//
// How to run:
//
// facedetectYN [camera ID] [model file]
//
// 		go run ./cmd/facedetectYN/main.go 0 face_detection_yunet_2023mar.onnx
//

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfacedetectYN [camera ID] [model file]")
		return
	}

	// parse args
	deviceID := os.Args[1]
	modelFile := os.Args[2]

	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect YN")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// prepare faces matrix
	faces := gocv.NewMat()
	defer faces.Close()

	// colors for the rect faces detected
	red := color.RGBA{255, 0, 0, 0}
	green := color.RGBA{0, 255, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}
	yellow := color.RGBA{255, 255, 0, 1}
	pink := color.RGBA{255, 105, 180, 0}

	detector := gocv.NewFaceDetectorYN(modelFile, "", image.Pt(200, 200))
	defer detector.Close()

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		sz := img.Size()
		detector.SetInputSize(image.Pt(sz[1], sz[0]))

		detector.Detect(img, &faces)

		if faces.Rows() < 1 {
			// no faces detected
			// show the captured frame anyway
			window.IMShow(img)
			if window.WaitKeyEx(1) >= 0 {
				break
			}
			continue
		}

		for r := 0; r < faces.Rows(); r++ {

			x0 := int(faces.GetFloatAt(r, 0))
			y0 := int(faces.GetFloatAt(r, 1))
			x1 := x0 + int(faces.GetFloatAt(r, 2))
			y1 := y0 + int(faces.GetFloatAt(r, 3))

			faceRect := image.Rect(x0, y0, x1, y1)

			rightEye := image.Pt(
				int(faces.GetFloatAt(r, 4)),
				int(faces.GetFloatAt(r, 5)),
			)

			leftEye := image.Pt(
				int(faces.GetFloatAt(r, 6)),
				int(faces.GetFloatAt(r, 7)),
			)

			noseTip := image.Pt(
				int(faces.GetFloatAt(r, 8)),
				int(faces.GetFloatAt(r, 9)),
			)

			rightMouthCorner := image.Pt(
				int(faces.GetFloatAt(r, 10)),
				int(faces.GetFloatAt(r, 11)),
			)

			leftMouthCorner := image.Pt(
				int(faces.GetFloatAt(r, 12)),
				int(faces.GetFloatAt(r, 13)),
			)

			gocv.Rectangle(&img, faceRect, green, 1)
			gocv.Circle(&img, rightEye, 1, blue, 1)
			gocv.Circle(&img, leftEye, 1, red, 1)
			gocv.Circle(&img, noseTip, 1, green, 1)
			gocv.Circle(&img, rightMouthCorner, 1, pink, 1)
			gocv.Circle(&img, leftMouthCorner, 1, yellow, 1)

		}
		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKeyEx(1) >= 0 {
			break
		}
	}
}
