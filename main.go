package main

import (
	"fmt"

	"github.com/sadnamSakib/goml/tabular"
)

func main() {
	// column := tabular.NewSeries(1, 3.1, "a", 4, 5)
	// fmt.Println(column.String())
	// column.Sort()
	// fmt.Println(column.String())
	// sum, err := column.Sum()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(sum)
	// fmt.Println(column.Concat())
	df, err := tabular.Read_CSV("tabular/people.csv", true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(df))
	fmt.Println(df.Head())
	err = tabular.Write_CSV(df, "newcsv.csv")
	if err != nil {
		fmt.Println(err)
	}

}
