package numerics

import (
	"fmt"
	"testing"
)

func TestNewIntArray(t *testing.T) {
	a := NewArray("[1, 2, 3, 4, 5]", IntType)
	if a.Len() != 5 {
		t.Errorf("Expected length 5, got %d", a.Len())
	}
	if a.Get(0).(int64) != 1 {
		t.Errorf("Expected 1, got %d", a.Get(0).(int64))
	}
	if a.Get(4).(int64) != 5 {
		t.Errorf("Expected 5, got %d", a.Get(4).(int64))
	}
	if a.IsOfType(IntType) == false {
		t.Error("Expected true, got false")
	}
	a.Set(0, 10)
	if a.Get(0).(int64) != 10 {
		t.Errorf("Expected 10, got %d", a.Get(0).(int64))
	}

}
func TestNewFloatArray(t *testing.T) {
	a := NewArray("[1.1, 2.2, 3.3, 4.4, 5.5]", FloatType)
	if a.Len() != 5 {
		t.Errorf("Expected length 5, got %d", a.Len())
	}
	if a.Get(0).(float64) != 1.1 {
		t.Errorf("Expected 1.1, got %f", a.Get(0).(float64))
	}
	if a.Get(4).(float64) != 5.5 {
		t.Errorf("Expected 5.5, got %f", a.Get(4).(float64))
	}
	if a.IsOfType(FloatType) == false {
		t.Error("Expected true, got false")
	}
	a.Set(0, 10.1)
	if a.Get(0).(float64) != 10.1 {
		t.Errorf("Expected 10.1, got %f", a.Get(0).(float64))
	}
}
func TestNewBoolArray(t *testing.T) {
	a := NewArray("[true, false, true, false, true]", BoolType)
	if a.Len() != 5 {
		t.Errorf("Expected length 5, got %d", a.Len())
	}
	if a.Get(0).(bool) != true {
		t.Errorf("Expected true, got %t", a.Get(0).(bool))
	}
	if a.Get(4).(bool) != true {
		t.Errorf("Expected true, got %t", a.Get(4).(bool))
	}
	if a.IsOfType(BoolType) == false {
		t.Error("Expected true, got false")
	}
	a.Set(0, false)
	if a.Get(0).(bool) != false {
		t.Errorf("Expected false, got %t", a.Get(0).(bool))
	}

}
func TestNewNestedArray(t *testing.T) {
	a := NewArray("[[1, 2], [3, 4], [5, 6]]", IntType)
	if a.Len() != 3 {
		t.Errorf("Expected length 3, got %d", a.Len())
	}
	if len(a.Get(0).([]Element)) != 2 {
		t.Errorf("Expected length 2, got %d", len(a.Get(0).([]Element)))
	}
	if a.Get(0).([]Element)[0].Get().(int64) != 1 {
		t.Errorf("Expected 1, got %d", a.Get(0).([]Element)[0].Get().(int64))
	}
	if a.IsOfType(ArrayType) == false {
		t.Error("Expected true, got false , type is ", a.GetType())
	}
	a.Set(0, NewArray("[10, 20]", IntType))
	if a.Get(0).([]Element)[0].Get().(int64) != 10 {
		t.Errorf("Expected 10, got %d", a.Get(0).([]Element)[0].Get().(int64))
	}
}

func TestArrayMax(t *testing.T) {
	a := NewArray("[1.1, 2.2, 3.3, 4.4, 5.5]", FloatType)
	if a.Max() != 5.5 {
		t.Errorf("Expected 5.5, got %f", a.Max())
	}
}

func TestArrayMin(t *testing.T) {
	a := NewArray("[1.1, 2.2, 3.3, 4.4, 5.5]", FloatType)
	if a.Min() != 1.1 {
		t.Errorf("Expected 1.1, got %f", a.Min())
	}
}

func TestArrayCopy(t *testing.T) {
	a := NewArray("[1.1, 2.2, 3.3, 4.4, 5.5]", FloatType)
	b := a.Copy()
	if a.Len() != b.Len() {
		t.Errorf("Expected length %d, got %d", a.Len(), b.Len())
	}
	if a.Get(0).(float64) != b.Get(0).(float64) {
		t.Errorf("Expected %f, got %f", a.Get(0).(float64), b.Get(0).(float64))
	}
}

