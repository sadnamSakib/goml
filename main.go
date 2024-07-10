package main

import (
	"fmt"

	"github.com/sadnamSakib/goml/tabular"
)

func main() {
	column := tabular.NewSeries(1, 3, "a", 4, 5)
	fmt.Println(column.String())
	column.Sort()
	fmt.Println(column.String())
}
