package gocv

import (
	"image"
	"image/color"
	"math"
	"testing"
)

func TestCvtColor(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in CvtColor test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img, dest, ColorBGRAToGray)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid convert in CvtColor test")
	}
}

func TestBilateralFilter(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in BilateralFilter test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	BilateralFilter(img, dest, 1, 2.0, 3.0)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid BilateralFilter test")
	}
}

func TestBlur(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in GaussianBlur test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Blur(img, dest, image.Pt(3, 3))
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Blur test")
	}
}

func TestDilate(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Dilate test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))

	Dilate(img, dest, kernel)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Dilate test")
	}
}

func TestMoments(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in Moments test")
	}
	defer img.Close()

	result := Moments(img, true)
	if len(result) < 1 {
		t.Errorf("Invalid Moments test: %v", result)
	}
}

func TestFindContours(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	res := FindContours(img, RetrievalExternal, ChainApproxSimple)
	if len(res) < 1 {
		t.Error("Invalid FindContours test")
	}

	area := ContourArea(res[0])
	if area != 127280.0 {
		t.Errorf("Invalid ContourArea test: %f", area)
	}

	r := BoundingRect(res[0])
	if r.Min.X != 0 || r.Max.Y != 320 {
		t.Errorf("Invalid BoundingRect test: %v", r)
	}
}

func TestErode(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Erode test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))

	Erode(img, dest, kernel)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Erode test")
	}
}

func TestMorphologyEx(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in MorphologyEx test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))

	MorphologyEx(img, dest, MorphOpen, kernel)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid MorphologyEx test")
	}
}

func TestGaussianBlur(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in GaussianBlur test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	GaussianBlur(img, dest, image.Pt(23, 23), 30, 50, 4)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Blur test")
	}
}

func TestLaplacian(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Laplacian test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Laplacian(img, dest, MatTypeCV16S, 1, 1, 0, BorderDefault)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Laplacian test")
	}
}

func TestScharr(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Scharr test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Scharr(img, dest, MatTypeCV16S, 1, 0, 0, 0, BorderDefault)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Scharr test")
	}
}

func TestMedianBlur(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in MedianBlur test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	MedianBlur(img, dest, 1)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid MedianBlur test")
	}
}

func TestCanny(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Canny(img, dest, 50, 150)
	if dest.Empty() {
		t.Error("Empty Canny test")
	}
	if img.Rows() != dest.Rows() {
		t.Error("Invalid Canny test rows")
	}
	if img.Cols() != dest.Cols() {
		t.Error("Invalid Canny test cols")
	}
}

func TestGoodFeaturesToTrackAndCornerSubPix(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in GoodFeaturesToTrack test")
	}
	defer img.Close()

	corners := NewMat()
	defer corners.Close()

	GoodFeaturesToTrack(img, corners, 500, 0.01, 10)
	if corners.Empty() {
		t.Error("Empty GoodFeaturesToTrack test")
	}
	if corners.Rows() != 205 {
		t.Errorf("Invalid GoodFeaturesToTrack test rows: %v", corners.Rows())
	}
	if corners.Cols() != 1 {
		t.Errorf("Invalid GoodFeaturesToTrack test cols: %v", corners.Cols())
	}

	tc := NewTermCriteria(Count|EPS, 20, 0.03)

	CornerSubPix(img, corners, image.Pt(10, 10), image.Pt(-1, -1), tc)
	if corners.Empty() {
		t.Error("Empty CornerSubPix test")
	}
	if corners.Rows() != 205 {
		t.Errorf("Invalid CornerSubPix test rows: %v", corners.Rows())
	}
	if corners.Cols() != 1 {
		t.Errorf("Invalid CornerSubPix test cols: %v", corners.Cols())
	}
}

func TestHoughCircles(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughCircles test")
	}
	defer img.Close()

	circles := NewMat()
	defer circles.Close()

	HoughCircles(img, circles, 3, 5.0, 5.0)
	if circles.Empty() {
		t.Error("Empty HoughCircles test")
	}
	if circles.Rows() != 1 {
		t.Errorf("Invalid HoughCircles test rows: %v", circles.Rows())
	}
	if circles.Cols() < 330 || circles.Cols() > 334 {
		t.Errorf("Invalid HoughCircles test cols: %v", circles.Cols())
	}
}

