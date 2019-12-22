// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core420 -lopencv_face420 -lopencv_videoio420 -lopencv_imgproc420 -lopencv_highgui420 -lopencv_imgcodecs420 -lopencv_objdetect420 -lopencv_features2d420 -lopencv_video420 -lopencv_dnn420 -lopencv_xfeatures2d420 -lopencv_plot420 -lopencv_tracking420 -lopencv_img_hash420 -lopencv_calib3d420
*/
import "C"
