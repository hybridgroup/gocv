// What it does:
//
// This example draws two examples, an atom and a rook, based on:
// https://docs.opencv.org/2.4/doc/tutorials/core/basic_geometric_drawing/basic_geometric_drawing.html.
//
// How to run:
//
// 		go run ./cmd/basic-drawing/main.go
//
// +build example

package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

var w = 400

func main() {
	windowA := gocv.NewWindow("basic drawing: atom")
	windowR := gocv.NewWindow("basic drawing: rook")
	defer windowA.Close()
	defer windowR.Close()

	atom := gocv.NewMatWithSize(w, w, gocv.MatTypeCV8UC3)
	defer atom.Close()

	rook := gocv.NewMatWithSize(w, w, gocv.MatTypeCV8UC3)
	defer rook.Close()

	black := color.RGBA{0, 0, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}
	red := color.RGBA{255, 0, 0, 0}
	white := color.RGBA{255, 255, 255, 0}
	yellow := color.RGBA{255, 255, 0, 0}

	// draw the atom
	gocv.Ellipse(&atom, image.Pt(w/2., w/2.), image.Pt(w/4.0, w/16.0), 90., 0, 360, blue, 2)
	gocv.Ellipse(&atom, image.Pt(w/2., w/2.), image.Pt(w/4.0, w/16.0), 0., 0, 360, blue, 2)
	gocv.Ellipse(&atom, image.Pt(w/2., w/2.), image.Pt(w/4.0, w/16.0), 45., 0, 360, blue, 2)
	gocv.Ellipse(&atom, image.Pt(w/2., w/2.), image.Pt(w/4.0, w/16.0), -45., 0, 360, blue, 2)
	gocv.Circle(&atom, image.Pt(w/2., w/2.), w/32., red, -1)

	// draw the rook
	points := [][]image.Point{
		{
			image.Pt(w/4., 7*w/8.),
			image.Pt(3*w/4., 7*w/8.),
			image.Pt(3*w/4., 13*w/16.),
			image.Pt(11*w/16., 13*w/16.),
			image.Pt(19*w/32., 3*w/8.),
			image.Pt(3*w/4., 3*w/8.),
			image.Pt(3*w/4., w/8.),
			image.Pt(26*w/40., w/8.),
			image.Pt(26*w/40., w/4.),
			image.Pt(22*w/40., w/4.),
			image.Pt(22*w/40., w/8.),
			image.Pt(18*w/40., w/8.),
			image.Pt(18*w/40., w/4.),
			image.Pt(14*w/40., w/4.),
			image.Pt(14*w/40., w/8.),
			image.Pt(w/4., w/8.),
			image.Pt(w/4., 3*w/8.),
			image.Pt(13*w/32., 3*w/8.),
			image.Pt(5*w/16., 13*w/16.),
			image.Pt(w/4., 13*w/16.),
		},
	}
	gocv.FillPoly(&rook, points, white)
	gocv.Rectangle(&rook, image.Rect(0, 7*w/8.0, w, w), yellow, -1)
	gocv.Line(&rook, image.Pt(0, 15*w/16), image.Pt(w, 15*w/16), black, 2)
	gocv.Line(&rook, image.Pt(w/4, 7*w/8), image.Pt(w/4, w), black, 2)
	gocv.Line(&rook, image.Pt(w/2, 7*w/8), image.Pt(w/2, w), black, 2)
	gocv.Line(&rook, image.Pt(3*w/4, 7*w/8), image.Pt(3*w/4, w), black, 2)

	for {
		windowA.IMShow(atom)
		windowR.IMShow(rook)

		if windowA.WaitKey(10) >= 0 || windowR.WaitKey(10) >= 0 {
			break
		}
	}
}
