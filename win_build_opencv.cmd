@echo off

REM Save the starting directory
set "START_DIR=%CD%"

REM Set install directory, default to C:\opencv if not provided
set "INSTALL_DIR=%~1"
if "%INSTALL_DIR%"=="" set "INSTALL_DIR=C:\opencv"

REM Always quote paths to handle spaces
if not exist "%INSTALL_DIR%\build" mkdir "%INSTALL_DIR%\build"

REM Change to build directory
cd /D "%INSTALL_DIR%\build"

REM Set enable_shared default
set "enable_shared=ON"
if /I "%2"=="static" (
  echo Build static opencv
  set "enable_shared=OFF"
)

REM Use variables for version
set "OPENCV_VERSION=4.13.0"

REM Run CMake with quoted paths
cmake "%INSTALL_DIR%\opencv-%OPENCV_VERSION%" -G "MinGW Makefiles" -B"%INSTALL_DIR%\build" -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH="%INSTALL_DIR%\opencv_contrib-%OPENCV_VERSION%\modules" -DBUILD_SHARED_LIBS=%enable_shared% -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=ON -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=ON -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -Wno-dev

if errorlevel 1 (
    echo CMake configuration failed!
    cd /D "%START_DIR%"
    exit /b 1
)

cmake --build . --target install
if errorlevel 1 (
    echo Build failed!
    cd /D "%START_DIR%"
    exit /b 1
)

REM Return to the original directory
cd /D "%START_DIR%"