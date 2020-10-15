###################
#  GoCV + OpenCV  #
###################
FROM gocv/opencv:4.5.0
LABEL maintainer="hybridgroup"

ENV GOPATH /go
WORKDIR $GOPATH

RUN go get -u -d gocv.io/x/gocv

WORKDIR ${GOPATH}/src/gocv.io/x/gocv/cmd/version/

RUN go build -o gocv_version -i main.go

CMD ["./gocv_version"]
