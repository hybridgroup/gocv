if not exist "C:\opencv" mkdir "C:\opencv"
if not exist "C:\opencv\build" mkdir "C:\opencv\build"

bitsadmin /transfer opencv https://github.com/opencv/opencv/archive/3.4.2.zip c:\opencv\opencv-3.4.2.zip
7z x c:\opencv\opencv-3.4.2.zip -oc:\opencv -y
del c:\opencv\opencv-3.4.2.zip /q

bitsadmin /transfer opencv-contrib https://github.com/opencv/opencv_contrib/archive/3.4.2.zip c:\opencv\opencv_contrib-3.4.2.zip
7z x c:\opencv\opencv_contrib-3.4.2.zip -oc:\opencv -y
del c:\opencv\opencv_contrib-3.4.2.zip /q

cd C:\opencv\build
set PATH=%PATH%;C:\Program Files (x86)\CMake\bin;C:\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\bin
cmake C:\opencv\opencv-3.4.2 -G "MinGW Makefiles" -BC:\opencv\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=C:\opencv\opencv_contrib-3.4.2\modules -DBUILD_SHARED_LIBS=ON -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DCPU_DISPATCH= -Wno-dev
mingw32-make
mingw32-make install
rmdir c:\opencv\opencv-3.4.2 /s /q
rmdir c:\opencv\opencv_contrib-3.4.2 /s /q
