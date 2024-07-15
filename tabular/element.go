package tabular

import "reflect"

type Element interface {
	Set(interface{})
	Get() interface{}
	IsNaN() bool
	Type() reflect.Type
}
