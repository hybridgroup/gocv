//go:build !customenv && !opencvstatic

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11 -DNDEBUG
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core4120 -lopencv_face4120 -lopencv_videoio4120 -lopencv_imgproc4120 -lopencv_highgui4120 -lopencv_imgcodecs4120 -lopencv_objdetect4120 -lopencv_features2d4120 -lopencv_video4120 -lopencv_dnn4120 -lopencv_xfeatures2d4120 -lopencv_plot4120 -lopencv_tracking4120 -lopencv_img_hash4120 -lopencv_calib3d4120 -lopencv_bgsegm4120 -lopencv_photo4120 -lopencv_aruco4120 -lopencv_wechat_qrcode4120 -lopencv_ximgproc4120
*/
import "C"
