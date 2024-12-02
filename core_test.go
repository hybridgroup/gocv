package gocv

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"runtime"
	"strings"
	"testing"
)

func TestMat(t *testing.T) {
	mat := NewMat()
	defer mat.Close()
	if !mat.Empty() {
		t.Error("New Mat should be empty")
	}
}

func TestMatClosed(t *testing.T) {
	mat := NewMat()
	mat.Close()
	if !mat.Closed() {
		t.Error("Closed Mat should be closed")
	}
}

func TestMatWithSizes(t *testing.T) {
	t.Run("create mat with multidimensional array", func(t *testing.T) {
		sizes := []int{100, 100, 100}
		mat := NewMatWithSizes(sizes, MatTypeCV8U)
		defer mat.Close()
		if mat.Empty() {
			t.Error("NewMatWithSizes should not be empty")
		}

		for i, val := range mat.Size() {
			if val != sizes[i] {
				t.Errorf("NewMatWithSizes incorrect size: %v\n", mat.Size())
			}
		}

		if mat.Rows() != -1 {
			t.Errorf("NewMatWithSizes incorrect row count: %v\n", mat.Rows())
		}

		if mat.Cols() != -1 {
			t.Errorf("NewMatWithSizes incorrect col count: %v\n", mat.Cols())
		}

		if mat.Channels() != 1 {
			t.Errorf("NewMatWithSizes incorrect channels count: %v\n", mat.Channels())
		}

		if mat.Type() != MatTypeCV8U {
			t.Errorf("NewMatWithSizes incorrect type: %v\n", mat.Type())
		}

		if mat.Total() != 100*100*100 {
			t.Errorf("NewMatWithSizes incorrect total: %v\n", mat.Total())
		}
	})

	t.Run("create 2x3x3 multidimensional array with 3 channels and scalar", func(t *testing.T) {
		sizes := []int{2, 3, 3}
		s := NewScalar(255.0, 105.0, 180.0, 0.0)
		mat := NewMatWithSizesWithScalar(sizes, MatTypeCV32FC3, s)
		defer mat.Close()
		if mat.Empty() {
			t.Error("NewMatWithSizesWithScalar should not be empty")
		}

		for i, val := range mat.Size() {
			if val != sizes[i] {
				t.Errorf("NewMatWithSizesWithScalar incorrect size: %v\n", mat.Size())
			}
		}

		if mat.Rows() != -1 {
			t.Errorf("NewMatWithSizesWithScalar incorrect row count: %v\n", mat.Rows())
		}

		if mat.Cols() != -1 {
			t.Errorf("NewMatWithSizesWithScalar incorrect col count: %v\n", mat.Cols())
		}

		if mat.Channels() != 3 {
			t.Errorf("NewMatWithSizesWithScalar incorrect channels count: %v\n", mat.Channels())
		}

		if mat.Type() != MatTypeCV32FC3 {
			t.Errorf("NewMatWithSizesWithScalar incorrect type: %v\n", mat.Type())
		}

		if mat.Total() != 2*3*3 {
			t.Errorf("NewMatWithSizesWithScalar incorrect total: %v\n", mat.Total())
		}

		matChans := Split(mat)
		scalar := []float32{255.0, 105.0, 180.0}
		for c := 0; c < mat.Channels(); c++ {
			for x := 0; x < sizes[0]; x++ {
				for y := 0; y < sizes[1]; y++ {
					for z := 0; z < sizes[2]; z++ {
						if s := matChans[c].GetFloatAt3(x, y, z); s != scalar[c] {
							t.Errorf("NewMatWithSizesWithScalar incorrect scalar: %v\n", s)
						}
					}
				}
			}
		}
		for _, ch := range matChans {
			ch.Close()
		}
	})

	t.Run("create 1x2x3 multidimensional array with 3 channel and data", func(t *testing.T) {
		sizes := []int{1, 2, 3}

		// generate byte array
		s := NewScalar(255.0, 123.0, 55.0, 0.0)
		mat1 := NewMatWithSizesWithScalar(sizes, MatTypeCV32FC2, s)
		defer mat1.Close()
		b := mat1.ToBytes()

		mat, err := NewMatWithSizesFromBytes(sizes, MatTypeCV32FC2, b)
		defer mat.Close()
		if err != nil {
			t.Errorf("NewMatWithSizesFromBytes %v\n", err)
		}
		if mat.Empty() {
			t.Error("NewMatWithSizesFromBytes should not be empty")
		}

		mat2, err := NewMatWithSizesFromBytes(sizes, MatTypeCV32FC2, nil)
		defer mat2.Close()
		if err == nil {
			t.Error("NewMatWithSizesFromBytes should return error with empty bytes")
		}

		b1 := mat.ToBytes()
		if !bytes.Equal(b, b1) {
			t.Error("NewMatWithSizesFromBytes byte arrays not equal")
		}

		for i, val := range mat.Size() {
			if val != sizes[i] {
				t.Errorf("NewMatWithSizesFromBytes incorrect size: %v\n", mat.Size())
			}
		}

		if mat.Rows() != -1 {
			t.Errorf("NewMatWithSizesFromBytes incorrect row count: %v\n", mat.Rows())
		}

		if mat.Cols() != -1 {
			t.Errorf("NewMatWithSizesFromBytes incorrect col count: %v\n", mat.Cols())
		}

		if mat.Channels() != 2 {
			t.Errorf("NewMatWithSizesFromBytes incorrect channels count: %v\n", mat.Channels())
		}

		if mat.Type() != MatTypeCV32FC2 {
			t.Errorf("NewMatWithSizesFromBytes incorrect type: %v\n", mat.Type())
		}

		if mat.Total() != 1*2*3 {
			t.Errorf("NewMatWithSizesFromBytes incorrect total: %v\n", mat.Total())
		}

		matChans := Split(mat)
		scalar := []float32{255.0, 123.0, 55.0}
		for c := 0; c < mat.Channels(); c++ {
			for x := 0; x < sizes[0]; x++ {
				for y := 0; y < sizes[1]; y++ {
					for z := 0; z < sizes[2]; z++ {
						if s := matChans[c].GetFloatAt3(x, y, z); s != scalar[c] {
							t.Errorf("NewMatWithSizesFromBytes incorrect value: %v\n", s)
						}
					}
				}
			}
		}
		for _, ch := range matChans {
			ch.Close()
		}
	})
}

func TestMatFromBytesWithEmptyByteSlice(t *testing.T) {
	_, err := NewMatFromBytes(600, 800, MatTypeCV8U, []byte{})
	if err == nil {
		t.Error("TestMatFromBytesWithEmptyByteSlise: " +
			"must fail because of an empty byte slice")
	}
	if !strings.Contains(err.Error(), ErrEmptyByteSlice.Error()) {
		t.Errorf("TestMatFromBytesWithEmptyByteSlice: "+
			"error must contain the following description: "+
			"%v, but have: %v", ErrEmptyByteSlice, err)
	}
}

func TestMatFromBytesSliceGarbageCollected(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	m, err := NewMatFromBytes(2, 5, MatTypeCV8U, data)
	if err != nil {
		t.Error("TestMatFromBytesSliceGarbageCollected: " +
			"failed to create Mat")
	}
	defer m.Close()

	// Force garbage collection. As data is not used after this, its backing array should
	// be collected.
	runtime.GC()

	v := m.GetUCharAt(0, 0)
	if v != 0 {
		t.Errorf("TestMatFromBytesSliceGarbageCollected: "+
			"unexpected value. Want %d, got %d.", 0, v)
	}
}

func TestMatWithSize(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat.Close()
	if mat.Empty() {
		t.Error("NewMatWithSize should not be empty")
	}

	if mat.Rows() != 101 {
		t.Errorf("NewMatWithSize incorrect row count: %v\n", mat.Rows())
	}

	if mat.Cols() != 102 {
		t.Errorf("NewMatWithSize incorrect col count: %v\n", mat.Cols())
	}

	if mat.Channels() != 1 {
		t.Errorf("NewMatWithSize incorrect channels count: %v\n", mat.Channels())
	}

	if mat.Type() != 0 {
		t.Errorf("NewMatWithSize incorrect type: %v\n", mat.Type())
	}

	if mat.Step() != 102 {
		t.Errorf("NewMatWithSize incorrect step count: %v\n", mat.Step())
	}
}

func TestMatWithSizeFromScalar(t *testing.T) {
	s := NewScalar(255.0, 105.0, 180.0, 0.0)
	mat := NewMatWithSizeFromScalar(s, 2, 3, MatTypeCV8UC3)
	defer mat.Close()
	if mat.Empty() {
		t.Error("NewMatWithSizeFromScalar should not be empty")
	}

	if mat.Rows() != 2 {
		t.Errorf("NewMatWithSizeFromScalar incorrect row count: %v\n", mat.Rows())
	}

	if mat.Cols() != 3 {
		t.Errorf("NewMatWithSizeFromScalar incorrect col count: %v\n", mat.Cols())
	}

	if mat.Channels() != 3 {
		t.Errorf("NewMatWithSizeFromScalar incorrect channels count: %v\n", mat.Channels())
	}

	if mat.Type() != 16 {
		t.Errorf("NewMatWithSizeFromScalar incorrect type: %v\n", mat.Type())
	}

	if mat.Total() != 6 {
		t.Errorf("incorrect total: %v\n", mat.Total())
	}

	if mat.Step() != 9 {
		t.Errorf("NewMatWithSizeFromScalar incorrect step count: %v\n", mat.Step())
	}

	sz := mat.Size()
	if sz[0] != 2 && sz[1] != 3 {
		t.Errorf("NewMatWithSize incorrect size: %v\n", sz)
	}

	matChans := Split(mat)
	scalarByte := []byte{255, 105, 180}
	for c := 0; c < mat.Channels(); c++ {
		for i := 0; i < mat.Rows(); i++ {
			for j := 0; j < mat.Cols(); j++ {
				if s := matChans[c].GetUCharAt(i, j); s != scalarByte[c] {
					t.Errorf("NewMatWithSizeFromScalar incorrect scalar: %v\n", s)
				}
			}
		}
	}
	for _, i := range matChans {
		i.Close()
	}
}

func TestMatFromPtr(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat.Close()
	pmat, _ := mat.FromPtr(11, 12, MatTypeCV8U, 10, 10)
	defer pmat.Close()

	if pmat.Rows() != 11 {
		t.Errorf("Mat copy incorrect row count: %v\n", pmat.Rows())
	}

	if pmat.Cols() != 12 {
		t.Errorf("Mat copy incorrect col count: %v\n", pmat.Cols())
	}
}

func TestMatClone(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat.Close()
	clone := mat.Clone()
	defer clone.Close()
	if clone.Rows() != 101 {
		t.Errorf("Mat clone incorrect row count: %v\n", clone.Rows())
	}

	if clone.Cols() != 102 {
		t.Errorf("Mat clone incorrect col count: %v\n", clone.Cols())
	}
}

func TestMatCopyTo(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat.Close()
	copy := NewMat()
	defer copy.Close()

	mat.CopyTo(&copy)
	if copy.Rows() != 101 {
		t.Errorf("Mat copy incorrect row count: %v\n", copy.Rows())
	}

	if copy.Cols() != 102 {
		t.Errorf("Mat copy incorrect col count: %v\n", copy.Cols())
	}
}

func TestMatCopyToWithMask(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat.Close()
	mask := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mask.Close()
	diff := NewMat()
	defer diff.Close()

	mat.SetUCharAt(0, 0, 255)
	mat.SetUCharAt(0, 1, 255)

	mask.SetUCharAt(0, 0, 255)

	copy := NewMat()
	defer copy.Close()

	mat.CopyToWithMask(&copy, mask)
	if copy.Rows() != 101 {
		t.Errorf("Mat copy incorrect row count: %v\n", copy.Rows())
	}

	if copy.Cols() != 102 {
		t.Errorf("Mat copy incorrect col count: %v\n", copy.Cols())
	}

	if copy.GetUCharAt(0, 0) != 255 || copy.GetUCharAt(0, 1) != 0 {
		t.Errorf("Mask failed to apply to source image")
	}

	Compare(mat, copy, &diff, CompareEQ)
	if CountNonZero(diff) == 0 {
		t.Errorf("Mat CopyToWithMask incorrect diff: %v\n", CountNonZero(diff))
	}
}

