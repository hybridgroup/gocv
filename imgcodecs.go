package opencv3

/*
#include <stdlib.h>
#include "imgcodecs.h"
*/
import "C"
import (
	"unsafe"
)

// IMRead reads an image file into a Mat
func IMRead(name string, flags int) Mat {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return Mat{p: C.Image_IMRead(cName, C.int(flags))}
}
