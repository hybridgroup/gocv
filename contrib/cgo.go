//go:build !customenv && !opencvstatic

package contrib

// Changes here should be mirrored in gocv/cgo.go and cuda/cgo.go

/*
#cgo !windows pkg-config: opencv5
#cgo CXXFLAGS:   --std=c++17 -DNDEBUG
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_stitching500 -lopencv_bgsegm500 -lopencv_bioinspired500 -lopencv_ccalib500 -lopencv_dnn_objdetect500 -lopencv_dnn_superres500 -lopencv_dpm500 -lopencv_face500 -lopencv_fuzzy500 -lopencv_gapi500 -lopencv_hfs500 -lopencv_img_hash500 -lopencv_intensity_transform500 -lopencv_line_descriptor500 -lopencv_quality500 -lopencv_rapid500 -lopencv_reg500 -lopencv_rgbd500 -lopencv_ptcloud500 -lopencv_signal500 -lopencv_structured_light500 -lopencv_calib500 -lopencv_phase_unwrapping500 -lopencv_superres500 -lopencv_optflow500 -lopencv_surface_matching500 -lopencv_videostab500 -lopencv_wechat_qrcode500 -lopencv_objdetect500 -lopencv_xfeatures2d500 -lopencv_shape500 -lopencv_ximgproc500 -lopencv_xobjdetect500 -lopencv_xphoto500 -lopencv_photo500 -lopencv_xstereo500 -lopencv_tracking500 -lopencv_highgui500 -lopencv_datasets500 -lopencv_videoio500 -lopencv_video500 -lopencv_text500 -lopencv_imgcodecs500 -lopencv_features500 -lopencv_dnn500 -lopencv_stereo500 -lopencv_plot500 -lopencv_ml500 -lopencv_imgproc500 -lopencv_geometry500 -lopencv_flann500 -lopencv_core500
*/
import "C"
