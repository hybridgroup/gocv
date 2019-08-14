// +build !customenv,!openvino

package cuda

// Changes here should be mirrored in gocv/cgo.go and contrib/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core411 -lopencv_face411 -lopencv_videoio411 -lopencv_imgproc411 -lopencv_highgui411 -lopencv_imgcodecs411 -lopencv_objdetect411 -lopencv_features2d411 -lopencv_video411 -lopencv_dnn411 -lopencv_xfeatures2d411 -lopencv_plot411 -lopencv_tracking411 -lopencv_img_hash411 -lopencv_calib3d411
*/
import "C"
