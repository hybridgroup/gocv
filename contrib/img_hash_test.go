package contrib

import (
	"errors"
	"testing"

	"gocv.io/x/gocv"
)

const (
	testImage  = "../images/space_shuttle.jpg"
	testImage2 = "../images/toy.jpg"
)

func compute(path string, hash ImgHashBase) (*gocv.Mat, error) {
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		return nil, errors.New("Invalid input")
	}
	defer img.Close()

	dst := gocv.NewMat()
	hash.Compute(img, &dst)
	if dst.Empty() {
		dst.Close()
		return nil, errors.New("Empty output")
	}

	return &dst, nil
}

func testHash(t *testing.T, hash ImgHashBase) {
	result, err := compute(testImage, hash)
	if err != nil {
		t.Error(err)
	}
	defer result.Close()

	t.Logf("%T: %x", hash, result.ToBytes())

	// Load second image and make sure it doesn't compare as identical
	result2, err := compute(testImage2, hash)
	if err != nil {
		t.Error(err)
	}
	defer result2.Close()

	similar := hash.Compare(*result, *result2)
	t.Logf("%T: similarity %g", hash, similar)
	// The range and meaning of this value varies between algorithms, and
	// there doesn't seem to be a well defined set of default thresholds, so
	// "anything but zero" is the minimum smoke test.
	if similar == 0 {
		t.Error("Image similarity is zero?")
	}
}

func TestHashes(t *testing.T) {
	t.Run("PHash", func(t *testing.T) { testHash(t, PHash{}) })
	t.Run("AverageHash", func(t *testing.T) { testHash(t, AverageHash{}) })
	t.Run("BlockMeanHash", func(t *testing.T) { testHash(t, BlockMeanHash{}) })
	t.Run("ColorMomentHash", func(t *testing.T) { testHash(t, ColorMomentHash{}) })
	t.Run("MarrHidlrethHash", func(t *testing.T) { testHash(t, NewMarrHildrethHash()) })
	t.Run("RadialVarianceHash", func(t *testing.T) { testHash(t, NewRadialVarianceHash()) })
}

func BenchmarkCompute(b *testing.B) {
	img := gocv.IMRead(testImage, gocv.IMReadColor)
	if img.Empty() {
		b.Error("Invalid input")
	}
	defer img.Close()
	b.ResetTimer()

	compute := func(b *testing.B, hash ImgHashBase) {
		for i := 0; i < b.N; i++ {
			dst := gocv.NewMat()
			hash.Compute(img, &dst)
			if dst.Empty() {
				b.Error("Empty output")
				dst.Close()
				return
			}
			dst.Close()
		}
	}

	b.Run("PHash", func(b *testing.B) { compute(b, PHash{}) })
	b.Run("AverageHash", func(b *testing.B) { compute(b, AverageHash{}) })
	b.Run("BlockMeanHash", func(b *testing.B) { compute(b, BlockMeanHash{}) })
	b.Run("ColorMomentHash", func(b *testing.B) { compute(b, ColorMomentHash{}) })
	b.Run("MarrHidlrethHash", func(b *testing.B) { compute(b, NewMarrHildrethHash()) })
	b.Run("RadialVarianceHash", func(b *testing.B) { compute(b, NewRadialVarianceHash()) })
}

func BenchmarkCompare(b *testing.B) {
	compare := func(b *testing.B, hash ImgHashBase) {
		result1, err := compute(testImage, hash)
		if err != nil {
			b.Error(err)
		}
		defer result1.Close()

		result2, err := compute(testImage2, hash)
		if err != nil {
			b.Error(err)
		}
		defer result2.Close()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hash.Compare(*result1, *result2)
		}
	}

	b.Run("PHash", func(b *testing.B) { compare(b, PHash{}) })
	b.Run("AverageHash", func(b *testing.B) { compare(b, AverageHash{}) })
	b.Run("BlockMeanHash", func(b *testing.B) { compare(b, BlockMeanHash{}) })
	b.Run("ColorMomentHash", func(b *testing.B) { compare(b, ColorMomentHash{}) })
	b.Run("MarrHidlrethHash", func(b *testing.B) { compare(b, NewMarrHildrethHash()) })
	b.Run("RadialVarianceHash", func(b *testing.B) { compare(b, NewRadialVarianceHash()) })
}
