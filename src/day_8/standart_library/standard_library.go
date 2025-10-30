package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Day 8: Standard Library Packages ===")

	// TASK 1: Using strings package
	// Create a string "hello world"
	// Convert to uppercase and check if it contains "WORLD"
	// Print the results
	str := "hello world"
	str = strings.ToUpper(str)
	println(strings.Contains(str, "WORLD"))

	// TASK 2: Using math package
	// Calculate square root of 25
	// Calculate 2 raised to the power of 8
	// Find the maximum of 10 and 20
	// Print all results
	println(math.Sqrt(25))
	println(math.Pow(2, 8))
	println(math.Max(10, 20))

	// TASK 3: Using time package
	// Get current time
	// Print current hour, minute, and second
	// Format time as "YYYY-MM-DD"
	timeNow := time.Now()
	fmt.Printf("Current hour : %v\n", timeNow.Hour())
	fmt.Printf("Current hour : %v\n", timeNow.Minute())
	fmt.Printf("Current hour : %v\n", timeNow.Second())
	fmt.Println(timeNow.Format("2006-01-02"))

	// TASK 4: Using strings package
	// Split "apple,banana,orange" by comma
	// Join ["Hello", "World"] with space
	// Print both results
	arr := strings.Split("apple,banana,orange", ",")
	fmt.Println(arr)
	joint := []string{"Hello", "World"}
	helloWorld := strings.Join(joint, " ")
	println(helloWorld)

	// TASK 5: Using math package with loops
	// Calculate square root of numbers 1 to 10
	// Print each result
	for i := 1; i <= 10; i++ {
		fmt.Printf("Square root of %v is %v\n", i, math.Sqrt(float64(i)))
	}
}
