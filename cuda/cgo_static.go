//go:build !customenv && static
// +build !customenv,static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo470 -lopencv_tracking470 -lopencv_superres470 -lopencv_stitching470 -lopencv_optflow470 -lopencv_gapi470 -lopencv_face470 -lopencv_dpm470 -lopencv_dnn_objdetect470 -lopencv_ccalib470 -lopencv_bioinspired470 -lopencv_bgsegm470 -lopencv_aruco470 -lopencv_xobjdetect470 -lopencv_ximgproc470 -lopencv_xfeatures2d470 -lopencv_videostab470 -lopencv_video470 -lopencv_structured_light470 -lopencv_shape470 -lopencv_rgbd470 -lopencv_rapid470 -lopencv_objdetect470 -lopencv_mcc470 -lopencv_highgui470 -lopencv_datasets470 -lopencv_calib3d470 -lopencv_videoio470 -lopencv_text470 -lopencv_line_descriptor470 -lopencv_imgcodecs470 -lopencv_img_hash470 -lopencv_hfs470 -lopencv_fuzzy470 -lopencv_features2d470 -lopencv_dnn_superres470 -lopencv_dnn470 -lopencv_xphoto470 -lopencv_wechat_qrcode470 -lopencv_surface_matching470 -lopencv_reg470 -lopencv_quality470 -lopencv_plot470 -lopencv_photo470 -lopencv_phase_unwrapping470 -lopencv_ml470 -lopencv_intensity_transform470 -lopencv_imgproc470 -lopencv_flann470 -lopencv_core470 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
