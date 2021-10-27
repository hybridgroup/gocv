//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core454 -lopencv_face454 -lopencv_videoio454 -lopencv_imgproc454 -lopencv_highgui454 -lopencv_imgcodecs454 -lopencv_objdetect454 -lopencv_features2d454 -lopencv_video454 -lopencv_dnn454 -lopencv_xfeatures2d454 -lopencv_plot454 -lopencv_tracking454 -lopencv_img_hash454 -lopencv_calib3d454 -lopencv_bgsegm454 -lopencv_photo454 -lopencv_aruco454
*/
import "C"
