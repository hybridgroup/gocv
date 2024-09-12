package gocv

import (
	"testing"
	"unsafe"
)

type mouseHandlerUserData struct {
	name string
}

func mouseHandler(event int, x int, y int, flags int, userdata interface{}) {}

func TestMouseHandler(t *testing.T) {
	windowName := "mouse"

	w := NewWindow(windowName)
	defer w.Close()

	udata := mouseHandlerUserData{
		name: "gocv",
	}

	w.SetMouseHandler(mouseHandler, &udata)
	go_onmouse_dispatcher(1, 2, 3, 4, unsafe.Pointer(&windowName))

}
