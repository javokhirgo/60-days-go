package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== Day 3: Advanced Functions Practice ===")

	// TASK 1: Variadic Function
	// Create a function that takes any number of integers and returns their sum
	// func sum(numbers ...int) int

	// TASK 2: Function with Multiple Parameters
	// Create a function that takes 3 numbers and returns the largest one
	// func findMax(a, b, c int) int

	// TASK 3: Function with String Parameters
	// Create a function that concatenates two strings: func concat(str1, str2 string) string

	// TASK 4: Error Handling Best Practices
	// Create a function that validates age: func validateAge(age int) error
	// Return different error messages for different invalid cases

	// TASK 5: Function Composition
	// Create helper functions and use them together
	// func isEven(n int) bool
	// func square(n int) int
	// Use them together to square even numbers only
}

func sum(numbers ...int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}
func findMax(a, b, c int) int {
	max := a
	if max < b {
		max = b
	}
	if max < c {
		max = c
	}
	return max
}
func concat(str1, str2 string) string {
	return str1 + str2
}
func validateAge(age int) error {
	if age < 0 {
		return errors.New("Age can't be negative ")
	} else if age == 0 {
		return errors.New("Age can't be zero")
	}
	return nil
}
func isEven(n int) bool {
	return n%2 == 0
}
func square(n int) int {
	if isEven(n) {
		return n * n
	} else {
		return 0
	}
}
