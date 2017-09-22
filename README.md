# Go OpenCV3

Go bindings for the [OpenCV 3](http://opencv.org/) computer vision package.

Supports the latest OpenCV v3.3

## How to use

This example opens a capture device and output window, and then displays the camera in the window:


```go
package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	deviceID := 0
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(int(deviceID)); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
		return
	}

	window := opencv3.NewWindow("Capture")

	img := opencv3.NewMat()
	defer img.Delete()

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
```

## How to install OpenCV 3.x

### Ubuntu/Linux

You will need to install from source.

### OS X

Instructions needed...

### Windows

Instructions needed...

## How to build/run code

### Ubuntu/Linux

In order to build/run Go code that uses this package, you will need to specify the location for the includes and libs for your OpenCV3 installation.

Once way to find out is to use the `pkg-config` tools like this:

```
$ pkg-config --cflags opencv
-I/opt/intel/computer_vision_sdk_2017.0.113/opencv

$ pkg-config --libs opencv        
-L/opt/intel/computer_vision_sdk_2017.0.113/opencv/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_pvl -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d
```

Once you have this info, you can build or run the Go code that consumes it by populating the needed `CGO_CPPFLAGS` and `CGO_LDFLAGS` env vars.

```
$ export CGO_CPPFLAGS="-I/opt/intel/computer_vision_sdk_2017.0.113/opencv/include" CGO_LDFLAGS="-L/opt/intel/computer_vision_sdk_2017.0.113/opencv/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_pvl -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"

$ go run ./examples/showinfo.go 
go-opencv3 version: 0.0.1
opencv lib version: 3.3.0-cvsdk.604
```

### OS X

Instructions here...

### Windows

Instructions here...

## Why this project exists

The [https://github.com/go-opencv/go-opencv](https://github.com/go-opencv/go-opencv) package for Go and OpenCV does not support any version above OpenCV 2.x, and work on adding support for OpenCV 3 has stalled mostly due to the complexity of SWIG.

This package uses C-style wrapper around the OpenCV 3 C++ classes to avoid having to deal with applying SWIG to a huge existing codebase.

The mappings are intended to match as close as possible to the original OpenCV project structure, to make it easier to find where to add further support.

For example, the [OpenCV `videoio` module](https://github.com/opencv/opencv/tree/master/modules/videoio) wrappers can be found in this project in the `videoio.*` files.

This package was influenced by the blog post https://medium.com/@peterleyssens/using-opencv-3-from-golang-5510c312a3c and the repo at https://github.com/sensorbee/opencv thank you!

## License

Licensed under the Apache 2.0 license. Copyright (c) 2017 The Hybrid Group.
