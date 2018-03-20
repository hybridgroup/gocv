// What it does:
//
// This example shows, how to use pose detection,
// using a simplified pretrained network from http://posefs1.perception.cs.cmu.edu/OpenPose.
//
// (Simplified here means: the original model also has detailed hand data, access to it was removed (for speed) in the resp. prototxt
//  also, while the models still have the data to detect poses for multiple persons, we restrict it to a single one here)
//
// Before using this, you'll need to download a pretrained caffemodel, and the respective prototxt:
//
// There is a model trained on the COCO dataset with 19 body parts, and 17 connections between them:
//
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/pose/coco/pose_iter_440000.caffemodel
// https://github.com/opencv/opencv_extra/blob/master/testdata/dnn/openpose_pose_coco.prototxt
//
// Also, there's one based on MPI data with 16 parts and 14 connections:
//
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/pose/mpi/pose_iter_160000.caffemodel
// https://raw.githubusercontent.com/opencv/opencv_extra/master/testdata/dnn/openpose_pose_mpi.prototxt
//
// (beware: ~200 mb download for any the caffemodels)
//
// See here for an explanation of the body parts:
// https://github.com/CMU-Perceptual-Computing-Lab/openpose/blob/master/src/openpose/pose/poseParameters.cpp#L7-L45
//
// How to run it:
//
// 		find an image with a *single* person in it, whole body, if possible, then:
//
// 		go run ./cmd/dnn-advanced/pose.go [imgpath] [protofile] [modelfile] [threshold(optional)]
//
// +build example

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"

	"gocv.io/x/gocv"
)

func main() {
	// parse args
	if len(os.Args) < 4 {
		fmt.Println("How to run:\ndnn-pose [imgpath] [protofile] [modelfile] [threshold(optional)]")
		return
	}
	imgpath := os.Args[1]
	proto := os.Args[2]
	model := os.Args[3]

	thresh := 0.1
	if len(os.Args) > 4 {
		thresh, _ = strconv.ParseFloat(os.Args[4], 64)
	}

	// load the image
	img := gocv.IMRead(imgpath, 1)
	if img.Empty() {
		fmt.Printf("Error loading image from: %v\n", imgpath)
		return
	}
	defer img.Close()

	// show initial image, we'll overlay the pose drawing later
	window := gocv.NewWindow("Dnn Pose Estimation with gocv")
	defer window.Close()
	window.IMShow(img)
	window.WaitKey(10)

	// open DNN model
	net := gocv.ReadNetFromCaffe(proto, model)
	if net.Empty() {
		fmt.Printf("Error reading network model from : %v %v\n", proto, model)
		return
	}
	defer net.Close()

	// convert image Mat to network blob
	blob := gocv.BlobFromImage(img, 1.0/255, image.Pt(368, 368), gocv.NewScalar(0, 0, 0, 0), false, false)
	defer blob.Close()

	// feed the blob into the classifier
	net.SetInput(blob, "")

	// run a forward pass through the network
	fmt.Println("Please be patient, we'll be crunching a lot of numbers here !")
	detBlob := net.Forward("")
	defer detBlob.Close()

	// careful, it is: (N,C,H,W)
	siz := gocv.GetBlobSize(detBlob)
	nparts := int(siz.Val2)

	// adjust params for the resp. model we're using
	pidx, npairs := 0, 0
	if nparts == 19 {
		// COCO
		pidx = 0
		npairs = 17
	} else if nparts == 16 {
		// MPI
		pidx = 1
		npairs = 14
	} else {
		// if you ever get here, you'll probably got the wrong prototxt
		fmt.Println("there should be 19 body parts for the COCO model and 16 for the MPI one, but this model has ", nparts, " parts.")
		return
	}

	// connection lookup table,
	// each connection is a bone in the skeleton
	// (a connection between 2 body parts)
	POSE_PAIRS := [2][17][2]int{
		{ // COCO
			{1, 2}, {1, 5}, {2, 3},
			{3, 4}, {5, 6}, {6, 7},
			{1, 8}, {8, 9}, {9, 10},
			{1, 11}, {11, 12}, {12, 13},
			{1, 0}, {0, 14},
			{14, 16}, {0, 15}, {15, 17}},
		{ // MPI
			{0, 1}, {1, 2}, {2, 3},
			{3, 4}, {1, 5}, {5, 6},
			{6, 7}, {1, 14}, {14, 8}, {8, 9},
			{9, 10}, {14, 11}, {11, 12}, {12, 13},
			{-1, -1}, {-1, -1}, {-1, -1}}}

	// find the location of the body parts,
	// (nose, neck, left foot, etc)
	// there is a "heatmap" (probability, where it'll be)
	// for each of them inside a seperate channel of the dnn output
	W := float64(img.Cols())
	H := float64(img.Rows())
	fmt.Println(W, H)
	points := make([]image.Point, nparts)
	for i := 0; i < nparts; i++ {
		// get the heatmap slice for the corresponding body part
		heatmap := gocv.GetBlobChannel(detBlob, 0, i)
		defer heatmap.Close()

		// the original OpenPose code tries to find all the local maxima.
		// to simplify this sample, we restrict it to a single person, and
		// just use the global maximum
		_, confidence, _, pm := gocv.MinMaxLoc(heatmap)

		// scale to image size
		p := image.Point{0, 0} // default for: not enough confidence
		if float64(confidence) > thresh {
			p.X = int((W * float64(pm.X)) / siz.Val4)
			p.Y = int((H * float64(pm.Y)) / siz.Val3)
		}
		points[i] = p
	}

	// finally, draw the skeleton point pairs:
	red := color.RGBA{255, 0, 0, 0}
	green := color.RGBA{0, 255, 0, 0}
	for i := 0; i < npairs; i++ {
		p1 := points[POSE_PAIRS[pidx][i][0]]
		p2 := points[POSE_PAIRS[pidx][i][1]]
		if p1.X == 0 || p1.Y == 0 || p2.X == 0 || p2.Y == 0 {
			// not enough confidence for one of the end points, see above
			continue
		}
		gocv.Line(&img, p1, p2, green, 3)
		gocv.Circle(&img, p1, 3, red, -1)
		gocv.Circle(&img, p2, 3, red, -1)
	}

	window.IMShow(img)
	window.WaitKey(0)
}
