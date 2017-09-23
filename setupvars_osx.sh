#!/bin/bash
export CGO_CPPFLAGS="-I/usr/local/Cellar/opencv/3.3.0/include -I/usr/local/Cellar/opencv/3.3.0/include/opencv2"
export CGO_CXXFLAGS="--std=c++1z -stdlib=libc++"
export CGO_LDFLAGS="-L/usr/local/Cellar/opencv/3.3.0/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"
