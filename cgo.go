//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core481 -lopencv_face481 -lopencv_videoio481 -lopencv_imgproc481 -lopencv_highgui481 -lopencv_imgcodecs481 -lopencv_objdetect481 -lopencv_features2d481 -lopencv_video481 -lopencv_dnn481 -lopencv_xfeatures2d481 -lopencv_plot481 -lopencv_tracking481 -lopencv_img_hash481 -lopencv_calib3d481 -lopencv_bgsegm481 -lopencv_photo481 -lopencv_aruco481 -lopencv_wechat_qrcode481 -lopencv_ximgproc481
*/
import "C"
