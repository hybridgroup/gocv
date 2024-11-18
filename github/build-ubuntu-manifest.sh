#!/bin/bash -xe

[ -z "${OPENCV_VERSION}" ] && echo "Need to set OPENCV_VERSION" && exit 1
[ -z "${UBUNTU_VERSION}" ] && echo "Need to set UBUNTU_VERSION" && exit 1

MANIFEST="774915305292.dkr.ecr.us-west-2.amazonaws.com/ubuntu-with-opencv:ubuntu-${UBUNTU_VERSION}-opencv-${OPENCV_VERSION}"
IMAGE_ARM="774915305292.dkr.ecr.us-west-2.amazonaws.com/ubuntu-with-opencv:ubuntu-${UBUNTU_VERSION}-opencv-${OPENCV_VERSION}-arm64"
IMAGE_AMD="774915305292.dkr.ecr.us-west-2.amazonaws.com/ubuntu-with-opencv:ubuntu-${UBUNTU_VERSION}-opencv-${OPENCV_VERSION}-amd64"
aws ecr get-login --no-include-email --region us-west-2 | bash
docker manifest rm "${MANIFEST}" || true
docker manifest create "${MANIFEST}" \
    "${IMAGE_ARM}" \
    "${IMAGE_AMD}"
docker manifest annotate --arch arm64 "${MANIFEST}" \
    "${IMAGE_ARM}"
docker manifest annotate --arch amd64 "${MANIFEST}" \
    "${IMAGE_AMD}"
docker manifest push "${MANIFEST}"
