// +build !customenv,static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo452 -lopencv_tracking452 -lopencv_superres452 -lopencv_stitching452 -lopencv_optflow452 -lopencv_gapi452 -lopencv_face452 -lopencv_dpm452 -lopencv_dnn_objdetect452 -lopencv_ccalib452 -lopencv_bioinspired452 -lopencv_bgsegm452 -lopencv_aruco452 -lopencv_xobjdetect452 -lopencv_ximgproc452 -lopencv_xfeatures2d452 -lopencv_videostab452 -lopencv_video452 -lopencv_structured_light452 -lopencv_shape452 -lopencv_rgbd452 -lopencv_rapid452 -lopencv_objdetect452 -lopencv_mcc452 -lopencv_highgui452 -lopencv_datasets452 -lopencv_calib3d452 -lopencv_videoio452 -lopencv_text452 -lopencv_line_descriptor452 -lopencv_imgcodecs452 -lopencv_img_hash452 -lopencv_hfs452 -lopencv_fuzzy452 -lopencv_features2d452 -lopencv_dnn_superres452 -lopencv_dnn452 -lopencv_xphoto452 -lopencv_surface_matching452 -lopencv_reg452 -lopencv_quality452 -lopencv_plot452 -lopencv_photo452 -lopencv_phase_unwrapping452 -lopencv_ml452 -lopencv_intensity_transform452 -lopencv_imgproc452 -lopencv_flann452 -lopencv_core452 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
