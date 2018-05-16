# Using the Intel OpenVINO toolkit

The Intel OpenVINO toolkit is a set of tools and libraries for computer vision applications. It also includes a complete build of OpenCV installed with the Intel OpenVINO. It uses computer vision and imaging algorithms developed at Intel.

GoCV support for the Intel OpenVINO Photography Vision Library (PVL) can be found here in the "gocv.io/x/gocv/openvino/pvl" package. Check out the README.md in the `pvl` directory for more information.

## How to install the Intel OpenVINO toolkit

You will need to install various dependencies before you will be able to run the Intel OpenVINO installer:

```
sudo apt-get update
sudo apt-get install build-essential ffmpeg cmake checkinstall pkg-config yasm libjpeg-dev curl imagemagick gedit mplayer unzip libpng12-dev libcairo2-dev libpango1.0-dev libgtk2.0-dev libgstreamer0.10-dev libswscale.dev libavcodec-dev libavformat-dev
```

### Installing OpenCL Support

If you also want to use the OpenCL support for GPU-based hardware acceleration, you must install the OpenCL runtime. First, install the dependencies:

```
sudo apt-get update
sudo apt-get install build-essential ffmpeg cmake checkinstall pkg-config yasm libjpeg-dev curl imagemagick gedit mplayer unzip libpng12-dev libcairo2-dev libpango1.0-dev libgtk2.0-dev libgstreamer0.10-dev libswscale.dev libavcodec-dev libavformat-dev
```

Next, obtain the OpenCL runtime package:

```
wget http://registrationcenter-download.intel.com/akdlm/irc_nas/11396/SRB5.0_linux64.zip
unzip SRB5.0_linux64.zip -d SRB5.0_linux64
cd SRB5.0_linux64
```

Last, install the OpenCL runtime:

```
sudo apt-get install xz-utils
mkdir intel-opencl
tar -C intel-opencl -Jxf intel-opencl-r5.0-63503.x86_64.tar.xz
tar -C intel-opencl -Jxf intel-opencl-devel-r5.0-63503.x86_64.tar.xz
tar -C intel-opencl -Jxf intel-opencl-cpu-r5.0-63503.x86_64.tar.xz
sudo cp -R intel-opencl/* /
sudo ldconfig
```

### Installing Intel OpenVINO toolkit

The most recent version of the Intel OpenVINO toolkit is currently R1. You can obtain it from here:

https://software.intel.com/en-us/openvino-toolkit

One you have downloaded the compressed file, unzip the contents, and then run the `install.sh` program within the extracted directory.

## How to build/run code

Setup the environment for the Intel OpenVINO toolkit, by running the `setupvars.sh` program included with OpenVINO:

```
source /opt/intel/computer_vision_sdk/bin/setupvars.sh
```

Then set the needed other exports for building/running GoCV code by running the `env.sh` that is in the GoCV `openvino` directory:

```
source openvino/env.sh
```

You only need to do these two steps one time per session. Once you have run them, you do not need to run them again until you close your terminal window.

Now you can run the version command example to make sure you are compiling/linking against Intel OpenVINO:

```
$ go run ./cmd/version/main.go 
gocv version: 0.11.0
opencv lib version: 3.4.1-cvsdk_2018_1.0.5
```

Examples that use the Intel OpenVINO toolkit can be found in the `cmd/openvino` directory of this repository.
