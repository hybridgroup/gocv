package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"

func newAsyncArray(p C.AsyncArray) AsyncArray {
	return AsyncArray{p: p}
}
