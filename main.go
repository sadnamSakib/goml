package main

import (
	"fmt"

	"github.com/sadnamSakib/goml/regressor"
	"github.com/sadnamSakib/goml/tabular"
)

func main() {
	df, err := tabular.Read_CSV("Mobile-Price-Prediction-cleaned_data.csv", true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(df.Head())
	df = df.SortBy("RAM")
	fmt.Println(df.Head())

	regr, err := regressor.LinearRegression(df, []string{"Ratings", "Battery_Power", "RAM"}, "Price")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(regr.Predict(4.8, 4500, 16))
	fmt.Println("Correlation with Battery Power: ", regr.Correlation("Battery_Power"))
	fmt.Println("Correlation with Ratings: ", regr.Correlation("Ratings"))
	fmt.Println("Correlation with RAM: ", regr.Correlation("RAM"))
	fmt.Println("Correlation with ROM: ", regr.Correlation("ROM"))
	fmt.Println("Correlation with Mobile Size: ", regr.Correlation("Mobile_Size"))
	fmt.Println("Correlation with Primary Cam: ", regr.Correlation("Primary_Cam"))
	fmt.Println("Correlation with Selfi Cam: ", regr.Correlation("Selfi_Cam"))
	regr.Plot2D("Battry_Power")

}
