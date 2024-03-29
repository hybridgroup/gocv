//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core490 -lopencv_face490 -lopencv_videoio490 -lopencv_imgproc490 -lopencv_highgui490 -lopencv_imgcodecs490 -lopencv_objdetect490 -lopencv_features2d490 -lopencv_video490 -lopencv_dnn490 -lopencv_xfeatures2d490 -lopencv_plot490 -lopencv_tracking490 -lopencv_img_hash490 -lopencv_calib3d490 -lopencv_bgsegm490 -lopencv_photo490 -lopencv_aruco490 -lopencv_wechat_qrcode490 -lopencv_ximgproc490
*/
import "C"
