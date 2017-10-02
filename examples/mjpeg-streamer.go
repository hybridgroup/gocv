// What it does:
//
// This example opens a video capture device, then streams MJPEG from it.
//
// How to run:
//
// mjoeg-streamer [camera ID]
//
// 		go run ./examples/mjoeg-streamer.go 1
//
// +build example

package main

import (
	"log"
	"net/http"

	"fmt"
	"os"
	"strconv"

	opencv3 ".."
	"github.com/saljam/mjpeg"
)

var (
	deviceID int
	err      error
	webcam   opencv3.VideoCapture
	img      opencv3.Mat

	stream *mjpeg.Stream
)

func capture() {
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		buf, _ := opencv3.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmjpeg-streamer [camera ID]")
		return
	}

	// parse args
	deviceID, _ = strconv.Atoi(os.Args[1])

	// open webcam
	webcam, err = opencv3.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// prepare image matrix
	img = opencv3.NewMat()
	defer img.Close()

	// create the mjpeg stream
	stream = mjpeg.NewStream()

	// start capturing
	go capture()

	// start http server
	http.Handle("/webcam", stream)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
