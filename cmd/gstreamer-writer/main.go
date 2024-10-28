/*
	GStreamer Video Writer

What it does:

	Captures video from webcam and outputs the frames to a gocv VideoWriter

How to run:

	go run main.go

To receive the video feed run in a terminal:

	gst-launch-1.0 udpsrc port=5000 ! application/x-rtp, encoding-name=MJPG ! rtpjpegdepay ! jpegdec !  videoconvert ! xvimagesink
*/
package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const (
	gstreamerPipe = "appsrc ! videoconvert ! video/x-raw,format=YUY2,width=640,height=480,framerate=30/1 ! jpegenc ! rtpjpegpay ! udpsink host=127.0.0.1 port=5000"
)

func main() {

	cap, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		panic(err)
	}
	defer cap.Close()
	cap.Set(gocv.VideoCaptureFrameWidth, 640.0)
	cap.Set(gocv.VideoCaptureFrameHeight, 480.0)

	wrt, err := gocv.VideoWriterFileWithAPI(gstreamerPipe,
		gocv.VideoCaptureGstreamer, "YUY2", 30.0, 640, 480, true)
	if err != nil {
		panic(err)
	}
	defer wrt.Close()

	if !wrt.IsOpened() {
		panic("wrt is not opened")
	}

	frame := gocv.NewMat()
	defer frame.Close()

	for {
		cap.Read(&frame)

		if frame.Empty() {
			fmt.Println("empty frame")
			continue
		}

		if err := wrt.Write(frame); err != nil {
			fmt.Println(err)
		}

	}

}
