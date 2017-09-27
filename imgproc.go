package opencv3

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"

// Rectangle draws a rectangle using to target image Mat.
func Rectangle(img Mat, r Rect) {
	cRect := C.struct_Rect{
		x:      C.int(r.X),
		y:      C.int(r.Y),
		width:  C.int(r.Width),
		height: C.int(r.Height),
	}

	C.Rectangle(img.p, cRect)
}
