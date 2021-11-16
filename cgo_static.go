//go:build !customenv && static
// +build !customenv,static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo454 -lopencv_tracking454 -lopencv_superres454 -lopencv_stitching454 -lopencv_optflow454 -lopencv_gapi454 -lopencv_face454 -lopencv_dpm454 -lopencv_dnn_objdetect454 -lopencv_ccalib454 -lopencv_bioinspired454 -lopencv_bgsegm454 -lopencv_aruco454 -lopencv_xobjdetect454 -lopencv_ximgproc454 -lopencv_xfeatures2d454 -lopencv_videostab454 -lopencv_video454 -lopencv_structured_light454 -lopencv_shape454 -lopencv_rgbd454 -lopencv_rapid454 -lopencv_objdetect454 -lopencv_mcc454 -lopencv_highgui454 -lopencv_datasets454 -lopencv_calib3d454 -lopencv_videoio454 -lopencv_text454 -lopencv_line_descriptor454 -lopencv_imgcodecs454 -lopencv_img_hash454 -lopencv_hfs454 -lopencv_fuzzy454 -lopencv_features2d454 -lopencv_dnn_superres454 -lopencv_dnn454 -lopencv_xphoto454 -lopencv_surface_matching454 -lopencv_reg454 -lopencv_quality454 -lopencv_plot454 -lopencv_photo454 -lopencv_phase_unwrapping454 -lopencv_ml454 -lopencv_intensity_transform454 -lopencv_imgproc454 -lopencv_flann454 -lopencv_core454 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
