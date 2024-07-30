package numerics

import (
	"math/rand"
	"strconv"
	"strings"
)

func parser(s string, d dtype) []Element {
	var i []Element
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	s = strings.TrimSpace(s)

	elements := splitter(s)
	for _, e := range elements {
		e = strings.TrimSpace(e)
		if strings.HasPrefix(e, "[") {
			i = append(i, Element{dtype: ArrayType, value: parser(e, d)})
		} else {
			if d == IntType {
				v, _ := strconv.ParseInt(e, 10, 64)
				elem := Element{dtype: d, value: v}
				i = append(i, elem)
			} else if d == FloatType {
				v, _ := strconv.ParseFloat(e, 64)
				elem := Element{dtype: d, value: v}
				i = append(i, elem)
			} else if d == BoolType {
				v, _ := strconv.ParseBool(e)
				elem := Element{dtype: d, value: v}
				i = append(i, elem)
			} else {
				elem := Element{dtype: d, value: e}
				i = append(i, elem)
			}
		}
	}
	return i
}

func NewArray(s string, d dtype) Array {
	var a Array = parser(s, d)
	return a
}

func Linspace(start, end, num int) Array {
	var a Array = NewArray("", FloatType)
	step := 0
	if num-1 != 0 {
		step = (end - start) / (num - 1)
	}
	for i := 0; i < num; i++ {
		a = append(a, Element{dtype: FloatType, value: float64(start + step*i)})
	}
	return a
}

func Random(start, end, size int) Array {
	var a Array = NewArray("", IntType)
	for i := 0; i < size; i++ {
		a = append(a, Element{dtype: IntType, value: int64(start + rand.Intn(end-start))})
	}
	return a
}

func Zeros(size ...int) Array {
	if len(size) == 1 {
		return Linspace(1, 1, size[0])
	} else if len(size) == 2 {
		zeroElement := Element{dtype: FloatType, value: 0}
		var Arrays Array
		for i := 0; i < size[0]; i++ {
			a := []Element{}
			for j := 0; j < size[1]; j++ {
				a = append(a, zeroElement)
			}
			Arrays.Append(Element{
				dtype: ArrayType, value: a,
			})

		}
		return Arrays
	} else {
		panic("Zeros function only accepts 1 or 2 arguments")
	}
}
