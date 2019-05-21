// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core410 -lopencv_face410 -lopencv_videoio410 -lopencv_imgproc410 -lopencv_highgui410 -lopencv_imgcodecs410 -lopencv_objdetect410 -lopencv_features2d410 -lopencv_video410 -lopencv_dnn410 -lopencv_xfeatures2d410 -lopencv_plot410 -lopencv_tracking410 -lopencv_img_hash410 -lopencv_calib3d410
*/
import "C"
