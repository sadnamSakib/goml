# **goml**
A lightweight module in go which supports basic data analysis, multi-dimensional array functions and regression operations.

# **Context Diagram**
<br>
<img src="Technical Diagrams.jpg" alt="vContext Diagram of the goml module" width="1000"/>
        <br><br>

# **Structure**
- `goml` : Module containing the following packages
    - `tabular` : Package which contains data retrieval , manipulation and analysis functions similar to pandas in python.
        - `Types`
            - Dataframe
                - head
                - tail
                - iloc
                - loc
                - drop
                - dropna
                - fillna
                - sort
                - groupBy
                - sortBy
                - concat
                - shape
                - apply
                - filter
                - unique
                - count
                - max
                - min
                - sum
                - mean
            - Series
                - `NewSeries`
                - `Append`
                - `String`
                - `Len`
                - `Sort`
                - `SortCopy`
                - Sum
                - Mean
                - Apply
                - Min
                - Max
                - Unique
                - Filter
        - `Functions`
            - readCsv
            - writeCsv
            
    - `numerics` : Package which contains multi dimensional array and matrix operations similar to numpy in python.
        - `Types`
            - Array
            - Matrix
        - `Functions`
            - array
            - matrix
        - `Methods`
            - shape
            - reshape
            - transpose
            - add
            - subtract
            - multiply
            - divide
            - dot
            - inverse
            - determinant
            - trace
            - eigenValues
            - eigenVectors
            - svd
            - adjoint
            - rank
            - pca
            - sort
            - unique
            - filter
            - apply
            - max
            - min
            - sum
            - mean

    - `regression` : Package which contains regression operations similar to sklearn in python.
        - `Types`
            - LinearRegression
            - LogisticRegression
        - `Functions`
            - linearRegression
            - logisticRegression
        - `Methods`
            - fit
            - predict
            - score
            - coefficients
            - intercept
            - mse
            - rmse
            - r2
            - confusionMatrix
            - accuracy
            - precision
            - recall
            - f1
            - roc
            - auc

# **Goals**
- To provide a lightweight module in go which supports basic data analysis, multi-dimensional array functions and regression operations.
- To provide a module which can be used for data analysis and machine learning in go.
- To use the concepts of go routines and channels to provide a faster and efficient module.
- To use the features of go to provide a module which is easy to use and understand.