func TestMatToBytes(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	b := mat1.ToBytes()
	if len(b) != 101*102 {
		t.Errorf("Mat bytes incorrect length: %v\n", len(b))
	}

	copy, err := NewMatFromBytes(101, 102, MatTypeCV8U, b)
	if err != nil {
		t.Error(err.Error())
	}
	defer copy.Close()
	if copy.Rows() != 101 {
		t.Errorf("Mat from bytes incorrect row count: %v\n", copy.Rows())
	}
	if copy.Cols() != 102 {
		t.Errorf("Mat region incorrect col count: %v\n", copy.Cols())
	}

	mat2 := NewMatWithSize(101, 102, MatTypeCV16S)
	defer mat2.Close()
	b = mat2.ToBytes()
	if len(b) != 101*102*2 {
		t.Errorf("Mat bytes incorrect length: %v\n", len(b))
	}

	mat3 := NewMatFromScalar(NewScalar(255.0, 105.0, 180.0, 0.0), MatTypeCV8UC3)
	defer mat3.Close()
	b = mat3.ToBytes()
	if len(b) != 3 {
		t.Errorf("Mat bytes incorrect length: %v\n", len(b))
	}
	if bytes.Compare(b, []byte{255, 105, 180}) != 0 {
		t.Errorf("Mat bytes unexpected values: %v\n", b)
	}
}

func TestMatEye(t *testing.T) {
	data := []byte{1, 0, 0, 1}
	e := Eye(2, 2, MatTypeCV8U)

	if bytes.Compare(e.ToBytes(), data) != 0 {
		t.Errorf("Mat bytes are not equal")
	}
	e.Close()

	data2 := []byte{1, 0, 0, 0, 1, 0}
	e2 := Eye(2, 3, MatTypeCV8U)
	if bytes.Compare(e2.ToBytes(), data2) != 0 {
		t.Errorf("Mat bytes are not equal")
	}

	val := e2.GetUCharAt(0, 2)
	if val != 0 {
		t.Errorf("Mat bytes unexpected value at [0,2]: %v\n", val)
	}
	e2.Close()
}

func TestMatZeros(t *testing.T) {
	expected := NewMatWithSize(2, 3, MatTypeCV8U)
	z := Zeros(2, 3, MatTypeCV8U)

	if bytes.Compare(z.ToBytes(), expected.ToBytes()) != 0 {
		t.Errorf("Mat bytes are not equal")
	}

	expected.Close()
	z.Close()

	expected2 := NewMatWithSize(2, 3, MatTypeCV64F)
	z2 := Zeros(2, 3, MatTypeCV64F)

	if bytes.Compare(z2.ToBytes(), expected2.ToBytes()) != 0 {
		t.Errorf("Mat bytes are not equal")
	}
	expected2.Close()
	z2.Close()
}

func TestMatOnes(t *testing.T) {
	expected := NewMatWithSizeFromScalar(Scalar{Val1: 1}, 2, 3, MatTypeCV8U)
	o := Ones(2, 3, MatTypeCV8U)
	if bytes.Compare(o.ToBytes(), expected.ToBytes()) != 0 {
		t.Errorf("Mat bytes are not equal")
	}
	defer expected.Close()
	defer o.Close()

	expected2 := NewMatWithSizeFromScalar(Scalar{Val1: 1}, 2, 1, MatTypeCV64F)
	o2 := Ones(2, 1, MatTypeCV64F)

	if bytes.Compare(o2.ToBytes(), expected2.ToBytes()) != 0 {
		t.Errorf("Mat bytes are not equal")
	}
	expected2.Close()
	o2.Close()
}

func TestMatDataPtr(t *testing.T) {
	const (
		rows = 101
		cols = 102
	)
	t.Run("Uint8", func(t *testing.T) {
		testPoints := []struct {
			row int
			col int
			val uint8
		}{
			{row: 0, col: 0, val: 10},
			{row: 30, col: 31, val: 20},
			{row: rows - 1, col: cols - 1, val: 30},
		}

		mat1 := NewMatWithSize(rows, cols, MatTypeCV8U)
		defer mat1.Close()

		b, err := mat1.DataPtrUint8()
		if err != nil {
			t.Error(err)
		}
		if len(b) != 101*102 {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		for _, p := range testPoints {
			mat1.SetUCharAt(p.row, p.col, p.val)

			if got := b[p.row*cols+p.col]; got != p.val {
				t.Errorf("Expected %d,%d = %d, but it was %d", p.row, p.col, p.val, got)
			}
		}

		mat2 := NewMatWithSize(3, 9, MatTypeCV32F)
		defer mat2.Close()
		b, err = mat2.DataPtrUint8()
		if err != nil {
			t.Error(err)
		}
		if len(b) != 3*9*4 {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		mat3 := mat1.Region(image.Rect(25, 25, 75, 75))
		defer mat3.Close()
		_, err = mat3.DataPtrUint8()
		if err == nil {
			t.Errorf("Expected error.")
		}
	})
	t.Run("Int8", func(t *testing.T) {
		testPoints := []struct {
			row int
			col int
			val int8
		}{
			{row: 0, col: 0, val: 10},
			{row: 30, col: 31, val: 20},
			{row: rows - 1, col: cols - 1, val: 30},
		}

		mat1 := NewMatWithSize(101, 102, MatTypeCV8S)
		defer mat1.Close()

		b, err := mat1.DataPtrInt8()
		if err != nil {
			t.Error(err)
		}

		if len(b) != rows*cols {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		for _, p := range testPoints {
			mat1.SetSCharAt(p.row, p.col, p.val)

			if got := b[p.row*cols+p.col]; got != p.val {
				t.Errorf("Expected %d,%d = %d, but it was %d", p.row, p.col, p.val, got)
			}
		}

		mat2 := NewMatWithSize(3, 9, MatTypeCV32F)
		defer mat2.Close()
		b, err = mat2.DataPtrInt8()
		if err != nil {
			t.Error(err)
		}
		if len(b) != 3*9*4 {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		mat3 := mat1.Region(image.Rect(25, 25, 75, 75))
		defer mat3.Close()
		_, err = mat3.DataPtrInt8()
		if err == nil {
			t.Errorf("Expected error.")
		}
	})
	t.Run("Uint16", func(t *testing.T) {
		testPoints := []struct {
			row int
			col int
			val uint16
		}{
			{row: 0, col: 0, val: 10},
			{row: 30, col: 31, val: 20},
			{row: rows - 1, col: cols - 1, val: 30},
		}

		mat1 := NewMatWithSize(rows, cols, MatTypeCV16U)
		defer mat1.Close()

		b, err := mat1.DataPtrUint16()
		if err != nil {
			t.Error(err)
		}
		if len(b) != rows*cols {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		for _, p := range testPoints {
			mat1.SetShortAt(p.row, p.col, int16(p.val))

			if got := b[p.row*cols+p.col]; got != p.val {
				t.Errorf("Expected %d,%d = %d, but it was %d", p.row, p.col, p.val, got)
			}
		}

		mat2 := NewMatWithSize(3, 9, MatTypeCV32F)
		defer mat2.Close()
		_, err = mat2.DataPtrUint16()
		if err == nil {
			t.Errorf("Expected error.")
		}

		mat3 := mat1.Region(image.Rect(25, 25, 75, 75))
		defer mat3.Close()
		_, err = mat3.DataPtrUint16()
		if err == nil {
			t.Errorf("Expected error.")
		}
	})
	t.Run("Int16", func(t *testing.T) {
		testPoints := []struct {
			row int
			col int
			val int16
		}{
			{row: 0, col: 0, val: 10},
			{row: 30, col: 31, val: 20},
			{row: rows - 1, col: cols - 1, val: 30},
		}

		mat1 := NewMatWithSize(rows, cols, MatTypeCV16S)
		defer mat1.Close()

		b, err := mat1.DataPtrInt16()
		if err != nil {
			t.Error(err)
		}
		if len(b) != rows*cols {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		for _, p := range testPoints {
			mat1.SetShortAt(p.row, p.col, p.val)

			if got := b[p.row*cols+p.col]; got != p.val {
				t.Errorf("Expected %d,%d = %d, but it was %d", p.row, p.col, p.val, got)
			}
		}

		mat2 := NewMatWithSize(3, 9, MatTypeCV32F)
		defer mat2.Close()
		_, err = mat2.DataPtrInt16()
		if err == nil {
			t.Errorf("Expected error.")
		}

		mat3 := mat1.Region(image.Rect(25, 25, 75, 75))
		defer mat3.Close()
		_, err = mat3.DataPtrInt16()
		if err == nil {
			t.Errorf("Expected error.")
		}
	})
	t.Run("Float32", func(t *testing.T) {
		testPoints := []struct {
			row int
			col int
			val float32
		}{
			{row: 0, col: 0, val: 10.5},
			{row: 30, col: 31, val: 20.5},
			{row: rows - 1, col: cols - 1, val: 30.5},
		}

		mat1 := NewMatWithSize(rows, cols, MatTypeCV32F)
		defer mat1.Close()

		b, err := mat1.DataPtrFloat32()
		if err != nil {
			t.Error(err)
		}
		if len(b) != rows*cols {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		for _, p := range testPoints {
			mat1.SetFloatAt(p.row, p.col, p.val)

			if got := b[p.row*cols+p.col]; got != p.val {
				t.Errorf("Expected %d,%d = %f, but it was %f", p.row, p.col, p.val, got)
			}
		}

		mat2 := NewMatWithSize(3, 9, MatTypeCV16S)
		defer mat2.Close()
		_, err = mat2.DataPtrFloat32()
		if err == nil {
			t.Errorf("Expected error.")
		}

		mat3 := mat1.Region(image.Rect(25, 25, 75, 75))
		defer mat3.Close()
		_, err = mat3.DataPtrFloat32()
		if err == nil {
			t.Errorf("Expected error.")
		}
	})
	t.Run("Float64", func(t *testing.T) {
		testPoints := []struct {
			row int
			col int
			val float64
		}{
			{row: 0, col: 0, val: 10.5},
			{row: 30, col: 31, val: 20.5},
			{row: rows - 1, col: cols - 1, val: 30.5},
		}

		mat1 := NewMatWithSize(rows, cols, MatTypeCV64F)
		defer mat1.Close()

		b, err := mat1.DataPtrFloat64()
		if err != nil {
			t.Error(err)
		}
		if len(b) != rows*cols {
			t.Errorf("Mat bytes incorrect length: %v\n", len(b))
		}

		for _, p := range testPoints {
			mat1.SetDoubleAt(p.row, p.col, p.val)

			if got := b[p.row*cols+p.col]; got != p.val {
				t.Errorf("Expected %d,%d = %f, but it was %f", p.row, p.col, p.val, got)
			}
		}

		mat2 := NewMatWithSize(3, 9, MatTypeCV16S)
		defer mat2.Close()
		_, err = mat2.DataPtrFloat64()
		if err == nil {
			t.Errorf("Expected error.")
		}

		mat3 := mat1.Region(image.Rect(25, 25, 75, 75))
		defer mat3.Close()
		_, err = mat3.DataPtrFloat64()
		if err == nil {
			t.Errorf("Expected error.")
		}
	})
}

func TestMatRegion(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV8U)
	defer mat.Close()
	region := mat.Region(image.Rect(20, 25, 80, 75))
	defer region.Close()
	if region.Rows() != 50 {
		t.Errorf("Mat region incorrect row count: %v\n", region.Rows())
	}

	if region.Cols() != 60 {
		t.Errorf("Mat region incorrect col count: %v\n", region.Cols())
	}
}

func TestMatReshape(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV8UC4)
	defer mat.Close()

	r := mat.Reshape(1, 1)
	defer r.Close()
	if r.Rows() != 1 {
		t.Errorf("Mat reshape incorrect row count: %v\n", r.Rows())
	}

	if r.Cols() != 40000 {
		t.Errorf("Mat reshape incorrect col count: %v\n", r.Cols())
	}
}

func TestMatPatchNaNs(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV32F)
	defer mat.Close()

	mat.PatchNaNs()
	if mat.Empty() {
		t.Error("TestMatPatchNaNs error.")
	}
}

func TestMatConvert(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV32F)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	src.ConvertTo(&dst, MatTypeCV16S)
	if dst.Empty() {
		t.Error("TestConvert dst should not be empty.")
	}
}

func TestMatConvertWithParams(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	src.ConvertToWithParams(&dst, MatTypeCV32F, 1.0/255.0, 0.0)
	if dst.Empty() {
		t.Error("TestConvertWithParams dst should not be empty.")
	}
}

func TestMatConvertFp16(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV32F)
	defer src.Close()
	dst := src.ConvertFp16()
	defer dst.Close()
	if dst.Empty() {
		t.Error("TestConvertFp16 dst should not be empty.")
	}
}

func TestMatSqrt(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV32F)
	defer src.Close()

	dst := src.Sqrt()
	defer dst.Close()
	if dst.Empty() {
		t.Error("TestSqrt dst should not be empty.")
	}
}

func TestMatMean(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV8U)
	defer mat.Close()
	mean := mat.Mean()
	if mean.Val1 != 0 {
		t.Errorf("Mat Mean incorrect Val1")
	}
}

