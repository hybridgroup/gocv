.ONESHELL:
.PHONY: test examples deps download build clean

test:
	go test .

EXAMPLES := $(shell ls examples/*.go | sed -e 's/examples\///')

examples:
	for example in $(EXAMPLES) ; do \
		go build -o /tmp/$$example examples/$$example ; \
	done ; \

deps:
	sudo apt-get update
	sudo apt-get install build-essential
	sudo apt-get install cmake git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev
	sudo apt-get install libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libjasper-dev libdc1394-22-dev

download:
	mkdir /tmp/opencv
	cd /tmp/opencv
	wget -O opencv.zip https://github.com/opencv/opencv/archive/3.3.0.zip
	unzip opencv.zip
	wget -O opencv_contrib.zip https://github.com/opencv/opencv_contrib/archive/3.3.0.zip
	unzip opencv_contrib.zip

build:
	cd /tmp/opencv/opencv-3.3.0
	mkdir build
	cd build
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D INSTALL_JAVA_EXAMPLES=OFF -D INSTALL_PYTHON_EXAMPLES=OFF -D INSTALL_C_EXAMPLES=OFF -D OPENCV_EXTRA_MODULES_PATH=/tmp/opencv/opencv_contrib-3.3.0/modules -D BUILD_EXAMPLES=OFF ..
	make -j4
	sudo make install
	sudo ldconfig

clean:
	cd ~
	rm -rf /tmp/opencv

install: download build clean
