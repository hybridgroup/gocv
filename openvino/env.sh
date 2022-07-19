export CGO_CXXFLAGS="--std=c++11"
export CGO_CPPFLAGS="-I${INTEL_OPENVINO_DIR}/extras/opencv/include -I${INTEL_OPENVINO_DIR}/runtime/include -I${INTEL_OPENVINO_DIR}/runtime/include/ie"
export CGO_LDFLAGS="-L${INTEL_OPENVINO_DIR}/extras/opencv/lib -L${INTEL_OPENVINO_DIR}/runtime/lib/intel64 -lpthread -ldl -lopenvino -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_calib3d -lopencv_photo"
export PKG_CONFIG_PATH=/usr/lib64/pkgconfig
