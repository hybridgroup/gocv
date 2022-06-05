//go:build !customenv && static
// +build !customenv,static

package contrib

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo460 -lopencv_tracking460 -lopencv_superres460 -lopencv_stitching460 -lopencv_optflow460 -lopencv_gapi460 -lopencv_face460 -lopencv_dpm460 -lopencv_dnn_objdetect460 -lopencv_ccalib460 -lopencv_bioinspired460 -lopencv_bgsegm460 -lopencv_aruco460 -lopencv_xobjdetect460 -lopencv_ximgproc460 -lopencv_xfeatures2d460 -lopencv_videostab460 -lopencv_video460 -lopencv_structured_light460 -lopencv_shape460 -lopencv_rgbd460 -lopencv_rapid460 -lopencv_objdetect460 -lopencv_mcc460 -lopencv_highgui460 -lopencv_datasets460 -lopencv_calib3d460 -lopencv_videoio460 -lopencv_text460 -lopencv_line_descriptor460 -lopencv_imgcodecs460 -lopencv_img_hash460 -lopencv_hfs460 -lopencv_fuzzy460 -lopencv_features2d460 -lopencv_dnn_superres460 -lopencv_dnn460 -lopencv_xphoto460 -lopencv_wechat_qrcode460 -lopencv_surface_matching460 -lopencv_reg460 -lopencv_quality460 -lopencv_plot460 -lopencv_photo460 -lopencv_phase_unwrapping460 -lopencv_ml460 -lopencv_intensity_transform460 -lopencv_imgproc460 -lopencv_flann460 -lopencv_core460 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
