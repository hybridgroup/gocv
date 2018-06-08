// +build !customenv,!openvino

package contrib

// Changes here should be mirrored in gocv/cgo.go.

/*
#cgo !windows pkg-config: opencv
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core342 -lopencv_face342 -lopencv_videoio342 -lopencv_imgproc342 -lopencv_highgui342 -lopencv_imgcodecs342 -lopencv_objdetect342 -lopencv_features2d342 -lopencv_video342 -lopencv_dnn342 -lopencv_xfeatures2d342 -lopencv_plot342 -lopencv_tracking342 -lopencv_img_hash342
*/
import "C"
