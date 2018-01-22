@echo off
IF EXIST C:\opencv\build\install\include\ (
    ECHO Configuring GoCV env for OpenCV.
    set CGO_CPPFLAGS=-IC:\opencv\build\install\include
    set CGO_LDFLAGS=-LC:\opencv\build\install\x64\mingw\lib -lopencv_core340 -lopencv_face340 -lopencv_videoio340 -lopencv_imgproc340 -lopencv_highgui340 -lopencv_imgcodecs340 -lopencv_objdetect340 -lopencv_features2d340 -lopencv_video340 -lopencv_dnn340 -lopencv_xfeatures2d340
) ELSE (
    ECHO ERROR: Unable to locate OpenCV for GoCV configuration.
)
