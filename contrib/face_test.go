package contrib

import (
	"image"
	"math"
	"testing"

	"gocv.io/x/gocv"
)

func TestLBPHFaceRecognizer_Methods(t *testing.T) {
	model := NewLBPHFaceRecognizer()
	if model == nil {
		t.Errorf("Invalid NewLBPHFaceRecognizer call %v", model)
	}

	model.Empty()

	labels := []int{1, 1, 1, 1, 2, 2, 2, 2}
	images := []gocv.Mat{
		gocv.IMRead("./att_faces/s1/1.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s1/2.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s1/3.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s1/4.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s2/1.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s2/2.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s2/3.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s2/4.pgm", gocv.IMReadGrayScale),
	}
	model.Train(images, labels)

	sample := gocv.IMRead("./att_faces/s2/5.pgm", gocv.IMReadGrayScale)
	label := model.Predict(sample)
	if label != 2 {
		t.Errorf("Invalid simple predict! label: %d", label)
	}
	resp := model.PredictExtendedResponse(sample)
	if resp.Label != 2 {
		t.Errorf("Invalid extended result predict! label: %d", resp.Label)
	}

	// set wrong threshold
	model.SetThreshold(0.0)
	label = model.Predict(sample)
	if label != -1 {
		t.Errorf("Invalid set wrong threshold! label: %d", label)
	}

	//// set good threshold
	model.SetThreshold(math.MaxFloat32)
	// set wrong radius
	model.SetRadius(0)
	label = model.Predict(sample)
	if label == 2 {
		t.Errorf("Invalid set wrong radius! label: %d", label)
	}

	neighbors := model.GetNeighbors()
	if neighbors == 0 {
		t.Errorf("Invalid get neighbors! n: %d", neighbors)
	}

	model.SetRadius(1)
	model.SetNeighbors(8)
	label = model.Predict(sample)
	if label != 2 {
		t.Errorf("Invalid set neighbors! label: %d", label)
	}

	// add new data
	sample = gocv.IMRead("./att_faces/s3/10.pgm", gocv.IMReadGrayScale)
	newLabels := []int{3, 3, 3, 3, 3, 3}
	newImages := []gocv.Mat{
		gocv.IMRead("./att_faces/s3/1.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s3/2.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s3/3.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s3/4.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s3/5.pgm", gocv.IMReadGrayScale),
		gocv.IMRead("./att_faces/s3/6.pgm", gocv.IMReadGrayScale),
	}
	model.Update(newImages, newLabels)
	label = model.Predict(sample)
	if label != 3 {
		t.Errorf("Invalid new data update: %d", label)
	}

	// test save and load
	fName := "data.yaml"
	model.SaveFile(fName)
	modelNew := NewLBPHFaceRecognizer()
	modelNew.LoadFile(fName)
	label = modelNew.Predict(sample)
	if label != 3 {
		t.Errorf("Invalid loaded data: %d", label)
	}
}

func TestLBPHFaceRecognizer_SetGridY_GetGridY(t *testing.T) {
	facer := NewLBPHFaceRecognizer()
	facer.SetGridY(5)

	if v := facer.GetGridY(); v != 5 {
		t.Errorf("got %d want 5", v)
	}
}

func TestLBPHFaceRecognizer_SetGrid_GetGrid(t *testing.T) {
	facer := NewLBPHFaceRecognizer()
	defer facer.Close()
	facer.SetGrid(image.Point{7, 5})

	if p := facer.GetGrid(); p.X != 7 || p.Y != 5 {
		t.Errorf("got %+v want {7, 5}", p)
	}
}
