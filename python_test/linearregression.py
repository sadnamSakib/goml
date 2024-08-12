import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression

import matplotlib.pyplot as plt

# Read the CSV file
data = pd.read_csv("../Student_Performance.csv")
print(data.head())

# Extract the independent variables (X) and the dependent variable (y)
X = data[["Previous Scores"]]
y = data["Performance Index"]

# Create an instance of the LinearRegression model
model = LinearRegression()

# Fit the model to the data
model.fit(X, y)

# Predict the dependent variable
step_size = 0.1  # Define the step size

y_pred = model.predict(X)

# Plot the actual and predicted values
plt.scatter(X["Previous Scores"], y)
plt.plot(X["Previous Scores"], y_pred, color="red")
plt.xlabel("Previous Scores")
plt.ylabel("Performance Index")
plt.title("Actual vs Predicted Performance Index")
plt.show()

# Make a prediction for a specific rating and battery power
# rating = 4.2
# battery_power = 2000
# prediction = model.predict([[rating, battery_power]])
# print(
#     f"The predicted price for a rating of {rating} and battery power of {battery_power} is: {prediction[0]}"
# )
