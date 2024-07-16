package tabular

import (
	"errors"
	"fmt"
	"strings"
)

type floatElements []floatElement

var EmptyFloatElementsError error = errors.New("empty floatElements")

func (s *floatElements) Len() int {
	return len(*s)
}

func (s *floatElements) String() string {
	var sb strings.Builder

	for i := 0; i < s.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		if (*s)[i].IsNaN() {
			sb.WriteString("NaN")
		} else {
			sb.WriteString(fmt.Sprintf("%v", (*s)[i].value))
		}
	}
	return sb.String()
}

func (s *floatElements) Get(i int) interface{} {
	return ((*s)[i].Get()).(float64)
}
func (s *floatElements) Head() string {
	length := len(*s)
	if length > 5 {
		length = 5
	}
	var sb strings.Builder
	for i := 0; i < length; i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		if (*s)[i].IsNaN() {
			sb.WriteString("NaN")
		} else {
			sb.WriteString(fmt.Sprintf("%v", (*s)[i].value))
		}
	}
	return sb.String()

}

func (s *floatElements) Min() Element {
	if s.Len() == 0 {
		panic("Cannot get min value from an empty column")
	}

	min := (*s)[0]
	for _, v := range *s {
		if v.value < min.value {
			min = v
		}
	}
	return &min
}

func (s *floatElements) Max() Element {
	if s.Len() == 0 {
		panic("Cannot get max value from an empty column")
	}
	max := (*s)[0]
	for _, v := range *s {
		if v.value > max.value {
			max = v
		}
	}
	return &max
}
func (s *floatElements) IsNan(i int) bool {
	return (*s)[i].IsNaN()

}
func (s *floatElements) Tail() string {
	length := len(*s)
	if length > 5 {
		length = 5
	}
	var sb strings.Builder
	for i := length - 1; i >= 0; i-- {
		if i < length-1 {
			sb.WriteString(",")
		}
		if (*s)[i].IsNaN() {
			sb.WriteString("NaN")
		} else {
			sb.WriteString(fmt.Sprintf("%v", (*s)[i].value))
		}
	}
	return sb.String()

}
func (s *floatElements) Less(i, j int) bool {
	return (*s)[i].value < (*s)[j].value
}

var _ Elements = (*floatElements)(nil)
