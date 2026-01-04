@echo off

REM Save the starting directory
set "START_DIR=%CD%"

REM Set install directory, default to C:\opencv if not provided
set "INSTALL_DIR=%~1"
if "%INSTALL_DIR%"=="" set "INSTALL_DIR=C:\opencv"

if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

echo Downloading OpenCV sources
echo.
echo For monitoring the download progress please check the %INSTALL_DIR% directory.
echo.

REM This is why there is no progress bar:
REM https://github.com/PowerShell/PowerShell/issues/2138

REM Downloading: opencv-4.13.0.zip [91MB]
powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri 'https://github.com/opencv/opencv/archive/4.13.0.zip' -OutFile \"%INSTALL_DIR%\opencv-4.13.0.zip\""
echo Extracting...
powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path \"%INSTALL_DIR%\opencv-4.13.0.zip\" -DestinationPath \"%INSTALL_DIR%\""
del "%INSTALL_DIR%\opencv-4.13.0.zip" /q
echo.

echo Downloading: opencv_contrib-4.13.0.zip [58MB]
powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri 'https://github.com/opencv/opencv_contrib/archive/4.13.0.zip' -OutFile \"%INSTALL_DIR%\opencv_contrib-4.13.0.zip\""
echo Extracting...
powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path \"%INSTALL_DIR%\opencv_contrib-4.13.0.zip\" -DestinationPath \"%INSTALL_DIR%\""
del "%INSTALL_DIR%\opencv_contrib-4.13.0.zip" /q
echo.

echo Done with downloading and extracting sources.
echo.

@echo on

REM Return to the original directory
chdir /D "%START_DIR%"