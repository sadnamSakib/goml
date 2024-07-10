package tabular

import (
	"cmp"
	"fmt"
	"strings"
)

type Series []interface{}

func NewSeries(values ...interface{}) Series {
	var s Series
	for _, value := range values {
		s.Append(value)
	}
	return s
}

func (s *Series) Append(val interface{}) {
	*s = append(*s, val)
}

func (s Series) String() string {
	var sb strings.Builder
	for i, v := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}

func (s Series) Len() int {
	return len(s)
}

func less(i, j interface{}) bool {

	if fmt.Sprintf("%T", i) != fmt.Sprintf("%T", j) {
		return cmp.Compare(fmt.Sprintf("%v", i), fmt.Sprintf("%v", j)) < 0
	}

	switch i.(type) {
	case int:
		return cmp.Compare(i.(int), j.(int)) < 0
	case float64:
		return cmp.Compare(i.(float64), j.(float64)) < 0
	case string:
		return cmp.Compare(i.(string), j.(string)) < 0
	case bool:
		if i.(bool) == j.(bool) {
			return false
		} else if !i.(bool) {
			return true
		} else {
			return false
		}
	default:
		return cmp.Compare(fmt.Sprintf("%v", i), fmt.Sprintf("%v", j)) < 0
	}
}

func (s *Series) Sort(function ...func(a, b interface{}) bool) {
	if len(function) > 1 {
		panic("too many functions")
	}
	if len(function) == 0 {
		function = append(function, less)
	}
	sortFunc := function[0]
	mergeSort(s, 0, s.Len()-1, sortFunc)

}

func (s Series) SortCopy(function ...func(a, b interface{}) bool) Series {
	result := make(Series, len(s))
	copy(result, s)
	result.Sort(function...)
	return result
}

func merge(s *Series, left, mid, right int, sortFunc func(a, b interface{}) bool) {
	n1 := mid - left + 1
	n2 := right - mid
	L := make(Series, n1)
	R := make(Series, n2)
	for i := 0; i < n1; i++ {
		L[i] = (*s)[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = (*s)[mid+1+i]
	}
	i := 0
	j := 0
	k := left
	for i < n1 && j < n2 {
		if sortFunc(L[i], R[j]) {
			(*s)[k] = L[i]
			i++
		} else {
			(*s)[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		(*s)[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		(*s)[k] = R[j]
		j++
		k++
	}
}

func mergeSort(s *Series, start, end int, sortFunc func(a, b interface{}) bool) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(s, start, mid, sortFunc)
		mergeSort(s, mid+1, end, sortFunc)
		merge(s, start, mid, end, sortFunc)
	}
}
