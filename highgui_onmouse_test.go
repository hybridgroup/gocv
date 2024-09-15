// Do not run these tests on mac OS X. They fail with errors suggesting the GUI
// should only be touched from the main thread.
//go:build !darwin
// +build !darwin

package gocv

import (
	"testing"
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
}
