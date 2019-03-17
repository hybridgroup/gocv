// What it does:
//
// 	This program outputs the current OpenCV library version and CUDA version the console.
//
// How to run:
//
// 		go run ./cmd/cuda/main.go
//
// +build example

package main

import (
	"fmt"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/cuda"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Println("cuda information:")
	devices := cuda.GetCudaEnabledDeviceCount()
	for i := 0; i < devices; i++ {
		fmt.Print("  ")
		cuda.PrintShortCudaDeviceInfo(i)
	}
}
