// +build example
//
// Do not build by default.
//
// This example outputs the current OpenCV library version to the console.
//
// how to run:
// 		go run ./examples/showinfo.go
//
package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	fmt.Printf("go-opencv3 version: %s\n", opencv3.Version())
	fmt.Printf("opencv lib version: %s\n", opencv3.OpenCVVersion())
}
