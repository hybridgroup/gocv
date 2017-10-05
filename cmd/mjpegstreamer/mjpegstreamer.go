// Package mjpegstreamer opens a video capture device, then streams MJPEG from it.
// Once running point your browser to the hostname/port you passed in the
// command line (for example http://localhost:8080) and you should see
// the live video stream.
package mjpegstreamer

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hybridgroup/gocv"
	"github.com/saljam/mjpeg"
)

var (
	deviceID int
	err      error
	webcam   gocv.VideoCapture
	img      gocv.Mat

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

		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
	}
}

func Run(deviceID int, host string) {
	// open webcam
	webcam, err = gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// prepare image matrix
	img = gocv.NewMat()
	defer img.Close()

	// create the mjpeg stream
	stream = mjpeg.NewStream()

	// start capturing
	go capture()

	// start http server
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(host, nil))
}
