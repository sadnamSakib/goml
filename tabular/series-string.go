package tabular

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type stringElements []stringElement

var EmptystringElementsError error = errors.New("empty *stringElements")

func (s *stringElements) Len() int {
	return (len(*s))
}

func (s *stringElements) String() string {
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
func (s *stringElements) Get(i int) interface{} {
	return ((*s)[i].Get()).(string)
}
func (s *stringElements) IsNan(i int) bool {
	return (*s)[i].IsNaN()

}

func (s *stringElements) Sort(lessFuncs ...func(a, b int) bool) {
	if len(lessFuncs) == 0 {
		sort.Slice(*s, func(i, j int) bool {
			return (*s)[i].value < (*s)[j].value
		})
	} else {
		sort.Slice(*s, lessFuncs[0])
	}
}

func (s *stringElements) Min() Element {
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

func (s *stringElements) Max() Element {
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

func (s *stringElements) Head() string {
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

var _ Elements = (*stringElements)(nil)
