// What it does:
//
// This example uses the VideoCapture class to capture video
// then displays the ASCII representation of the video frames.
//
// How to run:
//
// 		go run ./cmd/asciicam/main.go
//

package main

import (
	"fmt"
	"os"

	"github.com/subeshb1/wasm-go-image-to-ascii/convert"
	"gocv.io/x/gocv"
)

const (
	// FixedWidth is the fixed width of the ASCII image
	FixedWidth = 80

	// FixedHeight is the fixed height of the ASCII image
	FixedHeight = 40
)

var buf gocv.Mat

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tcaptest [camera ID]")
		return
	}

	// parse args
	deviceID := os.Args[1]

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// streaming, capture from webcam
	buf = gocv.NewMat()
	defer buf.Close()

	asciiConverter := convert.NewImageConverter()
	opts := &convert.Options{FixedWidth: FixedWidth, FixedHeight: FixedHeight, Colored: true}

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&buf); !ok {
			fmt.Printf("Device error: %v\n", deviceID)
			continue
		}
		if buf.Empty() {
			continue
		}

		img, err := buf.ToImage()
		if err != nil {
			fmt.Printf("Error converting mat to image: %v\n", err)
			continue
		}

		// clear screen
		fmt.Print("\033[H\033[2J")

		// print the ASCII representation of the frame
		fmt.Println(asciiConverter.Image2ASCIIString(img, opts))
	}
}
