// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core453 -lopencv_face453 -lopencv_videoio453 -lopencv_imgproc453 -lopencv_highgui453 -lopencv_imgcodecs453 -lopencv_objdetect453 -lopencv_features2d453 -lopencv_video453 -lopencv_dnn453 -lopencv_xfeatures2d453 -lopencv_plot453 -lopencv_tracking453 -lopencv_img_hash453 -lopencv_calib3d453 -lopencv_bgsegm453 -lopencv_photo453
*/
import "C"
