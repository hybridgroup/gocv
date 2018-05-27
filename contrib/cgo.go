// +build !customenv

package contrib

// Changes here should be mirrored in gocv/cgo.go.

/*
#cgo !windows pkg-config: opencv
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core341 -lopencv_face341 -lopencv_videoio341 -lopencv_imgproc341 -lopencv_highgui341 -lopencv_imgcodecs341 -lopencv_objdetect341 -lopencv_features2d341 -lopencv_video341 -lopencv_dnn341 -lopencv_xfeatures2d341 -lopencv_plot341 -lopencv_tracking341 -lopencv_img_hash341
*/
import "C"