func TestLinspace(t *testing.T) {
	a := Linspace(1, 10, 10)
	if a.Len() != 10 {
		t.Errorf("Expected length 10, got %d", a.Len())
	}
	if a.Get(0).(float64) != 1 {
		t.Errorf("Expected 1, got %f", a.Get(0).(float64))
	}
	if a.Get(9).(float64) != 10 {
		t.Errorf("Expected 10, got %f", a.Get(9).(float64))
	}
}
func TestRandom(t *testing.T) {
	a := Random(1, 10, 10)
	if a.Len() != 10 {
		t.Errorf("Expected length 10, got %d", a.Len())
	}
	if a.Get(0).(int64) < 1 || a.Get(0).(int64) > 10 {
		t.Errorf("Expected between 1 and 10, got %d", a.Get(0).(int64))
	}
}

func TestZeros(t *testing.T) {
	a := Zeros(3, 3)
	if a.Len() != 3 {
		t.Errorf("Expected length 3, got %d", a.Len())
	}

}

func TestNewMatrix(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4.3, 5.2, 6]", FloatType)
	row3 := NewArray("[true, false, true]", BoolType)
	row4 := NewArray("[1,2,3]", IntType)
	m, err := NewMatrix(4, 3, []Array{row1, row2, row3, row4})
	if err != nil {
		t.Error(err)
	}
	if m.RowNum() != 4 {
		t.Errorf("Expected 4 rows, got %d", m.RowNum())
	}
	if m.ColNum() != 3 {
		t.Errorf("Expected 3 columns, got %d", m.ColNum())
	}
	if m.Get(0, 0) != 1.0 {
		t.Errorf("Expected 1, got %f", m.Get(0, 0))
	}
	if m.Get(1, 0) != 4.3 {
		t.Errorf("Expected 4.3, got %f", m.Get(1, 0))
	}
}
func TestGetRows(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4.3, 5.2, 6]", FloatType)
	row3 := NewArray("[true, false, true]", BoolType)
	row4 := NewArray("[1,2,3]", IntType)
	m, err := NewMatrix(4, 3, []Array{row1, row2, row3, row4})
	if err != nil {
		t.Error(err)
	}
	rows := m.GetRows(1, 2)
	if rows.RowNum() != 1 {
		t.Errorf("Expected 2 rows, got %d", rows.RowNum())
	}
	if rows.ColNum() != 3 {
		t.Errorf("Expected 4 columns, got %d", rows.ColNum())
	}
	row := rows.GetRow(0)
	if row.Get(0, 0) != 4.3 {
		t.Errorf("Expected 4.3, got %f", row.Get(0, 0))
	}
}

func TestGetCols(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4.3, 5.2, 6]", FloatType)
	row3 := NewArray("[true, false, true]", BoolType)
	row4 := NewArray("[1,2,3]", IntType)
	m, err := NewMatrix(4, 3, []Array{row1, row2, row3, row4})
	if err != nil {
		t.Error(err)
	}
	cols := m.GetColumns(1, 2)
	if cols.RowNum() != 4 {
		t.Errorf("Expected 4 rows, got %d", cols.RowNum())
	}
	if cols.ColNum() != 1 {
		t.Errorf("Expected 2 columns, got %d", cols.ColNum())
	}
	col := cols.GetColumn(0)
	if col.Get(0, 0) != 2.0 {
		t.Errorf("Expected 2.0, got %f", col.Get(0, 0))
	}
}

func TestTranspose(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4.3, 5.2, 6]", FloatType)
	row3 := NewArray("[true, false, true]", BoolType)
	row4 := NewArray("[1,2,3]", IntType)
	m, _ := NewMatrix(4, 3, []Array{row1, row2, row3, row4})
	row1 = NewArray("[1, 4.3, 1, 1]", FloatType)
	row2 = NewArray("[2, 5.2, 0, 2]", FloatType)
	row3 = NewArray("[3, 6, 1, 3]", FloatType)
	mT, _ := NewMatrix(3, 4, []Array{row1, row2, row3})
	if !Equal(m.Transpose(), mT) {
		t.Error("Transpose not equal")
	}
}

func TestAppendX(t *testing.T) {
	rows1 := NewArray("[1, 2, 3]", IntType)
	rows2 := NewArray("[4,5,6]", FloatType)
	rows3 := NewArray("[7,8,9]", IntType)

	x, _ := NewMatrix(3, 3, []Array{rows1, rows2, rows3})
	columnOfOnes, _ := NewMatrix(3, 1, []Array{NewArray("[1]", IntType), NewArray("[1]", IntType), NewArray("[1]", IntType)})
	x, _ = AppendX(columnOfOnes, x)

	rows1 = NewArray("[1,1, 2, 3]", IntType)
	rows2 = NewArray("[1,4,5,6]", FloatType)
	rows3 = NewArray("[1,7,8,9]", IntType)
	expectedX, _ := NewMatrix(3, 4, []Array{rows1, rows2, rows3})
	if !Equal(x, expectedX) {
		t.Error("AppendX not equal")
	}

}

