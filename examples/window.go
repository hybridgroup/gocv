package main

import (
	"time"

	opencv3 ".."
)

func main() {
	opencv3.NewWindow("Hello")

	for {
		opencv3.WaitKey(1)
		time.Sleep(100 * time.Microsecond)
	}
}
