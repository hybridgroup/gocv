# Go OpenCV3

Go bindings for the [OpenCV 3](http://opencv.org/) computer vision package.

Supports the latest OpenCV v3.3

## How to use

This example opens a video capture device using device "0". It also uses the CascadeClassifier class, after having loaded an external data file containing the classifier data. The program grabs each frame from the video, then uses the classifier to detect faces. If any faces are found, it draws a green rectangle around each one, then displays the image in an output window:

```go
package main

import (
	"fmt"

	opencv3 "github.com/hybridgroup/go-opencv3"
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

	classifier := opencv3.NewCascadeClassifier()
	classifier.Load("data/haarcascade_frontalface_default.xml")

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d\n", len(rects))
		if len(rects) > 0 {
			opencv3.DrawRectsToImage(img, rects)
		}

		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
```

## How to install OpenCV 3.x

### Ubuntu/Linux

You can use `make` to install OpenCV 3.3 with the handy `Makefile` included with this repo. If you already have installed it, you do not need to do so again. Also the installation performed by the `Makefile` may remove options that you have already installed.

#### Install required packages

First, you need to update the system, and install any required packages:

		make deps

#### Download source

Next, download the OpenCV 3.3 and OpenCV Contrib source code:

		make download

#### Build

Build and install everything. This will take quite a while:

		make build

#### Cleanup extra files

After the installation is complete, you can remove the extra files and folders:

		make cleanup

### OS X

You can install OpenCV 3.3 using Homebrew:

		brew install opencv

### Windows

Instructions needed...

## How to build/run code

### Ubuntu/Linux

In order to build/run Go code that uses this package, you will need to specify the location for the includes and libs for your OpenCV3 installation.

For most Linux installations, you can simply run the script:

		./setupvars.sh

Now you should be able to build or run any of the examples:

		go run ./examples/showinfo.go

The showinfo.go program should output the following:

		go-opencv3 version: 0.0.1
		opencv lib version: 3.3.0
		
### Other Linux installations

One way to find out thelocations for your includes and libs is to use the `pkg-config` tool like this:

		pkg-config --cflags opencv                                            
		-I/usr/local/include/opencv -I/usr/local/include

		pkg-config --libs opencv
		-L/usr/local/lib -lopencv_stitching -lopencv_superres -lopencv_videostab -lopencv_photo -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dpm -lopencv_face -lopencv_freetype -lopencv_fuzzy -lopencv_img_hash -lopencv_line_descriptor -lopencv_optflow -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_dnn -lopencv_plot -lopencv_ml -lopencv_xfeatures2d -lopencv_shape -lopencv_video -lopencv_ximgproc -lopencv_calib3d -lopencv_features2d -lopencv_highgui -lopencv_videoio -lopencv_flann -lopencv_xobjdetect -lopencv_imgcodecs -lopencv_objdetect -lopencv_xphoto -lopencv_imgproc -lopencv_core

Once you have this info, you can build or run the Go code that consumes it by populating the needed `CGO_CPPFLAGS` and `CGO_LDFLAGS` ENV vars.

For example:

		export CGO_CPPFLAGS="-I/usr/local/include" 
		export CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"

Please note that you will need to run these 2 lines of code one time in your current session in order to build or run the code, in order to setup the needed ENV variables.

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
