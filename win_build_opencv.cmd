@echo off

REM Set install directory, default to C:\opencv if not provided
set "INSTALL_DIR=%~1"
if "%INSTALL_DIR%"=="" set "INSTALL_DIR=C:\opencv"

if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"
if not exist "%INSTALL_DIR%\build" mkdir "%INSTALL_DIR%\build"

echo Downloading OpenCV sources
echo.
echo For monitoring the download progress please check the %INSTALL_DIR% directory.
echo.

REM This is why there is no progress bar:
REM https://github.com/PowerShell/PowerShell/issues/2138

echo Downloading: opencv-4.13.0.zip [91MB]
powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri https://github.com/opencv/opencv/archive/4.13.0.zip -OutFile %INSTALL_DIR%\opencv-4.13.0.zip"
echo Extracting...
powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path %INSTALL_DIR%\opencv-4.13.0.zip -DestinationPath %INSTALL_DIR%"
del %INSTALL_DIR%\opencv-4.13.0.zip /q
echo.

echo Downloading: opencv_contrib-4.13.0.zip [58MB]
powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri https://github.com/opencv/opencv_contrib/archive/4.13.0.zip -OutFile %INSTALL_DIR%\opencv_contrib-4.13.0.zip"
echo Extracting...
powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path %INSTALL_DIR%\opencv_contrib-4.13.0.zip -DestinationPath %INSTALL_DIR%"
del %INSTALL_DIR%\opencv_contrib-4.13.0.zip /q
echo.

echo Done with downloading and extracting sources.
echo.

@echo on

cd /D %INSTALL_DIR%\build
if [%2]==[static] (
  echo Build static opencv
  set enable_shared=OFF
) else (
  set enable_shared=ON
)
cmake %INSTALL_DIR%\opencv-4.13.0 -G "MinGW Makefiles" -B%INSTALL_DIR%\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=%INSTALL_DIR%\opencv_contrib-4.13.0\modules -DBUILD_SHARED_LIBS=%enable_shared% -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=ON -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=ON -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -Wno-dev
cmake --build . --target install
rmdir %INSTALL_DIR%\opencv-4.13.0 /s /q
rmdir %INSTALL_DIR%\opencv_contrib-4.13.0 /s /q
chdir /D %GOPATH%\src\gocv.io\x\gocv
