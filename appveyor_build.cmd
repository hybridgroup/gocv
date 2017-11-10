if not exist "%APPVEYOR_BUILD_FOLDER%\opencv\%BUILD_DIR%" mkdir "%APPVEYOR_BUILD_FOLDER%\opencv\%BUILD_DIR%"

cd opencv

cmake -G "%BUILD_ENV%" -H"." -B"%BUILD_DIR%" -DOPENCV_EXTRA_MODULES_PATH=../opencv_contrib/modules -DBUILD_SHARED_LIBS=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -Wno-dev

cd %BUILD_DIR%

cmake --build . --config Release

cd ..\..
cd

xcopy "%APPVEYOR_BUILD_FOLDER%\opencv\%BUILD_DIR%\bin\Release\*.dll" .\cv2 /I
xcopy "%APPVEYOR_BUILD_FOLDER%\LICENSE*.txt" .\cv2 /I

dir
