// What it does:
//
// This example uses Brute-Force Matching
// with SIFT Descriptors and Ratio Test
//
// How to run:
//
// feature-matching [query image] [training image]
//
// 		go run ./cmd/feature-matching/main.go input.jpg training.jpg
//

package main

import (
	"fmt"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: feature-matching /path/to/query /path/to/train")
		panic("error: no files provided")
	}

	// opening query image
	query := gocv.IMRead(os.Args[1], gocv.IMReadGrayScale)
	defer query.Close()

	// opening train image
	train := gocv.IMRead(os.Args[2], gocv.IMReadGrayScale)
	defer train.Close()

	// creating new SIFT
	sift := gocv.NewSIFT()
	defer sift.Close()

	// detecting and computing keypoints using SIFT method
	kp1, des1 := sift.DetectAndCompute(query, gocv.NewMat())
	kp2, des2 := sift.DetectAndCompute(train, gocv.NewMat())

	// finding K best matches for each descriptor
	bf := gocv.NewBFMatcher()
	matches := bf.KnnMatch(des1, des2, 2)

	// application of ratio test
	var good []gocv.DMatch
	for _, m := range matches {
		if len(m) > 1 {
			if m[0].Distance < 0.75*m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	// matches color
	c1 := color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 0,
	}

	// point color
	c2 := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}

	// creating empty mask
	mask := make([]byte, 0)

	// new matrix for output image
	out := gocv.NewMat()
	// drawing matches
	gocv.DrawMatches(query, kp1, train, kp2, good, &out, c1, c2, mask, gocv.DrawDefault)

	// creating output window with result
	window := gocv.NewWindow("Output")
	window.IMShow(out)
	defer window.Close()

	window.WaitKey(0)
}
