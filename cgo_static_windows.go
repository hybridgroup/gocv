//go:build !customenv && opencvstatic && windows

package gocv

// Changes here should be mirrored in contrib/cgo_static_windows.go and cuda/cgo_static_windows.go.

/*
#cgo CXXFLAGS:   --std=c++11 -DNDEBUG
#cgo CPPFLAGS:   -IC:/opencv/build/install/include
#cgo LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo4130 -lopencv_tracking4130 -lopencv_superres4130 -lopencv_stitching4130 -lopencv_optflow4130 -lopencv_gapi4130 -lopencv_face4130 -lopencv_dpm4130 -lopencv_dnn_objdetect4130 -lopencv_ccalib4130 -lopencv_bioinspired4130 -lopencv_bgsegm4130 -lopencv_aruco4130 -lopencv_xobjdetect4130 -lopencv_ximgproc4130 -lopencv_xfeatures2d4130 -lopencv_videostab4130 -lopencv_video4130 -lopencv_structured_light4130 -lopencv_shape4130 -lopencv_rgbd4130 -lopencv_rapid4130 -lopencv_objdetect4130 -lopencv_mcc4130 -lopencv_highgui4130 -lopencv_datasets4130 -lopencv_calib3d4130 -lopencv_videoio4130 -lopencv_text4130 -lopencv_line_descriptor4130 -lopencv_imgcodecs4130 -lopencv_img_hash4130 -lopencv_hfs4130 -lopencv_fuzzy4130 -lopencv_features2d4130 -lopencv_dnn_superres4130 -lopencv_dnn4130 -lopencv_xphoto4130 -lopencv_wechat_qrcode4130 -lopencv_surface_matching4130 -lopencv_reg4130 -lopencv_quality4130 -lopencv_plot4130 -lopencv_photo4130 -lopencv_phase_unwrapping4130 -lopencv_ml4130 -lopencv_intensity_transform4130 -lopencv_imgproc4130 -lopencv_flann4130 -lopencv_core4130 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
