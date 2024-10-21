package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Check if two command-line arguments are provided (input and output file names)
	if len(os.Args) != 2 {
		fmt.Println("Math skills usage: please enter <input file>")
		os.Exit(1) // Prints exit status 1. The program did not run successfully as the required arguments were not provided.
	}

	// Get the input & output file names from command-line arguemts
	inputFile := os.Args[1]
	// outputFile := os.Args[2]
	fmt.Println("Input filename: ", inputFile)
	// fmt.Println("Output filename: ", outputFile)

	// Read the contents from the input file as bytes: []byte var inputData
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err) // open <filename>: no such file or directory
		os.Exit(1)
	}

	// Convert the []byte inputData to string for line-by-line parsing.
	inputStr := string(inputData)
	// fmt.Println("Input data:\n", inputStr) // Print the string of inputData

	// Splitting lines: The strings.Split function separates each line of the file into a slice; split based on newlines
	lines := strings.Split(inputStr, "\n")
	// fmt.Println("Lines []string:", lines) // Print the lines []string

	// Declare the slice to hold the numbers
	var numbers []float64

	// Loop through each line, convert it to a float, and store it
	for _, line := range lines {

		// Trim leading and trailing whitespace
		line = strings.TrimSpace(line)
		// fmt.Println("Line:", line) // Print the line

		// Skip empty lines
		if line == "" {
			continue
		}

		// Parsing numbers: strconv.ParseFloat convert each line into a floating-point number.
		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error converting line to number:", err)
			continue
		}

		// Append the number the slice of numbers []float64
		numbers = append(numbers, number)
		// fmt.Println("Numbers slice:", numbers, "Number added to []:", number)
	}

	// Perform some math operation on the numbers
	// counter := 0
	l := len(numbers)

	// Calculate the Median:
	// Sort the slice of numbers.
	sort.Float64s(numbers)
	var median float64
	if l%2 == 1 {
		// Odd number of elements, take the middle one
		median = numbers[l/2]
	} else {
		// Even number of elements, take the average of the two middle ones
		median = (numbers[l/2-1] + numbers[l/2]) / 2
	}

	// Variance: Measures how far a set of numbers are spread out from their average (mean).
	// Variance is expressed in squared units of the original data.
	var variance float64
	var sqDiffSum float64

	// Step 1: Calculate the mean (average) of the dataset.
	var sum float64
	var mean float64
	for _, num := range numbers {
		sum += num
	}
	if len(numbers) > 0 {
		mean = sum / float64(l)
		// fmt.Println("Average: ", mean)
	} else {
		fmt.Println("No valid numbers found in the file.")
	}

	// Step 2: Calculate squared differences
	for _, num := range numbers {
		sqDiff := (num - mean) * (num - mean)
		sqDiffSum += sqDiff
	}

	// Step 3: Calculate variance. (float64(l-1) Use sample variance)
	variance = sqDiffSum / float64(l) // Use population variance

	// Standard Deviation: The square root of the variance.
	// It gives a measure of variability in the same units as the data.
	stdDev := math.Sqrt(variance)

	// Prepare the output string with all the results. Note %.2f is two floating points, %v is integer
	answers := fmt.Sprintf("Average: %.0f\nMedian: %.0f\nVariance: %.0f\nStandard Deviation: %.0f\n",
		mean, median, variance, stdDev)
	// 	math.Round(mean), math.Round(median), math.Round(variance), math.Round(stdDev))

	// Write the modified contents to the output file, 0644 6=rw Owner, 4=r Group, Others
	/* err = os.WriteFile(outputFile, []byte(answers), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	} */

	// Print the modified content to the terminal
	fmt.Println("Calculations:\n", strings.TrimSpace(answers))
}
