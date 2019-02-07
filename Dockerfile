FROM ubuntu:16.04 AS opencv
LABEL maintainer="hybridgroup"

RUN apt-get update && apt-get install -y --no-install-recommends \
            git build-essential cmake pkg-config unzip libgtk2.0-dev \
            curl ca-certificates libcurl4-openssl-dev libssl-dev \
            libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev \
            libjpeg-dev libpng-dev libtiff-dev libdc1394-22-dev && \
            rm -rf /var/lib/apt/lists/*

ARG OPENCV_VERSION="4.0.1"
ENV OPENCV_VERSION $OPENCV_VERSION

RUN curl -Lo opencv.zip https://github.com/opencv/opencv/archive/${OPENCV_VERSION}.zip && \
            unzip -q opencv.zip && \
            curl -Lo opencv_contrib.zip https://github.com/opencv/opencv_contrib/archive/${OPENCV_VERSION}.zip && \
            unzip -q opencv_contrib.zip && \
            rm opencv.zip opencv_contrib.zip && \
            cd opencv-${OPENCV_VERSION} && \
            mkdir build && cd build && \
            cmake -D CMAKE_BUILD_TYPE=RELEASE \
                  -D CMAKE_INSTALL_PREFIX=/usr/local \
                  -D OPENCV_EXTRA_MODULES_PATH=../../opencv_contrib-${OPENCV_VERSION}/modules \
                  -D WITH_JASPER=OFF \
                  -D BUILD_DOCS=OFF \
                  -D BUILD_EXAMPLES=OFF \
                  -D BUILD_TESTS=OFF \
                  -D BUILD_PERF_TESTS=OFF \
                  -D BUILD_opencv_java=NO \
                  -D BUILD_opencv_python=NO \
                  -D BUILD_opencv_python2=NO \
                  -D BUILD_opencv_python3=NO \
                  -D OPENCV_GENERATE_PKGCONFIG=ON .. && \
            make -j $(nproc --all) && \
            make preinstall && make install && ldconfig && \
            cd / && rm -rf opencv*

#################
#  Go + OpenCV  #
#################
FROM opencv AS gocv
LABEL maintainer="hybridgroup"

ARG GOVERSION="1.11.2"
ENV GOVERSION $GOVERSION

RUN apt-get update && apt-get install -y --no-install-recommends \
            git software-properties-common && \
            curl -Lo go${GOVERSION}.linux-amd64.tar.gz https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz && \
            tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz && \
            rm go${GOVERSION}.linux-amd64.tar.gz && \
            rm -rf /var/lib/apt/lists/*

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN go get -u -d gocv.io/x/gocv && go run ${GOPATH}/src/gocv.io/x/gocv/cmd/version/main.go
