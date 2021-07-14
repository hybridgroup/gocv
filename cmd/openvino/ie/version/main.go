// What it does:
//
// 	This program outputs the current OpenVINO IE library version to the console.
//
// How to run:
//
// 		go run ./cmd/openvino/ie/version/main.go
//

package main

import (
	"fmt"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/openvino/ie"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("OpenVINO Inference Engine version: %s\n", ie.Version())
}
