// What it does:
//
// This example shows how to find color in an image
//
// How to run:
//
// 		go run ./cmd/img-color/main.go ./images/circles.jpg
//
// +build example

package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\timage-color [imgfile]")
		return
	}

	filename := os.Args[1]

	windowO := gocv.NewWindow("original image")
	defer windowO.Close()

	windowM := gocv.NewWindow("mask image")
	defer windowM.Close()

	img := gocv.IMRead(filename, gocv.IMReadColor)
	defer img.Close()

	hsv := gocv.NewMat()
	defer hsv.Close()

	gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)

	mask := gocv.NewMat()
	defer mask.Close()

	lowerBlack := gocv.NewMatFromScalar(gocv.NewScalar(0.0, 0.0, 0.0, 0.0), gocv.MatTypeCV8U)
	defer lowerBlack.Close()

	upperBlack := gocv.NewMatFromScalar(gocv.NewScalar(0.0, 0.0, 0.0, 0.0), gocv.MatTypeCV8U)
	defer upperBlack.Close()

	gocv.InRange(hsv, lowerBlack, upperBlack, &mask)

	windowO.IMShow(img)
	windowM.IMShow(mask)

	windowO.WaitKey(0)
	windowM.WaitKey(0)
}
