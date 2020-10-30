// What it does:
//
// This example shows how to find chessboard patterns in an image
//
// How to run:
//
// 		go run ./cmd/find-chessboard/main.go ./images/chessboard_4x6.png
//
// +build example

package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tfind-chessboard [imgfile]")
		return
	}

	filename := os.Args[1]
	img := gocv.IMRead(filename, gocv.IMReadColor)

	if img.Empty() {
		fmt.Printf("Error reading chessboard image")
		return
	}
	defer img.Close()

	corners := gocv.NewMat()
	defer corners.Close()
	sz := image.Point{X: 4, Y: 6}
	found := gocv.FindChessboardCorners(img, sz, &corners, 0)
	if found == false {
		fmt.Printf("chessboard pattern not found")
		return
	}
	if corners.Empty() {
		fmt.Printf("corners mat is empty")
		return
	}

	fmt.Printf("Corners Found. Size: %+v Rows: %+v Cols: %+v\n", corners.Size(), corners.Rows(), corners.Cols())
	clone := img.Clone()
	defer clone.Close()
	gocv.DrawChessboardCorners(&clone, sz, corners, found)
	if clone.Empty() {
		fmt.Printf("Error writing to chessboard image")
		return
	}

	window := gocv.NewWindow("Chessboards")
	defer window.Close()
	for {
		window.IMShow(clone)

		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
