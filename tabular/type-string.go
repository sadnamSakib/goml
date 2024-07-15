package tabular

import "reflect"

type stringElement struct {
	value string
	nan   bool
}

func (e *stringElement) Set(v interface{}) {
	if v == nil {
		e.nan = true
		return
	}
	e.value = v.(string)
}
func (e *stringElement) Get() interface{} {
	return e.value
}
func (e *stringElement) IsNaN() bool {
	return e.nan
}
func (e *stringElement) Type() reflect.Type {
	return reflect.TypeOf(e.value)
}

var _ Element = (*stringElement)(nil)
