// What it does:
//
// This example uses the Window class to open an image file, and then display
// the image in a Window class.
//
// How to run:
//
// 		go run ./examples/window.go /home/ron/Pictures/mcp23017.jpg
//
// +build example

package main

import (
	"os"

	gocv ".."
)

func main() {
	filename := os.Args[1]
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)

	for {
		window.IMShow(img)
		gocv.WaitKey(1)
	}
}
