
if not exist "C:\opencv" mkdir "C:\opencv"
if not exist "C:\opencv\build" mkdir "C:\opencv\build"
if not exist "C:\opencv\testdata" mkdir "C:\opencv\testdata"

appveyor DownloadFile https://github.com/opencv/opencv/archive/4.5.4.zip -FileName c:\opencv\opencv-4.5.4.zip
7z x c:\opencv\opencv-4.5.4.zip -oc:\opencv -y
del c:\opencv\opencv-4.5.4.zip /q
appveyor DownloadFile https://github.com/opencv/opencv_contrib/archive/4.5.4.zip -FileName c:\opencv\opencv_contrib-4.5.4.zip
7z x c:\opencv\opencv_contrib-4.5.4.zip -oc:\opencv -y
del c:\opencv\opencv_contrib-4.5.4.zip /q
cd C:\opencv\build
set PATH=C:\Perl\site\bin;C:\Perl\bin;C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPowerShell\v1.0\;C:\Program Files\7-Zip;C:\Program Files\Microsoft\Web Platform Installer\;C:\Tools\PsTools;C:\Program Files (x86)\CMake\bin;C:\go\bin;C:\Tools\NuGet;C:\Program Files\LLVM\bin;C:\Tools\curl\bin;C:\ProgramData\chocolatey\bin;C:\Program Files (x86)\Yarn\bin;C:\Users\appveyor\AppData\Local\Yarn\bin;C:\Program Files\AppVeyor\BuildAgent\
set PATH=%PATH%;C:\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev0\mingw64\bin
dir C:\opencv
cmake C:\opencv\opencv-4.5.4 -G "MinGW Makefiles" -BC:\opencv\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=C:\opencv\opencv_contrib-4.5.4\modules -DBUILD_SHARED_LIBS=ON -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=OFF -DCPU_DISPATCH= -DBUILD_opencv_gapi=OFF -DOPENCV_GENERATE_PKGCONFIG=ON -DOPENCV_ENABLE_NONFREE=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -DWITH_TBB=ON -Wno-dev
mingw32-make -j%NUMBER_OF_PROCESSORS%
mingw32-make install
appveyor DownloadFile https://raw.githubusercontent.com/opencv/opencv_extra/master/testdata/dnn/bvlc_googlenet.prototxt -FileName C:\opencv\testdata\bvlc_googlenet.prototxt
appveyor DownloadFile http://dl.caffe.berkeleyvision.org/bvlc_googlenet.caffemodel -FileName C:\opencv\testdata\bvlc_googlenet.caffemodel
appveyor DownloadFile https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip -FileName C:\opencv\testdata\inception5h.zip
appveyor DownloadFile https://github.com/onnx/models/raw/master/vision/classification/inception_and_googlenet/googlenet/model/googlenet-9.onnx -FileName C:\opencv\testdata\googlenet-9.onnx
7z x C:\opencv\testdata\inception5h.zip -oC:\opencv\testdata tensorflow_inception_graph.pb -y
rmdir c:\opencv\opencv-4.5.4 /s /q
rmdir c:\opencv\opencv_contrib-4.5.4 /s /q
