package tabular

import (
	"os"
	"testing"
)

const filePath = "../Mobile-Price-Prediction-cleaned_data.csv"
const headers = true

func TestReadCSV(t *testing.T) {
	df, err := Read_CSV(filePath, headers)
	if err != nil {
		t.Error(err)
	}
	if df.GetRowNum() != 807 {
		t.Errorf("Expected 807 rows, got %d", df.GetRowNum())
	}
	if df.GetColNum() != 8 {
		t.Errorf("Expected 8 columns, got %d", df.GetColNum())
	}
}
func TestWriteCSV(t *testing.T) {
	df, err := Read_CSV(filePath, headers)
	if err != nil {
		t.Error(err)
	}
	err = Write_CSV(df, "test.csv")
	if err != nil {
		t.Error(err)
	}
	df2, err := Read_CSV("test.csv", headers)
	if err != nil {
		t.Error(err)
	}
	if df.GetRowNum() != df2.GetRowNum() {
		t.Errorf("Expected %d rows, got %d", df.GetRowNum(), df2.GetRowNum())
	}
	if df.GetColNum() != df2.GetColNum() {
		t.Errorf("Expected %d columns, got %d", df.GetColNum(), df2.GetColNum())
	}
	err = os.Remove("test.csv")
	if err != nil {
		t.Error(err)
	}
}

func TestSeries(t *testing.T) {
	t.Run("Test Read_CSV", TestReadCSV)
	t.Run("Test Write_CSV", TestWriteCSV)

}
