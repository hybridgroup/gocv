// What it does:
//
// This example shows how to perform pose detection using models from OpenPose, an open source
// human body, hand, and facial keypoint detector.
//
// For more information about OpenPose, please go to:
// https://github.com/CMU-Perceptual-Computing-Lab/openpose
//
// Before using running this example, you'll need to download a pretrained Caffe model,
// and the respective prototxt config file.
//
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/pose/coco/pose_iter_440000.caffemodel
// https://raw.githubusercontent.com/opencv/opencv_extra/master/testdata/dnn/openpose_pose_coco.prototxt
//
// You can also try the hand pose model:
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/hand/pose_iter_102000.caffemodel
// https://raw.githubusercontent.com/CMU-Perceptual-Computing-Lab/openpose/master/models/hand/pose_deploy.prototxt
//
//
// How to run:
//
// go run ./cmd/dnn-pose-detection/main.go 0 ~/Downloads/pose_iter_440000.caffemodel ~/Downloads/openpose_pose_coco.prototxt openvino fp16
//
// +build example

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

var net *gocv.Net
var images chan *gocv.Mat
var poses chan [][]image.Point
var pose [][]image.Point

func main() {
	if len(os.Args) < 4 {
		fmt.Println("How to run:\ndnn-pose-detection [videosource] [modelfile] [configfile] ([backend] [device])")
		return
	}

	// parse args
	deviceID := os.Args[1]
	proto := os.Args[2]
	model := os.Args[3]
	backend := gocv.NetBackendDefault
	if len(os.Args) > 4 {
		backend = gocv.ParseNetBackend(os.Args[4])
	}

	target := gocv.NetTargetCPU
	if len(os.Args) > 5 {
		target = gocv.ParseNetTarget(os.Args[5])
	}

	// open capture device
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("DNN Pose Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	// open OpenPose model
	n := gocv.ReadNet(model, proto)
	net = &n
	if net.Empty() {
		fmt.Printf("Error reading network model from : %v %v\n", model, proto)
		return
	}
	defer net.Close()
	net.SetPreferableBackend(gocv.NetBackendType(backend))
	net.SetPreferableTarget(gocv.NetTargetType(target))

	fmt.Printf("Start reading device: %v\n", deviceID)

	images = make(chan *gocv.Mat, 1)
	poses = make(chan [][]image.Point)

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Error cannot read device %v\n", deviceID)
		return
	}

	processFrame(&img)

	go performDetection()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		select {
		case pose = <-poses:
			// we've received the next pose from channel, so send next image frame for detection
			processFrame(&img)

		default:
			// show current frame without blocking, so do nothing here
		}

		drawPose(&img)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

func processFrame(i *gocv.Mat) {
	frame := gocv.NewMat()
	i.CopyTo(&frame)
	images <- &frame
}

// performDetection analyzes the results from the detector network.
// the result is an array of "heatmaps" which are the probability
// of a body part being in location x,y
func performDetection() {
	for {
		// get next frame from channel
		frame := <-images

		// convert image Mat to 368x368 blob that the pose detector can analyze
		blob := gocv.BlobFromImage(*frame, 1.0/255.0, image.Pt(368, 368), gocv.NewScalar(0, 0, 0, 0), false, false)

		// feed the blob into the detector
		net.SetInput(blob, "")

		// run a forward pass thru the network
		prob := net.Forward("")

		var midx int
		s := prob.Size()
		nparts, h, w := s[1], s[2], s[3]

		// find out, which model we have
		switch nparts {
		case 19:
			// COCO body
			midx = 0
			nparts = 18 // skip background
		case 16:
			// MPI body
			midx = 1
			nparts = 15 // skip background
		case 22:
			// hand
			midx = 2
		default:
			fmt.Println("there should be 19 parts for the COCO model, 16 for MPI, or 22 for the hand model")
			return
		}

		// find the most likely match for each part
		pts := make([]image.Point, 22)
		for i := 0; i < nparts; i++ {
			pts[i] = image.Pt(-1, -1)
			heatmap, _ := prob.FromPtr(h, w, gocv.MatTypeCV32F, 0, i)

			_, maxVal, _, maxLoc := gocv.MinMaxLoc(heatmap)
			if maxVal > 0.1 {
				pts[i] = maxLoc
			}
			heatmap.Close()
		}

		// determine scale factor
		sX := int(float32(frame.Cols()) / float32(w))
		sY := int(float32(frame.Rows()) / float32(h))

		// create the results array of pairs of points with the lines that best fit
		// each body part, e.g.
		// [[point A for body part 1, point B for body part 1],
		//  [point A for body part 2, point B for body part 2], ...]
		results := [][]image.Point{}
		for _, p := range PosePairs[midx] {
			a := pts[p[0]]
			b := pts[p[1]]

			// high enough confidence in this pose?
			if a.X <= 0 || a.Y <= 0 || b.X <= 0 || b.Y <= 0 {
				continue
			}

			// scale to image size
			a.X *= sX
			a.Y *= sY
			b.X *= sX
			b.Y *= sY

			results = append(results, []image.Point{a, b})
		}
		prob.Close()
		blob.Close()
		frame.Close()

		// send pose results in channel
		poses <- results
	}
}

func drawPose(frame *gocv.Mat) {
	for _, pts := range pose {
		gocv.Line(frame, pts[0], pts[1], color.RGBA{0, 255, 0, 0}, 2)
		gocv.Circle(frame, pts[0], 3, color.RGBA{0, 0, 200, 0}, -1)
		gocv.Circle(frame, pts[1], 3, color.RGBA{0, 0, 200, 0}, -1)
	}
}

// PosePairs is a table of the body part connections in the format [model_id][pair_id][from/to]
// For details please see:
// https://github.com/CMU-Perceptual-Computing-Lab/openpose/blob/master/doc/output.md
//
var PosePairs = [3][20][2]int{
	{ // COCO body
		{1, 2}, {1, 5}, {2, 3},
		{3, 4}, {5, 6}, {6, 7},
		{1, 8}, {8, 9}, {9, 10},
		{1, 11}, {11, 12}, {12, 13},
		{1, 0}, {0, 14},
		{14, 16}, {0, 15}, {15, 17},
	},
	{ // MPI body
		{0, 1}, {1, 2}, {2, 3},
		{3, 4}, {1, 5}, {5, 6},
		{6, 7}, {1, 14}, {14, 8}, {8, 9},
		{9, 10}, {14, 11}, {11, 12}, {12, 13},
	},
	{ // hand
		{0, 1}, {1, 2}, {2, 3}, {3, 4}, // thumb
		{0, 5}, {5, 6}, {6, 7}, {7, 8}, // pinkie
		{0, 9}, {9, 10}, {10, 11}, {11, 12}, // middle
		{0, 13}, {13, 14}, {14, 15}, {15, 16}, // ring
		{0, 17}, {17, 18}, {18, 19}, {19, 20}, // small
	}}
