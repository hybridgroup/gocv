# Go OpenCV3

Go bindings for the [OpenCV 3](http://opencv.org/) computer vision package.

Supports the latest OpenCV v3.3

Uses C-style wrapper around the OpenCV 3 C++ classes to avoid having to deal with applying SWIG to a huge existing codebase.

The mappings are intended to match as close as possible to the original OpenCV project structure.

For example, the [OpenCV `videoio` module](https://github.com/opencv/opencv/tree/master/modules/videoio) wrappers can be found in this project in the `videoio.*` files.

Based on concepts & code from the blog post https://medium.com/@peterleyssens/using-opencv-3-from-golang-5510c312a3c and the repo at https://github.com/sensorbee/opencv

## How to build/run

You will need to specify the location for the includes and libs for your OpenCV3 installation.

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
