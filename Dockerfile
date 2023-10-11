# to build this docker image:
#   docker build .
FROM ghcr.io/hybridgroup/opencv:4.8.1

ENV GOPATH /go

COPY . /go/src/gocv.io/x/gocv/

WORKDIR /go/src/gocv.io/x/gocv
RUN go build -tags example -o /build/gocv_version -i ./cmd/version/

CMD ["/build/gocv_version"]
