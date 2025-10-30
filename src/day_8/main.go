package main

import (
	"day_8/calculator"
	"fmt"
)

func main() {
	fmt.Println("=== Day 8: Using Custom Packages ===")

	// TASK 1: Use Add function from calculator package
	// Call calculator.Add(5, 3)
	// Print the result
	fmt.Println("Add:", calculator.Add(5, 3))

	// TASK 2: Use Subtract function
	// Call calculator.Subtract(10, 4)
	// Print the result
	fmt.Println("Subtract:", calculator.Subtract(10, 4))

	// TASK 3: Use Multiply function
	// Call calculator.Multiply(6, 7)
	// Print the result
	fmt.Println("Multiply:", calculator.Multiply(6, 7))

	// TASK 4: Use Divide function
	// Call calculator.Divide(15, 3)
	// Print the result
	fmt.Println("Divide:", calculator.Divide(15, 3))

	// TASK 5: Use the constant
	// Print calculator.MaxOperations
	fmt.Println("MaxOperations:", calculator.MaxOperations)
}
