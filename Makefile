.ONESHELL:
.PHONY: test deps download build clean astyle cmds docker

# GoCV version to use.
GOCV_VERSION?="v0.39.0"

# OpenCV version to use.
OPENCV_VERSION?=4.10.0

# Go version to use when building Docker image
GOVERSION?=1.22.4

# Temporary directory to put files into.
TMP_DIR?=/tmp/

# Build shared or static library
BUILD_SHARED_LIBS?=ON

# Package list for each well-known Linux distribution
RPMS=cmake curl wget git gtk2-devel libpng-devel libjpeg-devel libtiff-devel tbb tbb-devel libdc1394-devel unzip gcc-c++
DEBS=unzip wget build-essential cmake curl git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libdc1394-22-dev libharfbuzz-dev libfreetype6-dev
DEBS_BOOKWORM=unzip wget build-essential cmake curl git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbbmalloc2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libharfbuzz-dev libfreetype6-dev
DEBS_UBUNTU_JAMMY=unzip wget build-essential cmake curl git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libdc1394-dev libharfbuzz-dev libfreetype6-dev
JETSON=build-essential cmake git unzip pkg-config libjpeg-dev libpng-dev libtiff-dev libavcodec-dev libavformat-dev libswscale-dev libgtk2.0-dev libcanberra-gtk* libxvidcore-dev libx264-dev libgtk-3-dev libtbb2 libtbb-dev libdc1394-22-dev libv4l-dev v4l-utils libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev libavresample-dev libvorbis-dev libxine2-dev libfaac-dev libmp3lame-dev libtheora-dev libopencore-amrnb-dev libopencore-amrwb-dev libopenblas-dev libatlas-base-dev libblas-dev liblapack-dev libeigen3-dev gfortran libhdf5-dev protobuf-compiler libprotobuf-dev libgoogle-glog-dev libgflags-dev

explain:
	@echo "For quick install with typical defaults of both OpenCV and GoCV, run 'make install'"

# Detect Linux distribution
distro_deps=
ifneq ($(shell which dnf 2>/dev/null),)
	distro_deps=deps_fedora
else
ifneq ($(shell which apt-get 2>/dev/null),)
ifneq ($(shell cat /etc/os-release 2>/dev/null | grep "Jammy Jellyfish"),)
	distro_deps=deps_ubuntu_jammy
else
ifneq ($(shell cat /etc/debian_version 2>/dev/null | grep "12."),)
	distro_deps=deps_debian_bookworm
else
	distro_deps=deps_debian
endif
endif
else
ifneq ($(shell which yum 2>/dev/null),)
	distro_deps=deps_rh_centos
endif
endif
endif

# Install all necessary dependencies.
deps: $(distro_deps)

deps_rh_centos:
	sudo yum -y install pkgconfig $(RPMS)

deps_fedora:
	sudo dnf -y install pkgconf-pkg-config $(RPMS)

deps_debian_bookworm:
	sudo apt-get -y update
	sudo apt-get -y install $(DEBS_BOOKWORM)

deps_debian:
	sudo apt-get -y update
	sudo apt-get -y install $(DEBS)

deps_ubuntu_jammy:
	sudo apt-get -y update
	sudo apt-get -y install $(DEBS_UBUNTU_JAMMY)

deps_jetson:
	sudo sh -c "echo '/usr/local/cuda/lib64' >> /etc/ld.so.conf.d/nvidia-tegra.conf"
	sudo ldconfig
	sudo apt-get -y update
	sudo apt-get -y install $(JETSON)

# Download OpenCV source tarballs.
download:
	rm -rf $(TMP_DIR)opencv
	mkdir $(TMP_DIR)opencv
	cd $(TMP_DIR)opencv
	curl -Lo opencv.zip https://github.com/opencv/opencv/archive/refs/tags/$(OPENCV_VERSION).zip
	unzip -q opencv.zip
	curl -Lo opencv_contrib.zip https://github.com/opencv/opencv_contrib/archive/refs/tags/$(OPENCV_VERSION).zip
	unzip -q opencv_contrib.zip
	rm opencv.zip opencv_contrib.zip
	cd -

# Download openvino source tarballs.
download_openvino:
	sudo rm -rf /usr/local/dldt/
	sudo rm -rf /usr/local/openvino/
	sudo git clone https://github.com/openvinotoolkit/openvino -b 2019_R3.1 /usr/local/openvino/

