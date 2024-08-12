package numerics

import (
	"math"
	"reflect"
	"strings"
)

type Array []Element

func (a Array) Len() int { return len(a) }

func (a Array) Mean() float64 {
	if !a.IsOfType(FloatType) && !a.IsOfType(IntType) {
		panic("Array is not of type float or int ")
	}
	sum := 0.0
	for _, v := range a {
		if v.dtype == FloatType {
			sum += v.value.(float64)
		} else {
			sum += float64(v.value.(int))
		}
	}
	return sum / float64(a.Len())

}

func (a Array) Std() float64 {
	if !a.IsOfType(FloatType) && !a.IsOfType(IntType) {
		panic("Array is not of type float or int ")
	}
	mean := a.Mean()
	sum := 0.0
	for _, v := range a {
		sum += (v.value.(float64) - mean) * (v.value.(float64) - mean)
	}
	return math.Sqrt(sum / float64(a.Len()))
}

func splitter(s string) []string {
	a := []string{}
	bracketsOpen := 0
	temp := ""
	for _, c := range s {
		if c == '[' {
			bracketsOpen++
			temp += string(c)
		} else if c == ']' {
			bracketsOpen--
			temp += string(c)
			if bracketsOpen == 0 {
				a = append(a, temp)
				temp = ""
			}
		} else {
			if c == ',' && bracketsOpen == 0 {
				if temp != "" {
					a = append(a, temp)
				}
				temp = ""
			} else {
				temp += string(c)
			}

		}
	}
	if temp != "" {
		a = append(a, temp)
	}
	return a

}

func (a Array) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, val := range a {
		sb.WriteString(val.String())
	}
	s := (sb.String())[:sb.Len()-1]
	sb.Reset()
	sb.WriteString(s)
	sb.WriteString("]")
	return sb.String()
}

func (a *Array) Append(e Element) {
	*a = append(*a, e)
}

func (a Array) IsOfType(t dtype) bool {
	for _, v := range a {
		if v.dtype != t {
			return false
		}
	}
	return true
}

func (a Array) ScalarMultiplication(v float64) {
	if !a.IsOfType(FloatType) {
		panic("Array is not of type float")
	}
	for i := range a {
		(a)[i].value = (a)[i].value.(float64) * v
	}
}

func (a Array) ScalarAddition(v float64) {
	if !a.IsOfType(FloatType) {
		panic("Array is not of type float")
	}
	for i := range a {
		(a)[i].value = (a)[i].value.(float64) + v
	}
}

func (a Array) ScalarSubtraction(v float64) {
	if !a.IsOfType(FloatType) {
		panic("Array is not of type float")
	}
	for i := range a {
		(a)[i].value = (a)[i].value.(float64) - v
	}
}

func (a Array) ScalarDivision(v float64) {
	if !a.IsOfType(FloatType) {
		panic("Array is not of type float")
	}
	for i := range a {
		(a)[i].value = (a)[i].value.(float64) / v
	}
}
func (a Array) Max() float64 {
	if !a.IsOfType(FloatType) {
		panic("Array is not of type float")
	}
	max := a[0].value.(float64)
	for _, v := range a {
		if v.value.(float64) > max {
			max = v.value.(float64)
		}
	}
	return max
}
func (a Array) Min() float64 {
	if !a.IsOfType(FloatType) {
		panic("Array is not of type float")
	}
	min := a[0].value.(float64)
	for _, v := range a {
		if v.value.(float64) < min {
			min = v.value.(float64)
		}
	}
	return min
}

func (a Array) Copy() Array {
	b := make(Array, len(a))
	copy(b, a)
	return b
}

func (a Array) GetType() dtype {
	return a[0].dtype
}

func (a Array) Get(i int) interface{} {
	return a[i].value
}

func (a *Array) Set(i int64, v interface{}) {
	dTypeOfArray := a.GetType()
	dTypeOfValue := reflect.TypeOf(v)
	if dTypeOfArray == IntType && dTypeOfValue == reflect.TypeOf(int(int64(0))) {
		(*a)[i].value = int64(v.(int))

	} else if dTypeOfArray == FloatType && dTypeOfValue == reflect.TypeOf(float64(0)) {
		(*a)[i].value = v.(float64)

	} else if dTypeOfArray == BoolType && dTypeOfValue == reflect.TypeOf(true) {
		(*a)[i].value = v.(bool)

	} else if dTypeOfArray == ArrayType && dTypeOfValue == reflect.TypeOf(Array{}) {
		e := []Element{}
		for _, val := range v.(Array) {
			elem := Element{dtype: val.dtype, value: val.value}
			e = append(e, elem)
		}
		(*a)[i].value = e
	} else {
		panic("Type mismatch")
	}
}
