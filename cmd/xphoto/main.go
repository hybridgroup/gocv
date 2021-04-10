// What it does:
//
// This example uses the MergeMerten class to merge LDR image.
// then save it to an HDR image file on disk.
//
// How to run:
//
// saveimage [camera ID] [image file]
//
// 		go run ./cmd/saveimage/main.go 0 filename.jpg
//
// +build example

package main

import (
	"flag"
	"fmt"
	"os"
	// 	"strings"

	"gocv.io/x/gocv"
	// 	"gocv.io/x/gocv/contrib"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\thdrimage [image file 1] [image file 2] [image file 3] hdr_image.png")
		return
	}
	flag.Parse()
	if flag.NArg() < 3 {
		flag.Usage()
		return
	}

	// read images
	inputs := flag.Args()
	inputImages := make([]gocv.Mat, 3)

	for i := 0; i < 3; i++ {
		img := gocv.IMRead(inputs[i], gocv.IMReadColor)
		if img.Empty() {
			fmt.Printf("cannot read image %s\n", inputs[i])
			return
		}
		defer img.Close()
		inputImages[i] = img
	}

	saveFile := inputs[3]
	ouputImage := gocv.NewMat()
	defer ouputImage.Close()

	alignwtb := gocv.NewAlignMTBWithParams(3, 20, false)
	alignwtb.Process(inputImages, inputImages)

	mertens := gocv.NewMergeMertens()
	mertens.Process(inputImages, &ouputImage)

	graywoldwb := gocv.NewGrayworldWB()
	graywoldwb.SetSaturationThreshold(0.7)
	graywoldwb.BalanceWhite(ouputImage, &ouputImage)

	gocv.IMWrite(saveFile, ouputImage)

}
