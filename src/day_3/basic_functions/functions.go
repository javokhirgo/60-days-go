package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Day 3: Functions, Multiple Returns, Error Handling ===")

	// TASK 1: Basic Function
	// Create a function that adds two numbers and returns the result
	// Call it and print the result
	fmt.Printf("%v\n", add(4, 5))

	// TASK 2: Function with Multiple Returns
	// Create a function that divides two numbers and returns (result, remainder)
	// Handle the case where divisor is 0
	//		fmt.Printf("Divisor can't be 0")
	result, reminder := divide(6, 2)
	fmt.Printf("Result is : %v\nReminder is : %v\n", result, reminder)

	// TASK 3: Error Handling Function
	// Create a function that takes a number and returns (square, error)
	// Return error if number is negative
	square, error := square(0)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Square:", square)
	}

	// TASK 4: Named Return Values
	// Create a function with named returns: func calculate(x, y int) (sum, product int)
	// Use "naked return"

	// TASK 5: Error Handling with Multiple Returns
	// Create a function that calculates area of rectangle: func area(length, width float64) (float64, error)
	// Return error if length or width is negative or zero

}
func add(a, b int) int {
	return a + b
}
func divide(a, b int) (int, int) {
	if b == 0 {
		return 0, 0
	} else {
		return a / b, a % b
	}
}
func square(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Number is negative")
	} else {
		return math.Sqrt(a), nil
	}
}
func calculate(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return
}
func area(length, width float64) (float64, error) {
	if length <= 0 || width <= 0 {
		return 0, errors.New("Can't be negative ")
	} else {
		return length * width, nil
	}
}
