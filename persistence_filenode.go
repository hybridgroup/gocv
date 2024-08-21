package gocv

/*
#include <stdlib.h>
#include "persistence.h"

static char* getArrayItem(char** arr, int i) {
	return arr[i];
}
*/
import "C"

type FileNodeType int

const (
	FileNodeTypeNone     FileNodeType = 0
	FileNodeTypeInt      FileNodeType = 1
	FileNodeTypeReal     FileNodeType = 2
	FileNodeTypeFloat    FileNodeType = FileNodeTypeReal
	FileNodeTypeStr      FileNodeType = 3
	FileNodeTypeString   FileNodeType = FileNodeTypeStr
	FileNodeTypeSeq      FileNodeType = 4
	FileNodeTypeMap      FileNodeType = 5
	FileNodeTypeTypeMask FileNodeType = 7
	FileNodeTypeFlow     FileNodeType = 8
	FileNodeTypeUniform  FileNodeType = 8
	FileNodeTypeEmpty    FileNodeType = 16
	FileNodeTypeNamed    FileNodeType = 32
)

// FileNode is a wrapper for the OpenCV FileNode class
//
// Ref: https://docs.opencv.org/4.x/de/dd9/classcv_1_1FileNode.html
type FileNode struct {
	p C.FileNode
}

func (fn *FileNode) Empty() bool {
	return bool(C.FileNode_Empty(fn.p))
}

func (fn *FileNode) IsInt() bool {
	return bool(C.FileNode_IsInt(fn.p))
}

func (fn *FileNode) IsMap() bool {
	return bool(C.FileNode_IsMap(fn.p))
}

func (fn *FileNode) IsNamed() bool {
	return bool(C.FileNode_IsNamed(fn.p))
}

func (fn *FileNode) IsNone() bool {
	return bool(C.FileNode_IsNone(fn.p))
}

func (fn *FileNode) IsReal() bool {
	return bool(C.FileNode_IsReal(fn.p))
}

func (fn *FileNode) IsSeq() bool {
	return bool(C.FileNode_IsSeq(fn.p))
}

func (fn *FileNode) IsString() bool {
	return bool(C.FileNode_IsString(fn.p))
}

func (fn *FileNode) Keys() []string {

	c_keys_count := C.FileNode_KeysCount(fn.p)
	c_keys := C.FileNode_Keys(fn.p)
	defer C.FileNode_KeysFree(c_keys, c_keys_count)

	keys := make([]string, int(c_keys_count))

	for i := 0; i < int(c_keys_count); i++ {
		keys[i] = C.GoString(C.getArrayItem(c_keys, C.int(i)))
	}
	return keys
}

func (fn *FileNode) Close() {
	C.FileNode_Close(fn.p)
}