func TestMatMeanWithMask(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV8U)
	defer mat.Close()
	mask := NewMatWithSize(100, 100, MatTypeCV8U)
	defer mask.Close()
	mean := mat.MeanWithMask(mask)
	if mean.Val1 != 0 {
		t.Errorf("Mat Mean incorrect Val1")
	}
}

func TestLUT(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Source Mat in LUT test")
	}
	defer src.Close()

	lut := IMRead("images/lut.png", IMReadColor)
	if lut.Empty() {
		t.Error("Invalid read of LUT Mat in LUT test")
	}
	defer lut.Close()

	dst := NewMat()
	defer dst.Close()

	LUT(src, lut, &dst)
	if dst.Cols() != 400 || dst.Rows() != 343 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestMatAccessors(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	if mat.GetUCharAt(50, 50) != 0 {
		t.Errorf("GetUCharAt incorrect value: %v\n", mat.GetUCharAt(50, 50))
	}
	if mat.GetUCharAt3(50, 50, 0) != 0 {
		t.Errorf("GetUCharAt3 incorrect value: %v\n", mat.GetUCharAt3(50, 50, 0))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV8S)
	if mat.GetSCharAt(50, 50) != 0 {
		t.Errorf("GetSCharAt incorrect value: %v\n", mat.GetSCharAt(50, 50))
	}
	if mat.GetSCharAt3(50, 50, 0) != 0 {
		t.Errorf("GetSCharAt3 incorrect value: %v\n", mat.GetSCharAt3(50, 50, 0))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV16S)
	if mat.GetShortAt(50, 50) != 0 {
		t.Errorf("GetShortAt incorrect value: %v\n", mat.GetShortAt(50, 50))
	}
	if mat.GetShortAt3(50, 50, 0) != 0 {
		t.Errorf("GetShortAt3 incorrect value: %v\n", mat.GetShortAt3(50, 50, 0))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV32S)
	if mat.GetIntAt(50, 50) != 0 {
		t.Errorf("GetIntAt incorrect value: %v\n", mat.GetIntAt(50, 50))
	}
	if mat.GetIntAt3(50, 50, 0) != 0 {
		t.Errorf("GetIntAt3 incorrect value: %v\n", mat.GetIntAt3(50, 50, 0))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV32F)
	if mat.GetFloatAt(50, 50) != 0.0 {
		t.Errorf("GetFloatAt incorrect value: %v\n", mat.GetFloatAt(50, 50))
	}
	if mat.GetFloatAt3(50, 50, 0) != 0.0 {
		t.Errorf("GetFloatAt3 incorrect value: %v\n", mat.GetFloatAt3(50, 50, 0))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV64F)
	if mat.GetDoubleAt(50, 50) != 0.0 {
		t.Errorf("GetDoubleAt incorrect value: %v\n", mat.GetDoubleAt(50, 50))
	}
	if mat.GetDoubleAt3(50, 50, 0) != 0.0 {
		t.Errorf("GetDoubleAt3 incorrect value: %v\n", mat.GetDoubleAt3(50, 50, 0))
	}
	mat.Close()
}

func TestMatMutators(t *testing.T) {
	t.Run("SetTo", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(0, 0, 0, 0), 1, 1, MatTypeCV8U)
		mat.SetTo(NewScalar(255, 255, 255, 255))
		for z := 0; z < mat.Channels(); z++ {
			if mat.GetUCharAt3(0, 0, z) != 255 {
				t.Errorf("SetTo incorrect value: z=%v: %v\n", z, mat.GetUCharAt3(0, 0, z))
			}
		}
		mat.Close()
	})
	t.Run("SetUCharAt", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV8U)
		mat.SetUCharAt(50, 50, 25)
		if mat.GetUCharAt(50, 50) != 25 {
			t.Errorf("SetUCharAt incorrect value: %v\n", mat.GetUCharAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SetUCharAt3", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV8U)
		mat.SetUCharAt3(50, 50, 0, 25)
		if mat.GetUCharAt3(50, 50, 0) != 25 {
			t.Errorf("SetUCharAt3 incorrect value: %v\n", mat.GetUCharAt3(50, 50, 0))
		}
		mat.Close()
	})
	t.Run("SetSCharAt", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV8S)
		mat.SetSCharAt(50, 50, 25)
		if mat.GetSCharAt(50, 50) != 25 {
			t.Errorf("SetSCharAt incorrect value: %v\n", mat.GetSCharAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SetSCharAt3", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV8S)
		mat.SetSCharAt3(50, 50, 0, 25)
		if mat.GetSCharAt3(50, 50, 0) != 25 {
			t.Errorf("SetSCharAt3 incorrect value: %v\n", mat.GetSCharAt3(50, 50, 0))
		}
		mat.Close()
	})
	t.Run("SetShortAt", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV16S)
		mat.SetShortAt(50, 50, 25)
		if mat.GetShortAt(50, 50) != 25 {
			t.Errorf("SetShortAt incorrect value: %v\n", mat.GetShortAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SetShortAt3", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV16S)
		mat.SetShortAt3(50, 50, 0, 25)
		if mat.GetShortAt3(50, 50, 0) != 25 {
			t.Errorf("SetShortAt3 incorrect value: %v\n", mat.GetShortAt3(50, 50, 0))
		}
		mat.Close()
	})
	t.Run("SetIntAt", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV32S)
		mat.SetIntAt(50, 50, 25)
		if mat.GetIntAt(50, 50) != 25 {
			t.Errorf("SetIntAt incorrect value: %v\n", mat.GetIntAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SetIntAt3", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV32S)
		mat.SetIntAt3(50, 50, 0, 25)
		if mat.GetIntAt3(50, 50, 0) != 25 {
			t.Errorf("SetIntAt3 incorrect value: %v\n", mat.GetIntAt3(50, 50, 0))
		}
		mat.Close()
	})
	t.Run("SetFloatAt", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV32F)
		mat.SetFloatAt(50, 50, 25.0)
		if mat.GetFloatAt(50, 50) != 25 {
			t.Errorf("SetFloatAt incorrect value: %v\n", mat.GetFloatAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SetFloatAt3", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV32F)
		mat.SetFloatAt3(50, 50, 0, 25.0)
		if mat.GetFloatAt3(50, 50, 0) != 25 {
			t.Errorf("SetFloatAt incorrect value: %v\n", mat.GetFloatAt3(50, 50, 0))
		}
		mat.Close()
	})
	t.Run("SetDoubleAt", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV64F)
		mat.SetDoubleAt(50, 50, 25.0)
		if mat.GetDoubleAt(50, 50) != 25.0 {
			t.Errorf("SetDoubleAt incorrect value: %v\n", mat.GetDoubleAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SetDoubleAt3", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV64F)
		mat.SetDoubleAt3(50, 50, 0, 25.0)
		if mat.GetDoubleAt3(50, 50, 0) != 25.0 {
			t.Errorf("SetDoubleAt3 incorrect value: %v\n", mat.GetDoubleAt3(50, 50, 0))
		}
		mat.Close()
	})
	t.Run("AddUChar", func(t *testing.T) {
		mat := NewMatWithSize(101, 102, MatTypeCV8U)
		mat.AddUChar(42)
		if mat.GetUCharAt(50, 50) != 42 {
			t.Errorf("AddUChar incorrect value: %v\n", mat.GetUCharAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SubtractUChar", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(42.0, 0, 0, 0), 101, 102, MatTypeCV8U)
		mat.SubtractUChar(40)
		if mat.GetUCharAt(50, 50) != 2 {
			t.Errorf("SubtractUChar incorrect value: %v\n", mat.GetUCharAt(50, 50))
		}
		mat.Close()
	})
	t.Run("MultiplyUChar", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(5.0, 0, 0, 0), 101, 102, MatTypeCV8U)
		mat.MultiplyUChar(5)
		if mat.GetUCharAt(50, 50) != 25 {
			t.Errorf("MultiplyUChar incorrect value: %v\n", mat.GetUCharAt(50, 50))
		}
		mat.Close()
	})
	t.Run("DivideUChar", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(25.0, 0, 0, 0), 101, 102, MatTypeCV8U)
		mat.DivideUChar(5)
		if mat.GetUCharAt(50, 50) != 5 {
			t.Errorf("DivideUChar incorrect value: %v\n", mat.GetUCharAt(50, 50))
		}
		mat.Close()
	})
	t.Run("AddFloat", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(30.0, 0, 0, 0), 101, 102, MatTypeCV32F)
		mat.AddFloat(1.0)
		if mat.GetFloatAt(50, 50) != 31.0 {
			t.Errorf("AddFloat incorrect value: %v\n", mat.GetFloatAt(50, 50))
		}
		mat.Close()
	})
	t.Run("SubtractFloat", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(30.0, 0, 0, 0), 101, 102, MatTypeCV32F)
		mat.SubtractFloat(1.0)
		if mat.GetFloatAt(50, 50) != 29.0 {
			t.Errorf("SubtractFloat incorrect value: %v\n", mat.GetFloatAt(50, 50))
		}
		mat.Close()
	})
	t.Run("MultiplyFloat", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(30.0, 0, 0, 0), 101, 102, MatTypeCV32F)
		mat.MultiplyFloat(2.0)
		if mat.GetFloatAt(50, 50) != 60.0 {
			t.Errorf("MultiplyFloat incorrect value: %v\n", mat.GetFloatAt(50, 50))
		}
		mat.Close()
	})
	t.Run("DivideFloat", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(30.0, 0, 0, 0), 101, 102, MatTypeCV32F)
		mat.DivideFloat(2.0)
		if mat.GetFloatAt(50, 50) != 15.0 {
			t.Errorf("DivideFloat incorrect value: %v\n", mat.GetFloatAt(50, 50))
		}
		mat.Close()
	})
	t.Run("MultiplyMatrix", func(t *testing.T) {
		mat := NewMatWithSizeFromScalar(NewScalar(30.0, 0, 0, 0), 2, 1, MatTypeCV32F)
		mat2 := NewMatWithSizeFromScalar(NewScalar(30.0, 0, 0, 0), 1, 2, MatTypeCV32F)
		mat3 := mat.MultiplyMatrix(mat2)
		for i := 0; i < mat3.Cols(); i++ {
			for j := 0; j < mat3.Rows(); j++ {
				if mat3.GetFloatAt(i, j) != 900.0 {
					t.Errorf("MultiplyMatrix incorrect value: %v\n", mat3.GetFloatAt(i, j))
				}
			}
		}
		mat.Close()
		mat2.Close()
		mat3.Close()
	})
}

func TestMatAbsDiff(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	AbsDiff(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatAbsDiff dest mat3 should not be empty.")
	}
}

func TestMatAdd(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	Add(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatAdd dest mat3 should not be empty.")
	}
}

func TestMatAddWeighted(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	AddWeighted(mat1, 2.0, mat2, 3.0, 4.0, &mat3)
	if mat3.Empty() {
		t.Error("TestMatAddWeighted dest mat3 should not be empty.")
	}
}

func TestMatBitwiseOperations(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	BitwiseAnd(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseAnd dest mat3 should not be empty.")
	}

	BitwiseOr(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseOr dest mat3 should not be empty.")
	}

	BitwiseXor(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseXor dest mat3 should not be empty.")
	}

	BitwiseNot(mat1, &mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseNot dest mat3 should not be empty.")
	}

}

func TestMatBitwiseOperationsWithMasks(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	mat4 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat4.Close()
	BitwiseAndWithMask(mat1, mat2, &mat3, mat4)
	if mat3.Empty() {
		t.Error("TestMatBitwiseAndWithMask dest mat3 should not be empty.")
	}

	BitwiseOrWithMask(mat1, mat2, &mat3, mat4)
	if mat3.Empty() {
		t.Error("TestMatBitwiseOrWithMask dest mat3 should not be empty.")
	}

	BitwiseXorWithMask(mat1, mat2, &mat3, mat4)
	if mat3.Empty() {
		t.Error("TestMatBitwiseXorWithMask dest mat3 should not be empty.")
	}
	BitwiseNotWithMask(mat1, &mat3, mat4)
	if mat3.Empty() {
		t.Error("TestMatBitwiseNotWithMask dest mat3 should not be empty.")
	}
}

func TestMatInRange(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	lb := NewMatFromScalar(NewScalar(20.0, 100.0, 100.0, 0.0), MatTypeCV8U)
	defer lb.Close()
	ub := NewMatFromScalar(NewScalar(20.0, 100.0, 100.0, 0.0), MatTypeCV8U)
	defer ub.Close()
	dst := NewMat()
	defer dst.Close()
	InRange(mat1, lb, ub, &dst)
	if dst.Empty() {
		t.Error("TestMatAddWeighted dest mat3 should not be empty.")
	}
}

