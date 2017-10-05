// Package version outputs the current OpenCV library version to the console.
package version

import (
	"fmt"

	"github.com/hybridgroup/gocv"
)

func Run() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
}