# Build openvino.
build_openvino_package:
	cd /usr/local/openvino/inference-engine
	sudo git submodule init
	sudo git submodule update --recursive
	sudo ./install_dependencies.sh
	sudo mv -f thirdparty/clDNN/common/intel_ocl_icd/6.3/linux/Release thirdparty/clDNN/common/intel_ocl_icd/6.3/linux/RELEASE
	sudo mkdir build
	cd build
	sudo rm -rf *
	sudo cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D ENABLE_VPU=ON -D ENABLE_MKL_DNN=ON -D ENABLE_CLDNN=ON ..
	sudo $(MAKE) -j $(shell nproc --all --ignore 1)
	sudo touch VERSION
	sudo mkdir -p src/ngraph
	sudo cp thirdparty/ngraph/src/ngraph/version.hpp src/ngraph
	cd -

# Build OpenCV.
build:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -D WITH_TBB=ON -DOPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV on Raspbian with ARM hardware optimizations.
build_raspi:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
ifneq ($(shell uname -m | grep "aarch64"),)
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=OFF -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D ENABLE_NEON=ON -D WITH_JASPER=OFF -D WITH_TBB=ON -D OPENCV_GENERATE_PKGCONFIG=ON -D WITH_FREETYPE=ON ..
else
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=OFF -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D ENABLE_NEON=ON -D ENABLE_VFPV3=ON -D WITH_JASPER=OFF -D OPENCV_GENERATE_PKGCONFIG=ON -D WITH_FREETYPE=ON ..
endif
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV on Raspberry pi zero which has ARMv6.
build_raspi_zero:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=OFF -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D ENABLE_VFPV2=ON -D WITH_JASPER=OFF -D OPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV for NVidia Jetson with CUDA.
build_jetson:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE \
		-D CMAKE_INSTALL_PREFIX=/usr/local \
		-D EIGEN_INCLUDE_PATH=/usr/include/eigen3 \
		-D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} \
		-D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules \
		-D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=OFF -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO \
		-D WITH_OPENCL=OFF \
		-D WITH_CUDA=ON \
		-D CUDA_ARCH_BIN=5.3 \
		-D CUDA_ARCH_PTX="" \
		-D WITH_CUDNN=ON \
		-D WITH_CUBLAS=ON \
		-D ENABLE_FAST_MATH=ON \
		-D CUDA_FAST_MATH=ON \
		-D OPENCV_DNN_CUDA=ON \
		-D ENABLE_NEON=ON \
		-D WITH_QT=OFF \
		-D WITH_OPENMP=ON \
		-D WITH_OPENGL=ON \
		-D BUILD_TIFF=ON \
		-D WITH_FFMPEG=ON \
		-D WITH_GSTREAMER=ON \
		-D WITH_TBB=ON \
		-D BUILD_TBB=ON \
		-D BUILD_TESTS=OFF \
		-D WITH_EIGEN=ON \
		-D WITH_V4L=ON \
		-D WITH_LIBV4L=ON \
		-D OPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV with non-free contrib modules.
build_nonfree:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -D WITH_TBB=ON -DOPENCV_GENERATE_PKGCONFIG=ON -DOPENCV_ENABLE_NONFREE=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV with openvino.
build_openvino:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D ENABLE_CXX11=ON -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D WITH_INF_ENGINE=ON -D InferenceEngine_DIR=/usr/local/dldt/inference-engine/build -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -D WITH_TBB=ON -DOPENCV_GENERATE_PKGCONFIG=ON -DOPENCV_ENABLE_NONFREE=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV with cuda.
build_cuda:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -D WITH_TBB=ON -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_CUDA=ON -DENABLE_FAST_MATH=1 -DCUDA_FAST_MATH=1 -DWITH_CUBLAS=ON -DCUDA_TOOLKIT_ROOT_DIR=/usr/local/cuda/ -DBUILD_opencv_cudacodec=OFF -D WITH_CUDNN=ON -D OPENCV_DNN_CUDA=ON -D CUDA_GENERATION=Auto -DOPENCV_ENABLE_NONFREE=ON -D WITH_GSTREAMER=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV statically linked
build_static:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=OFF -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -DWITH_JASPER=OFF -DWITH_QT=OFF -DWITH_GTK=OFF -DWITH_FFMPEG=OFF -DWITH_TIFF=OFF -DWITH_WEBP=OFF -DWITH_PNG=OFF -DWITH_1394=OFF -DWITH_OPENJPEG=OFF -DOPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Build OpenCV with cuda.
build_all:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -j $(shell nproc --all --ignore 1) -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D ENABLE_CXX11=ON -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D WITH_INF_ENGINE=ON -D InferenceEngine_DIR=/usr/local/dldt/inference-engine/build -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=ON -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -D WITH_TBB=ON -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_CUDA=ON -DENABLE_FAST_MATH=1 -DCUDA_FAST_MATH=1 -DWITH_CUBLAS=1 -DCUDA_TOOLKIT_ROOT_DIR=/usr/local/cuda/ -DBUILD_opencv_cudacodec=OFF -D WITH_CUDNN=ON -D OPENCV_DNN_CUDA=ON -D CUDA_GENERATION=Auto ..
	$(MAKE) -j $(shell nproc --all --ignore 1)
	$(MAKE) preinstall
	cd -