func TestMatInRangeWithScalar(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	lb := NewScalar(20.0, 100.0, 100.0, 0.0)
	ub := NewScalar(20.0, 100.0, 100.0, 0.0)
	dst := NewMat()
	defer dst.Close()
	InRangeWithScalar(mat1, lb, ub, &dst)
	if dst.Empty() {
		t.Error("TestMatAddWeighted dest mat3 should not be empty.")
	}
}

func TestMatDCT(t *testing.T) {
	src := NewMatWithSize(64, 64, MatTypeCV32F)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	DCT(src, &dst, DftForward)
	if dst.Empty() {
		t.Error("TestMatDCT dst should not be empty.")
	}
}

func TestMatDFT(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV32F)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	m := GetOptimalDFTSize(101)
	n := GetOptimalDFTSize(102)
	if m != 108 {
		t.Errorf("TestMatOptimalDFT dst error: %d", m)
	}

	if n != 108 {
		t.Errorf("TestMatOptimalDFT dst error: %d", n)
	}

	DFT(src, &dst, DftForward)
	if dst.Empty() {
		t.Error("TestMatDFT dst should not be empty.")
	}
}

func TestMatDivide(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	Divide(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatDivide dest mat3 should not be empty.")
	}
}

func TestMeanStdDev(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	dstStdDev := NewMat()
	defer dstStdDev.Close()
	MeanStdDev(src, &dst, &dstStdDev)
	if dst.Empty() {
		t.Error("TestMeanStdDev dst should not be empty.")
	}
	if dstStdDev.Empty() {
		t.Error("TestMeanStdDev dstStdDev should not be empty.")
	}
}

func TestMatMerge(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV8U)
	defer src.Close()
	src2 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer src2.Close()
	src3 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer src3.Close()
	dst := NewMat()
	defer dst.Close()
	Merge([]Mat{src, src2, src3}, &dst)
	if dst.Empty() {
		t.Error("TestMatMerge dst should not be empty.")
	}
}

func TestMatMulSpectrums(t *testing.T) {
	a := NewMatWithSize(101, 102, MatTypeCV32F)
	defer a.Close()
	b := NewMatWithSize(101, 102, MatTypeCV32F)
	defer b.Close()
	dst := NewMat()
	defer dst.Close()
	MulSpectrums(a, b, &dst, 0)
	if dst.Empty() {
		t.Error("TestMatMulSpectrums dst should not be empty.")
	}
	dst2 := NewMat()
	defer dst2.Close()
	//test with dftrows flag (the only flag accepted in addition to 0)
	MulSpectrums(a, b, &dst2, DftRows)
	if dst2.Empty() {
		t.Error("TestMatMulSpectrums dst should not be empty.")
	}
}

func TestMatMultiply(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV64F)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV64F)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	Multiply(mat1, mat2, &mat3)
	if mat3.Empty() {
		t.Error("TestMatMultiply dest mat3 should not be empty.")
	}

	// since this is a single channel Mat, only the first value in the scalar is used
	mat4 := NewMatWithSizeFromScalar(NewScalar(2.0, 0.0, 0.0, 0.0), 101, 102, MatTypeCV64F)
	defer mat4.Close()
	mat5 := NewMatWithSizeFromScalar(NewScalar(3.0, 0.0, 0.0, 0.0), 101, 102, MatTypeCV64F)
	defer mat5.Close()
	Multiply(mat4, mat5, &mat3)
	if mat3.Empty() {
		t.Error("TestMatMultiply dest mat3 should not be empty.")
	}
	if mat3.GetDoubleAt(0, 0) != 6.0 {
		t.Error("TestMatMultiply invalue value in dest mat3.")
	}
}

func TestMatMultiplyWithParams(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV64F)
	defer mat1.Close()
	mat2 := NewMatWithSize(101, 102, MatTypeCV64F)
	defer mat2.Close()
	mat3 := NewMat()
	defer mat3.Close()
	MultiplyWithParams(mat1, mat2, &mat3, 0.5, -1)
	if mat3.Empty() {
		t.Error("TestMatMultiplyWithParams dest mat3 should not be empty.")
	}

	// since this is a single channel Mat, only the first value in the scalar is used
	mat4 := NewMatWithSizeFromScalar(NewScalar(2.0, 0.0, 0.0, 0.0), 101, 102, MatTypeCV64F)
	defer mat4.Close()
	mat5 := NewMatWithSizeFromScalar(NewScalar(3.0, 0.0, 0.0, 0.0), 101, 102, MatTypeCV64F)
	defer mat5.Close()
	MultiplyWithParams(mat4, mat5, &mat3, 2.0, -1)
	if mat3.Empty() {
		t.Error("TestMatMultiplyWithParams dest mat3 should not be empty.")
	}
	if mat3.GetDoubleAt(0, 0) != 12.0 {
		t.Error("TestMatMultiplyWithParams invalue value in dest mat3.")
	}
}

func TestMatNormalize(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	Normalize(src, &dst, 0.0, 255.0, NormMinMax)
	if dst.Empty() {
		t.Error("TestMatNormalize dst should not be empty.")
	}
}

func TestMatPerspectiveTransform(t *testing.T) {
	src := NewMatWithSize(100, 1, MatTypeCV32F+MatChannels2)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	tm := NewMatWithSize(3, 3, MatTypeCV32F)
	defer tm.Close()
	PerspectiveTransform(src, &dst, tm)
	if dst.Empty() {
		t.Error("PerspectiveTransform error")
	}
}

func TestMatSolve(t *testing.T) {
	a := NewMatWithSize(3, 3, MatTypeCV32F)
	defer a.Close()
	b := NewMatWithSize(3, 1, MatTypeCV32F)
	defer b.Close()
	solve := NewMat()
	defer solve.Close()

	testPoints := []struct {
		x2 float32
		x  float32
		c  float32
		y  float32
	}{
		{x2: 1, x: 1, c: 1, y: 0},
		{x2: 0, x: 0, c: 1, y: 2},
		{x2: 9, x: 3, c: 1, y: 2},
	}

	for row, p := range testPoints {
		a.SetFloatAt(row, 0, p.x2)
		a.SetFloatAt(row, 1, p.x)
		a.SetFloatAt(row, 2, p.c)

		b.SetFloatAt(row, 0, p.y)
	}

	solved := Solve(a, b, &solve, SolveDecompositionLu)

	if !solved {
		t.Errorf("TestMatSolve could not solve linear equations")
	}

	if solve.GetFloatAt(0, 0) != 1 || solve.GetFloatAt(1, 0) != -3 || solve.GetFloatAt(2, 0) != 2 {
		t.Errorf("TestMatSolve incorrect results: got %v expected %v, got %v expected %v, got %v expected %v",
			solve.GetFloatAt(0, 0), 1,
			solve.GetFloatAt(1, 0), -3,
			solve.GetFloatAt(2, 0), 2)
	}
}

func TestSolveCubic(t *testing.T) {
	coeffs := NewMatWithSize(1, 4, MatTypeCV32F)
	defer coeffs.Close()
	roots := NewMat()
	defer roots.Close()

	coeffs.SetFloatAt(0, 0, 2.0)
	coeffs.SetFloatAt(0, 1, 3.0)
	coeffs.SetFloatAt(0, 2, -11.0)
	coeffs.SetFloatAt(0, 3, -6.0)

	rootsCount := SolveCubic(coeffs, &roots)

	expectedRootsCount := 3
	if rootsCount != expectedRootsCount {
		t.Errorf("TestSolveCubic incorrect numbers of roots %d, expected %d", rootsCount, expectedRootsCount)
	}

	if roots.GetFloatAt(0, 0) != -3.0 || roots.GetFloatAt(0, 1) != 2.0 || roots.GetFloatAt(0, 0) != -3.0 {
		t.Errorf("TestSolveCubic incorrect roots: got %f expected %f, got %f expected %f, got %f expected %f",
			roots.GetFloatAt(0, 0), -3.0,
			roots.GetFloatAt(0, 1), -0.5,
			roots.GetFloatAt(0, 0), -3.0)
	}
}

func TestSolvePoly(t *testing.T) {
	coeffs := NewMatWithSize(1, 3, MatTypeCV32F)
	defer coeffs.Close()
	roots := NewMat()
	defer roots.Close()

	// x² - 14x + 49 = 0
	coeffs.SetFloatAt(0, 0, 49.0)
	coeffs.SetFloatAt(0, 1, -14.0)
	coeffs.SetFloatAt(0, 2, 1)

	diffError := SolvePoly(coeffs, &roots, 300)

	diffTolerance := 1.0e-61
	if diffError > diffTolerance {
		t.Errorf("TestSolvePoly was not exact, got an error of %e and should have been less than %f", diffError, diffTolerance)
	}

	if roots.GetFloatAt(0, 0) != 7.0 {
		t.Errorf("TestSolvePoly incorrect roots: got %f expected %f",
			roots.GetFloatAt(0, 0), 7.0)
	}
}

func TestMatReduceToSingleRow(t *testing.T) {
	rows := 2
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col+1))
		}
	}

	Reduce(src, &dst, 0, ReduceSum, MatTypeCV32F)

	sz := dst.Size()
	if sz[0] != 1 && sz[1] != 3 {
		t.Errorf("TestMatReduceToSingleRow incorrect size: %v\n", sz)
	}

	if dst.GetFloatAt(0, 0) != 2 || dst.GetFloatAt(0, 1) != 4 || dst.GetFloatAt(0, 2) != 6 {
		t.Errorf("TestMatReduceToSingleRow incorrect reduce result: %v at (0, 0) expected 2, %v at (0, 1) expected 4, %v at (0, 2) expected 6",
			dst.GetFloatAt(0, 0), dst.GetFloatAt(0, 1), dst.GetFloatAt(0, 2))
	}
}

func TestMatReduceToSingleColumn(t *testing.T) {
	rows := 2
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col+1))
		}
	}

	Reduce(src, &dst, 1, ReduceSum, MatTypeCV32F)

	sz := dst.Size()
	if sz[0] != 3 && sz[1] != 1 {
		t.Errorf("TestMatReduceToSingleColumn incorrect size: %v\n", sz)
	}

	if dst.GetFloatAt(0, 0) != 6 || dst.GetFloatAt(1, 0) != 6 {
		t.Errorf("TestMatReduceToSingleColumn incorrect reduce result: %v at (0, 0) expected 6, %v at (1, 0) expected 6",
			dst.GetFloatAt(0, 0), dst.GetFloatAt(1, 0))
	}
}

func TestMatReduceArgMax(t *testing.T) {
	rows := 2
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col+1))
		}
	}

	ReduceArgMax(src, &dst, 1, false)

	sz := dst.Size()
	if sz[0] != 2 && sz[1] != 1 {
		t.Errorf("TestMatReduceArgMax incorrect size: %v\n", sz)
	}

	if dst.GetUCharAt(0, 0) != 2 || dst.GetUCharAt(1, 0) != 2 {
		t.Errorf("TestMatReduceArgMax incorrect reduce result: %v at (0, 0) expected 1, %v at (1, 0) expected 1",
			dst.GetUCharAt(0, 0), dst.GetUCharAt(1, 0))
	}
}

func TestMatReduceArgMin(t *testing.T) {
	rows := 2
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col+1))
		}
	}

	ReduceArgMin(src, &dst, 1, false)

	sz := dst.Size()
	if sz[0] != 2 && sz[1] != 1 {
		t.Errorf("TestMatReduceArgMax incorrect size: %v\n", sz)
	}

	if dst.GetUCharAt(0, 0) != 0 || dst.GetUCharAt(1, 0) != 0 {
		t.Errorf("TestMatReduceArgMax incorrect reduce result: %v at (0, 0) expected 0, %v at (1, 0) expected 0",
			dst.GetUCharAt(0, 0), dst.GetUCharAt(1, 0))
	}
}

func TestRepeat(t *testing.T) {
	rows := 1
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col))
		}
	}

	dst := NewMat()
	defer dst.Close()
	Repeat(src, 3, 1, &dst)

	size := dst.Size()
	expectedRows := 3
	expectedCols := 3

	if size[0] != expectedRows || size[1] != expectedCols {
		t.Errorf("TestRepeat incorrect size, got y=%d x=%d, expected y=%d x=%d.", size[0], size[1], expectedRows, expectedCols)
	}

	for row := 0; row < expectedRows; row++ {
		for col := 0; col < expectedCols; col++ {

			result := dst.GetUCharAt(row, col)

			if result != uint8(col) {
				t.Errorf("TestRepeat dst at row=%d col=%d should be %d and got %d.", row, col, col, result)
			}
		}
	}
}

