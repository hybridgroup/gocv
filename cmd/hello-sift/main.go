package main

import (
	"fmt"
	"image/color"
	"os"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

// you can run this example from root of this repository
// with go run ./cmd/hello-sift /path/to/querry /path/to/train
func main() {
	if len(os.Args) != 3 {
		println("Usage: app /path/to/querry /path/to/train")
	}

	querry := gocv.IMRead(os.Args[1], gocv.IMReadGrayScale)
	defer querry.Close()

	train := gocv.IMRead(os.Args[2], gocv.IMReadGrayScale)
	defer train.Close()

	println("images opened")

	sift := contrib.NewSIFT()
	defer sift.Close()

	println("sift created")

	kp1, des1 := sift.DetectAndCompute(querry, gocv.NewMat())
	kp2, des2 := sift.DetectAndCompute(train, gocv.NewMat())

	println("detect and compute done")

	bf := gocv.NewBFMatcher()
	matches := bf.Knntch(des1, des2, 2)

	println("matches created")

	var good []gocv.DMatch

	for _, m := range matches {
		if len(m) > 1 {
			if m[0].Distance < 0.75*m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	println("good selected")

	fmt.Println(len(good))

	gocv.DrawKeyPoints(querry, kp1, &querry, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)

	gocv.DrawKeyPoints(train, kp2, &train, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)

	// Temporary solution until DrawKeyPointsBGRA comes into mainstream
	gocv.CvtColor(querry, &querry, gocv.ColorBGRToRGBA)
	gocv.CvtColor(train, &train, gocv.ColorBGRToRGBA)

	var out gocv.Mat

	c1 := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}

	c2 := color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 0,
	}

	var mask [len(matches)]char

	gocv.DrawMatches(train, kp1, querry, kp2, good, &out, c1, c2, mask, gocv.DrawDefault)

	window1 := gocv.NewWindow("Query")
	window1.IMShow(querry)

	window2 := gocv.NewWindow("Train")
	window2.IMShow(train)

	window1.WaitKey(0)
	window2.WaitKey(0)

}
