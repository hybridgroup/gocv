# Using the Intel OpenVINO Inference Engine

The Intel OpenVINO Inference Engine is a set of libraries for executing convolutional neural networks.

GoCV support for the Intel OpenVINO Inference Engine will be able to be found here in the "gocv.io/x/gocv/openvino/ie" package.

## How It Works

Support in GoCV for the Intel OpenVINO Inference Engine requires version R2+ in order to work. Sinec that version is not yet released, you can install OpenVINO version R1, obtain OpenCV 3.4.2-dev and then compile as described here:

https://github.com/opencv/opencv/wiki/Intel%27s-Deep-Learning-Inference-Engine-backend

You will likely need to specify the modules to be pre-loaded in order to use the OpenVINO IE:

    export LD_LIBRARY_PATH=/opt/intel/computer_vision_sdk/deployment_tools/inference_engine/samples/build/intel64/Release/lib/:/opt/intel/computer_vision_sdk/deployment_tools/inference_engine/external/mkltiny_lnx/lib/:$LD_LIBRARY_PATH

    export LD_PRELOAD=/opt/intel/computer_vision_sdk/deployment_tools/inference_engine/external/mkltiny_lnx/lib/libiomp5.so:/opt/intel/computer_vision_sdk/deployment_tools/inference_engine/external/cldnn/lib/libclDNN64.so

## How to use

This loads a Caffe model, and then uses OpenVINO inference engine to prepare it for execution on the GPU:

```go
net := gocv.ReadNet("/path/to/your/model.caffemodel", "/path/to/your/config.proto")
if net.Empty() {
    fmt.Println("Error reading network model")
    return
}

net.SetPreferableBackend(gocv.NetBackendType("openvino"))
net.SetPreferableTarget(gocv.NetTargetType("fp16"))
```
