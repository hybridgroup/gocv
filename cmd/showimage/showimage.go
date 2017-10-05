// Package shoiwiumage uses the Window class to open an image file and display it
package showimage

import (
	"github.com/hybridgroup/gocv"
)

func Run(filename string) {
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	for {
		window.IMShow(img)
		gocv.WaitKey(1)
	}
}
