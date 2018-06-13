// What it does:
//
// This example uses pre-trained deep neural network models from OpenPose to perform body pose detection.
// it can be used for body pose detection, using either the COCO model(18 parts):
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/pose/coco/pose_iter_440000.caffemodel
// https://raw.githubusercontent.com/opencv/opencv_extra/master/testdata/dnn/openpose_pose_coco.prototxt
//
// or the MPI model(16 parts):
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/pose/mpi/pose_iter_160000.caffemodel
// https://raw.githubusercontent.com/opencv/opencv_extra/master/testdata/dnn/openpose_pose_mpi_faster_4_stages.prototxt
//
// (to simplify this sample, the body models are restricted to a single person.)
//
// you can also try the hand pose model:
// http://posefs1.perception.cs.cmu.edu/OpenPose/models/hand/pose_iter_102000.caffemodel
// https://raw.githubusercontent.com/CMU-Perceptual-Computing-Lab/openpose/master/models/hand/pose_deploy.prototxt
//
//
// How to run:
//
// 		go run ./cmd/dnn-pose-detection/main.go [videosource] [modelfile] [configfile] ([backend] [device])
//
// +   build example

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

// connection table, in the format [model_id][pair_id][from/to]
// please look at the nice explanation at the bottom of:
// https://github.com/CMU-Perceptual-Computing-Lab/openpose/blob/master/doc/output.md
//
//PosePairs := make(map[string]int)

var POSE_PAIRS = [3][20][2]int{
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
	net := gocv.ReadNet(model, proto)
	if net.Empty() {
		fmt.Printf("Error reading network model from : %v %v\n", model, proto)
		return
	}
	defer net.Close()
	net.SetPreferableBackend(gocv.NetBackendType(backend))
	net.SetPreferableTarget(gocv.NetTargetType(target))

	fmt.Printf("Start reading camera device: %v\n", deviceID)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image Mat to 368x368 blob that the pose detector can analyze
		blob := gocv.BlobFromImage(img, 1.0/255.0, image.Pt(368, 368), gocv.NewScalar(0, 0, 0, 0), false, false)

		// feed the blob into the detector
		net.SetInput(blob, "")

		// run a forward pass thru the network
		prob := net.Forward("")

		performDetection(&img, prob)

		prob.Close()
		blob.Close()

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

// performDetection analyzes the results from the detector network.
// the result is an array of "heatmaps" which are the probability
// of a body part being in location x,y
func performDetection(frame *gocv.Mat, results gocv.Mat) {
	var midx, npairs int
	s := results.Size()
	nparts := s[1]
	h := s[2]
	w := s[3]

	// find out, which model we have
	switch nparts {
	case 19:
		// COCO body
		midx = 0
		npairs = 17
		nparts = 18 // skip background
	case 16:
		// MPI body
		midx = 1
		npairs = 14
	case 22:
		// hand
		midx = 2
		npairs = 20
	default:
		fmt.Println("there should be 19 parts for the COCO model, 16 for MPI, or 22 for the hand model")
		return
	}

	points := make([]image.Point, 22)
	for i := 0; i < nparts; i++ {
		pt := image.Pt(-1, -1)
		heatmap, _ := results.FromPtr(h, w, gocv.MatTypeCV32F, 0, i)

		// 1 maximum per heatmap
		_, maxVal, _, maxLoc := gocv.MinMaxLoc(heatmap)
		if maxVal > 0.1 {
			pt = maxLoc
		}
		points[i] = pt
	}

	// connect body parts and draw
	sX := float32(frame.Cols()) / float32(w)
	sY := float32(frame.Rows()) / float32(h)

	for i := 0; i < npairs; i++ {
		a := points[POSE_PAIRS[midx][i][0]]
		b := points[POSE_PAIRS[midx][i][1]]

		// we did not find enough confidence before
		if a.X <= 0 || a.Y <= 0 || b.X <= 0 || b.Y <= 0 {
			continue
		}

		// scale to image size
		a.X *= int(sX)
		a.Y *= int(sY)
		b.X *= int(sX)
		b.Y *= int(sY)

		gocv.Line(frame, a, b, color.RGBA{0, 255, 0, 0}, 2)
		gocv.Circle(frame, a, 3, color.RGBA{0, 0, 200, 0}, -1)
		gocv.Circle(frame, b, 3, color.RGBA{0, 0, 200, 0}, -1)
	}
}
