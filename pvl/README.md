# Using the Intel Photographic Vision Library

Info here...

Setup main Intel SDK env:

```
source /opt/intel/computer_vision_sdk_2017.0.113/bin/setupvars.sh
```


Then set the needed other exports:

```
export CGO_CPPFLAGS="-I/opt/intel/computer_vision_sdk_2017.0.113/opencv/include" CGO_LDFLAGS="-L/opt/intel/computer_vision_sdk_2017.0.113/opencv/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_pvl -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"
```

Run the `showinfo.go` example to make sure you are compiling/linking against the Intel CV SDK:

```
$ go run ./examples/showinfo.go 
go-opencv3 version: 0.0.1
opencv lib version: 3.3.0-cvsdk.604
```
