// What it does:
//
// This example uses a deep neural network to perform object detection
// of whatever is in front of the camera.
//
// Download the model file from:
// ...
//
// Also, you will need the prototxt file:
// ...
//
// And the words text file with the class descriptions:
//
//
// How to run:
//
// 		go run ./cmd/dnn-detection/main.go 0 [modelfile] [configfile] [descriptionsfile] [backend] [device]
//
// +build example

package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

// readDescriptions reads the descriptions from a file
// and returns a slice of its lines.
func readDescriptions(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("How to run:\ndnn-detection [camera ID] [modelfile] [configfile] [descriptionsfile] ([backend] [device])")
		return
	}

	// parse args
	deviceID := os.Args[1]
	proto := os.Args[2]
	model := os.Args[3]
	//descr := os.Args[4]
	// descriptions, err := readDescriptions(descr)
	// if err != nil {
	// 	fmt.Printf("Error reading descriptions file: %v\n", descr)
	// 	return
	// }

	backend := gocv.NetBackendDefault
	if len(os.Args) > 5 {
		backend = gocv.ParseNetBackend(os.Args[5])
	}

	target := gocv.NetTargetCPU
	if len(os.Args) > 6 {
		target = gocv.ParseNetTarget(os.Args[6])
	}

	// open capture device
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("DNN Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	// open DNN classifier
	net := gocv.ReadNet(model, proto)
	if net.Empty() {
		fmt.Printf("Error reading network model from : %v %v\n", model, proto)
		return
	}
	defer net.Close()
	net.SetPreferableBackend(gocv.NetBackendType(backend))
	net.SetPreferableTarget(gocv.NetTargetType(target))

	// status := "Ready"
	statusColor := color.RGBA{0, 255, 0, 0}
	fmt.Printf("Start reading camera device: %v\n", deviceID)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image Mat to 300x300 blob that the classifier can analyze
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(300, 300), gocv.NewScalar(104, 177, 123, 0), false, false)

		// feed the blob into the classifier
		net.SetInput(blob, "")

		// run a forward pass thru the network
		probs := net.ForwardLayers([]string{"detection_out"})

		// make sure we have valid output layer type
		outs := net.GetUnconnectedOutLayers()
		layer := net.GetLayer(outs[1])
		layerType := layer.GetType()
		layer.Close()
		if layerType != "DetectionOutput" {
			fmt.Printf("Error unknown output layer type %v\n", layerType)
			continue
		}

		for i := 0; i < probs[0].Total(); i += 7 {
			confidence := probs[0].GetFloatAt(0, i+2)
			if confidence > 0.5 {
				classId := probs[0].GetFloatAt(0, i+1) - 1
				left := int(probs[0].GetFloatAt(0, i+3) * float32(img.Cols()))
				top := int(probs[0].GetFloatAt(0, i+4) * float32(img.Rows()))
				right := int(probs[0].GetFloatAt(0, i+5) * float32(img.Cols()))
				bottom := int(probs[0].GetFloatAt(0, i+6) * float32(img.Rows()))
				fmt.Println(classId, confidence, left, top, right, bottom)
				gocv.Rectangle(&img, image.Rect(left, top, right, bottom), statusColor, 2)
			}
		}

		probs[0].Close()
		blob.Close()

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