# Cleanup temporary build files.
clean:
	go clean --cache
	rm -rf $(TMP_DIR)opencv

# Cleanup old library files.
sudo_pre_install_clean:
ifneq (,$(wildcard /usr/local/lib/libopencv*))
	sudo rm -rf /usr/local/lib/cmake/opencv4/
	sudo rm -rf /usr/local/lib/libopencv*
	sudo rm -rf /usr/local/lib/pkgconfig/opencv*
	sudo rm -rf /usr/local/include/opencv*
else
ifneq (,$(wildcard /usr/local/lib64/libopencv*))
	sudo rm -rf /usr/local/lib64/cmake/opencv4/
	sudo rm -rf /usr/local/lib64/libopencv*
	sudo rm -rf /usr/local/lib64/pkgconfig/opencv*
	sudo rm -rf /usr/local/include/opencv*
else
ifneq (,$(wildcard /usr/local/lib/aarch64-linux-gnu/libopencv*))
	sudo rm -rf /usr/local/lib/aarch64-linux-gnu/cmake/opencv4/
	sudo rm -rf /usr/local/lib/aarch64-linux-gnu/libopencv*
	sudo rm -rf /usr/local/lib/aarch64-linux-gnu/pkgconfig/opencv*
	sudo rm -rf /usr/local/include/opencv*
endif
endif
endif

# Do everything.
install: deps download sudo_pre_install_clean build sudo_install clean verify

# Do everything on Raspbian.
install_raspi: deps download sudo_pre_install_clean build_raspi sudo_install clean verify

# Do everything on the raspberry pi zero.
install_raspi_zero: deps download sudo_pre_install_clean build_raspi_zero sudo_install clean verify

# Do everything on Jetson.
install_jetson: deps download sudo_pre_install_clean build_jetson sudo_install clean verify

# Do everything with cuda.
install_cuda: deps download sudo_pre_install_clean build_cuda sudo_install clean verify verify_cuda

# Do everything with openvino.
install_openvino: deps download download_openvino sudo_pre_install_clean build_openvino_package sudo_install_openvino build_openvino sudo_install clean verify_openvino

# Do everything statically.
install_static: deps download sudo_pre_install_clean build_static sudo_install clean verify_static

# Do everything with non-free modules from cpencv_contrib.
install_nonfree: deps download sudo_pre_install_clean build_nonfree sudo_install clean verify

# Do everything with openvino and cuda.
install_all: deps download download_openvino sudo_pre_install_clean build_openvino_package sudo_install_openvino build_all sudo_install clean verify_openvino verify_cuda

# Install system wide.
sudo_install:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)/build
	sudo $(MAKE) install
	sudo ldconfig
	cd -

# Install system wide.
sudo_install_openvino:
	cd /usr/local/openvino/inference-engine/build
	sudo $(MAKE) install
	sudo ldconfig
	cd -

# Build a minimal Go app to confirm gocv works.
verify:
	go run ./cmd/version/main.go

# Build a minimal Go app to confirm gocv works with statically built OpenCV.
verify_static:
	go run -tags static ./cmd/version/main.go

# Build a minimal Go app to confirm gocv cuda works.
verify_cuda:
	go run ./cmd/cuda/main.go

# Build a minimal Go app to confirm gocv openvino works.
verify_openvino:
	go run -tags openvino ./cmd/version/main.go

# Runs tests.
# This assumes env.sh was already sourced.
# pvt is not tested here since it requires additional depenedences.
test:
	go test -tags matprofile . ./contrib

test_cuda:
	go test -tags matprofile ./cuda

docker:
	docker build --build-arg OPENCV_VERSION=$(OPENCV_VERSION) --build-arg GOVERSION=$(GOVERSION) .

astyle:
	astyle --project=.astylerc --recursive *.cpp,*.h


releaselog:
	git log --pretty=format:"%s" $(GOCV_VERSION)..HEAD

CMDS=basic-drawing caffe-classifier captest capwindow counter dnn-detection dnn-pose-detection dnn-style-transfer faceblur facedetect facedetect-from-url feature-matching find-chessboard find-circles find-lines hand-gestures hello img-similarity mjpeg-streamer motion-detect saveimage savevideo showimage ssd-facedetect tf-classifier tracking version xphoto
cmds:
	for cmd in $(CMDS) ; do \
		go build -o build/$$cmd cmd/$$cmd/main.go ;
	done ; \
