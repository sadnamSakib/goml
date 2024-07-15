package tabular

import "reflect"

type intElement struct {
	value int64
	nan   bool
}

func (e *intElement) Set(v interface{}) {
	if v == nil {
		e.nan = true
		return
	}
	e.value = v.(int64)
}
func (e *intElement) Get() interface{} {
	return e.value
}
func (e *intElement) IsNaN() bool {
	return e.nan
}
func (e *intElement) Type() reflect.Type {
	return reflect.TypeOf(e.value)
}

var _ Element = (*intElement)(nil)