func TestScaleAdd(t *testing.T) {
	rows := 2
	cols := 3
	src1 := NewMatWithSize(rows, cols, MatTypeCV64F)
	defer src1.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src1.SetDoubleAt(row, col, float64(col))
		}
	}

	src2 := NewMatWithSize(rows, cols, MatTypeCV64F)
	defer src2.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src2.SetDoubleAt(row, col, 1.0)
		}
	}

	dst := NewMat()
	defer dst.Close()

	alpha := 1.5
	ScaleAdd(src1, alpha, src2, &dst)

	if dst.Empty() {
		t.Error("TestScaleAdd dst should not be empty.")
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			expected := float64(col)*alpha + 1.0
			result := dst.GetDoubleAt(row, col)
			if result != expected {
				t.Errorf("TestScaleAdd dst at row=%d col=%d should be %f and got %f.", row, col, expected, result)
			}
		}
	}
}

func TestSetIdentity(t *testing.T) {
	rows := 4
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV64F)
	defer src.Close()
	scalar := 2.5
	SetIdentity(src, scalar)

	if src.Empty() {
		t.Error("TestSetIdentity src should not be empty.")
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			result := src.GetDoubleAt(row, col)
			expected := 0.0
			if row == col {
				expected = scalar
			}
			if result != expected {
				t.Errorf("TestSetIdentity src at row=%d col=%d should be %f and got %f.", row, col, expected, result)
			}
		}
	}
}

func TestMatSortEveryRowDescending(t *testing.T) {
	rows := 2
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col))
		}
	}

	dst := NewMat()
	defer dst.Close()

	flags := SortEveryRow + SortDescending
	Sort(src, &dst, flags)

	if dst.Empty() {
		t.Error("TestMatSortEveryRowDescending dst should not be empty.")
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			expected := cols - col - 1
			result := dst.GetUCharAt(row, col)
			if result != uint8(expected) {
				t.Errorf("TestMatSortEveryRowDescending dst at row=%d col=%d should be %d and got %d.", row, col, expected, result)
			}
		}
	}
}

func TestMatSortIdxEveryRowDescending(t *testing.T) {
	rows := 2
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col))
		}
	}

	dst := NewMat()
	defer dst.Close()
	flags := SortEveryRow + SortDescending
	SortIdx(src, &dst, flags)

	if dst.Empty() {
		t.Error("TestMatSortIdxEveryRowDescending dst should not be empty.")
	}
}

func TestMatSplit(t *testing.T) {
	src := IMRead("images/face.jpg", 1)
	defer src.Close()
	chans := Split(src)
	if len(chans) != src.Channels() {
		t.Error("Split Channel count differs")
	}
	dst := NewMat()
	defer dst.Close()
	Merge(chans, &dst)
	for _, ch := range chans {
		ch.Close()
	}
	diff := NewMat()
	defer diff.Close()
	AbsDiff(src, dst, &diff)
	sum := diff.Sum()
	if sum.Val1 != 0 || sum.Val2 != 0 || sum.Val3 != 0 {
		t.Error("Split/Merged images differ")
	}
}

func TestMatSubtract(t *testing.T) {
	src1 := IMRead("images/lut.png", 1)
	defer src1.Close()
	src2 := IMRead("images/lut.png", 1)
	defer src2.Close()
	dst := NewMat()
	defer dst.Close()
	Subtract(src1, src2, &dst)
	sum := dst.Sum()
	if sum.Val1 != 0 || sum.Val2 != 0 || sum.Val3 != 0 {
		t.Error("Sum of Subtracting equal images is not 0")
	}
}

func TestMatTrace(t *testing.T) {
	rows := 3
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()

	// Create and identity eye matrix
	for row := 0; row <= rows; row++ {
		for col := 0; col <= cols; col++ {
			if row == col {
				src.SetUCharAt(row, col, uint8(1))
			}
		}
	}

	trace := Trace(src)
	expected := NewScalar(3, 0, 0, 0)

	if trace.Val1 != expected.Val1 || trace.Val2 != expected.Val2 || trace.Val3 != expected.Val3 || trace.Val4 != expected.Val4 {
		t.Errorf("Trace values should be %v and was %v", expected, trace)
	}
}

func TestMatTransform(t *testing.T) {
	src := IMRead("images/lut.png", 1)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	tm := NewMatWithSize(4, 4, MatTypeCV8UC4)
	defer tm.Close()
	Transform(src, &dst, tm)
	if dst.Empty() {
		t.Error("Transform error")
	}
}

func TestMatTranspose(t *testing.T) {
	src := IMRead("images/lut.png", 1)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	Transpose(src, &dst)
	if dst.Empty() {
		t.Error("Transpose error")
	}
}

func TestMatTransposeND(t *testing.T) {
	rows := 1
	cols := 3
	src := NewMatWithSize(rows, cols, MatTypeCV8U)
	defer src.Close()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			src.SetUCharAt(row, col, uint8(col))
		}
	}

	dst := NewMat()
	defer dst.Close()
	TransposeND(src, []int{1, 0}, &dst)
	if dst.Empty() {
		t.Error("TransposeND error")
	}
}

func TestPolarToCart(t *testing.T) {
	magnitude := NewMatWithSize(101, 102, MatTypeCV32F)
	angle := NewMatWithSize(101, 102, MatTypeCV32F)
	x := NewMat()
	y := NewMat()

	PolarToCart(magnitude, angle, &x, &y, false)

	if x.Empty() || y.Empty() {
		t.Error("TestPolarToCart neither x nor y should be empty.")
	}

	x.Close()
	y.Close()
	magnitude.Close()
	angle.Close()
}

func TestMatPow(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV8U)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	power := 2.0
	Pow(src, power, &dst)

	if dst.Empty() {
		t.Error("TestMatPow dest should not be empty.")
	}
}

func TestMatSum(t *testing.T) {
	src := NewMatFromScalar(NewScalar(1, 2, 3, 4), MatTypeCV8UC4)
	defer src.Close()
	sum := src.Sum()
	if sum.Val1 != 1 || sum.Val2 != 2 || sum.Val3 != 3 || sum.Val4 != 4 {
		t.Error("Sum values do not match constructor")
	}
}

func TestTermCriteria(t *testing.T) {
	tc := NewTermCriteria(Count, 50, 2.0)
	if tc.p == nil {
		t.Error("TermCriteria has invalid value")
	}
}

func TestScalar(t *testing.T) {
	s := NewScalar(127.0, 255.0, 64.0, 0.0)
	if s.Val1 != 127.0 || s.Val2 != 255.0 || s.Val3 != 64.0 || s.Val4 != 0.0 {
		t.Error("Scalar has invalid value")
	}
}

func TestToCPoints(t *testing.T) {
	points := []image.Point{
		image.Pt(0, 0),
		image.Pt(1, 1),
	}

	cPoints := toCPoints(points)

	if int(cPoints.length) != len(points) {
		t.Error("Invalid C Points length")
	}
}

func TestToCStrings(t *testing.T) {
	strs := []string{
		"hello",
		"fellow",
		"CStrings",
	}

	cStrs := toCStrings(strs)

	if int(cStrs.length) != len(strs) {
		t.Error("Invalid CStrings length")
	}
}

func TestMatBatchDistance(t *testing.T) {
	src1 := NewMatWithSize(100, 100, MatTypeCV8U)
	src2 := NewMatWithSize(100, 100, MatTypeCV8U)
	mask := NewMatWithSize(100, 100, MatTypeCV8U)
	dist := NewMat()
	nidx := NewMat()
	BatchDistance(src1, src2, dist, -1, nidx, NormL2, 15, mask, 0, false)
	if dist.Empty() {
		t.Error("TestBatchDistance dst should not be empty.")
	}
	src1.Close()
	src2.Close()
	mask.Close()
	dist.Close()
	nidx.Close()
}

func TestMatBorderInterpolate(t *testing.T) {
	n := BorderInterpolate(1, 5, 1)
	if n == 0 {
		t.Error("TestBorderInterpolate dst should not be 0.")
	}
}

func TestMatCalcCovarMatrix(t *testing.T) {
	samples := NewMatWithSize(10, 10, MatTypeCV32F)
	covar := NewMat()
	mean := NewMat()
	CalcCovarMatrix(samples, &covar, &mean, CovarRows, MatTypeCV64F)
	if covar.Empty() {
		t.Error("TestCalcCovarMatrix dst should not be empty.")
	}
	samples.Close()
	covar.Close()
	mean.Close()
}

func TestMatCartToPolar(t *testing.T) {
	x := NewMatWithSize(100, 100, MatTypeCV32F)
	y := NewMatWithSize(100, 100, MatTypeCV32F)
	magnitude := NewMat()
	angle := NewMat()
	CartToPolar(x, y, &magnitude, &angle, false)
	if magnitude.Empty() || angle.Empty() {
		t.Error("TestCartToPolar neither magnitude nor angle should be empty.")
	}
	x.Close()
	y.Close()
	magnitude.Close()
	angle.Close()
}

func TestMatCheckRange(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat1.Close()
	ret := CheckRange(mat1)
	if !ret {
		t.Error("TestCheckRange error.")
	}
}

func TestMatCompleteSymm(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV32F)
	CompleteSymm(src, false)
	if src.Empty() {
		t.Error("TestCompleteSymm src should not be empty.")
	}
	src.Close()
}

func TestMatConvertScaleAbs(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV32F)
	dst := NewMat()
	ConvertScaleAbs(src, &dst, 1, 0)
	if dst.Empty() {
		t.Error("TestConvertScaleAbs dst should not be empty.")
	}
	src.Close()
	dst.Close()
}

func TestMatCopyMakeBorder(t *testing.T) {
	src := NewMatWithSize(100, 100, MatTypeCV32F)
	dst := NewMat()
	CopyMakeBorder(src, &dst, 10, 10, 10, 10, BorderReflect, color.RGBA{0, 0, 0, 0})
	if dst.Empty() {
		t.Error("TestCopyMakeBorder dst should not be empty.")
	}
	src.Close()
	dst.Close()
}

func TestMatDeterminant(t *testing.T) {
	mat1 := NewMatWithSize(101, 101, MatTypeCV32F)
	defer mat1.Close()
	ret := Determinant(mat1)
	if ret != 0 {
		t.Error("TestMatDeterminant error.")
	}
}

func TestMatEigen(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	eigenvalues := NewMat()
	eigenvectors := NewMat()
	Eigen(src, &eigenvalues, &eigenvectors)
	if eigenvectors.Empty() || eigenvalues.Empty() {
		t.Error("TestEigen should not have empty eigenvectors or eigenvalues.")
	}
	src.Close()
	eigenvectors.Close()
	eigenvalues.Close()
}

func TestMatEigenNonSymmetric(t *testing.T) {
	src := NewMatWithSizeFromScalar(NewScalar(0.1, 0.1, 0.1, 0.1), 10, 10, MatTypeCV32F)
	eigenvalues := NewMat()
	eigenvectors := NewMat()
	EigenNonSymmetric(src, &eigenvalues, &eigenvectors)
	if eigenvectors.Empty() || eigenvalues.Empty() {
		t.Error("TestEigenNonSymmetric should not have empty eigenvectors or eigenvalues.")
	}
	src.Close()
	eigenvectors.Close()
	eigenvalues.Close()
}

func TestPCABackProject(t *testing.T) {
	data := NewMatWithSize(3, 1, MatTypeCV32F)
	defer data.Close()
	data.SetFloatAt(0, 0, float32(-5))
	data.SetFloatAt(1, 0, float32(0))
	data.SetFloatAt(2, 0, float32(-10))

	mean := NewMatWithSize(1, 4, MatTypeCV32F)
	defer mean.Close()
	mean.SetFloatAt(0, 0, float32(2))
	mean.SetFloatAt(0, 1, float32(4))
	mean.SetFloatAt(0, 2, float32(4))
	mean.SetFloatAt(0, 3, float32(8))

	vectors := NewMatWithSizeFromScalar(NewScalar(0, 0, 0, 0), 1, 4, MatTypeCV32F)
	defer vectors.Close()
	vectors.SetFloatAt(0, 0, float32(0.2))
	vectors.SetFloatAt(0, 1, float32(0.4))
	vectors.SetFloatAt(0, 2, float32(0.4))
	vectors.SetFloatAt(0, 3, float32(0.8))

	result := NewMat()
	defer result.Close()

	PCABackProject(data, mean, vectors, &result)
	if result.Empty() {
		t.Error("PCABackProject should not have empty result.")
	}
}

