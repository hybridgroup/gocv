# Using the Intel OpenVINO Photography Vision Library

The Intel OpenVINO Photography Vision Library (PVL) is a set of extensions to OpenCV that is installed with the Intel CV SDK. It uses computer vision and imaging algorithms developed at Intel.

GoCV support for the PVL can be found here in the "gocv.io/x/gocv/openvino/pvl" package.

## How to use

```go
package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/openvino/pvl"
)

func main() {
	deviceID := 0

	// open webcam
	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}	
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("PVL")

	// prepare input image matrix
	img := gocv.NewMat()
	defer img.Close()

	// prepare grayscale image matrix
	imgGray := gocv.NewMat()
	defer imgGray.Close()
	
	// color to draw the rect for detected faces
	blue := color.RGBA(0, 0, 255, 0)

	// load PVL FaceDetector to recognize faces
	fd := pvl.NewFaceDetector()
	defer fd.Close()

	// enable tracking mode for more efficient tracking of video source
	fd.SetTrackingModeEnabled(true)

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image to grayscale for detection
		gocv.CvtColor(img, &imgGray, gocv.ColorBGR2GRAY);
	
		// detect faces
		faces := fd.DetectFaceRect(imgGray)
		fmt.Printf("found %d faces\n", len(faces))

		// draw a rectangle around each face on the original image
		for _, face := range faces {
			gocv.Rectangle(&img, face.Rectangle(), blue, 3)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)
	}
}
```

Some PVL examples are in the [cmd/openvino/pvl directory](../cmd/openvino/pvl) of this repo, in the form of some useful commands such as the [smile detector](../cmd/openvino/pvl/smiledetector).

## How to build/run code

Setup the environment for the Intel OpenVINO toolkit, by running the `setupvars.sh` program included with OpenVINO:

```
source /opt/intel/computer_vision_sdk/bin/setupvars.sh
```

Then set the needed other exports for building/running GoCV code by running the `env.sh` that is in the GoCV `openvino` directory:

You only need to do these two steps one time per session. Once you have run them, you do not need to run them again until you close your terminal window.

Now you can run the version command example to make sure you are compiling/linking against Intel OpenVINO:

```
$ go run ./cmd/version/main.go 
gocv version: 0.7.0
opencv lib version: 3.3.1-cvsdk_2017_R3.2
```

Examples that use the Intel OpenVINO toolkit can be found in the `cmd/openvino/pvl` directory of this repository.
