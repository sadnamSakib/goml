package tabular

type Elements interface {
	Len() int
	String() string
	Sort(...func(a, b int) bool)
	Min() Element
	Max() Element
	Head() string
	Get(i int) interface{}
	IsNan(i int) bool
}
