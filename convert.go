package opencv3

/*
#include "convert.h"
*/
import "C"
import (
	"unsafe"
)

func toByteArray(b []byte) C.struct_ByteArray {
	return C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}
}

// toGoBytes returns binary data. Serializing is depends on C/C++ implementation.
func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}
