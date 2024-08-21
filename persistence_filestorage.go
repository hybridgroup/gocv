package gocv

/*
#include <stdlib.h>
#include <stdbool.h>
#include "persistence.h"
*/
import "C"
import "unsafe"

type FileStorageMode int

const (
	FileStorageModeRead        FileStorageMode = 0
	FileStorageModeWrite       FileStorageMode = 1
	FileStorageModeAppend      FileStorageMode = 2
	FileStorageModeMemory      FileStorageMode = 4
	FileStorageModeFormatMask  FileStorageMode = (7 << 3)
	FileStorageModeFormatAuto  FileStorageMode = 0
	FileStorageModeFormatXml   FileStorageMode = (1 << 3)
	FileStorageModeFormatYaml  FileStorageMode = (2 << 3)
	FileStorageModeFormatJson  FileStorageMode = (3 << 3)
	FileStorageModeBase64      FileStorageMode = 64
	FileStorageModeWriteBase64 FileStorageMode = FileStorageModeBase64 | FileStorageModeWrite
)

type FileStorageState int

const (
	FileStorageStateUndefined     FileStorageState = 0
	FileStorageStateValueExpected FileStorageState = 1
	FileStorageStateNameExpected  FileStorageState = 2
	FileStorageStateInsideMap     FileStorageState = 4
)

// FileStorage is a wrapper for the OpenCV FileStorage class
//
// Ref: https://docs.opencv.org/4.x/da/d56/classcv_1_1FileStorage.html
type FileStorage struct {
	p C.FileStorage
}

func MewFileStorage() *FileStorage {
	return &FileStorage{p: C.FileStorage_Create()}
}

func NewFileStorageWithParams(filename string, flags FileStorageMode, encoding string) *FileStorage {
	c_filename := C.CString(filename)
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_filename))
	defer C.free(unsafe.Pointer(c_encoding))

	return &FileStorage{p: C.FileStorage_CreateWithParams(c_filename, C.int(flags), c_encoding)}
}

func (fs *FileStorage) Close() {
	fs.Release()
}

func (fs *FileStorage) Release() {
	C.FileStorage_Release(fs.p)
}

func (fs *FileStorage) ElName() string {
	c_str := C.FileStorage_ElName(fs.p)
	defer C.free(unsafe.Pointer(c_str))

	str := C.GoString(c_str)
	return str
}

func (fs *FileStorage) State() FileStorageState {
	state := C.FileStorage_State(fs.p)
	return FileStorageState(int(state))
}

func (fs *FileStorage) EndWriteStruct() {
	C.FileStorage_EndWriteStruct(fs.p)

}

func (fs *FileStorage) GetFormat() FileStorageMode {
	fmt := C.FileStorage_GetFormat(fs.p)
	return FileStorageMode(int(fmt))
}

func (fs *FileStorage) IsOpened() bool {
	b := C.FileStorage_IsOpened(fs.p)
	return bool(b)
}

func (fs *FileStorage) Open(filename string, flags FileStorageMode, encoding string) bool {
	c_filename := C.CString(filename)
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_filename))
	defer C.free(unsafe.Pointer(c_encoding))

	b := C.FileStorage_Open(fs.p, c_filename, C.int(flags), c_encoding)
	return bool(b)
}

func (fs *FileStorage) ReleaseAndGetString() string {
	c_str := C.FileStorage_ReleaseAndGetString(fs.p)
	defer C.free(unsafe.Pointer(c_str))

	str := C.GoString(c_str)
	return str
}

func (fs *FileStorage) StartWriteStruct(name string, flags FileNodeType, typeName string) {
	c_name := C.CString(name)
	c_typeName := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_name))
	defer C.free(unsafe.Pointer(c_typeName))

	C.FileStorage_StartWriteStruct(fs.p, c_name, C.int(flags), c_typeName)
}

func (fs *FileStorage) WriteMat(name string, mat Mat) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.FileStorage_WriteMat(fs.p, c_name, mat.p)
}

func (fs *FileStorage) WriteString(name string, val string) {
	c_name := C.CString(name)
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_name))
	defer C.free(unsafe.Pointer(c_val))

	C.FileStorage_WriteString(fs.p, c_name, c_val)
}

func (fs *FileStorage) WriteStringArray(name string, val []string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//c_val := []*C.char{}
	c_val := make([]*C.char, 0, len(val))

	for _, v := range val {
		c_val = append(c_val, C.CString(v))
	}
	defer func() {
		for _, p := range c_val {
			C.free(unsafe.Pointer(p))
		}
		//C.free(unsafe.Pointer(c_val))
		//possible leak?
	}()
	C.FileStorage_WriteStringArray(fs.p, c_name, &c_val[0], C.size_t(len(val)))
}

func (fs *FileStorage) WriteDouble(name string, val float32) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.FileStorage_WriteDouble(fs.p, c_name, C.double(val))
}

func (fs *FileStorage) WriteInt(name string, val int) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.FileStorage_WriteInt(fs.p, c_name, C.int(val))
}

func (fs *FileStorage) WriteComment(comment string, append bool) {
	c_comment := C.CString(comment)
	defer C.free(unsafe.Pointer(c_comment))

	C.FileStorage_WriteComment(fs.p, c_comment, C.bool(append))
}

func (fs *FileStorage) WriteRaw(fmt string, vec []byte) {
	c_fmt := C.CString(fmt)
	defer C.free(unsafe.Pointer(c_fmt))

	c_vec := C.CBytes(vec)
	defer C.free(c_vec)

	C.FileStorage_WriteRaw(fs.p, c_fmt, c_vec, C.size_t(len(vec)))

}

func (fs *FileStorage) GetFirstTopLevelNode() *FileNode {
	node_p := C.FileStorage_GetFirstTopLevelNode(fs.p)
	return &FileNode{p: node_p}
}

func (fs *FileStorage) GetNode(name string) *FileNode {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	node_p := C.FileStorage_GetNode(fs.p, c_name)

	return &FileNode{p: node_p}

}

func (fs *FileStorage) Root(streamIdx int) *FileNode {
	node_p := C.FileStorage_Root(fs.p, C.int(streamIdx))
	return &FileNode{p: node_p}
}
