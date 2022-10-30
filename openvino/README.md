# Using the Intel® Distribution of OpenVINO™ toolkit

The [Intel® Distribution of OpenVINO™ toolkit](https://software.intel.com/en-us/openvino-toolkit) is a set of tools and libraries for computer vision applications, that uses computer vision and imaging algorithms developed at Intel. It also includes a complete build of OpenCV 4.5.5. Users must build opencv from source with openvino support to use a newer opencv version (eg. 4.6.0).

GoCV supports using the Intel GPU or Intel OpenVINO Inference Engine as a backend for the OpenCV deep neural network (DNN) module. For details, please see:
https://github.com/hybridgroup/gocv/blob/release/openvino/ie/README.md

## Installing Intel OpenVINO toolkit

The most recent version of the Intel OpenVINO toolkit is currently 2022.1 LTS. You can obtain it from here:

https://software.intel.com/en-us/openvino-toolkit

One you have downloaded the compressed file, unzip the contents, and then run the `install.sh` program within the extracted directory.

## How to build/run code

Setup the environment for the Intel OpenVINO toolkit, by running the `setupvars.sh` program included with OpenVINO:

```
source /opt/intel/openvino_2022/setupvars.sh
```

Then set the needed other exports for building/running GoCV code by running the `env.sh` that is in the GoCV `openvino` directory:

```
source openvino/env.sh
```

You only need to do these two steps one time per session. Once you have run them, you do not need to run them again until you close your terminal window.

Now you can run the version command example to make sure you are compiling/linking against Intel OpenVINO:

```
$ go run -tags customenv  ./cmd/version/main.go
gocv version: 0.31.0
opencv lib version: 4.5.5-openvino
```

Note the use of `-tags customenv` is needed when using `go run`, `go build`, and `go test` with OpenVINO, so the CGo compiler can pickup the correct settings for the environment, and ignore the usual defaults.

Examples that use the Intel OpenVINO toolkit can be found in the `cmd/openvino` directory of this repository.
