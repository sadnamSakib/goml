package regressor

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/sadnamSakib/goml/numerics"
	"github.com/sadnamSakib/goml/tabular"
)

type Regressor struct {
	Theta   numerics.Array
	X       numerics.Matrix
	Y       numerics.Array
	Columns []string
}

func LinearRegression(df tabular.DataFrame, independent_columns []string, prediction_column string) (Regressor, error) {

	x, err := numerics.NewMatrix(
		len(independent_columns),
		df.RowNum(),
		df.GetColumnsAsArray(independent_columns...),
	)

	if err != nil {
		fmt.Println(err)
	}

	y := df.GetColumnsAsArray(prediction_column)[0]

	x = x.Transpose()
	mY, err := numerics.NewMatrix(1, y.Len(), []numerics.Array{y})
	if err != nil {
		fmt.Println(err)
		return Regressor{}, err
	}
	ones, err := numerics.NewMatrix(1, x.RowNum, []numerics.Array{numerics.Linspace(1, 1, x.RowNum)})
	if err != nil {
		fmt.Println(err)

		return Regressor{}, err
	}
	ones = ones.Transpose()

	x, err = numerics.AppendX(ones, x)
	if err != nil {
		fmt.Println(err)

		return Regressor{}, err
	}
	fmt.Println(x.RowNum, x.ColNum)

	mY = mY.Transpose()

	xT := x.Transpose()

	xTx, err := numerics.Multiply(xT, x)
	if err != nil {
		fmt.Println(err)

		return Regressor{}, err
	}

	xTxInv, err := xTx.Inverse()

	if err != nil {
		fmt.Println(err)
		return Regressor{}, err
	}
	xTy, err := numerics.Multiply(xT, mY)
	if err != nil {
		fmt.Println(err)
		return Regressor{}, err
	}

	theta, err := numerics.Multiply(xTxInv, xTy)
	if err != nil {
		fmt.Println(err)
		return Regressor{}, err
	}
	fmt.Println(theta)

	return Regressor{
		Theta:   theta.GetColumn(0),
		X:       x,
		Y:       y,
		Columns: independent_columns,
	}, nil
}

func (r Regressor) Predict(features ...float64) float64 {
	if len(features) != r.X.ColNum-1 {
		fmt.Println("Invalid number of feature values")
		return 0.0
	}
	Y := 0.0
	Y += r.Theta[0].Get().(float64)

	for i := 1; i < r.X.ColNum; i++ {
		Y += r.Theta[i].Get().(float64) * float64(features[i-1])
	}

	return Y
}

func (r Regressor) Plot2D(columnName string) {
	p := plot.New()
	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	pts := make(plotter.XYs, r.X.RowNum)
	points := 0
	for i := 0; i < r.X.RowNum; i++ {
		j := func() int {
			for j, v := range r.Columns {
				if v == columnName {

					return j + 1
				}
			}
			return 1
		}()
		valX := r.X.GetColumn(j)[i].Get().(float64)
		valY := r.Y[i].Get().(float64)

		pts[points] = plotter.XY{X: valX, Y: valY}
		points++
	}

	pointPlot, err := plotter.NewScatter(pts)
	if err != nil {
		fmt.Println(err)
	}
	p.Add(pointPlot)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "regression.png"); err != nil {
		fmt.Println(err)
	}

	pts = make(plotter.XYs, r.X.RowNum)
	for i := 0; i < r.X.RowNum; i++ {
		j := func() int {
			for j, v := range r.Columns {
				if v == columnName {

					return j + 1
				}
			}
			return 1
		}()
		valX := r.X.GetColumn(j)[i].Get().(float64)
		input := make([]float64, r.X.ColNum-1)
		for k := 1; k < r.X.ColNum; k++ {
			if k == j {
				input[j-1] = valX
				continue
			}
			input[k-1] = r.X.GetColumn(k).Mean()
		}
		valY := r.Predict(input...)
		pts[i] = plotter.XY{X: valX, Y: valY}
	}

	linePlot, err := plotter.NewLine(pts)
	if err != nil {
		fmt.Println(err)
	}
	p.Add(linePlot)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "regression.png"); err != nil {
		fmt.Println(err)
	}

}

func (r Regressor) Correlation(x string) float64 {
	xArray := func() numerics.Array {
		for i, v := range r.Columns {
			if v == x {
				return r.X.GetColumn(i + 1)
			}
		}
		return numerics.Array{}
	}()
	yArray := r.Y
	xMean := xArray.Mean()
	yMean := yArray.Mean()
	xStd := xArray.Std()
	yStd := yArray.Std()
	cov := 0.0
	for i := 0; i < xArray.Len(); i++ {
		cov += (xArray[i].Get().(float64) - xMean) * (yArray[i].Get().(float64) - yMean)
	}
	cov = cov / float64(xArray.Len())

	return cov / (xStd * yStd)
}
