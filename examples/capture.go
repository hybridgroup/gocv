package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	deviceID := 0
	webcam := opencv3.NewVideoCapture()
	defer webcam.Delete()

	if ok := webcam.OpenDevice(int(deviceID)); !ok {
		fmt.Printf("error opening device: %v\n", deviceID)
	}

	// streaming, capture from webcam
	buf := opencv3.NewMatVec3b()
	defer buf.Delete()
	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(buf); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if buf.Empty() {
			continue
		}

		fmt.Println("frame")
	}
}
