package regressor

import (
	"fmt"
	"testing"

	"github.com/sadnamSakib/goml/tabular"
)

const filePath1 = "../Mobile-Price-Prediction-cleaned_data.csv"
const filePath2 = "../Student_Performance.csv"
const headers = true

func TestLinearRegression(t *testing.T) {
	var df, _ = tabular.Read_CSV(filePath1, headers)
	var df_train, _ = df.TrainTestSplit(0.2)
	fmt.Println("Rows: ", df_train.RowNum())
	fmt.Println("Columns: ", df_train.ColNum())
	regr, err := LinearRegression(df_train, []string{"Ratings", "Battery_Power", "RAM", "Primary_Cam", "Mobile_Size", "Selfi_Cam", "ROM"}, "Price")
	if err != nil {
		t.Error(err)
	}
	test_size := 807.0 * 0.8

	if regr.RowNum() != int(test_size) {
		t.Errorf("Expected 807 rows, got %d", regr.RowNum())
	}
	if regr.ColNum() != 7 {
		t.Errorf("Expected 7 columns, got %d", regr.ColNum())
	}
}

func TestPredict1(t *testing.T) {
	var df, _ = tabular.Read_CSV(filePath1, headers)
	var df_train, df_test = df.TrainTestSplit(0.2)

	regr, err := LinearRegression(df_train, []string{"Ratings", "Battery_Power", "RAM", "Primary_Cam", "Mobile_Size", "Selfi_Cam", "ROM"}, "Price")
	if err != nil {
		t.Error(err)
	}
	rss := 0.0
	tss := 0.0
	m := df_test.Mean("Price")
	convertToFloat := func(i interface{}) float64 {
		switch i.(type) {
		case int64:
			return float64(i.(int64))
		case float64:
			return i.(float64)
		default:
			return 0.0
		}
	}

	for i := 0; i < df_test.RowNum(); i++ {

		predicted := regr.Predict((convertToFloat(df_test.Get(i, 0))), (convertToFloat(df_test.Get(i, 1))), (convertToFloat(df_test.Get(i, 2))), (convertToFloat(df_test.Get(i, 3))), (convertToFloat(df_test.Get(i, 4))), (convertToFloat(df_test.Get(i, 5))), (convertToFloat(df_test.Get(i, 6))))
		actual := convertToFloat(df_test.Get(i, 7))

		rss += (actual - predicted) * (actual - predicted)
		tss += (actual - m) * (actual - m)

	}
	r2 := 1 - (rss / tss)

	fmt.Printf("Accuracy: %f%%\n", (r2 * 100))
}

func TestPlotting1(t *testing.T) {
	var df, err = tabular.Read_CSV(filePath1, headers)
	if err != nil {
		t.Error(err)
	}

	regr, err := LinearRegression(df, []string{"Ratings", "Battery_Power", "RAM", "Primary_Cam", "Mobile_Size", "Selfi_Cam", "ROM"}, "Price")
	if err != nil {
		t.Error(err)
	}
	regr.Plot2D("Ratings")
}
func TestCorrelation1(t *testing.T) {
	var df, err = tabular.Read_CSV(filePath1, headers)
	if err != nil {
		t.Error(err)
	}

	regr, err := LinearRegression(df, []string{"Ratings", "Battery_Power", "RAM", "Primary_Cam", "Mobile_Size", "Selfi_Cam", "ROM"}, "Price")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Ratings Correlation Coefficient: ", regr.Correlation("Ratings"))
	fmt.Println("Battery Power Correlation Coefficient: ", regr.Correlation("Battery_Power"))
	fmt.Println("RAM Correlation Coefficient: ", regr.Correlation("RAM"))
	fmt.Println("Primary Cam Correlation Coefficient: ", regr.Correlation("Primary_Cam"))
	fmt.Println("Mobile Size Correlation Coefficient: ", regr.Correlation("Mobile_Size"))
	fmt.Println("Selfi Cam Correlation Coefficient: ", regr.Correlation("Selfi_Cam"))
	fmt.Println("ROM Correlation Coefficient: ", regr.Correlation("ROM"))

}

func TestPredict2(t *testing.T) {
	var df, _ = tabular.Read_CSV(filePath2, headers)

	var df_train, df_test = df.TrainTestSplit(0.2)

	regr, err := LinearRegression(df_train, []string{"Hours Studied", "Previous Scores", "Sleep Hours", "Sample Question Papers Practiced"}, "Performance Index")
	if err != nil {
		t.Error(err)
	}
	rss := 0.0
	tss := 0.0
	m := df_test.Mean("Performance Index")
	convertToFloat := func(i interface{}) float64 {
		switch i.(type) {
		case int64:
			return float64(i.(int64))
		case float64:
			return i.(float64)
		default:
			return 0.0
		}
	}

	for i := 0; i < df_test.RowNum(); i++ {

		predicted := regr.Predict((convertToFloat(df_test.Get(i, 0))), (convertToFloat(df_test.Get(i, 1))), (convertToFloat(df_test.Get(i, 3))), (convertToFloat(df_test.Get(i, 4))))
		actual := convertToFloat(df_test.Get(i, 5))
		// fmt.Println("Predicted: ", predicted, "Actual: ", actual, "Mean: ", m)
		rss += (actual - predicted) * (actual - predicted)
		tss += (actual - m) * (actual - m)

	}
	r2 := 1 - (rss / tss)

	fmt.Printf("Accuracy: %f%%\n", (r2 * 100))
}

func TestPlotting2(t *testing.T) {
	var df, err = tabular.Read_CSV(filePath2, headers)
	if err != nil {
		t.Error(err)
	}

	regr, err := LinearRegression(df, []string{"Hours Studied", "Previous Scores", "Sleep Hours", "Sample Question Papers Practiced"}, "Performance Index")
	if err != nil {
		t.Error(err)
	}
	regr.Plot2D("Hours Studied")
}

func TestCorrelation2(t *testing.T) {
	var df, err = tabular.Read_CSV(filePath2, headers)
	if err != nil {
		t.Error(err)
	}

	regr, err := LinearRegression(df, []string{"Hours Studied", "Previous Scores", "Sleep Hours", "Sample Question Papers Practiced"}, "Performance Index")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Previous Score Correlation Coefficient: ", regr.Correlation("Previous Scores"))
	fmt.Println("Hours Studied Correlation Coefficient: ", regr.Correlation("Hours Studied"))
	fmt.Println("Sleep Hours Correlation Coefficient: ", regr.Correlation("Sleep Hours"))
	fmt.Println("Sample Question Papers Practiced: ", regr.Correlation("Sample Question Papers Practiced"))

}

func TestRegressor(t *testing.T) {
	t.Run("TestLinearRegression", TestLinearRegression)
	t.Run("TestPredict1", TestPredict1)
	t.Run("TestPlotting1", TestPlotting1)
	t.Run("TestCorrelation1", TestCorrelation1)
	t.Run("TestPredict2", TestPredict2)
	t.Run("TestPlotting2", TestPlotting2)
	t.Run("TestCorrelation2", TestCorrelation2)
}
