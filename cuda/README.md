# Cuda (Experimental)

In order to use the cuda package, the cuda toolkit from nvidia needs to be installed on the host system. 

Please see https://docs.nvidia.com/cuda/index.html for more information.

Furthermore opencv must be compiled with cuda support.

## Compiling opencv with cuda

For now we have included the make target `install_cuda` that compiles opencv with cuda. (For more details on the compilation process please see the `Makefile`)

Simply issue the command `make install_cuda` and you should be good to go.

Then finally verify that it is all working 

    cd $GOPATH/src/gocv.io/x/gocv
	go run ./cmd/cuda/main.go
	
You should see something along the lines of:

    gocv version: 0.19.0
    cuda information:
      Device 0:  "GeForce MX150"  2003Mb, sm_61, Driver/Runtime ver.10.0/10.0