func TestPCACompute(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	// Set some source data so the PCA is done on a non-zero matrix.
	src.SetFloatAt(0, 0, 17)
	src.SetFloatAt(2, 1, 5)
	src.SetFloatAt(9, 9, 25)
	mean := NewMat()
	eigenvectors := NewMat()
	eigenvalues := NewMat()
	maxComponents := 2
	PCACompute(src, &mean, &eigenvectors, &eigenvalues, maxComponents)
	if mean.Empty() || eigenvectors.Empty() || eigenvalues.Empty() {
		t.Error("TestPCACompute should not have empty eigenvectors or eigenvalues.")
	}
	if eigenvectors.Rows() > maxComponents {
		t.Errorf("TestPCACompute unexpected numComponents, got=%d, want<=%d", eigenvectors.Rows(), maxComponents)
	}
	src.Close()
	mean.Close()
	eigenvectors.Close()
	eigenvalues.Close()
}

func TestPCAProject(t *testing.T) {
	data := NewMatWithSize(3, 4, MatTypeCV32F)
	defer data.Close()
	data.SetFloatAt(0, 0, float32(1))
	data.SetFloatAt(0, 1, float32(2))
	data.SetFloatAt(0, 2, float32(2))
	data.SetFloatAt(0, 3, float32(4))
	data.SetFloatAt(1, 0, float32(2))
	data.SetFloatAt(1, 1, float32(4))
	data.SetFloatAt(1, 2, float32(4))
	data.SetFloatAt(1, 3, float32(8))
	data.SetFloatAt(2, 0, float32(0))
	data.SetFloatAt(2, 1, float32(0))
	data.SetFloatAt(2, 2, float32(0))
	data.SetFloatAt(2, 3, float32(0))

	mean := NewMatWithSize(1, 4, MatTypeCV32F)
	defer mean.Close()
	mean.SetFloatAt(0, 0, float32(2))
	mean.SetFloatAt(0, 1, float32(4))
	mean.SetFloatAt(0, 2, float32(4))
	mean.SetFloatAt(0, 3, float32(8))

	vectors := NewMatWithSizeFromScalar(NewScalar(0, 0, 0, 0), 1, 4, MatTypeCV32F)
	defer vectors.Close()
	vectors.SetFloatAt(0, 0, float32(0.2))
	vectors.SetFloatAt(0, 1, float32(0.4))
	vectors.SetFloatAt(0, 2, float32(0.4))
	vectors.SetFloatAt(0, 3, float32(0.8))

	result := NewMat()
	defer result.Close()

	PCAProject(data, mean, vectors, &result)
	if result.Empty() {
		t.Error("PCABackProject should not have empty result.")
	}
}

func TestPSNR(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Source Mat in PSNR test")
	}
	defer src.Close()

	result := PSNR(src, src)
	if result == 0 {
		t.Error("Unexpected PSNR of 0")
	}
}

func TestSVBackSubst(t *testing.T) {
	w := NewMatWithSizeFromScalar(NewScalar(2, 0, 0, 0), 2, 2, MatTypeCV32F)
	defer w.Close()

	u := NewMatWithSizeFromScalar(NewScalar(4, 0, 0, 0), 2, 2, MatTypeCV32F)
	defer u.Close()

	vt := NewMatWithSizeFromScalar(NewScalar(2, 0, 0, 0), 2, 2, MatTypeCV32F)
	defer vt.Close()

	rhs := NewMatWithSizeFromScalar(NewScalar(1, 0, 0, 0), 2, 2, MatTypeCV32F)
	defer rhs.Close()

	dst := NewMat()
	defer dst.Close()

	SVBackSubst(w, u, vt, rhs, &dst)
	if dst.Empty() {
		t.Error("SVBackSubst should not have empty result.")
	}
}

func TestSVDecomp(t *testing.T) {
	src := NewMatWithSize(1, 4, MatTypeCV32F)
	defer src.Close()
	src.SetFloatAt(0, 0, float32(1))
	src.SetFloatAt(0, 1, float32(4))
	src.SetFloatAt(0, 2, float32(8))
	src.SetFloatAt(0, 3, float32(6))

	w := NewMat()
	defer w.Close()

	u := NewMat()
	defer u.Close()

	vt := NewMat()
	defer vt.Close()

	SVDecomp(src, &w, &u, &vt)
	if w.Empty() || u.Empty() || vt.Empty() {
		t.Error("SVDecomp should not have empty results.")
	}
}

func TestMatExp(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	dst := NewMat()
	Exp(src, &dst)
	if dst.Empty() {
		t.Error("TestExp dst should not be empty.")
	}
	src.Close()
	dst.Close()
}

func TestMatExtractChannel(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F+MatChannels3)
	dst := NewMat()
	ExtractChannel(src, &dst, 1)
	if dst.Empty() {
		t.Error("TestExtractChannel dst should not be empty.")
	}
	src.Close()
	dst.Close()
}

func TestMatFindNonZero(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV8U)
	defer src.Close()
	src.SetFloatAt(3, 3, 17)
	src.SetFloatAt(4, 4, 17)

	dst := NewMat()
	defer dst.Close()

	FindNonZero(src, &dst)

	if dst.Empty() {
		t.Error("TestMatFindNonZero dst should not be empty.")
	}
	if dst.Rows() != 2*2 {
		t.Error("TestMatFindNonZero didn't find all nonzero locations.")
	}
}

func TestMatFlip(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Flip(src, &dst, 0)

	if dst.Empty() {
		t.Error("TestMatFlip dst should not be empty.")
	}
	if dst.Rows() != src.Rows() {
		t.Error("TestMatFlip src and dst size should be same.")
	}
}

func TestMatPhase(t *testing.T) {
	x := NewMatFromScalar(NewScalar(1.2, 2.3, 3.4, 4.5), MatTypeCV32F)
	defer x.Close()

	y := NewMatFromScalar(NewScalar(5.6, 6.7, 7.8, 8.9), MatTypeCV32F)
	defer y.Close()

	angle := NewMatWithSize(4, 5, MatTypeCV32F)
	defer angle.Close()

	Phase(x, y, &angle, false)

	if angle.Empty() {
		t.Error("TestMatPhase angle should not be empty.")
	}

	if angle.Rows() != x.Rows() {
		t.Error("TestMatPhase x and angle size should be same.")
	}
}

func TestMatGemm(t *testing.T) {
	src1 := NewMatWithSize(3, 4, MatTypeCV32F)
	defer src1.Close()

	src2 := NewMatWithSize(4, 3, MatTypeCV32F)
	defer src2.Close()

	src3 := NewMat()
	defer src3.Close()

	dst := NewMat()
	defer dst.Close()

	Gemm(src1, src2, 1, src3, 0, &dst, 0)

	if dst.Empty() {
		t.Error("Gemm dst should not be empty.")
	}
	if dst.Rows() != src1.Rows() {
		t.Error("Gemm src and dst size should be same.")
	}
}

func TestMatHconcat(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Hconcat(src, src, &dst)

	if dst.Empty() {
		t.Error("TestMatHconcat dst should not be empty.")
	}
	if dst.Cols() != 2*src.Cols() {
		t.Error("TestMatHconcat dst.Cols should be 2 x src.Cols.")
	}
}

func TestMatVconcat(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Vconcat(src, src, &dst)

	if dst.Empty() {
		t.Error("TestMatVconcat dst should not be empty.")
	}
	if dst.Rows() != 2*src.Rows() {
		t.Error("TestMatVconcat dst.Cols should be 2 x src.Rows().")
	}
}

func TestRotate(t *testing.T) {
	src := NewMatWithSize(1, 2, MatTypeCV64F)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	Rotate(src, &dst, Rotate90Clockwise)
	if dst.Rows() != 2 {
		t.Errorf("expected rows: %d got %d", src.Cols(), dst.Rows())
	}

	dst2src := NewMat()
	defer dst2src.Close()

	Rotate(dst, &dst2src, Rotate90CounterClockwise)
	if dst2src.Rows() != 1 {
		t.Errorf("expected rows: %d got %d", src.Rows(), dst2src.Rows())
	}
}

func TestMatIdct(t *testing.T) {
	src := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	IDCT(src, &dst, 0)
	if dst.Empty() {
		t.Error("Idct dst should not be empty.")
	}
	if dst.Rows() != src.Rows() {
		t.Error("Idct src and dst size should be same.")
	}
}

func TestMatIdft(t *testing.T) {
	src := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	IDFT(src, &dst, 0, 0)
	if dst.Empty() {
		t.Error("Idct dst should not be empty.")
	}
	if dst.Rows() != src.Rows() {
		t.Error("Idct src and dst size should be same.")
	}
}

func TestMatInsertChannel(t *testing.T) {
	src := NewMatWithSize(4, 4, MatTypeCV8U)
	defer src.Close()

	dst := NewMatWithSize(4, 4, MatTypeCV8UC3)
	defer dst.Close()

	InsertChannel(src, &dst, 1)
	if dst.Channels() != 3 {
		t.Error("TestMatInsertChannel dst should change the channel count")
	}
}

func TestMatInvert(t *testing.T) {
	src := NewMatWithSize(4, 4, MatTypeCV32F) // only implemented for symm. Mats
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Invert(src, &dst, SolveDecompositionLu)
	if dst.Empty() {
		t.Error("Invert dst should not be empty.")
	}
}

func TestKMeans(t *testing.T) {
	src := NewMatWithSize(4, 4, MatTypeCV32F) // only implemented for symm. Mats
	defer src.Close()

	bestLabels := NewMat()
	defer bestLabels.Close()

	centers := NewMat()
	defer centers.Close()

	criteria := NewTermCriteria(Count, 10, 1.0)
	KMeans(src, 2, &bestLabels, criteria, 2, KMeansRandomCenters, &centers)
	if bestLabels.Empty() {
		t.Error("bla")
	}
}

func TestKMeansPoints(t *testing.T) {
	points := []image.Point{
		image.Pt(0, 0),
		image.Pt(1, 1),
	}

	pv := NewPointVectorFromPoints(points)
	defer pv.Close()

	bestLabels := NewMat()
	defer bestLabels.Close()
	centers := NewMat()
	defer centers.Close()

	criteria := NewTermCriteria(Count, 10, 1.0)
	KMeansPoints(pv, 2, &bestLabels, criteria, 2, KMeansRandomCenters, &centers)
	if bestLabels.Empty() || bestLabels.Size()[0] != len(points) {
		t.Error("Labels is not proper")
	}
}

func TestMatLog(t *testing.T) {
	src := NewMatWithSize(4, 3, MatTypeCV32F)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Log(src, &dst)
	if dst.Empty() {
		t.Error("Log dst should not be empty.")
	}
}

func TestMatMagnitude(t *testing.T) {
	src1 := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src1.Close()
	src2 := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src2.Close()

	dst := NewMat()
	defer dst.Close()

	Magnitude(src1, src2, &dst)
	if dst.Empty() {
		t.Error("Magnitude dst should not be empty.")
	}
}

func TestMatMahalanobis(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	defer src.Close()

	RandU(&src, Scalar{Val1: -128}, Scalar{Val1: 128})

	icovar := NewMatWithSize(10, 10, MatTypeCV32F)
	defer icovar.Close()
	mean := NewMatWithSize(1, 10, MatTypeCV32F)
	defer mean.Close()

	CalcCovarMatrix(src, &icovar, &mean, CovarRows|CovarNormal, MatTypeCV32F)
	icovar.Inv()

	line1 := src.Row(0)
	defer line1.Close()
	line2 := src.Row(1)
	defer line2.Close()

	result := Mahalanobis(line1, line2, icovar)
	if result == 0 {
		t.Error("Mahalanobis result should not be empty.")
	}
}

func TestMulTransposed(t *testing.T) {
	src := Eye(10, 10, MatTypeCV32FC1)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	MulTransposed(src, &dst, true)
	if dst.Empty() {
		t.Error("MulTransposed dst should not be empty.")
	}
}

func TestMatMax(t *testing.T) {
	src1 := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src1.Close()
	src2 := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src2.Close()

	dst := NewMat()
	defer dst.Close()

	Max(src1, src2, &dst)
	if dst.Empty() {
		t.Error("Max dst should not be empty.")
	}
}

func TestMatMin(t *testing.T) {
	src1 := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src1.Close()
	src2 := NewMatWithSize(4, 4, MatTypeCV32F)
	defer src2.Close()

	dst := NewMat()
	defer dst.Close()

	Min(src1, src2, &dst)
	if dst.Empty() {
		t.Error("Min dst should not be empty.")
	}
}

