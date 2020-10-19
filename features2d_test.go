package gocv

import (
	"image/color"
	"testing"
)

func TestAKAZE(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in AKAZE test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	ak := NewAKAZE()
	defer ak.Close()

	kp := ak.Detect(img)
	if len(kp) < 512 {
		t.Errorf("Invalid KeyPoint array in AKAZE test: %d", len(kp))
	}

	mask := NewMat()
	defer mask.Close()

	kp2, desc := ak.DetectAndCompute(img, mask)
	defer desc.Close()
	if len(kp2) < 512 {
		t.Errorf("Invalid KeyPoint array in AKAZE DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in AKAZE DetectAndCompute")
	}
}

func TestAgastFeatureDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in AgastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	ad := NewAgastFeatureDetector()
	defer ad.Close()

	kp := ad.Detect(img)
	if len(kp) < 2800 {
		t.Errorf("Invalid KeyPoint array in AgastFeatureDetector test: %d", len(kp))
	}
}

func TestBRISK(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in BRISK test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	br := NewBRISK()
	defer br.Close()

	kp := br.Detect(img)
	if len(kp) < 513 {
		t.Errorf("Invalid KeyPoint array in BRISK Detect: %d", len(kp))
	}

	mask := NewMat()
	defer mask.Close()

	kp2, desc := br.DetectAndCompute(img, mask)
	defer desc.Close()
	if len(kp2) != 1105 {
		t.Errorf("Invalid KeyPoint array in BRISK DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in AKAZE DetectAndCompute")
	}
}

func TestFastFeatureDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in FastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	fd := NewFastFeatureDetector()
	defer fd.Close()

	kp := fd.Detect(img)
	if len(kp) < 2690 {
		t.Errorf("Invalid KeyPoint array in FastFeatureDetector test: %d", len(kp))
	}
}

func TestFastFeatureDetectorWithParams(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in FastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	fd := NewFastFeatureDetectorWithParams(10, true, FastFeatureDetectorType916)
	defer fd.Close()

	kp := fd.Detect(img)
	if len(kp) < 2690 {
		t.Errorf("Invalid KeyPoint array in FastFeatureDetector test: %d", len(kp))
	}
}

func TestGFTTDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in GFTTDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	gft := NewGFTTDetector()
	defer gft.Close()

	kp := gft.Detect(img)
	if len(kp) < 512 {
		t.Errorf("Invalid KeyPoint array in GFTTDetector test: %d", len(kp))
	}
}

func TestKAZE(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in KAZE test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	k := NewKAZE()
	defer k.Close()

	kp := k.Detect(img)
	if len(kp) < 512 {
		t.Errorf("Invalid KeyPoint array in KAZE test: %d", len(kp))
	}

	mask := NewMat()
	defer mask.Close()

	kp2, desc := k.DetectAndCompute(img, mask)
	defer desc.Close()
	if len(kp2) < 512 {
		t.Errorf("Invalid KeyPoint array in KAZE DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in KAZE DetectAndCompute")
	}
}

func TestMSER(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in MSER test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	mser := NewMSER()
	defer mser.Close()

	kp := mser.Detect(img)
	if len(kp) != 232 && len(kp) != 234 && len(kp) != 261 {
		t.Errorf("Invalid KeyPoint array in MSER test: %d", len(kp))
	}
}

func TestORB(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in AgastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	od := NewORB()
	defer od.Close()

	kp := od.Detect(img)
	if len(kp) != 500 {
		t.Errorf("Invalid KeyPoint array in ORB test: %d", len(kp))
	}

	mask := NewMat()
	defer mask.Close()

	kp2, desc := od.DetectAndCompute(img, mask)
	defer desc.Close()
	if len(kp2) != 500 {
		t.Errorf("Invalid KeyPoint array in ORB DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in ORB DetectAndCompute")
	}
}

func TestSimpleBlobDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in SimpleBlobDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	bd := NewSimpleBlobDetector()
	defer bd.Close()

	kp := bd.Detect(img)
	if len(kp) != 2 {
		t.Errorf("Invalid KeyPoint array in SimpleBlobDetector test: %d", len(kp))
	}
}

func TestSimpleBlobDetectorWithParams(t *testing.T) {
	img := IMRead("images/circles.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in SimpleBlobDetector test")
	}
	defer img.Close()

	params := NewSimpleBlobDetectorParams()
	params.SetMaxArea(27500.0)

	bdp := NewSimpleBlobDetectorWithParams(params)
	defer bdp.Close()

	kp := bdp.Detect(img)
	if len(kp) != 4 {
		t.Errorf("Invalid KeyPoint array in SimpleBlobDetector test: %d", len(kp))
	}
}

func TestSimpleBlobDetectorParams(t *testing.T) {
	float64EqualityThreshold := 1e-5
	knownBlobColor := 235
	knownFilterByArea := false
	knownFilterByCircularity := true
	knownFilterByColor := false
	knownFilterByConvexity := false
	knownFilterByInertia := false
	knownMaxArea := 20000.0
	knownMaxCircularity := 0.99
	knownMaxConvexity := 0.98
	knownMaxInertiaRatio := 0.97
	knownMaxThreshold := 233.0
	knownMinArea := 230.0
	knownMinCircularity := 0.9
	knownMinConvexity := 0.89
	knownMinDistBetweenBlobs := 15.5
	knownMinInertiaRatio := 0.88
	knownMinRepeatability := 5
	knownMinThreshold := 200.0
	knownThresholdStep := 2.0

	params := NewSimpleBlobDetectorParams()
	params.SetBlobColor(knownBlobColor)
	params.SetFilterByArea(knownFilterByArea)
	params.SetFilterByCircularity(knownFilterByCircularity)
	params.SetFilterByColor(knownFilterByColor)
	params.SetFilterByConvexity(knownFilterByConvexity)
	params.SetFilterByInertia(knownFilterByInertia)
	params.SetMaxArea(knownMaxArea)
	params.SetMaxCircularity(knownMaxCircularity)
	params.SetMaxConvexity(knownMaxConvexity)
	params.SetMaxInertiaRatio(knownMaxInertiaRatio)
	params.SetMaxThreshold(knownMaxThreshold)
	params.SetMinArea(knownMinArea)
	params.SetMinCircularity(knownMinCircularity)
	params.SetMinConvexity(knownMinConvexity)
	params.SetMinDistBetweenBlobs(knownMinDistBetweenBlobs)
	params.SetMinInertiaRatio(knownMinInertiaRatio)
	params.SetMinRepeatability(knownMinRepeatability)
	params.SetMinThreshold(knownMinThreshold)
	params.SetThresholdStep(knownThresholdStep)

	if params.GetBlobColor() != knownBlobColor {
		t.Error("BlobColor incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetFilterByArea() != knownFilterByArea {
		t.Error("FilterByArea incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetFilterByCircularity() != knownFilterByCircularity {
		t.Error("FilterByCircularity incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetFilterByColor() != knownFilterByColor {
		t.Error("FilterByColor incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetFilterByConvexity() != knownFilterByConvexity {
		t.Error("FilterByConvexity incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetFilterByInertia() != knownFilterByInertia {
		t.Error("FilterByInertia incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetMaxArea() != knownMaxArea {
		t.Error("MaxArea incorrect in SimpleBlobDetectorParams test")
	}

	diffMaxCircularity := params.GetMaxCircularity() - knownMaxCircularity
	if diffMaxCircularity > float64EqualityThreshold {
		t.Errorf("DiffMaxCircularity greater than float64EqualityThreshold in SimpleBlobDetectorParams test. Diff: %f", diffMaxCircularity)
	}

	diffMaxConvexity := params.GetMaxConvexity() - knownMaxConvexity
	if diffMaxConvexity > float64EqualityThreshold {
		t.Errorf("DiffMaxConvexity greater than float64EqualityThreshold in SimpleBlobDetectorParams test. Diff: %f", diffMaxConvexity)
	}

	diffMaxInertiaRatio := params.GetMaxInertiaRatio() - knownMaxInertiaRatio
	if diffMaxInertiaRatio > float64EqualityThreshold {
		t.Errorf("DiffMaxInertiaRatio greater than float64EqualityThreshold in SimpleBlobDetectorParams test. Diff: %f", diffMaxInertiaRatio)
	}

	if params.GetMaxThreshold() != knownMaxThreshold {
		t.Error("MaxThreshold incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetMinArea() != knownMinArea {
		t.Error("MinArea incorrect in SimpleBlobDetectorParams test")
	}

	diffMinCircularity := params.GetMinCircularity() - knownMinCircularity
	if diffMinCircularity > float64EqualityThreshold {
		t.Errorf("DiffMinCircularity greater than float64EqualityThreshold in SimpleBlobDetectorParams test. Diff %f", diffMinCircularity)
	}

	diffMinConvexity := params.GetMinConvexity() - knownMinConvexity
	if diffMinConvexity > float64EqualityThreshold {
		t.Errorf("DiffMinConvexity greater than float64EqualityThreshold in SimpleBlobDetectorParams test. Diff: %f", diffMinConvexity)
	}

	if params.GetMinDistBetweenBlobs() != knownMinDistBetweenBlobs {
		t.Error("MinDistBetweenBlobs incorrect in SimpleBlobDetectorParams test")
	}

	diffMinInertiaRatio := params.GetMinInertiaRatio() - knownMinInertiaRatio
	if diffMinInertiaRatio > float64EqualityThreshold {
		t.Errorf("DiffMinInertiaRatio greater than float64EqualityThreshold in SimpleBlobDetectorParams test. Diff: %f", diffMinInertiaRatio)
	}

	if params.GetMinRepeatability() != knownMinRepeatability {
		t.Error("MinRepeatability incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetMinThreshold() != knownMinThreshold {
		t.Error("MinThreshold incorrect in SimpleBlobDetectorParams test")
	}

	if params.GetThresholdStep() != knownThresholdStep {
		t.Error("ThresholdStep incorrect in SimpleBlobDetectorParams test")
	}
}

func TestBFMatcher(t *testing.T) {
	descriptorFile := "images/sift_descriptor.png"
	desc1 := IMRead(descriptorFile, IMReadGrayScale)
	if desc1.Empty() {
		t.Error("descriptor one is empty in BFMatcher test")
	}
	defer desc1.Close()

	desc2 := IMRead(descriptorFile, IMReadGrayScale)
	if desc2.Empty() {
		t.Error("descriptor two is empty in BFMatcher test")
	}
	defer desc2.Close()

	bf := NewBFMatcher()
	defer bf.Close()

	k := 2
	dMatches := bf.KnnMatch(desc1, desc2, k)
	if len(dMatches) < 1 {
		t.Errorf("DMatches was excepted to have at least one element")
	}
	for i := range dMatches {
		if len(dMatches[i]) != k {
			t.Errorf("Length does not match k cluster amount in BFMatcher")
		}
	}

	bfParams := NewBFMatcherWithParams(NormHamming, false)
	defer bfParams.Close()

	dMatches = bfParams.KnnMatch(desc1, desc2, k)
	if len(dMatches) < 1 {
		t.Errorf("DMatches was excepted to have at least one element")
	}
	for i := range dMatches {
		if len(dMatches[i]) != k {
			t.Errorf("Length does not match k cluster amount in BFMatcher")
		}
	}
}

func TestFlannBasedMatcher(t *testing.T) {
	descriptorFile := "images/sift_descriptor.png"
	desc1 := IMRead(descriptorFile, IMReadGrayScale)
	if desc1.Empty() {
		t.Error("descriptor one is empty in FlannBasedMatcher test")
	}
	defer desc1.Close()
	desc1.ConvertTo(&desc1, MatTypeCV32F)

	desc2 := IMRead(descriptorFile, IMReadGrayScale)
	if desc2.Empty() {
		t.Error("descriptor two is empty in FlannBasedMatcher test")
	}
	defer desc2.Close()
	desc2.ConvertTo(&desc2, MatTypeCV32F)

	f := NewFlannBasedMatcher()
	defer f.Close()

	k := 2
	dMatches := f.KnnMatch(desc1, desc2, k)
	if len(dMatches) < 1 {
		t.Errorf("DMatches was excepted to have at least one element")
	}
	for i := range dMatches {
		if len(dMatches[i]) != k {
			t.Errorf("Length does not match k cluster amount in FlannBasedMatcher")
		}
	}
}

func TestDrawKeyPoints(t *testing.T) {
	keypointsFile := "images/simple.jpg"
	img := IMRead(keypointsFile, IMReadColor)
	if img.Empty() {
		t.Error("keypoints file is empty in DrawKeyPoints test")
	}
	defer img.Close()

	ffd := NewFastFeatureDetector()
	kp := ffd.Detect(img)

	simpleKP := NewMat()
	defer simpleKP.Close()
	DrawKeyPoints(img, kp, &simpleKP, color.RGBA{255, 0, 0, 0}, DrawDefault)

	if simpleKP.Rows() != img.Rows() || simpleKP.Cols() != img.Cols() {
		t.Error("Invalid DrawKeyPoints test")
	}
}

func TestDrawMatches(t *testing.T) {
	queryFile := "images/box.png"
	trainFile := "images/box_in_scene.png"

	query := IMRead(queryFile, IMReadGrayScale)
	train := IMRead(trainFile, IMReadGrayScale)

	if query.Empty() || train.Empty() {
		t.Error("at least one of files is empty in DrawMatches test")
	}

	defer query.Close()
	defer train.Close()

	sift := NewSIFT()
	defer sift.Close()

	m1 := NewMat()
	m2 := NewMat()
	defer m1.Close()
	defer m2.Close()

	kp1, des1 := sift.DetectAndCompute(query, m1)
	kp2, des2 := sift.DetectAndCompute(train, m2)
	defer des1.Close()
	defer des2.Close()

	bf := NewBFMatcher()
	defer bf.Close()
	matches := bf.KnnMatch(des1, des2, 2)

	if len(matches) == 0 {
		t.Error("no matches found in DrawMatches test")
	}

	var good []DMatch
	for _, m := range matches {
		if len(m) > 1 {
			if m[0].Distance < 0.75*m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	c := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}

	mask := make([]byte, 0)

	out := NewMat()
	defer out.Close()

	DrawMatches(query, kp1, train, kp2, good, &out, c, c, mask, DrawDefault)

	if out.Cols() != (query.Cols()+train.Cols()) || out.Rows() < train.Rows() || out.Rows() < query.Rows() {
		t.Error("Invalid DrawMatches test")
	}

	mask = make([]byte, len(good))

	smoke := NewMat()
	defer smoke.Close()

	DrawMatches(query, kp1, train, kp2, good, &smoke, c, c, mask, DrawDefault)
}

func TestSIFT(t *testing.T) {
	img := IMRead("./images/face.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in SIFT test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	si := NewSIFT()
	defer si.Close()

	kp := si.Detect(img)
	if len(kp) == 512 {
		t.Errorf("Invalid KeyPoint array in SIFT test: %d", len(kp))
	}

	mask := NewMat()
	defer mask.Close()

	kp2, desc := si.DetectAndCompute(img, mask)
	defer desc.Close()
	if len(kp2) == 512 {
		t.Errorf("Invalid KeyPoint array in SIFT DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in SIFT DetectAndCompute")
	}
}
