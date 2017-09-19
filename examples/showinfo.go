package main

import (
	"fmt"

	opencv3 ".."
)

func main() {
	fmt.Printf("go-opencv3 version: %s\n", opencv3.Version())
	fmt.Printf("opencv lib version: %s\n", opencv3.OpenCVVersion())
}
