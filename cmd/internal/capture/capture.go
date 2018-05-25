package capture

import (
	"strconv"

	"gocv.io/x/gocv"
)

func Open(deviceID string) (*gocv.VideoCapture, error) {
	id, err := strconv.Atoi(deviceID)
	if err == nil {
		return gocv.VideoCaptureDevice(int(id))
	}
	return gocv.VideoCaptureFile(deviceID)
}
