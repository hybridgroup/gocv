//go:build !customenv && opencvstatic && windows

package gocv

// Changes here should be mirrored in contrib/cgo_static_windows.go and cuda/cgo_static_windows.go.

/*
#cgo CXXFLAGS:   --std=c++11 -DNDEBUG
#cgo CPPFLAGS:   -IC:/opencv/build/install/include
#cgo LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo4120 -lopencv_tracking4120 -lopencv_superres4120 -lopencv_stitching4120 -lopencv_optflow4120 -lopencv_gapi4120 -lopencv_face4120 -lopencv_dpm4120 -lopencv_dnn_objdetect4120 -lopencv_ccalib4120 -lopencv_bioinspired4120 -lopencv_bgsegm4120 -lopencv_aruco4120 -lopencv_xobjdetect4120 -lopencv_ximgproc4120 -lopencv_xfeatures2d4120 -lopencv_videostab4120 -lopencv_video4120 -lopencv_structured_light4120 -lopencv_shape4120 -lopencv_rgbd4120 -lopencv_rapid4120 -lopencv_objdetect4120 -lopencv_mcc4120 -lopencv_highgui4120 -lopencv_datasets4120 -lopencv_calib3d4120 -lopencv_videoio4120 -lopencv_text4120 -lopencv_line_descriptor4120 -lopencv_imgcodecs4120 -lopencv_img_hash4120 -lopencv_hfs4120 -lopencv_fuzzy4120 -lopencv_features2d4120 -lopencv_dnn_superres4120 -lopencv_dnn4120 -lopencv_xphoto4120 -lopencv_wechat_qrcode4120 -lopencv_surface_matching4120 -lopencv_reg4120 -lopencv_quality4120 -lopencv_plot4120 -lopencv_photo4120 -lopencv_phase_unwrapping4120 -lopencv_ml4120 -lopencv_intensity_transform4120 -lopencv_imgproc4120 -lopencv_flann4120 -lopencv_core4120 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
