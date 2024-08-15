package cuda

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/pascaldekloe/goe/verify"
	"gocv.io/x/gocv"
)

func TestCanny_Detect(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in Canny test")
	}
	defer src.Close()

	cimg, dimg := NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	detector := NewCannyEdgeDetector(50, 100)
	defer detector.Close()

	cimg.Upload(src)
	detector.Detect(cimg, &dimg)
	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty Canny test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid Canny test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid Canny test cols")
	}
}

func TestHoughLines_Calc(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	canny := NewCannyEdgeDetector(100, 200)
	defer canny.Close()

	detector := NewHoughLinesDetectorWithParams(1, math.Pi/180, 50, true, 4096)
	defer detector.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.Upload(src)
	canny.Detect(cimg, &mimg)
	detector.Detect(mimg, &dimg)
	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}

	if dest.Rows() != 2 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}

	expected := map[float32]float32{
		21:  1.5707964,
		337: 0.034906585,
		85:  1.5707964,
		276: 0,
		329: 0.034906585,
	}

	actual := make(map[float32]float32)
	for i := 0; i < dest.Cols(); i += 2 {
		actual[dest.GetFloatAt(0, i)] = dest.GetFloatAt(0, i+1)
	}

	for k, v := range expected {
		s32 := strconv.FormatFloat(float64(k), 'f', -1, 32)
		verify.Values(t, s32, actual[k], v)
	}
}

func TestHoughLines_CalcWithStream(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	stream := NewStream()
	defer stream.Close()

	canny := NewCannyEdgeDetector(100, 200)
	defer canny.Close()

	detector := NewHoughLinesDetectorWithParams(1, math.Pi/180, 50, true, 4096)
	defer detector.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, stream)
	canny.DetectWithStream(cimg, &mimg, stream)
	detector.DetectWithStream(mimg, &dimg, stream)
	dimg.DownloadWithStream(&dest, stream)

	stream.WaitForCompletion()

	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}

	if dest.Rows() != 2 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}

	expected := map[float32]float32{
		21:  1.5707964,
		337: 0.034906585,
		85:  1.5707964,
		276: 0,
		329: 0.034906585,
	}

	actual := make(map[float32]float32)
	for i := 0; i < dest.Cols(); i += 2 {
		actual[dest.GetFloatAt(0, i)] = dest.GetFloatAt(0, i+1)
	}

	for k, v := range expected {
		s32 := strconv.FormatFloat(float64(k), 'f', -1, 32)
		verify.Values(t, s32, actual[k], v)
	}
}

func TestHoughSegment_Calc(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughSegment test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	canny := NewCannyEdgeDetector(50, 100)
	defer canny.Close()

	detector := NewHoughSegmentDetector(1, math.Pi/180, 150, 50)
	defer detector.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.Upload(src)
	canny.Detect(cimg, &mimg)
	detector.Detect(mimg, &dimg)
	fimg := dimg.Reshape(0, dimg.Cols())
	defer fimg.Close()
	fimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughSegment test")
	}

	if dest.Rows() != 5 {
		t.Errorf("Invalid HoughSegment test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughSegment test cols: %v", dest.Cols())
	}

	type point struct {
		X, Y int32
	}

	expected := map[point]point{
		{1, 21}:   {398, 21},
		{304, 21}: {10, 315},
	}

	actual := make(map[point]point)
	for i := 0; i < dest.Rows(); i += 4 {
		actual[point{dest.GetVeciAt(i, 0)[0], dest.GetVeciAt(i, 0)[1]}] =
			point{dest.GetVeciAt(i, 0)[2], dest.GetVeciAt(i, 0)[3]}
	}

	for k, v := range expected {
		verify.Values(t, fmt.Sprintf("%d %d", k.X, k.Y), actual[k], v)
	}
}

func TestHoughSegment_CalcWithStream(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughSegment test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	stream := NewStream()
	defer stream.Close()

	canny := NewCannyEdgeDetector(50, 100)
	defer canny.Close()

	detector := NewHoughSegmentDetector(1, math.Pi/180, 150, 50)
	defer detector.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, stream)
	canny.DetectWithStream(cimg, &mimg, stream)
	detector.DetectWithStream(mimg, &dimg, stream)
	fimg := dimg.Reshape(0, dimg.Cols())
	defer fimg.Close()
	fimg.DownloadWithStream(&dest, stream)

	stream.WaitForCompletion()

	if dest.Empty() {
		t.Error("Empty HoughSegment test")
	}

	if dest.Rows() != 5 {
		t.Errorf("Invalid HoughSegment test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughSegment test cols: %v", dest.Cols())
	}

	type point struct {
		X, Y int32
	}

	expected := map[point]point{
		{1, 21}:   {398, 21},
		{304, 21}: {10, 315},
	}

	actual := make(map[point]point)
	for i := 0; i < dest.Rows(); i += 4 {
		actual[point{dest.GetVeciAt(i, 0)[0], dest.GetVeciAt(i, 0)[1]}] =
			point{dest.GetVeciAt(i, 0)[2], dest.GetVeciAt(i, 0)[3]}
	}

	for k, v := range expected {
		verify.Values(t, fmt.Sprintf("%d %d", k.X, k.Y), actual[k], v)
	}
}

