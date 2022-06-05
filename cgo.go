//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core460 -lopencv_face460 -lopencv_videoio460 -lopencv_imgproc460 -lopencv_highgui460 -lopencv_imgcodecs460 -lopencv_objdetect460 -lopencv_features2d460 -lopencv_video460 -lopencv_dnn460 -lopencv_xfeatures2d460 -lopencv_plot460 -lopencv_tracking460 -lopencv_img_hash460 -lopencv_calib3d460 -lopencv_bgsegm460 -lopencv_photo460 -lopencv_aruco460 -lopencv_wechat_qrcode460 -lopencv_ximgproc460
*/
import "C"
