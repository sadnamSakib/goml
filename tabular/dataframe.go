package tabular

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type DataFrame struct {
	columns []Series
	rows    int
	cols    int
}

func rowSplitter(line string) []string {

	row := []string{}
	length := len(line)
	l, r := 0, 1
	for r < length {
		if line[l] == '"' {
			for r < length {
				if r < length-1 && line[r+1] == '"' && line[r] == '"' {
					r++
				} else if line[r] == '"' {
					break
				}
				r++
			}
			row = append(row, line[l+1:r])
			l = r + 1
			if l < length && line[l] == ',' {
				l++
			}
			r = l + 1
		} else if line[r] == ',' {
			row = append(row, line[l:r])
			l = r + 1
			r = l + 1
		} else {
			r++
		}
	}
	if l < length-1 {
		row = append(row, line[l:])

	}

	return row
}

func findIndex(columnNames []string, s string) int {
	for i, val := range columnNames {
		if val == s {
			return i
		}
	}
	return -1
}

func sortColumns(df *DataFrame, columnNames []string) {
	sort.Slice(df.columns, func(i, j int) bool {
		indexI := findIndex(columnNames, df.columns[i].Name)
		indexJ := findIndex(columnNames, df.columns[j].Name)
		return indexI < indexJ
	})
}

func Read_CSV(filename string, hasHeader ...bool) (DataFrame, error) {
	startingTime := time.Now()
	var df DataFrame = DataFrame{
		columns: []Series{},
		rows:    0,
		cols:    0,
	}

	var header bool
	var fixedNumberOfColumnsInAllRecords int
	var listOfColumns []string
	if len(hasHeader) == 0 {
		header = false
	} else {
		header = hasHeader[0]
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	madeHeader := 0
	rows := [][]string{}
	if header {
		scanner.Scan()
		listOfColumns = rowSplitter(scanner.Text())
		madeHeader = 1
		fixedNumberOfColumnsInAllRecords = len(listOfColumns)
		df.cols = fixedNumberOfColumnsInAllRecords
		for range listOfColumns {
			rows = append(rows, []string{})
		}
	}
	numRows := 0

	for scanner.Scan() {
		numRows++
		line := rowSplitter(scanner.Text())
		columnsInCurrentRecord := len(line)
		if madeHeader == 0 {
			col := []string{}
			for i := 0; i < columnsInCurrentRecord; i++ {
				col = append(col, string(rune('A'+i)))
			}
			listOfColumns = col
			madeHeader = 1
			fixedNumberOfColumnsInAllRecords = columnsInCurrentRecord
			df.cols = fixedNumberOfColumnsInAllRecords
			for range listOfColumns {
				rows = append(rows, []string{}) //
			}
		}

		for i := 0; i < fixedNumberOfColumnsInAllRecords; i++ {
			if i < columnsInCurrentRecord {
				rows[i] = append(rows[i], line[i])
			} else {
				rows[i] = append(rows[i], "")
			}
		}
	}
	df.rows = numRows

	seriesChannel := make(chan Series, len(listOfColumns))
	wg := sync.WaitGroup{}
	for i := range listOfColumns {
		wg.Add(1)
		go MakeSeries(listOfColumns[i], rows[i], &wg, seriesChannel)
	}
	go func() {
		wg.Wait()
		close(seriesChannel)
	}()

	for s := range seriesChannel {
		df.columns = append(df.columns, s)
	}
	sortColumns(&df, listOfColumns)
	endingTime := time.Now()
	fmt.Println("Time taken to read the file: ", endingTime.Sub(startingTime))

	return df, nil
}

func Write_CSV(df DataFrame, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var sb strings.Builder

	for key, val := range df.columns {

		sb.WriteString(val.Name)
		if key < df.cols-1 {

			sb.WriteString(",")
		}
	}
	sb.WriteString("\n")
	for i := 0; i < df.rows; i++ {
		for j, val := range df.columns {
			sb.WriteString(fmt.Sprintf("%v", val.elements.Get(i)))
			if j < df.cols-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString("\n")
	}
	file.WriteString(sb.String())
	return nil
}

func (df DataFrame) String() string {

	var sb strings.Builder
	for key, val := range df.columns {
		sb.WriteString(fmt.Sprintf("%s: ", key))
		sb.WriteString(val.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (df DataFrame) Head() string {
	var sb strings.Builder

	for i := 0; i < df.cols; i++ {
		sb.WriteString(fmt.Sprintf("%20s", df.columns[i].Name))

		sb.WriteString("|")

	}
	sb.WriteString("\n")
	for i := 0; i < df.cols; i++ {
		sb.WriteString(fmt.Sprintf("%s", "-----------------------"))

		sb.WriteString("|")

	}
	sb.WriteString("\n")
	for i := 0; i < 5; i++ {
		for _, val := range df.columns {
			sb.WriteString(fmt.Sprintf("%20v", val.elements.Get(i)))
			sb.WriteString("|")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (df DataFrame) Tail() string {
	var sb strings.Builder

	for i := 0; i < df.cols; i++ {
		sb.WriteString(fmt.Sprintf("%20s", df.columns[i].Name))

		sb.WriteString("|")

	}
	sb.WriteString("\n")
	for i := 0; i < df.cols; i++ {
		sb.WriteString(fmt.Sprintf("%s", "-----------------------"))

		sb.WriteString("|")

	}
	sb.WriteString("\n")
	endPoint := df.rows - 5
	if endPoint < 0 {
		endPoint = 0

	}
	for i := df.rows - 1; i > endPoint; i-- {
		for _, val := range df.columns {
			sb.WriteString(fmt.Sprintf("%20v", val.elements.Get(i)))
			sb.WriteString("|")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (df DataFrame) GetColumn(columnName string) Series {
	for _, val := range df.columns {
		if val.Name == columnName {
			return val
		}
	}
	return Series{}
}

func (df DataFrame) GetColumns(columnNames ...string) DataFrame {
	var newDF DataFrame = DataFrame{
		columns: []Series{},
		rows:    df.rows,
		cols:    len(columnNames),
	}
	for _, val := range columnNames {
		newDF.columns = append(newDF.columns, df.GetColumn(val))
	}
	return newDF
}

func (df DataFrame) GetRows(startingRow, endingRow int) DataFrame {
	var newDF DataFrame = DataFrame{
		columns: []Series{},
		rows:    endingRow - startingRow + 1,
		cols:    df.cols,
	}
	for _, val := range df.columns {
		newDF.columns = append(newDF.columns, val.GetRows(startingRow, endingRow))
	}
	return newDF
}

func (df DataFrame) SortBy(column string) DataFrame {
	var newDF = DataFrame{
		columns: []Series{},
		rows:    df.rows,
		cols:    df.cols,
	}
	for _, val := range df.columns {
		newDF.columns = append(newDF.columns, val.Copy())
	}

	sortByColumn := func(column string) Series {
		for _, val := range newDF.columns {
			if val.Name == column {
				return val
			}
		}
		return Series{}
	}(column)

	for i, val := range newDF.columns {
		newDF.columns[i] = val.SortBy(sortByColumn)
	}
	return newDF
}
