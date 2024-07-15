package tabular

import "reflect"

type boolElement struct {
	value bool
	nan   bool
}

func (e *boolElement) Set(v interface{}) {
	if v == nil {
		e.nan = true
		return
	}
	e.value = v.(bool)
}

func (e *boolElement) Get() interface{} {
	return e.value
}

func (e *boolElement) IsNaN() bool {
	return e.nan
}

func (e *boolElement) Type() reflect.Type {
	return reflect.TypeOf(e.value)
}

var _ Element = (*boolElement)(nil)
