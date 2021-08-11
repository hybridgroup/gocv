package contrib

import (
	"image"
	"image/color"
	"testing"

	"gocv.io/x/gocv"
)

func TestFreeTypeLoadFontData(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("../fonts/JetBrainsMono-Regular.ttf", 0)
}

func TestFreeTypeGetTextSize(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("../fonts/JetBrainsMono-Regular.ttf", 0)

	size, baseLine := ft.GetTextSize("test", 60, 2)

	if size.X != 140 {
		t.Error("Invalid text size width")
	}

	if size.Y != 46 {
		t.Error("Invalid text size height")
	}

	if baseLine != 1 {
		t.Errorf("invalid base. expected %d, actual %d", 1, baseLine)
	}
}

func TestFreeTypePutText(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("../fonts/JetBrainsMono-Regular.ttf", 0)

	img := gocv.NewMatWithSize(150, 500, gocv.MatTypeCV8UC3)
	if img.Empty() {
		t.Error("Invalid Mat")
	}
	defer img.Close()

	pt := image.Pt(80, 80)
	ft.PutText(&img, "Testing", pt, 60, color.RGBA{R: 255, G: 255, B: 255}, 2, 8, true)

	if img.Empty() {
		t.Error("Error in PutText test")
	}
}

func TestFreeTypeSetSplitNumber(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("../fonts/JetBrainsMono-Regular.ttf", 0)
	ft.SetSplitNumber(10)
}
