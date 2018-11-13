// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core400 -lopencv_face400 -lopencv_videoio400 -lopencv_imgproc400 -lopencv_highgui400 -lopencv_imgcodecs400 -lopencv_objdetect400 -lopencv_features2d400 -lopencv_video400 -lopencv_dnn400 -lopencv_xfeatures2d400 -lopencv_plot400 -lopencv_tracking400 -lopencv_img_hash400 -lopencv_calib3d400
*/
import "C"
