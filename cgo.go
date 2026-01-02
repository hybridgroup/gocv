//go:build !customenv && !opencvstatic

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11 -DNDEBUG
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core4130 -lopencv_face4130 -lopencv_videoio4130 -lopencv_imgproc4130 -lopencv_highgui4130 -lopencv_imgcodecs4130 -lopencv_objdetect4130 -lopencv_features2d4130 -lopencv_video4130 -lopencv_dnn4130 -lopencv_xfeatures2d4130 -lopencv_plot4130 -lopencv_tracking4130 -lopencv_img_hash4130 -lopencv_calib3d4130 -lopencv_bgsegm4130 -lopencv_photo4130 -lopencv_aruco4130 -lopencv_wechat_qrcode4130 -lopencv_ximgproc4130 -lopencv_mcc4130
*/
import "C"
