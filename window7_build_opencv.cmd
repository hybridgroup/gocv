echo window7_build_opencv.cmd

if not exist "C:\opencv" mkdir "C:\opencv"
if not exist "C:\opencv\build" mkdir "C:\opencv\build"



cd /D C:\opencv\build
set PATH=%PATH%;C:\Program Files (x86)\CMake\bin;C:\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\bin
cmake C:\opencv\opencv-4.0.0 -G "MinGW Makefiles" -BC:\opencv\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=C:\opencv\opencv_contrib-4.0.0\modules -DBUILD_SHARED_LIBS=ON -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -Wno-dev
mingw32-make -j%NUMBER_OF_PROCESSORS%
mingw32-make install
echo c:\opencv\opencv-4.0.0 /s /q
echo c:\opencv\opencv_contrib-4.0.0 /s /q
chdir /D %GOPATH%\src\gocv.io\x\gocv
