package contrib

//:testing

import (
	"gocv.io/x/gocv"
	"testing"
)

func TestBm3dDenoisingStepWithParams(t *testing.T) {
	// src = Input 8-bit or 16-bit 1-channel image.
	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC1)
	defer src.Close()
	dst1 := gocv.NewMat()
	defer dst1.Close()
	dst2 := gocv.NewMat()
	defer dst2.Close()

	Bm3dDenoisingStepWithParams(src, &dst1, &dst2,
		float32(1.0), int(4),
		int(16), int(2500),
		int(400), int(8),
		int(1), float32(2.0),
		gocv.NormL2,
		Bm3dAlgoStepAll,
		Bm3dTypeHaar)

	if src.Empty() || dst1.Rows() != src.Rows() || dst1.Cols() != src.Cols() {
		t.Error("Invlalid TestBm3dDenoisingStepWithParams test")
	}
}

func TestBm3dDenoisingWithParams(t *testing.T) {

	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC1)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	Bm3dDenoisingWithParams(src, &dst,
		float32(1.0), int(4),
		int(16), int(2500),
		int(400), int(8),
		int(1), float32(2.0),
		gocv.NormL2,
		Bm3dAlgoStepAll,
		Bm3dTypeHaar)

	if src.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() || dst.Type() != src.Type() {
		t.Error("Invlalid BalanceWhite test")
	}
}

func TestSetSaturationThreshold(t *testing.T) {

	grayworldwb := NewGrayworldWB()
	var saturation float32 = 0.7
	grayworldwb.SetSaturationThreshold(saturation)

	if grayworldwb.GetSaturationThreshold() < 0 {
		t.Error(" Invlalid SetSaturationThreshold test")
	}
}

func TestBalanceWhite(t *testing.T) {
	grayworldwb := NewGrayworldWB()
	defer grayworldwb.Close()

	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC3)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	grayworldwb.BalanceWhite(src, &dst)
	if src.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid BalanceWhite test")
	}
}

func TestNewLearningBasedWB(t *testing.T) {

	learningbasedwb := NewLearningBasedWB()
	defer learningbasedwb.Close()

	valueset := 2
	learningbasedwb.SetHistBinNum(valueset)
	learningbasedwb.SetRangeMaxVal(valueset)
	learningbasedwb.SetSaturationThreshold(float32(valueset))

	valuehistbinNum := learningbasedwb.GetHistBinNum()
	valuerangemaxval := learningbasedwb.GetRangeMaxVal()
	valuesaturation := learningbasedwb.GetSaturationThreshold()

	if valueset != valuehistbinNum {
		t.Error("Invalid TestNewLearningBasedWB : Set/Get HistBinNum test")
	}

	if valueset != valuerangemaxval {
		t.Error("Invalid TestNewLearningBasedWB : Set/Get RangeMaxVal test")
	}

	if valuesaturation < 0 {
		t.Error("Invalid TestNewLearningBasedWB : Set/Get SaturationThreshold test")
	}
}

func TestNewSimpleWB(t *testing.T) {

	simplewb := NewSimpleWB()
	defer simplewb.Close()

}

func TestNewTonemapDurand(t *testing.T) {

	tonemapdurand := NewTonemapDurand()
	defer tonemapdurand.Close()

	var valueset float32 = 2.05

	tonemapdurand.SetContrast(valueset)
	tonemapdurand.SetSaturation(valueset)
	tonemapdurand.SetSigmaColor(valueset)
	tonemapdurand.SetSigmaSpace(valueset)

	valuecontrast := tonemapdurand.GetContrast()
	valuestaturation := tonemapdurand.GetSaturation()
	valuesigmacolor := tonemapdurand.GetSigmaColor()
	valuesigmaspace := tonemapdurand.GetSigmaSpace()

	if valueset != valuecontrast {
		t.Error("Invalid result TestNewTonemapDurand : Set/Get Contrast test")
	}
	if valueset != valuestaturation {
		t.Error("Invalid result TestNewTonemapDurand : Set/Get Saturation test")
	}
	if valueset != valuesigmacolor {
		t.Error("Invalid result TestNewTonemapDurand : Set/Get SigmaColor test")
	}
	if valueset != valuesigmaspace {
		t.Error("Invalid result TestNewTonemapDurand : Set/Get SigmaSpace test")
	}

	tonemapdurand.SetGamma(valueset)
	valuegamma := tonemapdurand.GetGamma()

	if valueset != valuegamma {
		t.Error("Invalid result TestNewTonemapDurand : Set/Get Gamma test")
	}

}

func TestTonemapDurandProcess(t *testing.T) {
	tonemapdurand := NewTonemapDurand()
	defer tonemapdurand.Close()

	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV32FC3)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	tonemapdurand.Process(src, &dst)

	if src.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() || dst.Type() != gocv.MatTypeCV32FC3 {
		t.Error("Invlalid TestTonemapDurandProcess test")
	}
}

func TestOilPainting(t *testing.T) {

	jpgImageOilPainting := gocv.IMRead("../images/space_shuttle.jpg", gocv.IMReadColor)
	if jpgImageOilPainting.Empty() {
		t.Error("Invalid read of Source Mat in TestInpaint test")
	}
	defer jpgImageOilPainting.Close()

	t.Logf("Read of Source Mat in TestOilPainting test : %d x %d", jpgImageOilPainting.Cols(), jpgImageOilPainting.Rows())

	srcOilPainting := gocv.NewMat()
	defer srcOilPainting.Close()
	jpgImageOilPainting.ConvertTo(&srcOilPainting, gocv.MatTypeCV8UC3)

	dstOilPainting := gocv.NewMat()
	defer dstOilPainting.Close()

	OilPainting(srcOilPainting, &dstOilPainting, 2, 2)

	t.Logf("OilPainting : MAT %d <> %d : %d", dstOilPainting.Rows(), srcOilPainting.Rows(), dstOilPainting.Type())

	if srcOilPainting.Empty() || dstOilPainting.Rows() != srcOilPainting.Rows() || dstOilPainting.Cols() != srcOilPainting.Cols() {
		t.Error("Invlalid TestInpaint OilPainting test")
		return
	}

	gocv.IMWrite("testOilPainting.png", dstOilPainting)
}