func TestMatMinMaxIdx(t *testing.T) {
	src := NewMatWithSize(10, 10, MatTypeCV32F)
	defer src.Close()
	src.SetFloatAt(3, 3, 17)
	src.SetFloatAt(4, 4, 16)

	minVal, maxVal, _, _ := MinMaxIdx(src)

	if minVal != 0 {
		t.Error("TestMatMinMaxIdx minVal should be 0.")
	}
	if maxVal != 17 {
		t.Errorf("TestMatMinMaxIdx maxVal should be 17, was %f", maxVal)
	}
}

func TestMixChannels(t *testing.T) {
	bgra := NewMatWithSizeFromScalar(NewScalar(255, 0, 0, 255), 10, 10, MatTypeCV8UC4)
	defer bgra.Close()
	bgr := NewMatWithSize(bgra.Rows(), bgra.Cols(), MatTypeCV8UC3)
	defer bgr.Close()
	alpha := NewMatWithSize(bgra.Rows(), bgra.Cols(), MatTypeCV8UC1)
	defer alpha.Close()

	dst := []Mat{bgr, alpha}

	// bgra[0] -> bgr[2], bgra[1] -> bgr[1],
	// bgra[2] -> bgr[0], bgra[3] -> alpha[0]
	fromTo := []int{0, 2, 1, 1, 2, 0, 3, 3}

	MixChannels([]Mat{bgra}, dst, fromTo)

	bgrChans := Split(bgr)
	scalarByte := []byte{0, 0, 255}
	for c := 0; c < bgr.Channels(); c++ {
		for i := 0; i < bgr.Rows(); i++ {
			for j := 0; j < bgr.Cols(); j++ {
				if s := bgrChans[c].GetUCharAt(i, j); s != scalarByte[c] {
					t.Errorf("TestMixChannels incorrect bgr scalar: %v\n", s)
				}
			}
		}
	}

	for _, ch := range bgrChans {
		ch.Close()
	}

	alphaChans := Split(alpha)
	scalarByte = []byte{255}
	for c := 0; c < alpha.Channels(); c++ {
		for i := 0; i < alpha.Rows(); i++ {
			for j := 0; j < alpha.Cols(); j++ {
				if s := alphaChans[c].GetUCharAt(i, j); s != scalarByte[c] {
					t.Errorf("TestMixChannels incorrect alpha scalar: %v\n", s)
				}
			}
		}
	}
	for _, ch := range alphaChans {
		ch.Close()
	}
}

func TestGetVecfAt(t *testing.T) {
	var cases = []struct {
		m            Mat
		expectedSize int
	}{
		{NewMatWithSize(1, 1, MatTypeCV32FC1), 1},
		{NewMatWithSize(1, 1, MatTypeCV32FC2), 2},
		{NewMatWithSize(1, 1, MatTypeCV32FC3), 3},
		{NewMatWithSize(1, 1, MatTypeCV32FC4), 4},
	}

	for _, c := range cases {
		vec := c.m.GetVecfAt(0, 0)
		if len := len(vec); len != c.expectedSize {
			t.Errorf("TestGetVecfAt: expected %d, got: %d.", c.expectedSize, len)
		}
		c.m.Close()
	}
}

func TestGetVecdAt(t *testing.T) {
	var cases = []struct {
		m            Mat
		expectedSize int
	}{
		{NewMatWithSize(1, 1, MatTypeCV64FC1), 1},
		{NewMatWithSize(1, 1, MatTypeCV64FC2), 2},
		{NewMatWithSize(1, 1, MatTypeCV64FC3), 3},
		{NewMatWithSize(1, 1, MatTypeCV64FC4), 4},
	}

	for _, c := range cases {
		vec := c.m.GetVecdAt(0, 0)
		if len := len(vec); len != c.expectedSize {
			t.Errorf("TestGetVecdAt: expected %d, got: %d.", c.expectedSize, len)
		}
		c.m.Close()
	}
}

func TestGetVecbAt(t *testing.T) {
	var cases = []struct {
		m            Mat
		expectedSize int
	}{
		{NewMatWithSize(1, 1, MatTypeCV8UC1), 1},
		{NewMatWithSize(1, 1, MatTypeCV8UC2), 2},
		{NewMatWithSize(1, 1, MatTypeCV8UC3), 3},
		{NewMatWithSize(1, 1, MatTypeCV8UC4), 4},
	}

	for _, c := range cases {
		vec := c.m.GetVecbAt(0, 0)
		if len := len(vec); len != c.expectedSize {
			t.Errorf("TestGetVecbAt: expected %d, got: %d.", c.expectedSize, len)
		}
		c.m.Close()
	}
}

func TestGetVeciAt(t *testing.T) {
	var cases = []struct {
		m            Mat
		expectedSize int
	}{
		{NewMatWithSize(1, 1, MatTypeCV8UC1), 1},
		{NewMatWithSize(1, 1, MatTypeCV8UC2), 2},
		{NewMatWithSize(1, 1, MatTypeCV8UC3), 3},
		{NewMatWithSize(1, 1, MatTypeCV8UC4), 4},
	}

	for _, c := range cases {
		vec := c.m.GetVeciAt(0, 0)
		if len := len(vec); len != c.expectedSize {
			t.Errorf("TestGetVeciAt: expected %d, got: %d.", c.expectedSize, len)
		}
		c.m.Close()
	}
}

func TestGetTickFrequencyCount(t *testing.T) {
	freq := GetTickFrequency()
	if freq == 0 {
		t.Error("GetTickFrequency expected non zero.")
	}

	count := GetTickCount()
	if count == 0 {
		t.Error("GetTickCount expected non zero.")
	}
}

func TestMatT(t *testing.T) {
	var q = []float32{1, 3, 2, 4}
	src := NewMatWithSize(2, 2, MatTypeCV32F)
	defer src.Close()
	src.SetFloatAt(0, 0, 1)
	src.SetFloatAt(0, 1, 2)
	src.SetFloatAt(1, 0, 3)
	src.SetFloatAt(1, 1, 4)

	dst := src.T()
	defer dst.Close()

	ret, err := dst.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != q[i] {
			t.Errorf("MatT incorrect value: %v\n", ret[i])
		}
	}
}

func compareImages(img0, img1 image.Image) bool {
	bounds0 := img0.Bounds()
	bounds1 := img1.Bounds()
	dx0 := bounds0.Dx()
	dy0 := bounds0.Dy()
	if dx0 != bounds1.Dx() || dy0 != bounds1.Dy() {
		return false
	}
	xMin0 := bounds0.Min.X
	xMin1 := bounds1.Min.X
	yMin0 := bounds0.Min.Y
	yMin1 := bounds1.Min.Y
	for i := 0; i < dx0; i++ {
		for j := 0; j < dy0; j++ {
			point0 := img0.At(xMin0+i, yMin0+j)
			point1 := img1.At(xMin1+i, yMin1+j)
			r0, g0, b0, a0 := point0.RGBA()
			r1, g1, b1, a1 := point1.RGBA()
			r0 >>= 8
			g0 >>= 8
			b0 >>= 8
			a0 >>= 8
			r1 >>= 8
			g1 >>= 8
			b1 >>= 8
			a1 >>= 8
			if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
				return false
			}
		}
	}

	return true
}

func TestColRowRange(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	defer mat.Close()
	if mat.Empty() {
		t.Error("TestColRowRange should not be empty")
	}

	if mat.Rows() != 101 {
		t.Errorf("TestColRowRange incorrect row count: %v\n", mat.Rows())
	}

	if mat.Cols() != 102 {
		t.Errorf("TestColRowRange incorrect col count: %v\n", mat.Cols())
	}

	submatRow := mat.RowRange(0, 50)
	defer submatRow.Close()
	if submatRow.Rows() != 50 {
		t.Errorf("TestColRowRange incorrect submatRow count: %v\n", submatRow.Rows())
	}

	submatCols := mat.ColRange(0, 50)
	defer submatCols.Close()
	if submatCols.Cols() != 50 {
		t.Errorf("TestColRowRange incorrect submatCols count: %v\n", submatCols.Cols())
	}
}

func TestNormWithMats(t *testing.T) {
	mat1 := NewMatWithSize(100, 100, MatTypeCV8UC1)
	defer mat1.Close()

	mat2 := NewMatWithSize(100, 100, MatTypeCV8UC1)
	defer mat2.Close()

	d := NormWithMats(mat1, mat2, NormInf)
	if d != 0 {
		t.Fatal("expected 0")
	}
}

func Test_toGoStrings(t *testing.T) {
	goStrings := []string{"foo", "bar"}
	cStrings := toCStrings(goStrings)
	result := toGoStrings(cStrings)
	if len(goStrings) != len(result) {
		t.Errorf("TesttoGoStrings failed: length of converted string is not equal to original \n")
	}
	for i, s := range goStrings {
		if s != result[i] {
			t.Errorf("TesttoGoStrings failed: strings are not equal. expected=%s, actusal=%s", s, result[i])
		}
	}
}

func TestTheRNG(t *testing.T) {
	rng := TheRNG()
	if rng.p == nil {
		t.Errorf("got no rng")
	}
}

func TestSetRNGSeed(t *testing.T) {
	SetRNGSeed(123)
}

func TestRNG_Fill(t *testing.T) {
	rng := TheRNG()
	mat := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer mat.Close()
	rng.Fill(&mat, RNGDistNormal, 10, 20, false)
}

func TestRNG_Gaussian(t *testing.T) {
	rng := TheRNG()
	_ = rng.Gaussian(0.5)
}

func TestRNG_Next(t *testing.T) {
	rng := TheRNG()
	_ = rng.Next()
}

func TestRandN(t *testing.T) {
	mat := NewMatWithSize(5, 5, MatTypeCV8UC3)
	defer mat.Close()
	RandN(&mat, NewScalar(10, 10, 10, 10), NewScalar(20, 20, 20, 20))
}

func TestRandShuffle(t *testing.T) {
	mat := NewMatWithSize(5, 5, MatTypeCV8UC3)
	defer mat.Close()
	RandShuffle(&mat)
}

func TestRandShuffleWithParams(t *testing.T) {
	mat := NewMatWithSize(5, 5, MatTypeCV8UC3)
	defer mat.Close()
	RandShuffleWithParams(&mat, 1, TheRNG())
}

func TestRandU(t *testing.T) {
	mat := NewMatWithSize(5, 5, MatTypeCV8UC3)
	defer mat.Close()
	RandU(&mat, NewScalar(10, 10, 10, 10), NewScalar(20, 20, 20, 20))
}

func TestNewPointsVector(t *testing.T) {
	epv := NewPointsVector()
	defer epv.Close()

	if epv.Size() != 0 {
		t.Fatal("expected empty pointsvector size not 0")
	}

	pts := [][]image.Point{
		{
			image.Pt(10, 10),
			image.Pt(10, 20),
			image.Pt(20, 20),
			image.Pt(20, 10),
		},
	}

	psv := NewPointsVectorFromPoints(pts)
	defer psv.Close()

	if psv.IsNil() {
		t.Fatal("pointsvector pointer was nil")
	}

	if psv.Size() != 1 {
		t.Fatal("expected pointsvector size 1")
	}

	ipv := psv.At(10)
	if !ipv.IsNil() {
		t.Fatal("expected pointvector nil")
	}

	pv := psv.At(0)
	if pv.Size() != 4 {
		t.Fatal("expected pointvector size 4")
	}

	p := pv.At(0)
	if p != image.Pt(10, 10) {
		t.Fatal("invalid At() point")
	}

	p = pv.At(10)
	if p != image.Pt(0, 0) {
		t.Fatal("invalid At() point beyond range")
	}

	out := psv.ToPoints()
	if out[0][0] != image.Pt(10, 10) {
		t.Fatal("invalid ToPoints() point")
	}

	ps := []image.Point{
		image.Pt(10, 10),
		image.Pt(10, 20),
		image.Pt(20, 20),
		image.Pt(20, 10),
	}

	apv := NewPointVectorFromPoints(ps)
	defer apv.Close()

	psv.Append(apv)
	if psv.Size() != 2 {
		t.Fatal("unable to append to PointsVector")
	}
}

func TestNewPointVector(t *testing.T) {
	epv := NewPointVector()
	defer epv.Close()

	if epv.Size() != 0 {
		t.Fatal("expected empty pointvector size not 0")
	}

	pts := []image.Point{
		image.Pt(10, 10),
		image.Pt(10, 20),
		image.Pt(20, 20),
		image.Pt(20, 10),
	}

	pv := NewPointVectorFromPoints(pts)
	defer pv.Close()

	if pv.IsNil() {
		t.Fatal("pointvector pointer was nil")
	}

	if pv.Size() != 4 {
		t.Fatal("expected pointvector size 4")
	}

	p := pv.At(0)
	if p != image.Pt(10, 10) {
		t.Fatal("invalid point")
	}

	np := image.Pt(50, 50)

	pv.Append(np)
	if pv.Size() != 5 {
		t.Fatal("unable to append to PointVector")
	}

	mat := NewMatWithSize(4, 1, MatTypeCV32SC2)
	defer mat.Close()

	pvm := NewPointVectorFromMat(mat)
	defer pvm.Close()

	if pvm.Size() != 4 {
		t.Fatalf("expected size of NewPointVectorFromMat to be 4, was %d", pvm.Size())
	}
}

