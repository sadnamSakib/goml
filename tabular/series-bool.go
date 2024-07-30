package tabular

import (
	"errors"
	"fmt"
	"strings"
)

type boolElements []boolElement

var EmptyBoolElementsError error = errors.New("Empty column")

func (s *boolElements) Len() int {
	return len(*s)
}
func (s *boolElements) Head() string {
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

func (s *boolElements) String() string {
	var sb strings.Builder
	sb.WriteString("[")
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
	sb.WriteString("]")
	return sb.String()
}

func (s *boolElements) Min() Element {
	if s.Len() == 0 {
		panic("Cannot get min value from an empty column")
	}

	min := (*s)[0]
	for _, v := range *s {
		if v.value == false {
			min = v
		}
	}
	return &min
}

func (s *boolElements) Max() Element {
	if s.Len() == 0 {
		panic("Cannot get max value from an empty column")
	}

	max := (*s)[0]
	for _, v := range *s {
		if v.value == true {
			max = v
		}
	}
	return &max
}

func (s *boolElements) Get(i int) interface{} {
	return ((*s)[i].Get()).(bool)
}

func (s *boolElements) Less(i, j int) bool {
	return (*s)[i].value == false
}

func (s *boolElements) IsNan(i int) bool {
	return (*s)[i].IsNaN()

}

func (s *boolElements) Tail() string {
	length := len(*s)
	if length > 5 {
		length = 5
	}
	var sb strings.Builder
	for i := length; i > 0; i-- {
		if i < length {
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

var _ Elements = (*boolElements)(nil)
