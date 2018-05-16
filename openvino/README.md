# Using the Intel OpenVINO toolkit

The [Intel OpenVINO toolkit](https://software.intel.com/en-us/openvino-toolkit) is a set of tools and libraries for computer vision applications, that uses computer vision and imaging algorithms developed at Intel. It also includes a complete build of OpenCV.

GoCV support for the Intel OpenVINO Photography Vision Library (PVL) which can be found in the "gocv.io/x/gocv/openvino/pvl" package. Check out the README.md in the `pvl` directory for more information.

## Installing Intel OpenVINO toolkit

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
