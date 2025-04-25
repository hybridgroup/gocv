
param (
    [switch]$static = $false
 )
if ($static) {
  $BUILD_SHARED = "OFF"
} else {
  $BUILD_SHARED = "ON"
}

if (-not(Test-Path -path "c:\opencv\build")) {
  mkdir -p C:\opencv\build
}

echo "Downloading OpenCV sources" "" "For monitoring the download progress please check the C:\opencv directory."

cd C:\opencv

[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
Set-Variable ProgressPreference SilentlyContinue
if (-not(Test-Path -path "c:\opencv\opencv-4.11.0.zip")) {
  echo "Downloading: opencv-4.11.0.zip [91MB]"
  Invoke-WebRequest -Uri https://github.com/opencv/opencv/archive/4.11.0.zip -OutFile c:\opencv\opencv-4.11.0.zip
}

if (-not(Test-Path -path "c:\opencv\opencv-4.11.0")) {
  echo "Extracting..."
  Expand-Archive -Path c:\opencv\opencv-4.11.0.zip -DestinationPath c:\opencv
}

if (-not(Test-Path -path "c:\opencv\opencv_contrib-4.11.0.zip")) {
  echo "Downloading: opencv_contrib-4.11.0.zip [58MB]"
  Invoke-WebRequest -Uri https://github.com/opencv/opencv_contrib/archive/4.11.0.zip -OutFile c:\opencv\opencv_contrib-4.11.0.zip
}

if (-not(Test-Path -path "c:\opencv\opencv_contrib-4.11.0")) {
  echo "Extracting..."
  Expand-Archive -Path c:\opencv\opencv_contrib-4.11.0.zip -DestinationPath c:\opencv
}


$env:PATH = 'C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Program Files\Go\bin;C:\ProgramData\chocolatey\bin;C:\ProgramData\mingw64\mingw64\bin;C:\Program Files\CMake\bin;'


cd C:\opencv\build

cmake C:\opencv\opencv-4.11.0 -G "MinGW Makefiles" -BC:\opencv\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=C:\opencv\opencv_contrib-4.11.0\modules "-DBUILD_SHARED_LIBS=${BUILD_SHARED}" -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=ON -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=ON -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -Wno-dev
mingw32-make -j $env:NUMBER_OF_PROCESSORS
mingw32-make install

cd $PSScriptRoot
go run ./cmd/version/main.go
