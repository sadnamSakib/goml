package tabular

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"sync"

	"github.com/sadnamSakib/goml/numerics"
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
	return fmt.Sprintf("[%s]", s.elements.String())
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
		if s.elements == nil {
			elements := make(intElements, 0)
			s.elements = &elements
		}
		elements := s.elements.(*intElements)
		*elements = append(*elements, intElement{value: v})
	case reflect.Float64:
		v := val.(float64)
		if s.elements == nil {
			elements := make(floatElements, 0)
			s.elements = &elements
		}

		elements := s.elements.(*floatElements)
		*elements = append(*elements, floatElement{value: v})
	case reflect.String:
		v := val.(string)
		if s.elements == nil {
			elements := make(stringElements, 0)
			s.elements = &elements
		}
		elements := s.elements.(*stringElements)
		*elements = append(*elements, stringElement{value: v})
	case reflect.Bool:
		v := val.(bool)
		if s.elements == nil {
			elements := make(boolElements, 0)
			s.elements = &elements
		}
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
		fmt.Println("Default")
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

func (s *Series) GetRows(start, end int) Series {
	var newSeries Series
	newSeries.Name = s.Name
	newSeries.T = s.T
	for i := start; i < end; i++ {
		newSeries.Append(s.Get(i))
	}
	return newSeries
}

func (s *Series) Copy() Series {
	var newSeries Series = Series{
		Name: s.Name,
		T:    s.T,
	}
	switch s.T {
	case reflect.TypeOf(int64(0)):
		elements := make(intElements, s.Len())
		for i := 0; i < s.Len(); i++ {
			elements[i] = intElement{value: s.Get(i).(int64)}
		}
		newSeries.elements = &elements
	case reflect.TypeOf(float64(0)):
		elements := make(floatElements, s.Len())
		for i := 0; i < s.Len(); i++ {
			elements[i] = floatElement{value: s.Get(i).(float64)}
		}
		newSeries.elements = &elements
	case reflect.TypeOf(""):
		elements := make(stringElements, s.Len())
		for i := 0; i < s.Len(); i++ {
			elements[i] = stringElement{value: s.Get(i).(string)}
		}
		newSeries.elements = &elements
	case reflect.TypeOf(true):
		elements := make(boolElements, s.Len())
		for i := 0; i < s.Len(); i++ {
			elements[i] = boolElement{value: s.Get(i).(bool)}
		}
		newSeries.elements = &elements
	}

	return newSeries
}

func (s *Series) SortBy(column Series) Series {
	var newSeries = s.Copy()
	length := s.Len()
	switch s.T {
	case reflect.TypeOf(int64(0)):
		elements := make(intElements, length)
		for i := 0; i < length; i++ {
			elements[i] = intElement{value: s.Get(i).(int64)}
		}
		sort.Slice(elements, func(i, j int) bool {
			return column.elements.Less(i, j)
		})
		newSeries.elements = &elements
	case reflect.TypeOf(float64(0)):
		elements := make(floatElements, length)
		for i := 0; i < length; i++ {
			elements[i] = floatElement{value: s.Get(i).(float64)}
		}
		sort.Slice(elements, func(i, j int) bool {
			return column.elements.Less(i, j)
		})
		newSeries.elements = &elements
	case reflect.TypeOf(""):
		elements := make(stringElements, length)
		for i := 0; i < length; i++ {
			elements[i] = stringElement{value: s.Get(i).(string)}
		}
		sort.Slice(elements, func(i, j int) bool {
			return column.elements.Less(i, j)
		})
		newSeries.elements = &elements
	case reflect.TypeOf(true):
		elements := make(boolElements, length)
		for i := 0; i < length; i++ {
			elements[i] = boolElement{value: s.Get(i).(bool)}
		}
		sort.Slice(elements, func(i, j int) bool {
			return column.elements.Less(i, j)
		})
		newSeries.elements = &elements

	}

	return newSeries
}

func (s Series) Array() numerics.Array {
	var a numerics.Array
	switch s.T {
	case reflect.TypeOf(int64(0)):
		elements := s.elements.(*intElements)
		for _, val := range *elements {
			e := numerics.NewElement(float64(val.Get().(int64)), numerics.FloatType)
			a.Append(e)
		}
	case reflect.TypeOf(float64(0)):
		elements := s.elements.(*floatElements)
		for _, val := range *elements {
			e := numerics.NewElement(float64(val.Get().(float64)), numerics.FloatType)
			a.Append(e)
		}
	}
	return a
}

func (s Series) Mean() float64 {
	var sum float64
	switch s.T {
	case reflect.TypeOf(int64(0)):
		elements := s.elements.(*intElements)
		for _, val := range *elements {
			sum += float64(val.Get().(int64))
		}
	case reflect.TypeOf(float64(0)):
		elements := s.elements.(*floatElements)
		for _, val := range *elements {
			sum += val.Get().(float64)
		}
	}
	return sum / float64(s.Len())
}
