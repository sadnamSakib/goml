package tabular

type DataFrame map[string]Series

func NewDataFrame() DataFrame {
	return make(DataFrame)
}

func (df DataFrame) Append(colName string, col Series) {
	df[colName] = col
}
