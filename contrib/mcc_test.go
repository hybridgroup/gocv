//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_mcc)

package contrib

import (
	"math"
	"testing"

	"gocv.io/x/gocv"
)

const (
	macbethImage = "../images/macbeth.png"
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
	imgCopy := gocv.NewMat()
	img := gocv.IMRead(path, gocv.IMReadColor)
	img.CopyTo(&imgCopy)
	defer imgCopy.Close()

	if img.Empty() {
		t.Fatal("Invalid input: image is empty or could not be loaded. Check that ./images/macbeth.jpg exists and is a valid image.")
	}
	defer img.Close()

	detector := NewMccCCheckerDetector()
	defer detector.Close()

	res := detector.Process(img, MCC24)
	if !res {
		t.Error("Atleast one chart is expected to be detected got 0")
	}

	checkers := detector.GetListColorChecker()

	var EPS = 0.001

	for _, checker := range checkers {
		whitePatch := checker.GetColorCharts()[18*4 : 18*4+4]
		expected := make([]gocv.Point2f, 4)
		expected[0] = gocv.Point2f{X: 438.60522, Y: 465.06586}
		expected[1] = gocv.Point2f{X: 480.70596, Y: 461.58606}
		expected[2] = gocv.Point2f{X: 484.52277, Y: 502.17834}
		expected[3] = gocv.Point2f{X: 442.45264, Y: 505.96082}

		for idx, _ := range expected {

			if distPoint2f(whitePatch[idx], expected[idx]) > EPS {
				t.Errorf("White patch expected at %v got %v", expected[idx], whitePatch[idx])
			}
		}

		// Outputting for visual inspection
		cdraw := NewMccCCheckerDraw(checker, gocv.NewScalar(0, 250, 0, 255), 2)
		cdraw.Draw(img)
	}
	gocv.IMWrite("../images/macbeth-correct.png", img)
}

func distPoint2f(a, b gocv.Point2f) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}
