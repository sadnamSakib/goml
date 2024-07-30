package numerics

import (
	"fmt"
	"strings"
)

type dtype int

const (
	IntType dtype = iota
	FloatType
	StringType
	BoolType
	ArrayType
)

type Element struct {
	dtype dtype
	value interface{}
	nan   bool
}

func NewElement(v interface{}, t dtype) Element {
	return Element{
		dtype: t,
		value: v,
	}
}

func (e *Element) Set(v interface{}) {
	if v == nil {
		e.nan = true
		return
	}
	e.value = v
}

func (e *Element) Get() interface{} {
	return e.value
}

func (e *Element) IsNaN() bool {
	return e.nan
}

func (e *Element) Type() dtype {
	return e.dtype
}

func (e *Element) String() string {
	var sb strings.Builder
	switch e.dtype {
	case ArrayType:
		sb.WriteString("[")
		for _, e := range e.value.([]Element) {
			sb.WriteString(e.String())
		}
		s := (sb.String())[:sb.Len()-1]
		sb.Reset()
		sb.WriteString(s)
		sb.WriteString("]")
		sb.WriteString(",")
	default:
		sb.WriteString(fmt.Sprintf("%v", e.value))
		sb.WriteString(",")
	}
	return sb.String()
}