func TestHoughLines(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	HoughLines(img, dest, math.Pi/180, 1, 1)
	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}
	if dest.Rows() != 10411 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughLines test cols: %v", dest.Cols())
	}

	if dest.GetFloatAt(0, 0) != 0 && dest.GetFloatAt(0, 1) != 0 {
		t.Errorf("Invalid HoughLines first test element: %v, %v", dest.GetFloatAt(0, 0), dest.GetFloatAt(0, 1))
	}

	if dest.GetFloatAt(1, 0) != 0.99483764 && dest.GetFloatAt(1, 1) != 0 {
		t.Errorf("Invalid HoughLines second test element: %v, %v", dest.GetFloatAt(1, 0), dest.GetFloatAt(1, 1))
	}

	if dest.GetFloatAt(10409, 0) != -118.246056 && dest.GetFloatAt(10409, 1) != 2 {
		t.Errorf("Invalid HoughLines penultimate test element: %v, %v", dest.GetFloatAt(10409, 0), dest.GetFloatAt(10409, 1))
	}

	if dest.GetFloatAt(10410, 0) != -118.246056 && dest.GetFloatAt(10410, 1) != 2 {
		t.Errorf("Invalid HoughLines last test element: %v, %v", dest.GetFloatAt(10410, 0), dest.GetFloatAt(10410, 1))
	}
}

func TestHoughLinesP(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	HoughLinesP(img, dest, math.Pi/180, 1, 1)
	if dest.Empty() {
		t.Error("Empty HoughLinesP test")
	}
	if dest.Rows() != 435 {
		t.Errorf("Invalid HoughLinesP test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughLinesP test cols: %v", dest.Cols())
	}

	if dest.GetIntAt(0, 0) != 197 && dest.GetIntAt(0, 1) != 319 && dest.GetIntAt(0, 2) != 197 && dest.GetIntAt(0, 3) != 197 {
		t.Errorf("Invalid HoughLinesP first test element: %v, %v, %v, %v", dest.GetIntAt(0, 0), dest.GetIntAt(0, 1), dest.GetIntAt(0, 2), dest.GetIntAt(0, 3))
	}

	if dest.GetIntAt(1, 0) != 62 && dest.GetIntAt(1, 1) != 319 && dest.GetIntAt(1, 2) != 197 && dest.GetIntAt(1, 3) != 197 {
		t.Errorf("Invalid HoughLinesP second test element: %v, %v, %v, %v", dest.GetIntAt(1, 0), dest.GetIntAt(1, 1), dest.GetIntAt(1, 2), dest.GetIntAt(1, 3))
	}

	if dest.GetIntAt(433, 0) != 357 && dest.GetIntAt(433, 1) != 316 && dest.GetIntAt(433, 2) != 357 && dest.GetIntAt(433, 3) != 316 {
		t.Errorf("Invalid HoughLinesP penultimate test element: %v, %v, %v, %v", dest.GetIntAt(433, 0), dest.GetIntAt(433, 1), dest.GetIntAt(433, 2), dest.GetIntAt(433, 3))
	}

	if dest.GetIntAt(434, 0) != 43 && dest.GetIntAt(434, 1) != 316 && dest.GetIntAt(434, 2) != 43 && dest.GetIntAt(434, 3) != 316 {
		t.Errorf("Invalid HoughLinesP last test element: %v, %v, %v, %v", dest.GetIntAt(434, 0), dest.GetIntAt(434, 1), dest.GetIntAt(434, 2), dest.GetIntAt(434, 3))
	}
}

func TestThreshold(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Erode test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Threshold(img, dest, 25, 255, ThresholdBinary)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Threshold test")
	}
}

func TestDrawing(t *testing.T) {
	img := NewMatWithSize(150, 150, MatTypeCV8U)
	if img.Empty() {
		t.Error("Invalid Mat in Rectangle")
	}
	defer img.Close()

	ArrowedLine(img, image.Pt(50, 50), image.Pt(75, 75), color.RGBA{0, 0, 255, 0}, 3)
	Circle(img, image.Pt(60, 60), 20, color.RGBA{0, 0, 255, 0}, 3)
	Rectangle(img, image.Rect(50, 50, 75, 75), color.RGBA{0, 0, 255, 0}, 3)
	Line(img, image.Pt(50, 50), image.Pt(75, 75), color.RGBA{0, 0, 255, 0}, 3)

	if img.Empty() {
		t.Error("Error in Rectangle test")
	}
}

