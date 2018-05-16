export CGO_CPPFLAGS="-I${INTEL_CVSDK_DIR}/opencv/include" 
export CGO_LDFLAGS="-L${INTEL_CVSDK_DIR}/opencv/lib -lopencv_core -lopencv_pvl -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d"
