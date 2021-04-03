// +build !customenv

package cuda

// Changes here should be mirrored in gocv/cgo.go and contrib/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core452 -lopencv_face452 -lopencv_videoio452 -lopencv_imgproc452 -lopencv_highgui452 -lopencv_imgcodecs452 -lopencv_objdetect452 -lopencv_features2d452 -lopencv_video452 -lopencv_dnn452 -lopencv_xfeatures2d452 -lopencv_plot452 -lopencv_tracking452 -lopencv_img_hash452 -lopencv_calib3d452 -lopencv_bgsegm452
*/
import "C"
