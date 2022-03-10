//go:build !customenv && static
// +build !customenv,static

package contrib

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo455 -lopencv_tracking455 -lopencv_superres455 -lopencv_stitching455 -lopencv_optflow455 -lopencv_gapi455 -lopencv_face455 -lopencv_dpm455 -lopencv_dnn_objdetect455 -lopencv_ccalib455 -lopencv_bioinspired455 -lopencv_bgsegm455 -lopencv_aruco455 -lopencv_xobjdetect455 -lopencv_ximgproc455 -lopencv_xfeatures2d455 -lopencv_videostab455 -lopencv_video455 -lopencv_structured_light455 -lopencv_shape455 -lopencv_rgbd455 -lopencv_rapid455 -lopencv_objdetect455 -lopencv_mcc455 -lopencv_highgui455 -lopencv_datasets455 -lopencv_calib3d455 -lopencv_videoio455 -lopencv_text455 -lopencv_line_descriptor455 -lopencv_imgcodecs455 -lopencv_img_hash455 -lopencv_hfs455 -lopencv_fuzzy455 -lopencv_features2d455 -lopencv_dnn_superres455 -lopencv_dnn455 -lopencv_xphoto455 -lopencv_wechat_qrcode455 -lopencv_surface_matching455 -lopencv_reg455 -lopencv_quality455 -lopencv_plot455 -lopencv_photo455 -lopencv_phase_unwrapping455 -lopencv_ml455 -lopencv_intensity_transform455 -lopencv_imgproc455 -lopencv_flann455 -lopencv_core455 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
