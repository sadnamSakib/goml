package tabular

type Elements interface {
	Len() int
	String() string
	Min() Element
	Max() Element
	Head() string
	Tail() string
	Get(i int) interface{}
	Less(i, j int) bool
	IsNan(i int) bool
}
