# Cuda

In order to use the cuda package, the cuda toolkit from nvidia needs to be installed on the host system. 
 CUDA backend for DNN module requires CC (Compute Capability) 5.3 or higher. Check your GPU https://developer.nvidia.com/cuda-gpus

Please see https://docs.nvidia.com/cuda/index.html for more information.

Furthermore opencv must be compiled with cuda support.

GoCV supports using the Cuda as a backend for the OpenCV deep neural network (DNN) module.

## How to use

This code loads a Caffe model, and then uses Cuda to prepare it for execution on the GPU:

```go
net := gocv.ReadNet("/path/to/your/model.caffemodel", "/path/to/your/config.proto")
if net.Empty() {
    fmt.Println("Error reading network model")
    return
}

net.SetPreferableBackend(gocv.NetBackendType(gocv.NetBackendCUDA))
net.SetPreferableTarget(gocv.NetTargetType(gocv.NetTargetCUDA))
```

## Install Cuda
Download and install packages from https://developer.nvidia.com/cuda-downloads

	example 'cuda_10.2.89_440.33.01_linux.run'

Download and install packages from https://developer.nvidia.com/rdp/cudnn-archive

	example 'cuDNN Runtime Library for Ubuntu18.04 (Deb)' and 'cuDNN Developer Library for Ubuntu18.04 (Deb)'

## Compiling opencv with cuda

For now we have included the make target `install_cuda` that compiles opencv with cuda. (For more details on the compilation process please see the `Makefile`)

Simply issue the command `make install_cuda` and you should be good to go.

If you need static opencv libraries

	make install_cuda BUILD_SHARED_LIBS=OFF

Then finally verify that it is all working 

    cd $GOPATH/src/gocv.io/x/gocv
	go run ./cmd/cuda/main.go
	
You should see something along the lines of:

    gocv version: 0.19.0
    cuda information:
      Device 0:  "GeForce MX150"  2003Mb, sm_61, Driver/Runtime ver.10.0/10.0
