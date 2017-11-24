// What it does:
//
// This example uses the Window class to open an image file, and then display
// the image in a Window class.
//
// How to run:
//
// 		go run ./cmd/showimage/main.go /home/ron/Pictures/mcp23017.jpg
//
// +build example

package main

import (
	"os"

	"gocv.io/x/gocv"
)

func main() {
	filename := os.Args[1]
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)

	for {
		window.IMShow(img)
		window.WaitKey(1)
	}
}
