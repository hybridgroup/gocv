// +build !customenv

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core440 -lopencv_face440 -lopencv_videoio440 -lopencv_imgproc440 -lopencv_highgui440 -lopencv_imgcodecs440 -lopencv_objdetect440 -lopencv_features2d440 -lopencv_video440 -lopencv_dnn440 -lopencv_xfeatures2d440 -lopencv_plot440 -lopencv_tracking440 -lopencv_img_hash440 -lopencv_calib3d440 -lopencv_bgsegm440
*/
import "C"
