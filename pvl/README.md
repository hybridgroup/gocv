# Using the Intel Photography Vision Library

The Intel [Photography Vision Library (PVL)](https://software.intel.com/en-us/cvsdk-devguide-advanced-face-capabilities-in-intels-opencv) is a set of extensions to OpenCV that is installed with the Intel CV SDK. It uses computer vision and imaging algorithms developed at Intel.

GoCV support for the PVL can be found here in the "gocv.io/x/gocv/pvl" package.

## How to use

```go
package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/pvl"
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
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image to grayscale for detection
		gocv.CvtColor(img, imgGray, gocv.ColorBGR2GRAY);
	
		// detect faces
		faces := fd.DetectFaceRect(imgGray)
		fmt.Printf("found %d faces\n", len(faces))

		// draw a rectangle around each face on the original image
		for _, face := range faces {
			gocv.Rectangle(img, face.Rectangle(), blue, 3)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)
	}
}
```

Some PVL examples are in the [cmd/pvl directory](../cmd/pvl) of this repo, in the form of some useful commands such as the [smile detector](../cmd/pvl/smiledetector).

## How to install the Intel CV SDK

You can download the Intel CV SDK from here:

https://software.intel.com/en-us/computer-vision-sdk

## How to build/run code

Setup main Intel SDK env:

```
source /opt/intel/computer_vision_sdk_2017.0.113/bin/setupvars.sh
```

Then set the needed other exports:

```
export CGO_CPPFLAGS="-I${INTEL_CVSDK_DIR}/opencv/include" CGO_LDFLAGS="-L${INTEL_CVSDK_DIR}/opencv/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_pvl -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video"
```

Run the version command example to make sure you are compiling/linking against the Intel CV SDK:

```
$ go run ./cmd/version/main.go 
go-gocv version: 0.5.0
opencv lib version: 3.3.0-cvsdk.604
```
