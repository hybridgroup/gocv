//go:build !customenv && static
// +build !customenv,static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo480 -lopencv_tracking480 -lopencv_superres480 -lopencv_stitching480 -lopencv_optflow480 -lopencv_gapi480 -lopencv_face480 -lopencv_dpm480 -lopencv_dnn_objdetect480 -lopencv_ccalib480 -lopencv_bioinspired480 -lopencv_bgsegm480 -lopencv_aruco480 -lopencv_xobjdetect480 -lopencv_ximgproc480 -lopencv_xfeatures2d480 -lopencv_videostab480 -lopencv_video480 -lopencv_structured_light480 -lopencv_shape480 -lopencv_rgbd480 -lopencv_rapid480 -lopencv_objdetect480 -lopencv_mcc480 -lopencv_highgui480 -lopencv_datasets480 -lopencv_calib3d480 -lopencv_videoio480 -lopencv_text480 -lopencv_line_descriptor480 -lopencv_imgcodecs480 -lopencv_img_hash480 -lopencv_hfs480 -lopencv_fuzzy480 -lopencv_features2d480 -lopencv_dnn_superres480 -lopencv_dnn480 -lopencv_xphoto480 -lopencv_wechat_qrcode480 -lopencv_surface_matching480 -lopencv_reg480 -lopencv_quality480 -lopencv_plot480 -lopencv_photo480 -lopencv_phase_unwrapping480 -lopencv_ml480 -lopencv_intensity_transform480 -lopencv_imgproc480 -lopencv_flann480 -lopencv_core480 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
