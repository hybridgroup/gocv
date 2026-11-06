//go:build !customenv && !opencvstatic

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv5
#cgo CXXFLAGS:   --std=c++17 -DNDEBUG
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_stitching510 -lopencv_bgsegm510 -lopencv_bioinspired510 -lopencv_ccalib510 -lopencv_dnn_objdetect510 -lopencv_dnn_superres510 -lopencv_dpm510 -lopencv_face510 -lopencv_fuzzy510 -lopencv_gapi510 -lopencv_hfs510 -lopencv_img_hash510 -lopencv_intensity_transform510 -lopencv_line_descriptor510 -lopencv_quality510 -lopencv_rapid510 -lopencv_reg510 -lopencv_rgbd510 -lopencv_ptcloud510 -lopencv_signal510 -lopencv_structured_light510 -lopencv_calib510 -lopencv_phase_unwrapping510 -lopencv_superres510 -lopencv_optflow510 -lopencv_surface_matching510 -lopencv_videostab510 -lopencv_wechat_qrcode510 -lopencv_objdetect510 -lopencv_xfeatures2d510 -lopencv_shape510 -lopencv_ximgproc510 -lopencv_xobjdetect510 -lopencv_xphoto510 -lopencv_photo510 -lopencv_xstereo510 -lopencv_tracking510 -lopencv_highgui510 -lopencv_datasets510 -lopencv_videoio510 -lopencv_video510 -lopencv_text510 -lopencv_imgcodecs510 -lopencv_features510 -lopencv_dnn510 -lopencv_stereo510 -lopencv_plot510 -lopencv_ml510 -lopencv_imgproc510 -lopencv_geometry510 -lopencv_flann510 -lopencv_core510
*/
import "C"
