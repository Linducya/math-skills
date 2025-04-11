# Math Skills: Statistical Calculations

This project, written in **Go**, is designed to help you calculate key statistical measures from a dataset. The program reads data from a file and computes the following statistics:

- **Average**
- **Median**
- **Variance**
- **Standard Deviation**

## Objectives

The purpose of this project is to:
1. Develop a program that reads numerical data from a file.
2. Perform statistical calculations on the data.
3. Output the results in a clear and formatted manner.

## Instructions

### Input
Your program must read data from a file provided as an argument. The file will contain one numerical value per line, representing a statistical population. For example:
```
189
113
121
114
145
110
...
```

### Execution
To run your program, use a command similar to the following
1. Use the provided `data.txt` file or create your own input file with numerical data. Each number should be on a new line.
2. Run the program with the input file as an argument:
   ```bash
   $ go run main.go data.txt

### Output
After reading the file, the program must compute the requested statistics and print the results in the following format (example values shown):

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

Please note that the values should be rounded to the nearest integer.


## Testing

To ensure the correctness of your program, you can create your own test cases. Below are some example datasets and their expected results to help you verify your implementation.

### Example Test Cases

#### Test Case 1: Small Dataset
**Input (`data.txt`):**
```
10
20
30
40
50
```
**Expected Output:**
```
Average: 30
Median: 30
Variance: 200
Standard Deviation: 14
```

#### Test Case 2: Large Dataset
**Input (`data.txt`):**
```
10
20
30
40
50
60
70
80
90
100
```
**Expected Output:**
```
Average: 55
Median: 55
Variance: 825
Standard Deviation: 29
```

## Learning Outcomes

This project will help you learn about:
- Statistics and Mathematics
- Average
- Median
- Variance
- Standard Deviation
