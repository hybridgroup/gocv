//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_mcc)

package gocv

import (
	"reflect"
	"testing"
)

const (
	macbethImage = "./images/macbeth.jpg"
)

func TestMccDetectorParams(t *testing.T) {

	adaptiveThreshWinSizeMin := 23
	adaptiveThreshWinSizeMax := 153
	adaptiveThreshWinSizeStep := 16
	adaptiveThreshConstant := 7.0
	borderWidth := 0
	minContourLengthAllowed := 100
	minContourPointsAllowed := 4
	minImageSize := 1000
	minInterCheckerDistance := 10000
	minInterContourDistance := 100
	confidenceThreshold := 0.1
	findCandidatesApproxPolyDPEpsMultiplier := 0.05
	minContourSolidity := 0.9
	minContoursArea := 100.0
	minContoursAreaRate := 0.003
	B0factor := float32(1.25)
	maxError := float32(0.1)
	minGroupSize := uint(4)

	params := NewMccDetectorParameters()
	params.SetAdaptiveThreshWinSizeMin(adaptiveThreshWinSizeMin)
	params.SetAdaptiveThreshWinSizeMax(adaptiveThreshWinSizeMax)
	params.SetAdaptiveThreshWinSizeStep(adaptiveThreshWinSizeStep)
	params.SetAdaptiveThreshConstant(adaptiveThreshConstant)
	params.SetBorderWidth(borderWidth)
	params.SetMinContourLengthAllowed(minContourLengthAllowed)
	params.SetMinContourPointsAllowed(minContourPointsAllowed)
	params.SetMinImageSize(minImageSize)
	params.SetMinInterCheckerDistance(minInterCheckerDistance)
	params.SetMinInterContourDistance(minInterContourDistance)
	params.SetAdaptiveThreshConstant(adaptiveThreshConstant)
	params.SetConfidenceThreshold(confidenceThreshold)
	params.SetFindCandidatesApproxPolyDPEpsMultiplier(findCandidatesApproxPolyDPEpsMultiplier)
	params.SetMinContourSolidity(minContourSolidity)
	params.SetMinContoursArea(minContoursArea)
	params.SetMinContoursAreaRate(minContoursAreaRate)
	params.SetB0factor(B0factor)
	params.SetMaxError(maxError)
	params.SetMinGroupSize(minGroupSize)

	if params.GetAdaptiveThreshWinSizeMin() != adaptiveThreshWinSizeMin {
		t.Errorf("AdaptiveThreshWinSizeMin expected %v got %v", adaptiveThreshWinSizeMin, params.GetAdaptiveThreshWinSizeMin())
	}
	if params.GetAdaptiveThreshWinSizeMax() != adaptiveThreshWinSizeMax {
		t.Errorf("AdaptiveThreshWinSizeMax expected %v got %v", adaptiveThreshWinSizeMax, params.GetAdaptiveThreshWinSizeMax())
	}
	if params.GetAdaptiveThreshWinSizeStep() != adaptiveThreshWinSizeStep {
		t.Errorf("AdaptiveThreshWinSizeStep expected %v got %v", adaptiveThreshWinSizeStep, params.GetAdaptiveThreshWinSizeStep())
	}
	if params.GetAdaptiveThreshConstant() != adaptiveThreshConstant {
		t.Errorf("AdaptiveThreshConstant expected %v got %v", adaptiveThreshConstant, params.GetAdaptiveThreshConstant())
	}
	if params.GetBorderWidth() != borderWidth {
		t.Errorf("BorderWidth expected %v got %v", borderWidth, params.GetBorderWidth())
	}
	if params.GetMinContourLengthAllowed() != minContourLengthAllowed {
		t.Errorf("MinContourLengthAllowed expected %v got %v", minContourLengthAllowed, params.GetMinContourLengthAllowed())
	}
	if params.GetMinContourPointsAllowed() != minContourPointsAllowed {
		t.Errorf("MinContourPointsAllowed expected %v got %v", minContourPointsAllowed, params.GetMinContourPointsAllowed())
	}
	if params.GetMinImageSize() != minImageSize {
		t.Errorf("MinImageSize expected %v got %v", minImageSize, params.GetMinImageSize())
	}
	if params.GetMinInterCheckerDistance() != minInterCheckerDistance {
		t.Errorf("MinInterCheckerDistance expected %v got %v", minInterCheckerDistance, params.GetMinInterCheckerDistance())
	}
	if params.GetMinInterContourDistance() != minInterContourDistance {
		t.Errorf("MinInterContourDistance expected %v got %v", minInterContourDistance, params.GetMinInterContourDistance())
	}
	if params.GetAdaptiveThreshConstant() != adaptiveThreshConstant {
		t.Errorf("AdaptiveThreshConstant expected %v got %v", adaptiveThreshConstant, params.GetAdaptiveThreshConstant())
	}
	if params.GetConfidenceThreshold() != confidenceThreshold {
		t.Errorf("ConfidenceThreshold expected %v got %v", confidenceThreshold, params.GetConfidenceThreshold())
	}
	if params.GetFindCandidatesApproxPolyDPEpsMultiplier() != findCandidatesApproxPolyDPEpsMultiplier {
		t.Errorf("FindCandidatesApproxPolyDPEpsMultiplier expected %v got %v", findCandidatesApproxPolyDPEpsMultiplier, params.GetFindCandidatesApproxPolyDPEpsMultiplier())
	}
	if params.GetMinContourSolidity() != minContourSolidity {
		t.Errorf("MinContourSolidity expected %v got %v", minContourSolidity, params.GetMinContourSolidity())
	}
	if params.GetMinContoursArea() != minContoursArea {
		t.Errorf("MinContoursArea expected %v got %v", minContoursArea, params.GetMinContoursArea())
	}
	if params.GetMinContoursAreaRate() != minContoursAreaRate {
		t.Errorf("MinContoursAreaRate expected %v got %v", minContoursAreaRate, params.GetMinContoursAreaRate())
	}
	if params.GetB0factor() != B0factor {
		t.Errorf("B0factor expected %v got %v", B0factor, params.GetB0factor())
	}
	if params.GetMaxError() != maxError {
		t.Errorf("MaxError expected %v got %v", maxError, params.GetMaxError())
	}
	if params.GetMinGroupSize() != minGroupSize {
		t.Errorf("MinGroupSize expected %v got %v", minGroupSize, params.GetMinGroupSize())
	}
}

func TestProcess(t *testing.T) {
	path := macbethImage
	imgCopy := NewMat()
	img := IMRead(path, IMReadColor)
	img.CopyTo(&imgCopy)

	if img.Empty() {
		t.Fatal("Invalid input: image is empty or could not be loaded. Check that ./images/macbeth.jpg exists and is a valid image.")
	}
	defer func() {
		img.Close()
	}()

	detector := NewMccCCheckerDetector()
	defer func() {
		detector.Close()
	}()

	res := detector.Process(img, MCC24)
	if !res {
		t.Error("Atleast one chart is expected to be detected got 0")
	}

	checkers := detector.GetListColorChecker()

	for _, checker := range checkers {
		whitePatch := checker.GetColorCharts()[18]
		expected := Point2f{1220.626, 339.86987}
		if !reflect.DeepEqual(whitePatch, expected) {
			t.Errorf("White patch expected at %v got %v", expected, whitePatch)
		}
		// Outputting for visual inspection
		cdraw := NewMccCCheckerDraw(checker, NewScalar(0, 250, 0, 255), 2)
		cdraw.Draw(img)
	}
	IMWrite("./images/macbeth-correct.jpg", img)
}
