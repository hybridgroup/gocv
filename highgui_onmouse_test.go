package gocv

import (
	"fmt"
	"testing"
)

func mcb(event int, x int, y int, flags int, userdata interface{}) {
	name := *(userdata.(*string))
	fmt.Println(name, event, x, y, flags)
}

func TestMouseCallback(t *testing.T) {
	//t.Skip("TODO: figure out how to implement a test that can exercise the GUI")

	// Comment'd out just for the sake of testing
	// changes on this feature or until we find a
	// proper way to run this test.

	w := NewWindow("mouse")
	defer w.Close()

	name := "gocv"

	w.SetMouseCallback(mcb, &name)

	m := IMRead("images/face-detect.jpg", IMReadColor)
	defer m.Close()

outer_for:
	for {
		w.IMShow(m)
		switch w.WaitKey(5) {
		case 'q':
			break outer_for
		}
	}

}