func TestTemplateMatching_Match(t *testing.T) {
	imgScene := gocv.IMRead("../images/face.jpg", gocv.IMReadGrayScale)
	if imgScene.Empty() {
		t.Error("Invalid read of face.jpg in MatchTemplate test")
	}
	defer imgScene.Close()

	imgTemplate := gocv.IMRead("../images/toy.jpg", gocv.IMReadGrayScale)
	if imgTemplate.Empty() {
		t.Error("Invalid read of toy.jpg in MatchTemplate test")
	}
	defer imgTemplate.Close()

	cimg, timg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer timg.Close()
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	matcher := NewTemplateMatching(gocv.MatTypeCV8U, gocv.TmSqdiff)
	defer matcher.Close()

	cimg.Upload(imgScene)
	timg.Upload(imgTemplate)
	matcher.Match(cimg, timg, &dimg)
	dimg.Download(&dest)

	_, maxConfidence, _, _ := gocv.MinMaxLoc(dest)
	if maxConfidence < 0.95 {
		t.Errorf("Max confidence of %f is too low. MatchTemplate could not find template in scene.", maxConfidence)
	}
}

func TestDemosaicing(t *testing.T) {
	img := gocv.IMRead("../images/face.jpg", gocv.IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Demosaicing test")
	}
	defer img.Close()

	patterns := map[string]gocv.ColorConversionCode{
		"bg": gocv.ColorBayerBGToBGR,
		"gb": gocv.ColorBayerGBToBGR,
		"rg": gocv.ColorBayerRGToBGR,
		"gr": gocv.ColorBayerGRToBGR,
	}

	for pattern, code := range patterns {
		bayerImg, err := NewBayerFromMat(img, pattern)
		if bayerImg.Empty() {
			t.Error("Invalid conversion from Mat to Bayer in Demosaicing test")
		}
		if err != nil {
			t.Error(err)
		}

		cimg, dimg := NewGpuMat(), NewGpuMat()
		cimg.Upload(bayerImg)

		dest := gocv.NewMat()

		Demosaicing(cimg, &dimg, code)

		dimg.Download(&dest)
		if dest.Empty() || bayerImg.Rows() != dest.Rows() || bayerImg.Cols() != dest.Cols() {
			t.Error("Invalid convert in Demosaicing test")
		}

		bayerImg.Close()
		cimg.Close()
		dimg.Close()
		dest.Close()
	}
}

func TestDemosaicing_WithStream(t *testing.T) {
	img := gocv.IMRead("../images/face.jpg", gocv.IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Demosaicing test")
	}
	defer img.Close()

	patterns := map[string]gocv.ColorConversionCode{
		"bg": gocv.ColorBayerBGToBGR,
		"gb": gocv.ColorBayerGBToBGR,
		"rg": gocv.ColorBayerRGToBGR,
		"gr": gocv.ColorBayerGRToBGR,
	}

	for pattern, code := range patterns {
		stream := NewStream()

		bayerImg, err := NewBayerFromMat(img, pattern)
		if bayerImg.Empty() {
			t.Error("Invalid conversion from Mat to Bayer in DemosaicingWithStream test")
		}
		if err != nil {
			t.Error(err)
		}

		cimg, dimg := NewGpuMat(), NewGpuMat()
		dest := gocv.NewMat()

		cimg.UploadWithStream(bayerImg, stream)
		DemosaicingWithStream(cimg, &dimg, code, stream)
		dimg.DownloadWithStream(&dest, stream)

		stream.WaitForCompletion()

		if dest.Empty() || bayerImg.Rows() != dest.Rows() || bayerImg.Cols() != dest.Cols() {
			t.Error("Invalid convert in DemosaicingWithStream test")
		}

		bayerImg.Close()
		cimg.Close()
		dimg.Close()
		dest.Close()
		stream.Close()
	}
}

func NewBayerFromMat(src gocv.Mat, pattern string) (gocv.Mat, error) {
	dest := gocv.NewMatWithSize(src.Rows(), src.Cols(), gocv.MatTypeCV8UC1)

	switch pattern {
	case "bg":
		for y := 0; y < src.Rows(); y++ {
			for x := 0; x < src.Cols(); x++ {
				if (x+y)%2 != 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[1])
				} else if (x % 2) != 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[0])
				} else {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[2])
				}
			}
		}
	case "gb":
		for y := 0; y < src.Rows(); y++ {
			for x := 0; x < src.Cols(); x++ {
				if (x+y)%2 == 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[1])
				} else if (x % 2) == 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[0])
				} else {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[2])
				}
			}
		}
	case "rg":
		for y := 0; y < src.Rows(); y++ {
			for x := 0; x < src.Cols(); x++ {
				if (x+y)%2 != 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[1])
				} else if (x % 2) == 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[0])
				} else {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[2])
				}
			}
		}
	case "gr":
		for y := 0; y < src.Rows(); y++ {
			for x := 0; x < src.Cols(); x++ {
				if (x+y)%2 == 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[1])
				} else if (x % 2) != 0 {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[0])
				} else {
					dest.SetUCharAt(y, x, src.GetVecbAt(y, x)[2])
				}
			}
		}
	default:
		dest.Close()
		return gocv.Mat{}, fmt.Errorf("invalid pattern: %s", pattern)
	}

	return dest, nil
}
