package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	deviceID := 1
	vcap := opencv3.NewVideoCapture()
	defer vcap.Delete()

	if ok := vcap.OpenDevice(int(deviceID)); !ok {
		fmt.Errorf("error opening device: %v", deviceID)
	}
}
