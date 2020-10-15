// +build !customenv

package cuda

// Changes here should be mirrored in gocv/cgo.go and contrib/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core450 -lopencv_face450 -lopencv_videoio450 -lopencv_imgproc450 -lopencv_highgui450 -lopencv_imgcodecs450 -lopencv_objdetect450 -lopencv_features2d450 -lopencv_video450 -lopencv_dnn450 -lopencv_xfeatures2d450 -lopencv_plot450 -lopencv_tracking450 -lopencv_img_hash450 -lopencv_calib3d450 -lopencv_bgsegm450
*/
import "C"
