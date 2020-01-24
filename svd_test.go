package gocv

import (
	"testing"
)

func TestSVDCompute(t *testing.T) {
	var resultW = []float32{6.167493, 3.8214223}
	var resultU = []float32{-0.1346676, -0.99089086, 0.9908908, -0.1346676}
	var resultVt = []float32{0.01964448, 0.999807, -0.999807, 0.01964448}

	checkFunc := func(a []float32, b []float32) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	src := NewMatWithSize(2, 2, MatTypeCV32F)
	src.SetFloatAt(0, 0, 3.76956568)
	src.SetFloatAt(0, 1, -0.90478725)
	src.SetFloatAt(1, 0, 0.634576)
	src.SetFloatAt(1, 1, 6.10002347)
	defer src.Close()

	w := NewMat()
	defer w.Close()

	u := NewMat()
	defer u.Close()

	vt := NewMat()
	defer vt.Close()

	SVDCompute(src, &w, &u, &vt)

	dataW, err := w.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	if !checkFunc(resultW, dataW) {
		t.Error("w value is incorrect")
	}

	dataU, err := u.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	if !checkFunc(resultU, dataU) {
		t.Error("u value is incorrect")
	}

	dataVt, err := vt.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	if !checkFunc(resultVt, dataVt) {
		t.Error("vt value is incorrect")
	}
}
