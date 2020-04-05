// +build !customenv

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core430 -lopencv_face430 -lopencv_videoio430 -lopencv_imgproc430 -lopencv_highgui430 -lopencv_imgcodecs430 -lopencv_objdetect430 -lopencv_features2d430 -lopencv_video430 -lopencv_dnn430 -lopencv_xfeatures2d430 -lopencv_plot430 -lopencv_tracking430 -lopencv_img_hash430 -lopencv_calib3d430
*/
import "C"