func TestNewPoint2fVector(t *testing.T) {
	epv := NewPoint2fVector()
	defer epv.Close()

	if epv.Size() != 0 {
		t.Fatal("expected empty pointvector size not 0")
	}

	pts := []Point2f{
		{10.0, 10.0},
		{10.0, 20.0},
		{20.5, 21.5},
		{25.5, 30.5},
	}

	pv := NewPoint2fVectorFromPoints(pts)
	defer pv.Close()

	if pv.IsNil() {
		t.Fatal("point2fvector pointer was nil")
	}

	if pv.Size() != 4 {
		t.Fatal("expected point2fvector size 4")
	}

	p := pv.At(0)
	want := Point2f{10.0, 10.0}
	if p != want {
		t.Fatal("invalid point")
	}

	p = pv.At(10)
	nopoint := Point2f{0, 0}
	if p != nopoint {
		t.Fatal("invalid At() point beyond range")
	}

	out := pv.ToPoints()
	if len(out) != 4 && out[0] != want {
		t.Fatal("invalid ToPoints()")
	}

	mat := NewMatWithSize(4, 1, MatTypeCV32FC2)
	defer mat.Close()

	pvm := NewPoint2fVectorFromMat(mat)
	defer pvm.Close()

	if pvm.Size() != 4 {
		t.Fatalf("expected size of NewPoint2fVectorFromMat to be 4, was %d", pvm.Size())
	}
}

func TestNewPoints2fVector(t *testing.T) {
	epv := NewPoints2fVector()
	defer epv.Close()

	if epv.Size() != 0 {
		t.Fatal("expected empty points2fvector size not 0")
	}

	pts := [][]Point2f{
		{
			NewPoint2f(10.0, 10.0),
			NewPoint2f(10.0, 20.0),
			NewPoint2f(20.0, 20.0),
			NewPoint2f(20.0, 10.0),
		},
	}

	psv := NewPoints2fVectorFromPoints(pts)
	defer psv.Close()

	if psv.IsNil() {
		t.Fatal("points2fvector pointer was nil")
	}

	if psv.Size() != 1 {
		t.Fatal("expected points2fvector size 1")
	}

	ipv := psv.At(10)
	if !ipv.IsNil() {
		t.Fatal("expected pointvector nil")
	}

	pv := psv.At(0)
	if pv.Size() != 4 {
		t.Fatal("expected pointvector size 4")
	}

	p := pv.At(0)
	if p != NewPoint2f(10.0, 10.0) {
		t.Fatal("invalid At() point")
	}

	p = pv.At(10)
	if p != NewPoint2f(0, 0) {
		t.Fatal("invalid At() point beyond range")
	}

	out := psv.ToPoints()
	if out[0][0] != NewPoint2f(10.0, 10.0) {
		t.Fatal("invalid ToPoints() point")
	}

	ps := []Point2f{
		NewPoint2f(10, 10),
		NewPoint2f(10, 20),
		NewPoint2f(20, 20),
		NewPoint2f(20, 10),
	}

	apv := NewPoint2fVectorFromPoints(ps)
	defer apv.Close()

	psv.Append(apv)
	if psv.Size() != 2 {
		t.Fatal("unable to append to Points2fVector")
	}
}

func TestNewPoint3fVector(t *testing.T) {
	epv := NewPoint3fVector()
	defer epv.Close()

	if epv.Size() != 0 {
		t.Fatal("expected empty pointvector size not 0")
	}

	pts := []Point3f{
		{10.0, 10.0, 0.1},
		{10.0, 20.0, 1.0},
		{20.5, 21.5, 2.0},
	}

	pv := NewPoint3fVectorFromPoints(pts)
	defer pv.Close()

	pv.Append(NewPoint3f(25.5, 30.5, 3.0))

	if pv.IsNil() {
		t.Fatal("point3fvector pointer was nil")
	}

	if pv.Size() != 4 {
		t.Fatal("expected point3fvector size 4")
	}

	p := pv.At(0)
	want := Point3f{10.0, 10.0, 0.1}
	if p != want {
		t.Fatal("invalid point")
	}

	want2 := NewPoint3f(25.5, 30.5, 3.0)
	if pv.At(3) != want2 {
		t.Fatal("fail to append point to Point3fVector")
	}

	p = pv.At(10)
	nopoint := Point3f{0, 0, 0}
	if p != nopoint {
		t.Fatal("invalid At() point beyond range")
	}

	out := pv.ToPoints()
	if len(out) != 4 && out[0] != want {
		t.Fatal("invalid ToPoints()")
	}

	mat := NewMatWithSize(4, 1, MatTypeCV32FC3)
	defer mat.Close()

	pvm := NewPoint3fVectorFromMat(mat)
	defer pvm.Close()

	if pvm.Size() != 4 {
		t.Fatalf("expected size of NewPoint3fVectorFromMat to be 4, was %d", pvm.Size())
	}
}

func TestNewPoints3fVector(t *testing.T) {
	epv := NewPoints3fVector()
	defer epv.Close()

	if epv.Size() != 0 {
		t.Fatal("expected empty points3fvector size not 0")
	}

	pts := [][]Point3f{
		{
			NewPoint3f(10.0, 10.0, 0.1),
			NewPoint3f(10.0, 20.0, 0.2),
			NewPoint3f(20.0, 20.0, 0.3),
			NewPoint3f(20.0, 10.0, 0.4),
		},
	}

	psv := NewPoints3fVectorFromPoints(pts)
	defer psv.Close()

	if psv.IsNil() {
		t.Fatal("points3fvector pointer was nil")
	}

	if psv.Size() != 1 {
		t.Fatal("expected points3fvector size 1")
	}

	ipv := psv.At(10)
	if !ipv.IsNil() {
		t.Fatal("expected pointvector nil")
	}

	pv := psv.At(0)
	if pv.Size() != 4 {
		t.Fatal("expected pointvector size 4")
	}

	p := pv.At(0)
	if p != NewPoint3f(10.0, 10.0, 0.1) {
		t.Fatal("invalid At() point")
	}

	p = pv.At(10)
	if p != NewPoint3f(0, 0, 0) {
		t.Fatal("invalid At() point beyond range")
	}

	out := psv.ToPoints()
	if out[0][0] != NewPoint3f(10.0, 10.0, 0.1) {
		t.Fatal("invalid ToPoints() point")
	}

	ps := []Point3f{
		NewPoint3f(10, 10, 0.1),
		NewPoint3f(10, 20, 0.2),
		NewPoint3f(20, 20, 0.3),
		NewPoint3f(20, 10, 0.4),
	}

	apv := NewPoint3fVectorFromPoints(ps)
	defer apv.Close()

	psv.Append(apv)
	if psv.Size() != 2 {
		t.Fatal("unable to append to Points3fVector")
	}
}

func TestElemSize(t *testing.T) {
	m1 := NewMat()
	defer m1.Close()
	if m1.ElemSize() != 0 {
		t.Error("incorrect element size")
	}

	m2 := NewMatWithSize(2, 2, MatTypeCV16S)
	defer m2.Close()
	if m2.ElemSize() != 2 {
		t.Error("incorrect element size of MatTypeCV16S")
	}

	m3 := NewMatWithSize(2, 2, MatTypeCV16SC3)
	defer m3.Close()
	if m3.ElemSize() != 6 {
		t.Error("incorrect element size of MatTypeCV16SC3")
	}

	m4 := NewMatWithSize(2, 2, MatTypeCV32SC4)
	defer m4.Close()
	if m4.ElemSize() != 16 {
		t.Error("incorrect element size of MatTypeCV32SC4")
		return
	}
}

func TestSetThreadNumber(t *testing.T) {
	original := GetNumThreads()

	SetNumThreads(-1)
	if num := GetNumThreads(); num != original {
		t.Errorf("incorrect number of threads, got %d, want %d", num, original)
	}

	SetNumThreads(0)
	if num := GetNumThreads(); num < 1 {
		t.Errorf("incorrect number of threads, got %d, want at least 1", num)
	}

	SetNumThreads(1)
	if num := GetNumThreads(); num != 1 {
		t.Errorf("incorrect number of threads, got %d, want %d", num, 1)
	}

	SetNumThreads(original)
}

func TestMinMaxLoc(t *testing.T) {
	input := NewMatWithSize(2, 2, MatTypeCV32F)
	defer input.Close()
	input.SetFloatAt(0, 0, 1)
	input.SetFloatAt(0, 1, 2)
	input.SetFloatAt(1, 0, 3)
	input.SetFloatAt(1, 1, 4)
	minVal, maxVal, minLoc, maxLoc := MinMaxLoc(input)

	wantMinVal, wantMaxValue := float32(1.0), float32(4.0)
	if minVal != wantMinVal {
		t.Errorf("minVal got: %v, want %v", minVal, wantMinVal)
	}
	if maxVal != wantMaxValue {
		t.Errorf("maxVal got: %v, want %v", maxVal, wantMaxValue)
	}
	wantMinLoc, wantMaxLoc := image.Point{Y: 0, X: 0}, image.Point{Y: 1, X: 1}
	if minLoc != wantMinLoc {
		t.Errorf("minLoc got: %v, want %v", minLoc, wantMinLoc)
	}
	if maxLoc != wantMaxLoc {
		t.Errorf("maxLoc got: %v, want %v", maxLoc, wantMaxLoc)
	}
}

func TestMinMaxLocWithMask(t *testing.T) {
	input := NewMatWithSize(2, 2, MatTypeCV32F)
	defer input.Close()
	input.SetFloatAt(0, 0, 1)
	input.SetFloatAt(0, 1, 2)
	input.SetFloatAt(1, 0, 3)
	input.SetFloatAt(1, 1, 4)
	mask := NewMatWithSize(2, 2, MatTypeCV8U)
	defer mask.Close()
	mask.SetUCharAt(1, 0, 1)
	mask.SetUCharAt(1, 1, 1)
	minVal, maxVal, minLoc, maxLoc := MinMaxLocWithMask(input, mask)

	wantMinVal, wantMaxValue := float32(3.0), float32(4.0)
	if minVal != wantMinVal {
		t.Errorf("minVal got: %v, want %v", minVal, wantMinVal)
	}
	if maxVal != wantMaxValue {
		t.Errorf("maxVal got: %v, want %v", maxVal, wantMaxValue)
	}
	wantMinLoc, wantMaxLoc := image.Point{Y: 1, X: 0}, image.Point{Y: 1, X: 1}
	if minLoc != wantMinLoc {
		t.Errorf("minLoc got: %v, want %v", minLoc, wantMinLoc)
	}
	if maxLoc != wantMaxLoc {
		t.Errorf("maxLoc got: %v, want %v", maxLoc, wantMaxLoc)
	}
}

func TestNewRotatedRect(t *testing.T) {

	rr := NewRotatedRect(image.Pt(1, 1), 10, 10, 75.0)
	if rr.Angle != 75.0 {
		t.Errorf("NewRotatedRect not working as intended")
	}

}

func TestNewRotatedRect2f(t *testing.T) {

	pts := Point2f{
		X: 1.5,
		Y: 1.5,
	}

	rr := NewRotatedRect2f(pts, 10.5, 10.5, 75.0)
	if rr.Angle != 75.0 {
		t.Errorf("NewRotatedRect not working as intended")
	}

}

func TestNewMatFromPointVector(t *testing.T) {

	img := NewMatWithSize(320, 200, MatTypeCV32SC1)
	defer img.Close()

	size := img.Size()

	points := []image.Point{
		image.Pt(0, 0), image.Pt(0, size[0]-1),
		image.Pt(size[1]-1, size[0]-1),
		image.Pt(size[1]-1, 0),
	}

	pv := NewPointVectorFromPoints(points)
	defer pv.Close()

	m := NewMatFromPointVector(pv, false)
	defer m.Close()

	if m.Empty() {
		t.Error("Mat shlould not be empty")
	}

}

func TestNewMatFromPoint2fVector(t *testing.T) {

	pv2f := NewPoint2fVectorFromPoints([]Point2f{NewPoint2f(1.1, 2.2)})
	defer pv2f.Close()

	m := NewMatFromPoint2fVector(pv2f, false)
	defer m.Close()

	if m.Empty() {
		t.Error("Mat shlould not be empty")
	}

}
