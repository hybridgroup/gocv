//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core455 -lopencv_face455 -lopencv_videoio455 -lopencv_imgproc455 -lopencv_highgui455 -lopencv_imgcodecs455 -lopencv_objdetect455 -lopencv_features2d455 -lopencv_video455 -lopencv_dnn455 -lopencv_xfeatures2d455 -lopencv_plot455 -lopencv_tracking455 -lopencv_img_hash455 -lopencv_calib3d455 -lopencv_bgsegm455 -lopencv_photo455 -lopencv_aruco455 -lopencv_wechat_qrcode455
*/
import "C"
