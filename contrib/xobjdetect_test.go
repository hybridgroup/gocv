package contrib

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestWBDetector(t *testing.T) {

	img := gocv.IMRead("../images/face.jpg", gocv.IMReadAnyColor)
	if img.Empty() {
		t.Error("xobjdetect: cannot read image")
	}

	det := NewWBDetector()
	defer det.Close()

	det.Train("../images/gocvlogo.jpg", "../images/gocvlogo.png")

	det.Detect(img)

	fs := gocv.NewFileStorageWithParams("../testdata/WBDetector.json", gocv.FileStorageModeWrite, "utf-8")
	defer fs.Close()

	fs.StartWriteStruct("gocv", gocv.FileNodeTypeSeq, "model")

	det.Write(fs)

	fs.EndWriteStruct()

	node := fs.GetFirstTopLevelNode()

	det.Read(node)

}
