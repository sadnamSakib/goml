package numerics

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var invalidRowColumnError = errors.New("Invalid number of rows or columns")

type Matrix struct {
	rowNum int
	colNum int
	rows   []Array
}

func NewMatrix(row, col int, rows []Array) (Matrix, error) {

	if row <= 0 || col <= 0 || len(rows) != row {
		return Matrix{}, invalidRowColumnError
	}
	isOfTypeFloat := 0
	for i, row := range rows {
		if row.IsOfType(FloatType) {
			isOfTypeFloat++
		} else {
			rows[i] = convertToFloat(row)
		}
		if len(row) != col {
			return Matrix{}, invalidRowColumnError
		}

	}

	m := Matrix{
		rowNum: row,
		colNum: col,
		rows:   rows,
	}
	return m, nil
}

func (m Matrix) Copy() Matrix {

	rows := make([]Array, m.RowNum())
	for i := 0; i < len(m.rows); i++ {
		rows[i] = m.rows[i].Copy()
	}

	return Matrix{
		rowNum: m.RowNum(),
		colNum: m.ColNum(),
		rows:   rows,
	}

}

func (m Matrix) String() string {
	var sb strings.Builder
	for _, row := range m.rows {
		sb.WriteString(row.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m Matrix) GetRow(r int) Matrix {
	return Matrix{
		rowNum: 1,
		colNum: m.ColNum(),
		rows:   []Array{m.rows[r]},
	}
}
func (m Matrix) GetRows(s, e int) Matrix {
	newm := Matrix{
		rowNum: e - s,
		colNum: m.ColNum(),
		rows:   m.rows[s:e],
	}
	return newm
}
func (m Matrix) GetColumn(c int) Matrix {
	newm := Matrix{
		rowNum: m.RowNum(),
		colNum: 1,
		rows:   make([]Array, m.RowNum()),
	}
	for i := 0; i < m.RowNum(); i++ {
		newm.rows[i] = make([]Element, 1)
		newm.rows[i][0] = m.rows[i][c]
	}
	return newm
}

func (m Matrix) GetColumns(s, e int) Matrix {
	newM := Matrix{
		rowNum: m.RowNum(),
		colNum: e - s,
		rows:   make([]Array, m.RowNum()),
	}
	for i := 0; i < e-s; i++ {
		for j := 0; j < m.RowNum(); j++ {
			newM.rows[j] = append(newM.rows[j], m.rows[j][s+i])
		}
	}
	return newM
}

func (m Matrix) Transpose() Matrix {
	newM := Matrix{
		rowNum: m.ColNum(),
		colNum: m.RowNum(),
		rows:   make([]Array, m.ColNum()),
	}
	for i := 0; i < m.ColNum(); i++ {
		newM.rows[i] = make([]Element, m.RowNum())
		for j := 0; j < m.RowNum(); j++ {
			newM.rows[i][j] = m.rows[j][i]
		}
	}
	return newM
}

func (m Matrix) Minor(i, j int) Matrix {
	minor := Matrix{
		rowNum: m.RowNum() - 1,
		colNum: m.ColNum() - 1,
		rows:   make([]Array, m.RowNum()-1),
	}
	minorRow := 0
	for row := 0; row < m.RowNum(); row++ {
		if row == i {
			continue
		}
		minor.rows[minorRow] = make([]Element, m.ColNum()-1)
		minorCol := 0
		for col := 0; col < m.ColNum(); col++ {
			if col == j {
				continue
			}
			minor.rows[minorRow][minorCol] = m.rows[row][col]
			minorCol++
		}
		minorRow++
	}

	return minor
}

func (m Matrix) Adjoint() Matrix {
	m2 := Matrix{
		rowNum: m.RowNum(),
		colNum: m.ColNum(),
		rows:   make([]Array, m.RowNum()),
	}

	for i := 0; i < m.RowNum(); i++ {
		for j := 0; j < m.ColNum(); j++ {
			sign := 1.0
			if (i+j)%2 != 0 {
				sign = -sign
			}
			d, _ := m.Minor(i, j).Det()
			m2.rows[i].Append(Element{
				dtype: FloatType,
				value: sign * d,
			})
		}

	}
	return m2.Transpose()
}

func (m Matrix) Inverse() (Matrix, error) {
	if m.RowNum() != m.ColNum() {
		return Matrix{}, errors.New("Matrix is not square")
	}
	d, err := m.Det() // We get matrix is singular error here because of concurrency issue

	if err != nil {
		return Matrix{}, err
	}
	if d == 0.0 {
		fmt.Println("Hello")
		return Matrix{}, errors.New("Matrix is singular")

	}

	return m.Adjoint().ScalarMultiplication(1.0 / d), nil
}

func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

func getValueAsFloat64(e Element) float64 {
	if val, ok := e.value.(float64); ok {
		return val
	}
	if val, ok := e.value.(int64); ok {
		return float64(val)
	}
	return 0.0
}

func GaussianElim(A Matrix) (Matrix, error) {
	m := (A.Copy()) // The copy function is not copying the actual values

	n := m.RowNum()

	for j := 0; j < n; j++ {
		if getValueAsFloat64(m.rows[j][j]) == 0.0 {
			big := 0.0
			kRow := j

			for k := j + 1; k < n; k++ {
				if abs(getValueAsFloat64(m.rows[k][j])) > big {
					big = abs(getValueAsFloat64(m.rows[k][j]))
					kRow = k
				}
			}

			for l := j; l < n; l++ {
				dum := m.rows[j][l]
				m.rows[j][l] = m.rows[kRow][l]
				m.rows[kRow][l] = dum
			}
		}

		pivot := getValueAsFloat64(m.rows[j][j])

		if pivot == 0.0 {

			return Matrix{}, fmt.Errorf("matrix A is singular")
		}

		for i := j + 1; i < n; i++ {
			mult := getValueAsFloat64(m.rows[i][j]) / pivot
			for l := j; l < n; l++ {
				m.rows[i][l] = Element{
					dtype: m.rows[i][l].dtype,
					value: getValueAsFloat64(m.rows[i][l]) - mult*getValueAsFloat64(m.rows[j][l]),
				}
			}

		}
	}

	return m, nil
}

func (A Matrix) Det() (float64, error) {
	if A.ColNum() != A.RowNum() {
		return 0.0, errors.New("Not a square matrix")
	}
	U, err := GaussianElim(A)

	if err != nil {
		return 0.0, err
	}
	det := 1.0
	for i := 0; i < A.RowNum(); i++ {
		det *= getValueAsFloat64(U.rows[i][i])
	}

	return det, nil

}

func multiplyRowByColumn(row, col Array, e *Element) {
	l := row.Len()
	val := 0.0
	for i := 0; i < l; i++ {

		if row[i].dtype == FloatType {
			val += row[i].value.(float64) * col[i].value.(float64)
		} else {
			val += float64(row[i].value.(int64) * col[i].value.(int64))
		}

	}
	if e.dtype == FloatType {
		e.value = float64(val)
	} else {
		e.value = int64(val)
	}
}

func rowMultiplication(row, col int, m1, m2, m *Matrix, w *sync.WaitGroup, sem chan struct{}) {
	defer w.Done()

	defer func() { <-sem }()
	for j := 0; j < col; j++ {
		var r Array
		for i := 0; i < m1.ColNum(); i++ {
			r = append(r, m1.rows[row][i])
		}
		var c Array
		for i := 0; i < m2.RowNum(); i++ {
			c = append(c, m2.rows[i][j])
		}
		multiplyRowByColumn(r, c, &(m.rows[row][j]))
	}
}

func Multiply(m1 Matrix, m2 Matrix) (Matrix, error) {

	if m1.ColNum() != m2.RowNum() {
		return Matrix{}, errors.New("Number of columns on LHS does not match number of rows on RHS")
	}
	w := sync.WaitGroup{}
	m := Matrix{
		rowNum: m1.RowNum(),
		colNum: m2.ColNum(),
		rows:   make([]Array, m1.RowNum()),
	}
	sem := make(chan struct{}, m1.RowNum())
	for i := 0; i < m1.RowNum(); i++ {
		m.rows[i] = Zeros(m2.ColNum())
	}

	for i := 0; i < m1.RowNum(); i++ {
		w.Add(1)
		sem <- struct{}{}
		go rowMultiplication(i, m2.ColNum(), &m1, &m2, &m, &w, sem)
	}

	w.Wait()
	close(sem)

	return m, nil

}

func rowAddition(row int, m1, m2, m *Matrix, w *sync.WaitGroup) {
	w.Add(1)
	defer w.Done()
	for j := 0; j < m1.ColNum(); j++ {
		e := Element{}
		e.value = m1.rows[row][j].value.(float64) + m2.rows[row][j].value.(float64)
		e.dtype = FloatType
		m.rows[row].Append(e)
	}
}

func Subtract(m1, m2 Matrix) (Matrix, error) {
	if m1.RowNum() != m2.RowNum() || m1.ColNum() != m2.ColNum() {
		return Matrix{}, errors.New("Invalid number of rows or columns")
	}
	m := Matrix{
		rowNum: m1.RowNum(),
		colNum: m1.ColNum(),
		rows:   make([]Array, m1.RowNum()),
	}
	w := sync.WaitGroup{}
	sem := make(chan struct{}, runtime.NumCPU()*16)

	for i := 0; i < m1.RowNum(); i++ {
		sem <- struct{}{}
		go func(row int) {
			defer func() { <-sem }()
			rowSubtraction(row, &m1, &m2, &m, &w)
		}(i)
	}
	w.Wait()
	return m, nil
}

func rowSubtraction(row int, m1, m2, m *Matrix, w *sync.WaitGroup) {
	w.Add(1)
	defer w.Done()
	for j := 0; j < m1.ColNum(); j++ {
		e := Element{}
		if m1.rows[row][j].dtype == FloatType {
			e.value = m1.rows[row][j].value.(float64) + m2.rows[row][j].value.(float64)
			e.dtype = FloatType
		} else {
			e.value = m1.rows[row][j].value.(int64) + m2.rows[row][j].value.(int64)
			e.dtype = IntType
		}
		m.rows[row].Append(e)
	}
}

func Add(m1, m2 Matrix) (Matrix, error) {

	if m1.RowNum() != m2.RowNum() || m1.ColNum() != m2.ColNum() {
		return Matrix{}, errors.New("Invalid number of rows or columns")
	}
	m := Matrix{
		rowNum: m1.RowNum(),
		colNum: m1.ColNum(),
		rows:   make([]Array, m1.RowNum()),
	}
	w := sync.WaitGroup{}
	sem := make(chan struct{}, runtime.NumCPU()*16)

	for i := 0; i < m1.RowNum(); i++ {
		sem <- struct{}{}
		go func(row int) {
			defer func() { <-sem }()
			rowAddition(row, &m1, &m2, &m, &w)
		}(i)
	}
	w.Wait()
	return m, nil
}

func pow(val interface{}, p int) float64 {
	if val.(float64) == 0 {
		return 0
	}
	if p == 0 {
		return 1
	}
	if p == 1 {
		return val.(float64)
	}
	return val.(float64) * pow(val, p-1)
}

func Power(m1 Matrix, p int) (Matrix, error) {
	m := Matrix{
		rowNum: m1.RowNum(),
		colNum: m1.ColNum(),
		rows:   make([]Array, m1.RowNum()),
	}
	w := sync.WaitGroup{}
	sem := make(chan struct{}, runtime.NumCPU()*16)

	for i := 0; i < m1.RowNum(); i++ {
		sem <- struct{}{}
		go func(row int) {
			defer func() { <-sem }()
			w.Add(1)
			defer w.Done()
			for j := 0; j < m1.ColNum(); j++ {
				e := Element{}
				if m1.rows[row][j].dtype == FloatType {
					e.value = pow(m1.rows[row][j].value, p)
					e.dtype = FloatType
				} else {
					e.value = (pow(float64(m1.rows[row][j].value.(int64)), p))
					e.dtype = FloatType
				}
				m.rows[row].Append(e)
			}
		}(i)
	}
	w.Wait()
	return m, nil
}

func (m *Matrix) SetColumn(col int, a Array) {
	for i := 0; i < m.RowNum(); i++ {
		(*m).rows[i][col] = a[i]

	}

}

func AppendX(m1, m2 Matrix) (Matrix, error) {
	if m1.RowNum() != m2.RowNum() {
		return Matrix{}, errors.New("Invalid number of rows")
	}
	m := Matrix{
		rowNum: m1.RowNum(),
		colNum: m1.ColNum() + m2.ColNum(),
		rows:   make([]Array, m1.RowNum()),
	}

	for i := 0; i < m1.RowNum(); i++ {

		m.rows[i] = append(m1.rows[i], m2.rows[i]...)
	}

	return m, nil
}

func AppendY(m1, m2 Matrix) (Matrix, error) {
	if m1.ColNum() != m2.ColNum() {
		return Matrix{}, errors.New("Invalid number of rows")
	}
	m := Matrix{
		rowNum: m1.RowNum() + m2.RowNum(),
		colNum: m1.ColNum(),
		rows:   make([]Array, m1.RowNum()+m2.RowNum()),
	}
	for i := 0; i < m1.RowNum(); i++ {
		m.rows[i] = m1.rows[i]
	}
	for i := 0; i < m2.RowNum(); i++ {
		m.rows[m1.RowNum()+i] = m2.rows[i]
	}

	return m, nil
}

func (m Matrix) ScalarMultiplication(a float64) Matrix {
	for i := 0; i < m.RowNum(); i++ {
		for j := 0; j < m.ColNum(); j++ {
			if m.rows[i][j].dtype == FloatType {
				m.rows[i][j].value = m.rows[i][j].value.(float64) * a
			} else {
				m.rows[i][j].value = (m.rows[i][j].value.(int64) * int64(a))
			}
		}
	}
	return m
}

func (m Matrix) RowNum() int {
	return m.rowNum
}
func (m Matrix) ColNum() int {
	return m.colNum
}

func (m Matrix) Shape() (int, int) {
	return m.RowNum(), m.ColNum()
}

func (m Matrix) Get(i, j int) float64 {
	return m.rows[i][j].value.(float64)
}

func Equal(m1, m2 Matrix) bool {
	if m1.RowNum() != m2.RowNum() || m1.ColNum() != m2.ColNum() {
		return false
	}
	for i := range m1.rows {
		for j := range m1.rows[i] {
			if abs(m1.Get(i, j)-m2.Get(i, j)) > 0.000001 {
				return false
			}
		}
	}
	return true
}