func TestGetTextSize(t *testing.T) {
	size := GetTextSize("test", FontHersheySimplex, 1.2, 1)
	if size.X != 72 {
		t.Error("Invalid text size width")
	}

	if size.Y != 26 {
		t.Error("Invalid text size height")
	}
}
func TestPutText(t *testing.T) {
	img := NewMatWithSize(150, 150, MatTypeCV8U)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
	defer img.Close()

	pt := image.Pt(10, 10)
	PutText(img, "Testing", pt, FontHersheyPlain, 1.2, color.RGBA{255, 255, 255, 0}, 2)

	if img.Empty() {
		t.Error("Error in PutText test")
	}
}

func TestResize(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Resize test")
	}
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Resize(src, dst, image.Point{}, 0.5, 0.5, InterpolationDefault)
	if dst.Cols() != 200 || dst.Rows() != 172 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}

	Resize(src, dst, image.Pt(440, 377), 0, 0, InterpolationCubic)
	if dst.Cols() != 440 || dst.Rows() != 377 {
		t.Errorf("Expected dst size of 440x377 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestGetRotationMatrix2D(t *testing.T) {
	type args struct {
		center image.Point
		angle  float64
		scale  float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "90",
			args: args{image.Point{0, 0}, 90.0, 1.0},
			want: [][]float64{
				[]float64{6.123233995736766e-17, 1, 0},
				[]float64{-1, 6.123233995736766e-17, 0},
			},
		},
		{
			name: "45",
			args: args{image.Point{0, 0}, 45.0, 1.0},
			want: [][]float64{
				[]float64{0.7071067811865476, 0.7071067811865475, 0},
				[]float64{-0.7071067811865475, 0.7071067811865476, 0},
			},
		},
		{
			name: "0",
			args: args{image.Point{0, 0}, 0.0, 1.0},
			want: [][]float64{
				[]float64{1, 0, 0},
				[]float64{-0, 1, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRotationMatrix2D(tt.args.center, tt.args.angle, tt.args.scale)
			for row := 0; row < got.Rows(); row++ {
				for col := 0; col < got.Cols(); col++ {
					if got.GetDoubleAt(row, col) != tt.want[row][col] {
						t.Errorf("GetRotationMatrix2D() = %v, want %v at row:%v col:%v", got.GetDoubleAt(row, col), tt.want[row][col], row, col)
					}
				}
			}
		})
	}
}

func TestWarpAffine(t *testing.T) {
	type args struct {
		testPoint image.Point
		angle     float64
		scale     float64
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "300:200/45.0:1.0",
			args: args{image.Point{300, 200}, 45.0, 1.0},
			want: -8.407664e+21,
		},
		{
			name: "100:400/30.0:2.0",
			args: args{image.Point{100, 400}, 30.0, 2.0},
			want: -9.377235e-28,
		},
		{
			name: "500:150/120.0:1.0",
			args: args{image.Point{500, 150}, 120.0, 1.0},
			want: -3.5142978e-30,
		},
	}
	src := IMRead("images/face.jpg", IMReadGrayScale)
	if src.Empty() {
		panic("Invalid read of Mat in WarpAffine test")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mat := src.Clone()
			rot := GetRotationMatrix2D(image.Point{320, 240}, tt.args.angle, tt.args.scale)
			WarpAffine(mat, mat, rot, image.Point{640, 480})
			got := mat.GetFloatAt(tt.args.testPoint.X, tt.args.testPoint.Y)
			if got != tt.want {
				t.Errorf("WarpAffine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarpAffineWithParams(t *testing.T) {
	type args struct {
		testPoint image.Point
		angle     float64
		scale     float64
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "300:200/45.0:1.0",
			args: args{image.Point{300, 200}, 45.0, 1.0},
			want: -8.407664e+21,
		},
		{
			name: "100:400/30.0:2.0",
			args: args{image.Point{100, 400}, 30.0, 2.0},
			want: -9.377235e-28,
		},
		{
			name: "500:150/120.0:1.0",
			args: args{image.Point{500, 150}, 120.0, 1.0},
			want: -3.5142978e-30,
		},
	}
	src := IMRead("images/face.jpg", IMReadGrayScale)
	if src.Empty() {
		panic("Invalid read of Mat in WarpAffine test")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mat := src.Clone()
			rot := GetRotationMatrix2D(image.Point{320, 240}, tt.args.angle, tt.args.scale)
			WarpAffineWithParams(mat, mat, rot, image.Point{640, 480}, InterpolationLinear, BorderConstant, color.RGBA{0, 0, 0, 0})
			got := mat.GetFloatAt(tt.args.testPoint.X, tt.args.testPoint.Y)
			if got != tt.want {
				t.Errorf("WarpAffineWithParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
