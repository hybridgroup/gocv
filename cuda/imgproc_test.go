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

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	detector := NewCannyEdgeDetector(50, 100)
	defer detector.Close()

	dimg := detector.Detect(cimg)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

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

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	canny := NewCannyEdgeDetector(100, 200)
	defer canny.Close()

	mimg := canny.Detect(cimg)
	defer mimg.Close()

	detector := NewHoughLinesDetectorWithParams(1, math.Pi/180, 50, true, 4096)
	defer detector.Close()

	dimg := detector.Detect(mimg)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}

	if dest.Rows() != 2 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1588 {
		t.Errorf("Invalid HoughLines test cols: %v", dest.Cols())
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

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	canny := NewCannyEdgeDetector(100, 200)
	defer canny.Close()

	stream := NewStream()
	defer stream.Close()

	mimg := canny.DetectWithStream(cimg, stream)
	defer mimg.Close()

	detector := NewHoughLinesDetectorWithParams(1, math.Pi/180, 50, true, 4096)
	defer detector.Close()

	dimg := detector.DetectWithStream(mimg, stream)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}

	if dest.Rows() != 2 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1588 {
		t.Errorf("Invalid HoughLines test cols: %v", dest.Cols())
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

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	canny := NewCannyEdgeDetector(50, 100)
	defer canny.Close()

	mimg := canny.Detect(cimg)
	defer mimg.Close()

	detector := NewHoughSegmentDetector(1, math.Pi/180, 150, 50)
	defer detector.Close()

	dimg := detector.Detect(mimg)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughSegment test")
	}

	final := dest.Reshape(0, dest.Cols())
	defer final.Close()

	if final.Rows() != 5 {
		t.Errorf("Invalid HoughSegment test rows: %v", final.Rows())
	}
	if final.Cols() != 1 {
		t.Errorf("Invalid HoughSegment test cols: %v", final.Cols())
	}

	type point struct {
		X, Y int32
	}

	expected := map[point]point{
		{1, 21}:   {398, 21},
		{304, 21}: {10, 315},
	}

	actual := make(map[point]point)
	for i := 0; i < final.Rows(); i += 4 {
		actual[point{final.GetVeciAt(i, 0)[0], final.GetVeciAt(i, 0)[1]}] =
			point{final.GetVeciAt(i, 0)[2], final.GetVeciAt(i, 0)[3]}
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

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	canny := NewCannyEdgeDetector(50, 100)
	defer canny.Close()

	stream := NewStream()
	defer stream.Close()

	mimg := canny.DetectWithStream(cimg, stream)
	defer mimg.Close()

	detector := NewHoughSegmentDetector(1, math.Pi/180, 150, 50)
	defer detector.Close()

	dimg := detector.DetectWithStream(mimg, stream)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughSegment test")
	}

	final := dest.Reshape(0, dest.Cols())
	defer final.Close()

	if final.Rows() != 5 {
		t.Errorf("Invalid HoughSegment test rows: %v", final.Rows())
	}
	if final.Cols() != 1 {
		t.Errorf("Invalid HoughSegment test cols: %v", final.Cols())
	}

	type point struct {
		X, Y int32
	}

	expected := map[point]point{
		{1, 21}:   {398, 21},
		{304, 21}: {10, 315},
	}

	actual := make(map[point]point)
	for i := 0; i < final.Rows(); i += 4 {
		actual[point{final.GetVeciAt(i, 0)[0], final.GetVeciAt(i, 0)[1]}] =
			point{final.GetVeciAt(i, 0)[2], final.GetVeciAt(i, 0)[3]}
	}

	for k, v := range expected {
		verify.Values(t, fmt.Sprintf("%d %d", k.X, k.Y), actual[k], v)
	}
}
