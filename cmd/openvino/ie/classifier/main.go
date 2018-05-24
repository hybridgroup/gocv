// What it does:
//
// This example...
//
// How to run:
//
// 		go run ./cmd/openvino/ie/classifier/main.go 0 [xmlfile] [binfile] [descriptionsfile]
//
// +//build example

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/openvino/ie"
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
	if len(os.Args) < 4 {
		fmt.Println("How to run:\nclassifier [camera ID] [model] [weights] [descriptionsfile]")
		return
	}

	fmt.Println("IE Version:", ie.Version())

	// parse args
	deviceID, _ := strconv.Atoi(os.Args[1])
	model := os.Args[2]
	weights := os.Args[3]

	//descr := os.Args[4]
	// descriptions, err := readDescriptions(descr)
	// if err != nil {
	// 	fmt.Printf("Error reading descriptions file: %v\n", descr)
	// 	return
	// }

	// open capture device
	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("OpenVINO IE Classifier")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	pd := ie.NewInferenceEnginePlugin()
	defer pd.Close()

	nr := ie.NewCNNNetReader()
	defer nr.Close()

	nr.ReadNetwork(model)
	nr.ReadWeights(weights)

	net := nr.GetNetwork()
	defer net.Close()

	// status := "Ready"
	// statusColor := color.RGBA{0, 255, 0, 0}
	fmt.Printf("Start reading camera device: %v\n", deviceID)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image Mat to 224x224 blob that the classifier can analyze
		//blob := gocv.BlobFromImage(img, 1.0, image.Pt(224, 224), gocv.NewScalar(104, 117, 123, 0), false, false)

		// feed the blob into the classifier
		//net.SetInput(blob, "data")

		// run a forward pass thru the network
		//prob := net.Forward("prob")

		// reshape the results into a 1x1000 matrix
		//probMat := prob.Reshape(1, 1)

		// determine the most probable classification
		//_, maxVal, _, maxLoc := gocv.MinMaxLoc(probMat)

		// display classification
		//status = fmt.Sprintf("description: %v, maxVal: %v\n", descriptions[maxLoc.X], maxVal)
		//gocv.PutText(&img, status, image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, statusColor, 2)

		//blob.Close()
		//prob.Close()
		//probMat.Close()

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
