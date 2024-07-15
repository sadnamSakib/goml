package tabular

import (
	"fmt"
	"testing"
)

func testNewSeries(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}
	_, err := NewSeries(1, 2, 'a')
	if err == nil {
		t.Error("Expected Type Mismatch Error, got nil")
	}
}

func testAppend(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	s.Append(nil)
	if s.Len() != 4 {
		t.Errorf("Expected length 4, got %d", s.Len())
	}
	err := s.Append("a")
	if err == nil {
		t.Error("Expected Type Mismatch Error, got nil")

	}
}

func testString(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	if s.String() != "1, 2, 3" {
		t.Errorf("Expected '1, 2, 3', got %s", s.String())
	}
}

func testLen(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}
}

func testSort(t *testing.T) {
	s, _ := NewSeries(3, 2, 1)
	s.Sort()
	if s.String() != "1, 2, 3" {
		t.Errorf("Expected '1, 2, 3', got %s", s.String())
	}
	s, _ = NewSeries(3.1, 2.1, 1.1, nil)
	s.Sort()
	if s.String() != "<nil>, 1.1, 2.1, 3.1" {
		t.Errorf("Expected '<nil>, 1.1, 2.1, 3.1', got %s", s.String())
	}
	s, _ = NewSeries("c", "b", "a")
	s.Sort()
	if s.String() != "a, b, c" {
		t.Errorf("Expected 'a, b, c', got %s", s.String())
	}
	s, _ = NewSeries(true, false, true)
	s.Sort()
	if s.String() != "false, true, true" {
		t.Errorf("Expected 'false, true, true', got %s", s.String())
	}
}

func TestSortCopy(t *testing.T) {
	s, _ := NewSeries(3, 2, 1)
	v := s.SortCopy()
	if v.String() != "1, 2, 3" {
		t.Errorf("Expected '1, 2, 3', got %s", v.String())
	}
	if s.String() != "3, 2, 1" {
		t.Errorf("Expected '3, 2, 1', got %s. Original series changed.", s.String())
	}

}

func TestSum(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)

	sum, err := s.Sum()

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 6.0 {
		t.Errorf("Expected 6, got %v", sum)
	}

	s, _ = NewSeries(1, 2, nil)
	_, err = s.Sum()
	if err != nil {
		t.Errorf("Expected no error. Got %v.", err)
	}
	s, _ = NewSeries(1, 2, 3.1)
	sum, err = s.Sum()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 6.1 {
		t.Errorf("Expected 6.1, got %v", sum)
	}
	s, _ = NewSeries(1, 2, 3.1, nil)
	sum, err = s.Sum()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 6.1 {
		t.Errorf("Expected 6.1, got %v", sum)
	}
}

func TestMean(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	mean, err := s.Mean()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if mean != 2.0 {
		t.Errorf("Expected 2, got %v", mean)
	}
	s, _ = NewSeries(1, 2, "a")
	_, err = s.Mean()
	if err == nil {
		t.Error("Expected Type Mismatch Error, got nil")
	}
	s, _ = NewSeries(1, 2, nil)
	mean, err = s.Mean()
	if err != nil {
		t.Errorf("Expected no error. Got %v.", err)
	}
	if mean != 1.0 {
		t.Errorf("Expected 1, got %v", mean)

	}
	s, _ = NewSeries(1, 2, 3.1)
	mean, err = s.Mean()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if mean != 2.033333333333333 {
		t.Errorf("Expected 2.033333333333333, got %v", mean)
	}
	s, _ = NewSeries(1, 2, 3.1, nil)
	mean, err = s.Mean()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if mean != 1.525 {
		t.Errorf("Expected 1.525, got %v", mean)
	}
	s, _ = NewSeries("a", "b")
	_, err = s.Mean()
	if err == nil {
		t.Error("Expected Unsupported Type Error, got nil")
	}
}

