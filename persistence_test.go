package gocv

import (
	"testing"
)

func TestFileStorage(t *testing.T) {

	fs := NewFileStorageWithParams("testdata/filestorage", FileStorageModeWrite|FileStorageModeFormatJson, "utf-8")

	fs.StartWriteStruct("gocv", FileNodeTypeMap, "person")
	fs.ElName()
	fs.State()
	fs.GetFormat()
	fs.IsOpened()

	m := NewMat()
	defer m.Close()
	fs.WriteMat("mat", m)

	fs.WriteString("string", "string value")
	fs.WriteStringArray("stringArray", []string{"string", "array"})
	fs.WriteDouble("double", 3.1415927)
	fs.WriteInt("int", 42)
	fs.WriteComment("no comments", true)

	fs.EndWriteStruct()

	fs.StartWriteStruct("gocv2", FileNodeTypeSeq, "int")
	fs.WriteRaw("u", []byte{0, 0})
	fs.EndWriteStruct()

	fs.GetNode("gocv")
	fs.Root(0)

	fs.ReleaseAndGetString()

	fs = NewFileStorage()
	fs.Open("testdata/filestorage", FileStorageModeRead, "utf-8")

	fn := fs.GetFirstTopLevelNode()
	defer fn.Close()

	fn.Empty()
	fn.IsInt()
	fn.IsMap()
	fn.IsNamed()
	fn.IsNone()
	fn.IsReal()
	fn.IsSeq()
	fn.IsString()
	fn.Keys()
	fs.Release()
}
