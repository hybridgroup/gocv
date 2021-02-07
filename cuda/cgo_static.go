// +build !customenv,static,!windows

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
*/
import "C"
