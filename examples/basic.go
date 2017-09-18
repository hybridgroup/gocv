package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	deviceID := 0
	vcap := opencv3.NewVideoCapture()
	defer vcap.Delete()

	if ok := vcap.OpenDevice(int(deviceID)); !ok {
		fmt.Errorf("error opening device: %v", deviceID)
	}

	// streaming, capture from vcap
	buf := opencv3.NewMatVec3b()
	defer buf.Delete()
	fmt.Printf("start reading camera device: %v", deviceID)
	for {
		if ok := vcap.Read(buf); !ok {
			fmt.Errorf("cannot read a new file (device no: %d)", deviceID)
			return
		}
		if buf.Empty() {
			continue
		}

		fmt.Println("frame")
	}
}
