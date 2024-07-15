package tabular

import (
	"errors"
	"reflect"
	"strconv"
	"sync"
)

var UnsupportedTypeError = errors.New("unsupported type")

type Series struct {
	Name     string
	elements Elements
	T        reflect.Type
}

func MakeSeries(name string, v []string, wg *sync.WaitGroup, ch chan Series) {
	defer wg.Done()
	var t reflect.Type
	length := len(v)
	isBool := 0
	isFloat := 0
	isString := 0
	isInt := 0

	for _, val := range v {
		if val == "" || val == "NaN" {
			isBool++
			isFloat++
			isString++
			isInt++
			continue
		}
		if _, err := strconv.ParseBool(val); err == nil {
			isBool++
		}
		if _, err := strconv.ParseInt(val, 10, 64); err == nil {
			isInt++
		}
		if _, err := strconv.ParseFloat(val, 64); err == nil {
			isFloat++
		}
		isString++
	}
	if isBool == length {
		t = reflect.TypeOf(true)
	} else if isInt == length {
		t = reflect.TypeOf(int64(0))
	} else if isFloat == length {
		t = reflect.TypeOf(float64(0))
	} else {

		t = reflect.TypeOf("")
	}
	s, _ := new(name, v, t)

	ch <- s
}

func new(name string, v []string, t reflect.Type) (Series, error) {
	length := len(v)
	switch t.Kind() {
	case reflect.Int64:
		elements := make(intElements, length)
		for i, val := range v {
			if val == "" || val == "NaN" {
				elements[i].nan = true
				continue
			}
			elements[i].value, _ = strconv.ParseInt(val, 10, 64)

		}
		return Series{
			Name:     name,
			elements: &elements,
			T:        t,
		}, nil
	case reflect.Float64:
		elements := make(floatElements, length)
		for i, val := range v {
			if val == "" {
				elements[i].nan = true
				continue
			}
			elements[i].value, _ = strconv.ParseFloat(val, 64)
		}
		return Series{
			Name:     name,
			elements: &elements,
			T:        t,
		}, nil
	case reflect.String:
		elements := make(stringElements, length)
		for i, val := range v {
			if val == "" {
				elements[i].nan = true
				continue
			}
			elements[i].value = val
		}
		return Series{
			Name:     name,
			elements: &elements,
			T:        t,
		}, nil
	case reflect.Bool:
		elements := make(boolElements, length)
		for i, val := range v {
			if val == "" {
				elements[i].nan = true
				continue
			}
			elements[i].value, _ = strconv.ParseBool(val)
		}
		return Series{
			Name:     name,
			elements: &elements,
			T:        t,
		}, nil
	default:
		return Series{}, UnsupportedTypeError

	}
}

func (s *Series) Len() int {
	return s.elements.Len()
}
func (s *Series) String() string {
	return s.elements.String()
}
func (s *Series) Sort(lessFuncs ...func(a, b int) bool) {
	s.elements.Sort(lessFuncs...)
}
func (s *Series) Min() Element {
	return s.elements.Min()
}
func (s *Series) Max() Element {
	return s.elements.Max()

}
func (s *Series) Append(val interface{}) {
	switch s.T.Kind() {
	case reflect.Int64:
		v := val.(int64)
		elements := s.elements.(*intElements)
		*elements = append(*elements, intElement{value: v})
	case reflect.Float64:
		v := val.(float64)
		elements := s.elements.(*floatElements)
		*elements = append(*elements, floatElement{value: v})
	case reflect.String:
		v := val.(string)
		elements := s.elements.(*stringElements)
		*elements = append(*elements, stringElement{value: v})
	case reflect.Bool:
		v := val.(bool)
		elements := s.elements.(*boolElements)
		*elements = append(*elements, boolElement{value: v})
	default:
		return
	}
}
func (s *Series) Get(i int) interface{} {
	switch s.T.Kind() {
	case reflect.Int64:
		elements := s.elements.(*intElements)
		return (*elements)[i].Get()
	case reflect.Float64:
		elements := s.elements.(*floatElements)
		return (*elements)[i].Get()
	case reflect.String:
		elements := s.elements.(*stringElements)
		return (*elements)[i].Get()
	case reflect.Bool:
		elements := s.elements.(*boolElements)
		return (*elements)[i].Get()
	default:
		return nil
	}
}

func (s *Series) IsNaN(i int) bool {
	switch s.T.Kind() {
	case reflect.Int64:
		elements := s.elements.(*intElements)
		return (*elements)[i].IsNaN()
	case reflect.Float64:
		elements := s.elements.(*floatElements)
		return (*elements)[i].IsNaN()
	case reflect.String:
		elements := s.elements.(*stringElements)
		return (*elements)[i].IsNaN()
	case reflect.Bool:
		elements := s.elements.(*boolElements)
		return (*elements)[i].IsNaN()
	default:
		return false
	}
}

func (s *Series) Type() reflect.Type {
	return s.T
}

func (s *Series) getName() string {
	return s.Name
}

func (s *Series) setName(name string) {
	s.Name = name
}
func (s *Series) Set(i int, v interface{}) {
	switch s.T.Kind() {
	case reflect.Int64:
		elements := s.elements.(*intElements)
		(*elements)[i].Set(v)
	case reflect.Float64:
		elements := s.elements.(*floatElements)
		(*elements)[i].Set(v)
	case reflect.String:
		elements := s.elements.(*stringElements)
		(*elements)[i].Set(v)
	case reflect.Bool:
		elements := s.elements.(*boolElements)
		(*elements)[i].Set(v)
	default:
		return
	}
}
