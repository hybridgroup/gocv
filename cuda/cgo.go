//go:build !customenv && !opencvstatic

package cuda

// Changes here should be mirrored in gocv/cgo.go and contrib/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core4110 -lopencv_face4110 -lopencv_videoio4110 -lopencv_imgproc4110 -lopencv_highgui4110 -lopencv_imgcodecs4110 -lopencv_objdetect4110 -lopencv_features2d4110 -lopencv_video4110 -lopencv_dnn4110 -lopencv_xfeatures2d4110 -lopencv_plot4110 -lopencv_tracking4110 -lopencv_img_hash4110 -lopencv_calib3d4110 -lopencv_bgsegm4110 -lopencv_aruco4110 -lopencv_wechat_qrcode4110 -lopencv_ximgproc4110
*/
import "C"
