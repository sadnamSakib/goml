package tabular

import "reflect"

type floatElement struct {
	value float64
	nan   bool
}

func (e *floatElement) Set(v interface{}) {
	if v == nil {
		e.nan = true
		return
	}
	e.value = v.(float64)
}
func (e *floatElement) Get() interface{} {
	return e.value
}
func (e *floatElement) IsNaN() bool {
	return e.nan
}
func (e *floatElement) Type() reflect.Type {
	return reflect.TypeOf(e.value)
}

var _ Element = (*floatElement)(nil)
