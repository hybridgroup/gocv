#!/bin/bash -xe

[ -z "${OPENCV_VERSION}" ] && echo "Need to set OPENCV_VERSION" && exit 1
[ -z "${UBUNTU_VERSION}" ] && echo "Need to set UBUNTU_VERSION" && exit 1

# Build deploy docker image
GIT_COMMIT_ID=$(git rev-parse HEAD | cut -c1-7)
IMAGE="774915305292.dkr.ecr.us-west-2.amazonaws.com/ubuntu-with-opencv:ubuntu-${UBUNTU_VERSION}-opencv-${OPENCV_VERSION}-arm64"
AWS_CLI_VERSION=$(aws --version 2>&1 | cut -d " " -f1 | cut -d "/" -f2 | cut -c 1)

docker buildx build --no-cache --pull \
  --platform linux/arm64 \
  -t "${IMAGE}" \
  --build-arg OPENCV_VERSION=$OPENCV_VERSION \
  -f Dockerfile.umbo.ubuntu .
aws ecr get-login --no-include-email --region us-west-2 | bash
# Push docker images
docker push "${IMAGE}"
docker rmi -f "${IMAGE}"