func TestAppendY(t *testing.T) {
	rows1 := NewArray("[1, 2, 3]", IntType)
	rows2 := NewArray("[4,5,6]", FloatType)
	rows3 := NewArray("[7,8,9]", IntType)

	x, _ := NewMatrix(3, 3, []Array{rows1, rows2, rows3})
	rowOfOnes, _ := NewMatrix(1, 3, []Array{NewArray("[1,1,1]", IntType)})
	x, _ = AppendY(rowOfOnes, x)

	rows1 = NewArray("[1, 2, 3]", IntType)
	rows2 = NewArray("[4,5,6]", FloatType)
	rows3 = NewArray("[7,8,9]", IntType)
	rows4 := NewArray("[1,1,1]", IntType)
	expectedY, _ := NewMatrix(4, 3, []Array{rows4, rows1, rows2, rows3})
	if !Equal(x, expectedY) {
		t.Error("AppendY not equal")
	}

}

func TestMultiply(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4,5,6]", FloatType)
	row3 := NewArray("[7,8,9]", IntType)
	x, _ := NewMatrix(3, 3, []Array{row1, row2, row3})
	row1 = NewArray("[2,2 ,2 ]", IntType)
	row2 = NewArray("[3,3,3]", FloatType)
	row3 = NewArray("[4,4,4]", IntType)
	y, _ := NewMatrix(3, 3, []Array{row1, row2, row3})
	row1 = NewArray("[20, 20, 20]", FloatType)
	row2 = NewArray("[47, 47, 47]", FloatType)
	row3 = NewArray("[74, 74, 74]", FloatType)
	expected, _ := NewMatrix(3, 3, []Array{row1, row2, row3})
	result, _ := Multiply(x, y)
	if !Equal(result, expected) {
		t.Error("Multiplication not equal")
	}
}

func TestInverse(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4,5,7]", FloatType)
	row3 := NewArray("[8,9,12]", IntType)
	x, _ := NewMatrix(3, 3, []Array{row1, row2, row3})
	row1 = NewArray("[-3, 3, -1]", FloatType)
	row2 = NewArray("[8,-12,5]", FloatType)
	row3 = NewArray("[-4,7,-3]", FloatType)
	Ix, _ := NewMatrix(3, 3, []Array{row1, row2, row3})
	result, _ := x.Inverse()
	if !Equal(result, Ix) {
		fmt.Println(result.String())
		fmt.Println(Ix.String())
		t.Error("Inverse not equal")
	}

}

func TestDeterminant(t *testing.T) {
	row1 := NewArray("[1, 2, 3]", IntType)
	row2 := NewArray("[4,5,7]", FloatType)
	row3 := NewArray("[8,9,12]", IntType)
	x, _ := NewMatrix(3, 3, []Array{row1, row2, row3})
	d, _ := x.Det()
	abs := func(a float64) float64 {
		if a < 0 {
			return -a
		}
		return a
	}
	if func(a, b float64) bool {
		if abs(a-b) < 0.000001 {
			return false
		} else {
			return true
		}
	}(d, 1.000000) {
		t.Errorf("Expected %v, got %v", 1.0, d)
	}

}

func TestMatrix(t *testing.T) {
	t.Run("TestNewMatrix", TestNewMatrix)
	t.Run("TestGetRows", TestGetRows)
	t.Run("TestGetCols", TestGetCols)
	t.Run("TestTranspose", TestTranspose)
	t.Run("TestAppendX", TestAppendX)
	t.Run("TestAppendY", TestAppendY)
	t.Run("TestMultiply", TestMultiply)
	t.Run("TestInverse", TestInverse)
	t.Run("TestDeterminant", TestDeterminant)

}

func TestNumerics(t *testing.T) {
	t.Run("TestNewIntArray", TestNewIntArray)
	t.Run("TestNewFloatArray", TestNewFloatArray)
	t.Run("TestNewBoolArray", TestNewBoolArray)
	t.Run("TestNewNestedArray", TestNewNestedArray)
	t.Run("TestArrayMax", TestArrayMax)
	t.Run("TestArrayMin", TestArrayMin)
	t.Run("TestArrayCopy", TestArrayCopy)
	t.Run("TestLinspace", TestLinspace)
	t.Run("TestRandom", TestRandom)
	t.Run("TestZeros", TestZeros)

}
