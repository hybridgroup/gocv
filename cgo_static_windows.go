//go:build !customenv && static && windows

package gocv

// Changes here should be mirrored in contrib/cgo_static_windows.go and cuda/cgo_static_windows.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo CPPFLAGS:   -IC:/opencv/build/install/include
#cgo LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo4100 -lopencv_tracking4100 -lopencv_superres4100 -lopencv_stitching4100 -lopencv_optflow4100 -lopencv_gapi4100 -lopencv_face4100 -lopencv_dpm4100 -lopencv_dnn_objdetect4100 -lopencv_ccalib4100 -lopencv_bioinspired4100 -lopencv_bgsegm4100 -lopencv_aruco4100 -lopencv_xobjdetect4100 -lopencv_ximgproc4100 -lopencv_xfeatures2d4100 -lopencv_videostab4100 -lopencv_video4100 -lopencv_structured_light4100 -lopencv_shape4100 -lopencv_rgbd4100 -lopencv_rapid4100 -lopencv_objdetect4100 -lopencv_mcc4100 -lopencv_highgui4100 -lopencv_datasets4100 -lopencv_calib3d4100 -lopencv_videoio4100 -lopencv_text4100 -lopencv_line_descriptor4100 -lopencv_imgcodecs4100 -lopencv_img_hash4100 -lopencv_hfs4100 -lopencv_fuzzy4100 -lopencv_features2d4100 -lopencv_dnn_superres4100 -lopencv_dnn4100 -lopencv_xphoto4100 -lopencv_wechat_qrcode4100 -lopencv_surface_matching4100 -lopencv_reg4100 -lopencv_quality4100 -lopencv_plot4100 -lopencv_photo4100 -lopencv_phase_unwrapping4100 -lopencv_ml4100 -lopencv_intensity_transform4100 -lopencv_imgproc4100 -lopencv_flann4100 -lopencv_core4100 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
