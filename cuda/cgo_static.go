//go:build !customenv && static
// +build !customenv,static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo490 -lopencv_tracking490 -lopencv_superres490 -lopencv_stitching490 -lopencv_optflow490 -lopencv_gapi490 -lopencv_face490 -lopencv_dpm490 -lopencv_dnn_objdetect490 -lopencv_ccalib490 -lopencv_bioinspired490 -lopencv_bgsegm490 -lopencv_aruco490 -lopencv_xobjdetect490 -lopencv_ximgproc490 -lopencv_xfeatures2d490 -lopencv_videostab490 -lopencv_video490 -lopencv_structured_light490 -lopencv_shape490 -lopencv_rgbd490 -lopencv_rapid490 -lopencv_objdetect490 -lopencv_mcc490 -lopencv_highgui490 -lopencv_datasets490 -lopencv_calib3d490 -lopencv_videoio490 -lopencv_text490 -lopencv_line_descriptor490 -lopencv_imgcodecs490 -lopencv_img_hash490 -lopencv_hfs490 -lopencv_fuzzy490 -lopencv_features2d490 -lopencv_dnn_superres490 -lopencv_dnn490 -lopencv_xphoto490 -lopencv_wechat_qrcode490 -lopencv_surface_matching490 -lopencv_reg490 -lopencv_quality490 -lopencv_plot490 -lopencv_photo490 -lopencv_phase_unwrapping490 -lopencv_ml490 -lopencv_intensity_transform490 -lopencv_imgproc490 -lopencv_flann490 -lopencv_core490 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
