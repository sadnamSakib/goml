import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression

import matplotlib.pyplot as plt

# Read the CSV file
data = pd.read_csv("../Mobile-Price-Prediction-cleaned_data.csv")
print(data.head())

# Extract the independent variables (X) and the dependent variable (y)
X = data[["Ratings", "Battery_Power"]]
y = data["Price"]

# Create an instance of the LinearRegression model
model = LinearRegression()

# Fit the model to the data
model.fit(X, y)

# Predict the dependent variable
step_size = 0.1  # Define the step size

y_pred = model.predict(X)

# Plot the actual and predicted values
plt.scatter(X["Ratings"], y)
plt.scatter(X["Battery_Power"], y)
plt.plot(X["Ratings"], y_pred, color="red")
plt.plot(X["Battery_Power"], y_pred, color="blue")
plt.xlabel("Ratings and Battery Power")
plt.ylabel("Price")
plt.title("Linear Regression: Ratings and Battery Power vs Price")
plt.legend(["Ratings", "Battery Power"])
plt.show()

# Make a prediction for a specific rating and battery power
rating = 4.2
battery_power = 2000
prediction = model.predict([[rating, battery_power]])
print(
    f"The predicted price for a rating of {rating} and battery power of {battery_power} is: {prediction[0]}"
)
