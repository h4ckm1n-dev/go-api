import csv
import random
from datetime import datetime, timedelta


# Function to generate random dates
def random_date(start, end):
    return start + timedelta(days=random.randint(0, int((end - start).days)))


# Set the number of lines
num_lines = 100000

# Generate data
header = ["ID", "Name", "Value", "Category", "Date", "Quantity", "Price"]
categories = ["CategoryA", "CategoryB", "CategoryC"]
start_date = datetime.strptime("2023-01-01", "%Y-%m-%d")
end_date = datetime.strptime("2024-12-31", "%Y-%m-%d")

data = []
for i in range(1, num_lines + 1):
    data.append(
        [
            i,
            f"Item{i}",
            f"Value{i}",
            random.choice(categories),
            random_date(start_date, end_date).strftime("%Y-%m-%d"),
            random.randint(1, 50),
            round(random.uniform(5.0, 50.0), 2),
        ]
    )

# Write to CSV file
with open("data.csv", "w", newline="") as file:
    writer = csv.writer(file)
    writer.writerow(header)
    writer.writerows(data)

print(f"CSV file created successfully with {num_lines} lines.")
