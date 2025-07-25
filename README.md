# GoCV

[![GoCV](https://raw.githubusercontent.com/hybridgroup/gocv/release/images/gocvlogo.jpg)](http://gocv.io/)

[![Go Reference](https://pkg.go.dev/badge/gocv.io/x/gocv.svg)](https://pkg.go.dev/gocv.io/x/gocv)
[![Linux](https://github.com/hybridgroup/gocv/actions/workflows/linux.yml/badge.svg?branch=dev)](https://github.com/hybridgroup/gocv/actions/workflows/linux.yml)
[![macOS](https://github.com/hybridgroup/gocv/actions/workflows/macos.yml/badge.svg?branch=dev)](https://github.com/hybridgroup/gocv/actions/workflows/macos.yml)
[![Windows](https://github.com/hybridgroup/gocv/actions/workflows/windows.yml/badge.svg?branch=dev)](https://github.com/hybridgroup/gocv/actions/workflows/windows.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hybridgroup/gocv)](https://goreportcard.com/report/github.com/hybridgroup/gocv)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/hybridgroup/gocv/blob/release/LICENSE.txt)

The GoCV package provides Go language bindings for the [OpenCV 4](http://opencv.org/) computer vision library.

The GoCV package supports the latest releases of Go and OpenCV (v4.12.0) on Linux, Docker, macOS, and Windows. Our ongoing mission is help the Go programming language be a "first-class" client compatible with the latest developments in the OpenCV ecosystem.

GoCV supports [CUDA](https://en.wikipedia.org/wiki/CUDA) for hardware acceleration using Nvidia GPUs. Check out the [CUDA README](./cuda/README.md) for more info on how to use GoCV with OpenCV/CUDA.

GoCV also supports [Intel OpenVINO](https://software.intel.com/en-us/openvino-toolkit). Check out the [OpenVINO README](./openvino/README.md) for more info on how to use GoCV with the Intel OpenVINO toolkit.

## How to use

### Hello, video

This example opens a video capture device using device "0", reads frames, and shows the video in a GUI window:

```go
package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
```

### Face detect

![GoCV](https://raw.githubusercontent.com/hybridgroup/gocv/release/images/face-detect.jpg)

This is a more complete example that opens a video capture device using device "0". It also uses the CascadeClassifier class to load an external data file containing the classifier data. The program grabs each frame from the video, then uses the classifier to detect faces. If any faces are found, it draws a green rectangle around each one, then displays the video in an output window:

```go
package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
    // set to use a video capture device 0
    deviceID := 0

	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("data/haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)
	}
}
```

### More examples

There are examples in the [cmd directory](./cmd) of this repo in the form of various useful command line utilities, such as [capturing an image file](./cmd/saveimage), [streaming mjpeg video](./cmd/mjpeg-streamer), [counting objects that cross a line](./cmd/counter), and [using OpenCV with a DNN for face tracking](./cmd/facedetectYN).

## How to install

To install GoCV, you must first have the matching version of OpenCV installed on your system. The current release of GoCV requires OpenCV 4.12.0.

We have instructions for Linux, macOS, Windows, and other platform options as well.

### Linux

Please see our web site at https://gocv.io/getting-started/linux/

### macOS

Please see our web site at https://gocv.io/getting-started/macos/

### Windows

Please see our web site at https://gocv.io/getting-started/windows/

### Docker

Please see our web site at https://gocv.io/getting-started/docker/

### Android

There is some work in progress for running GoCV on Android using Gomobile. For information on how to install OpenCV/GoCV for Android, please see:
https://gist.github.com/ogero/c19458cf64bd3e91faae85c3ac8874120

See original discussion here:
https://github.com/hybridgroup/gocv/issues/235

## Profiling

Since memory allocations for images in GoCV are done through C based code, the go garbage collector will not clean all resources associated with a `Mat`. As a result, any `Mat` created _must_ be closed to avoid memory leaks.

To ease the detection and repair of the resource leaks, GoCV provides a `Mat` profiler that records when each `Mat` is created and closed. Each time a `Mat` is allocated, the stack trace is added to the profile. When it is closed, the stack trace is removed. See the [runtime/pprof documentation](https://golang.org/pkg/runtime/pprof/#Profile).

In order to include the MatProfile custom profiler, you MUST build or run your application or tests using the `-tags matprofile` build tag. For example:

    go run -tags matprofile cmd/version/main.go

You can get the profile's count at any time using:

```go
gocv.MatProfile.Count()
```

You can display the current entries (the stack traces) with:

```go
var b bytes.Buffer
gocv.MatProfile.WriteTo(&b, 1)
fmt.Print(b.String())
```

This can be very helpful to track down a leak. For example, suppose you have
the following nonsense program:

```go
package main

import (
	"bytes"
	"fmt"

	"gocv.io/x/gocv"
)

func leak() {
	gocv.NewMat()
}

func main() {
	fmt.Printf("initial MatProfile count: %v\n", gocv.MatProfile.Count())
	leak()

	fmt.Printf("final MatProfile count: %v\n", gocv.MatProfile.Count())
	var b bytes.Buffer
	gocv.MatProfile.WriteTo(&b, 1)
	fmt.Print(b.String())
}
```

Running this program produces the following output:

```
initial MatProfile count: 0
final MatProfile count: 1
gocv.io/x/gocv.Mat profile: total 1
1 @ 0x40b936c 0x40b93b7 0x40b94e2 0x40b95af 0x402cd87 0x40558e1
#	0x40b936b	gocv.io/x/gocv.newMat+0x4b	/go/src/gocv.io/x/gocv/core.go:153
#	0x40b93b6	gocv.io/x/gocv.NewMat+0x26	/go/src/gocv.io/x/gocv/core.go:159
#	0x40b94e1	main.leak+0x21			/go/src/github.com/dougnd/gocvprofexample/main.go:11
#	0x40b95ae	main.main+0xae			/go/src/github.com/dougnd/gocvprofexample/main.go:16
#	0x402cd86	runtime.main+0x206		/usr/local/Cellar/go/1.11.1/libexec/src/runtime/proc.go:201
```

We can see that this program would leak memory. As it exited, it had one `Mat` that was never closed. The stack trace points to exactly which line the allocation happened on (line 11, the `gocv.NewMat()`).

Furthermore, if the program is a long running process or if GoCV is being used on a web server, it may be helpful to install the HTTP interface )). For example:

```go
package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"gocv.io/x/gocv"
)

func leak() {
	gocv.NewMat()
}

func main() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			leak()
		}
	}()

	http.ListenAndServe("localhost:6060", nil)
}

```

This will leak a `Mat` once per second. You can see the current profile count and stack traces by going to the installed HTTP debug interface: [http://localhost:6060/debug/pprof/gocv.io/x/gocv.Mat](http://localhost:6060/debug/pprof/gocv.io/x/gocv.Mat?debug=1).

## How to contribute

Please take a look at our [CONTRIBUTING.md](./CONTRIBUTING.md) document to understand our contribution guidelines.

Then check out our [ROADMAP.md](./ROADMAP.md) document to know what to work on next.

## Why this project exists

The [https://github.com/go-opencv/go-opencv](https://github.com/go-opencv/go-opencv) package for Go and OpenCV did not support any version above OpenCV 2.x, and work on adding support for OpenCV 3 had stalled for over a year, mostly due to the complexity of [SWIG](http://swig.org/). That is why we started this project.

The GoCV package uses a C-style wrapper around the OpenCV 4 C++ classes to avoid having to deal with applying SWIG to a huge existing codebase. The mappings are intended to match as closely as possible to the original OpenCV project structure, to make it easier to find things, and to be able to figure out where to add support to GoCV for additional OpenCV image filters, algorithms, and other features.

For example, the [OpenCV `videoio` module](https://github.com/opencv/opencv/tree/master/modules/videoio) wrappers can be found in the GoCV package in the `videoio.*` files.

This package was inspired by the original https://github.com/go-opencv/go-opencv project, the blog post https://medium.com/@peterleyssens/using-opencv-3-from-golang-5510c312a3c and the repo at https://github.com/sensorbee/opencv thank you all!

## License

Licensed under the Apache 2.0 license. Copyright (c) 2017-2025 The Hybrid Group.

Logo generated by GopherizeMe - https://gopherize.me
