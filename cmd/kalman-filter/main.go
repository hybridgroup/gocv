// What it does:
//
// This example implements the kalman filter example from:
//
//	https://docs.opencv.org/4.6.0/de/d70/samples_2cpp_2kalman_8cpp-example.html#_a7
//
// Tracking of rotating point.
// Point moves in a circle and is characterized by a 1D state.
//
// state_k+1 = state_k + speed + process_noise N(0, 1e-5)
//
// The speed is constant.
// Both state and measurements vectors are 1D (a point angle).
// Measurement is the real state + gaussian noise N(0, 1e-1).
// The real and the measured points are connected with red line segment.
// The real and the estimated points are connected with yellow line segment.
// The real and the corrected estimated points are connected with green line segment.
//
// If Kalman filter works correctly, the yellow segment should be shorter than
// the red one andthe green segment should be shorter than the yellow one).
//
// Pressing any key (except ESC) will reset the tracking.
// Pressing ESC will stop the program.
//
// How to run:
//
//	go run ./cmd/kalman-filter/main.go
package main

import (
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

var (
	colorCyan   = color.RGBA{0, 255, 255, 1}
	colorYellow = color.RGBA{255, 255, 0, 1}
	colorGreen  = color.RGBA{0, 255, 0, 1}
	colorRed    = color.RGBA{255, 0, 0, 1}
	colorWhite  = color.RGBA{255, 255, 255, 1}
	calcPoint   = func(center gocv.Point2f, R, angle float32) gocv.Point2f {
		x := float32(math.Cos(float64(angle))) * R
		y := -float32(math.Sin(float64(angle))) * R
		return gocv.Point2f{X: center.X + x, Y: center.Y + y}
	}
)

func main() {
	var code int
	window := gocv.NewWindow("Kalman")
	img := gocv.NewMatWithSize(500, 500, gocv.MatTypeCV8UC3)
	kf := gocv.NewKalmanFilterWithParams(2, 1, 0, gocv.MatTypeCV32F)
	state := gocv.NewMatWithSize(2, 1, gocv.MatTypeCV32F)
	processNoise := gocv.NewMatWithSize(2, 1, gocv.MatTypeCV32F)
	measurement := gocv.Zeros(1, 1, gocv.MatTypeCV32F)
	for {
		img = gocv.Zeros(500, 500, gocv.MatTypeCV8UC3)
		state.SetFloatAt(0, 0, 0)
		state.SetFloatAt(0, 1, 2*math.Pi/6)
		transitionMatrix := gocv.NewMatWithSize(2, 2, gocv.MatTypeCV32F)
		transitionMatrix.SetFloatAt(0, 0, 1)
		transitionMatrix.SetFloatAt(0, 1, 1)
		transitionMatrix.SetFloatAt(1, 0, 0)
		transitionMatrix.SetFloatAt(1, 1, 1)
		kf.SetTransitionMatrix(transitionMatrix)

		gocv.SetIdentity(kf.GetMeasurementMatrix(), 1)
		gocv.SetIdentity(kf.GetProcessNoiseCov(), 1e-5)
		gocv.SetIdentity(kf.GetMeasurementNoiseCov(), 1e-1)
		gocv.SetIdentity(kf.GetErrorCovPost(), 1)

		statePost := kf.GetStatePost()
		gocv.RandN(&statePost, gocv.NewScalar(0, 0, 0, 0), gocv.NewScalar(0.1, 0.1, 0.1, 0.1))

		for {
			center := gocv.Point2f{X: float32(img.Cols()) * 0.5, Y: float32(img.Rows()) * 0.5}
			R := float32(img.Cols() / 3)
			stateAngle := state.GetFloatAt(0, 0)
			statePt := calcPoint(center, R, stateAngle)

			prediction := kf.Predict()
			predictAngle := prediction.GetFloatAt(0, 0)
			predictPt := calcPoint(center, R, predictAngle)

			// generate measurement
			measNoiseCov := kf.GetMeasurementNoiseCov()
			noise := float64(measNoiseCov.GetFloatAt(0, 0))
			gocv.RandN(&measurement, gocv.NewScalar(0, 0, 0, 0), gocv.NewScalar(noise, noise, noise, noise))
			measMatrix := kf.GetMeasurementMatrix()
			gocv.Add(measurement, measMatrix.MultiplyMatrix(state), &measurement)

			measAngle := measurement.GetFloatAt(0, 0)
			measPt := calcPoint(center, R, measAngle)

			// correct the state estimates based on measurements
			// updates statePost & errorCovPost
			statePost := kf.Correct(measurement)
			improvedAngle := statePost.GetFloatAt(0, 0)
			improvedPt := calcPoint(center, R, improvedAngle)

			// plot points
			img.MultiplyFloat(0.2)
			gocv.Circle(&img, image.Point{int(measPt.X), int(measPt.Y)}, 1, colorRed, 3)
			gocv.Circle(&img, image.Point{int(predictPt.X), int(predictPt.Y)}, 1, colorYellow, 3)
			gocv.Circle(&img, image.Point{int(improvedPt.X), int(improvedPt.Y)}, 1, colorGreen, 3)
			gocv.Circle(&img, image.Point{int(statePt.X), int(statePt.Y)}, 1, colorWhite, 4)

			// forecast one step
			test := transitionMatrix.MultiplyMatrix(kf.GetStatePost())
			newPt := calcPoint(center, R, test.GetFloatAt(0, 0))
			gocv.Circle(&img, image.Point{int(newPt.X), int(newPt.Y)}, 1, colorCyan, 6)

			gocv.Line(&img, image.Point{int(statePt.X), int(statePt.Y)}, image.Point{int(measPt.X), int(measPt.Y)}, colorRed, 1)
			gocv.Line(&img, image.Point{int(statePt.X), int(statePt.Y)}, image.Point{int(predictPt.X), int(predictPt.Y)}, colorYellow, 1)
			gocv.Line(&img, image.Point{int(statePt.X), int(statePt.Y)}, image.Point{int(improvedPt.X), int(improvedPt.Y)}, colorGreen, 1)

			noiseCovMat := kf.GetProcessNoiseCov()
			noiseCov := math.Sqrt(float64(noiseCovMat.GetFloatAt(0, 0)))
			gocv.RandN(&processNoise, gocv.Scalar{}, gocv.NewScalar(noiseCov, noiseCov, noiseCov, noiseCov))
			txm := kf.GetTransitionMatrix()
			gocv.Add(txm.MultiplyMatrix(state), processNoise, &state)

			window.IMShow(img)
			code = window.WaitKey(1000)
			if code > 0 {
				break
			}
		}
		if code == 27 || code == 'q' || code == 'Q' {
			break
		}
	}
}
