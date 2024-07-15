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

type DataFrame map[string]Series

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

func Read_CSV(filename string, hasHeader ...bool) (DataFrame, error) {
	startingTime := time.Now()
	var df DataFrame = make(map[string]Series)

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

		for range listOfColumns {
			rows = append(rows, []string{}) //
		}
	}

	for scanner.Scan() {
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
		df[s.Name] = s
	}

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
	columns := 0
	rows := -1
	listOfColumns := []string{}
	for key, val := range df {
		if rows == -1 {
			rows = val.Len()
		}
		listOfColumns = append(listOfColumns, key)
	}
	sort.Strings(listOfColumns)
	for i := 0; i < rows; i++ {
		for _, key := range listOfColumns {
			if columns > 0 {
				sb.WriteString(",")
			}
			if df[key].elements.IsNan(i) {
				sb.WriteString("NaN")
			} else {
				sb.WriteString(fmt.Sprintf("%v", df[key].elements.Get(i)))
			}
			columns++
		}
		sb.WriteString("\n")
		columns = 0

	}
	file.WriteString(sb.String())
	return nil
}

func (df DataFrame) String() string {

	var sb strings.Builder
	for key, val := range df {
		sb.WriteString(fmt.Sprintf("%s: ", key))
		sb.WriteString(val.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (df DataFrame) Head() string {
	var sb strings.Builder
	for _, val := range df {
		sb.WriteString(fmt.Sprintf("%s: ", val.Name))
		sb.WriteString(val.elements.Head())
		sb.WriteString("\n")
	}
	return sb.String()
}
