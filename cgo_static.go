//go:build !customenv && static
// +build !customenv,static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo481 -lopencv_tracking481 -lopencv_superres481 -lopencv_stitching481 -lopencv_optflow481 -lopencv_gapi481 -lopencv_face481 -lopencv_dpm481 -lopencv_dnn_objdetect481 -lopencv_ccalib481 -lopencv_bioinspired481 -lopencv_bgsegm481 -lopencv_aruco481 -lopencv_xobjdetect481 -lopencv_ximgproc481 -lopencv_xfeatures2d481 -lopencv_videostab481 -lopencv_video481 -lopencv_structured_light481 -lopencv_shape481 -lopencv_rgbd481 -lopencv_rapid481 -lopencv_objdetect481 -lopencv_mcc481 -lopencv_highgui481 -lopencv_datasets481 -lopencv_calib3d481 -lopencv_videoio481 -lopencv_text481 -lopencv_line_descriptor481 -lopencv_imgcodecs481 -lopencv_img_hash481 -lopencv_hfs481 -lopencv_fuzzy481 -lopencv_features2d481 -lopencv_dnn_superres481 -lopencv_dnn481 -lopencv_xphoto481 -lopencv_wechat_qrcode481 -lopencv_surface_matching481 -lopencv_reg481 -lopencv_quality481 -lopencv_plot481 -lopencv_photo481 -lopencv_phase_unwrapping481 -lopencv_ml481 -lopencv_intensity_transform481 -lopencv_imgproc481 -lopencv_flann481 -lopencv_core481 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
