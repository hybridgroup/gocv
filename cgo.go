// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go.

/*
#cgo !windows pkg-config: opencv
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core343 -lopencv_face343 -lopencv_videoio343 -lopencv_imgproc343 -lopencv_highgui343 -lopencv_imgcodecs343 -lopencv_objdetect343 -lopencv_features2d343 -lopencv_video343 -lopencv_dnn343 -lopencv_xfeatures2d343 -lopencv_plot343 -lopencv_tracking343 -lopencv_img_hash343 -lopencv_calib3d343
*/
import "C"
