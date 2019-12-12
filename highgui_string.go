package gocv

/*
#include <stdlib.h>
#include "highgui_gocv.h"
*/
import "C"

func (c WindowFlag) String() string {
	switch c {
	case WindowNormal:
		return "window-normal"
	case WindowFullscreen:
		return "window-fullscreen"
	case WindowFreeRatio:
		return "window-free-ratio"
	}
	return ""
}

func (c WindowPropertyFlag) String() string {
	switch c {
	case WindowPropertyFullscreen:
		return "window-property-fullscreen"
	case WindowPropertyAutosize:
		return "window-property-autosize"
	case WindowPropertyAspectRatio:
		return "window-property-aspect-ratio"
	case WindowPropertyOpenGL:
		return "window-property-opengl"
	case WindowPropertyVisible:
		return "window-property-visible"
	}
	return ""
}
