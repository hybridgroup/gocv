// how to use
// 		go run ./examples/window.go /home/ron/Pictures/mcp23017.jpg
//
package main

import (
	"os"
	"time"

	opencv3 ".."
)

func main() {
	filename := os.Args[1]
	window := opencv3.NewWindow("Hello")
	img := opencv3.IMRead(filename, 1)

	for {
		window.IMShow(img)
		opencv3.WaitKey(1)
		time.Sleep(100 * time.Microsecond)
	}
}
