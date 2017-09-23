// What it does:
//
// 	This program outputs the current OpenCV library version to the console.
//
// How to run:
//
// 		go run ./examples/showinfo.go
//
// +build example

package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	fmt.Printf("go-opencv3 version: %s\n", opencv3.Version())
	fmt.Printf("opencv lib version: %s\n", opencv3.OpenCVVersion())
}
