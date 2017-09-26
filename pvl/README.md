# Using the Intel Photography Vision Library

The Intel [Photography Vision Library (PVL)](https://software.intel.com/en-us/cvsdk-devguide-advanced-face-capabilities-in-intels-opencv) is a set of extensions to OpenCV that is installed with the Intel CV SDK. It uses computer vision and imaging algorithms developed at Intel.

## How to use

```go
package main

import (
	"fmt"

	opencv3 ".."
	pvl "../pvl"
)

func main() {
	deviceID := 0
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(deviceID); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	window := opencv3.NewWindow("PVL")

	img := opencv3.NewMat()
	defer img.Delete()

	fd := pvl.NewFaceDetector()
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

		faces := fd.DetectFaceRect(img)

		fmt.Printf("found %d\n", len(faces))
		if len(faces) > 0 {
			rects := []opencv3.Rect{faces[0].Rect()}
			opencv3.DrawRectsToImage(img, rects)
		}

		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
```

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
export CGO_CPPFLAGS="-I${INTEL_CVSDK_DIR}/opencv/include" CGO_LDFLAGS="-L${INTEL_CVSDK_DIR}/opencv/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_pvl -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"
```

Run the `showinfo.go` example to make sure you are compiling/linking against the Intel CV SDK:

```
$ go run ./examples/showinfo.go 
go-opencv3 version: 0.0.1
opencv lib version: 3.3.0-cvsdk.604
```