func TestConcat(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	if s.Concat() != "123" {
		t.Errorf("Expected '123', got %s", s.Concat())
	}
	s, _ = NewSeries("1", "2", "a")
	if s.Concat() != "12a" {
		t.Errorf("Expected '12a', got %s", s.Concat())
	}

}

func TestCount(t *testing.T) {
	s, _ := NewSeries(1, 2, 3, 3)
	if s.Count(1) != 1 {
		t.Errorf("Expected 1, got %d", s.Count(1))
	}
	if s.Count(4) != 0 {
		t.Errorf("Expected 0, got %d", s.Count(4))
	}
	if s.Count(3) != 2 {
		t.Errorf("Expected 2, got %d", s.Count(3))
	}

}

func TestContains(t *testing.T) {
	s, _ := NewSeries(1, 2, 3, 3)
	if !s.Contains(1) {
		t.Errorf("Expected true, got false")
	}
	if s.Contains(4) {
		t.Errorf("Expected false, got true")
	}
	if !s.Contains(3) {
		t.Errorf("Expected true, got false")
	}

}

func TestMinMax(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	min, err := s.Min()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if min != 1.0 {
		t.Errorf("Expected 1, got %v", min)
	}
	max, err := s.Max()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if max != 3.0 {
		t.Errorf("Expected 3, got %v", max)
	}
	s, _ = NewSeries(1, 2, nil)
	min, err = s.Min()
	if err != nil {
		t.Errorf("Expected no error. Got %v.", err)
	}
	if min != nil {
		t.Errorf("Expected nil, got %v", min)
	}
	max, err = s.Max()
	if err != nil {
		t.Errorf("Expected no error. Got %v.", err)
	}
	if max != 2.0 {
		t.Errorf("Expected 2, got %v", max)
	}
	s, _ = NewSeries(1, 2, 3.1)
	min, err = s.Min()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if min != 1.0 {
		t.Errorf("Expected 1.0, got %v", min)
	}
	max, err = s.Max()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if max != 3.1 {
		t.Errorf("Expected 3.1, got %v", max)
	}
	s, _ = NewSeries(1, 2, 3.1, nil)
	min, err = s.Min()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if min != nil {
		t.Errorf("Expected 1, got %v", min)
	}
	_, err = s.Max()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)

	}
}

func TestUnique(t *testing.T) {
	s, _ := NewSeries(1, 2, 3, 3)
	if s.Unique().String() != "1, 2, 3" {
		t.Errorf("Expected '1, 2, 3', got %s", s.Unique().String())
	}
}

func TestFilter(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	s.Filter(func(x interface{}) bool {
		return x.(float64) > 1.0
	})
	if s.String() != "2, 3" {
		t.Errorf("Expected '2, 3', got %s", s.String())
	}

}

func TestApply(t *testing.T) {
	s, _ := NewSeries(1, 2, 3)
	s.Apply(func(x interface{}) interface{} {
		return x.(float64) * 2.0
	})
	if s.String() != "2, 4, 6" {
		t.Errorf("Expected '2, 4, 6', got %s", s.String())
	}
}

func TestReadCSV(t *testing.T) {
	df, err := Read_CSV("people.csv", true)

	fmt.Println(df.String())
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

}

func TestSeries(t *testing.T) {
	t.Run("NewSeries", testNewSeries)
	t.Run("Append", testAppend)
	t.Run("String", testString)
	t.Run("Len", testLen)
	t.Run("Sort", testSort)
	t.Run("SortCopy", TestSortCopy)
	t.Run("Sum", TestSum)
	t.Run("Mean", TestMean)
	t.Run("Concat", TestConcat)
	t.Run("Count", TestCount)
	t.Run("Contains", TestContains)
	t.Run("MinMax", TestMinMax)
	t.Run("Unique", TestUnique)
	t.Run("Filter", TestFilter)
	t.Run("Apply", TestApply)
}

// func TestDataFrame(t *testing.T) {
// 	t.Run("ReadCSV", TestReadCSV)
// }
