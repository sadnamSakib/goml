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
	theta   numerics.Array
	x       numerics.Matrix
	y       numerics.Array
	columns []string
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
	fmt.Println(x.ColNum(), x.RowNum())
	fmt.Println(y.Len())
	x = x.Transpose()
	mY, err := numerics.NewMatrix(1, y.Len(), []numerics.Array{y})

	if err != nil {
		fmt.Println(err)
		return Regressor{}, err
	}
	ones, err := numerics.NewMatrix(1, x.RowNum(), []numerics.Array{numerics.Linspace(1, 1, x.RowNum())})
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
	var t numerics.Array
	for i := 0; i < theta.RowNum(); i++ {
		t = append(t, numerics.NewElement(theta.Get(i, 0), numerics.FloatType))
	}

	return Regressor{
		theta:   t,
		x:       x,
		y:       y,
		columns: independent_columns,
	}, nil
}

func (r Regressor) RowNum() int {
	return r.x.RowNum()
}

func (r Regressor) ColNum() int {
	return r.x.ColNum() - 1
}

func (r Regressor) Predict(features ...float64) float64 {

	if len(features) > r.x.ColNum()-1 {
		fmt.Println("Invalid number of feature values")
		return 0.0
	}
	Y := 0.0
	Y += r.theta[0].Get().(float64)

	for i := 1; i < r.x.ColNum(); i++ {
		Y += r.theta[i].Get().(float64) * float64(features[i-1])
	}

	return Y
}

func (r Regressor) Plot2D(columnName string) {
	p := plot.New()
	p.Title.Text = "Linear Regression"
	p.X.Label.Text = columnName
	p.Y.Label.Text = "Y"
	pts := make(plotter.XYs, r.x.RowNum())
	points := 0
	for i := 0; i < r.x.RowNum(); i++ {
		j := func() int {
			for j, v := range r.columns {
				if v == columnName {
					return j + 1
				}
			}
			return 1
		}()
		valX := r.x.GetColumn(j).Get(i, 0)
		valY := r.y[i].Get().(float64)

		pts[points] = plotter.XY{X: valX, Y: valY}
		points++
	}

	pointPlot, err := plotter.NewScatter(pts)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Add(pointPlot)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "regression.png"); err != nil {

		fmt.Println(err)
		return
	}

	pts = make(plotter.XYs, r.x.RowNum())
	for i := 0; i < r.x.RowNum(); i++ {
		j := func() int {
			for j, v := range r.columns {
				if v == columnName {

					return j + 1
				}
			}
			return 1
		}()
		valX := r.x.GetColumn(j).Get(i, 0)
		input := make([]float64, r.x.ColNum()-1)

		for k := 1; k < r.x.ColNum(); k++ {
			if k == j {
				input[j-1] = valX
				continue
			}
			input[k-1] = func(m numerics.Matrix) float64 {
				val := 0.0
				for i := 0; i < m.RowNum(); i++ {
					val += m.Get(i, 0)
				}
				return val / float64(m.RowNum())
			}(r.x.GetColumn(k))
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
		return
	}
	fmt.Println("Plot saved as regression.png")

}

func (r Regressor) Correlation(x string) float64 {
	xArray := func() numerics.Array {
		for i, v := range r.columns {
			if v == x {
				return func() numerics.Array {
					var a numerics.Array
					for j := 0; j < r.x.RowNum(); j++ {
						a.Append(numerics.NewElement(r.x.Get(j, i+1), numerics.FloatType))
					}
					return a
				}()
			}
		}
		return numerics.Array{}
	}()
	yArray := r.y
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

func (r Regressor) mean_squared_error() float64 {
	var mse float64
	for i := 0; i < r.x.RowNum(); i++ {
		predicted := r.Predict(r.x.Get(i, 1), r.x.Get(i, 2), r.x.Get(i, 3), r.x.Get(i, 4), r.x.Get(i, 5), r.x.Get(i, 6), r.x.Get(i, 7))
		actual := r.y[i].Get().(float64)
		mse += (actual - predicted) * (actual - predicted)
	}
	return mse / float64(r.x.RowNum())
}
