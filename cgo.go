//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core470 -lopencv_face470 -lopencv_videoio470 -lopencv_imgproc470 -lopencv_highgui470 -lopencv_imgcodecs470 -lopencv_objdetect470 -lopencv_features2d470 -lopencv_video470 -lopencv_dnn470 -lopencv_xfeatures2d470 -lopencv_plot470 -lopencv_tracking470 -lopencv_img_hash470 -lopencv_calib3d470 -lopencv_bgsegm470 -lopencv_photo470 -lopencv_aruco470 -lopencv_wechat_qrcode470 -lopencv_ximgproc470
*/
import "C"
