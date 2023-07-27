//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core480 -lopencv_face480 -lopencv_videoio480 -lopencv_imgproc480 -lopencv_highgui480 -lopencv_imgcodecs480 -lopencv_objdetect480 -lopencv_features2d480 -lopencv_video480 -lopencv_dnn480 -lopencv_xfeatures2d480 -lopencv_plot480 -lopencv_tracking480 -lopencv_img_hash480 -lopencv_calib3d480 -lopencv_bgsegm480 -lopencv_photo480 -lopencv_aruco480 -lopencv_wechat_qrcode480 -lopencv_ximgproc480
*/
import "C"
