//go:build !customenv && opencvstatic && windows

package gocv

// Changes here should be mirrored in contrib/cgo_static_windows.go and cuda/cgo_static_windows.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo CPPFLAGS:   -IC:/opencv/build/install/include
#cgo LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo4110 -lopencv_tracking4110 -lopencv_superres4110 -lopencv_stitching4110 -lopencv_optflow4110 -lopencv_gapi4110 -lopencv_face4110 -lopencv_dpm4110 -lopencv_dnn_objdetect4110 -lopencv_ccalib4110 -lopencv_bioinspired4110 -lopencv_bgsegm4110 -lopencv_aruco4110 -lopencv_xobjdetect4110 -lopencv_ximgproc4110 -lopencv_xfeatures2d4110 -lopencv_videostab4110 -lopencv_video4110 -lopencv_structured_light4110 -lopencv_shape4110 -lopencv_rgbd4110 -lopencv_rapid4110 -lopencv_objdetect4110 -lopencv_mcc4110 -lopencv_highgui4110 -lopencv_datasets4110 -lopencv_calib3d4110 -lopencv_videoio4110 -lopencv_text4110 -lopencv_line_descriptor4110 -lopencv_imgcodecs4110 -lopencv_img_hash4110 -lopencv_hfs4110 -lopencv_fuzzy4110 -lopencv_features2d4110 -lopencv_dnn_superres4110 -lopencv_dnn4110 -lopencv_xphoto4110 -lopencv_wechat_qrcode4110 -lopencv_surface_matching4110 -lopencv_reg4110 -lopencv_quality4110 -lopencv_plot4110 -lopencv_photo4110 -lopencv_phase_unwrapping4110 -lopencv_ml4110 -lopencv_intensity_transform4110 -lopencv_imgproc4110 -lopencv_flann4110 -lopencv_core4110 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
