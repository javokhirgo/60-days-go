package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Conditionals Practice ===")

	// TASK 1: Basic If/Else
	// Check if a number is positive, negative, or zero
	var n int
	fmt.Printf("Enter number : ")
	fmt.Scan(&n)
	var result string
	if n == 0 {
		result = "0"
	} else if n > 0 {
		result = "positive"
	} else {
		result = "negative"
	}
	fmt.Printf("Your number %v is %v\n", n, result)

	// TASK 2: If with Short Statement
	// Use the short if syntax: if x := someFunction(); x > 0 { ... }
	if x := 10; x > 9 {
		x = 11
	}

	// TASK 3: Switch with Numbers
	// Create a switch statement for day of week (1=Monday, 2=Tuesday, etc.)
	switch int(int(time.Now().Weekday())) {
	case 1:
		fmt.Printf("%v=Monday\n", int(time.Now().Weekday()))
	case 2:
		fmt.Printf("%v=Tuesday\n", int(time.Now().Weekday()))
	case 3:
		fmt.Printf("%v=Wednesday\n", int(time.Now().Weekday()))
	case 4:
		fmt.Printf("%v=Thursday\n", int(time.Now().Weekday()))
	case 5:
		fmt.Printf("%v=Friday\n", int(time.Now().Weekday()))
	case 6:
		fmt.Printf("%v=Saturday\n", int(time.Now().Weekday()))
	case 7:
		fmt.Printf("%v=Sunday\n", int(time.Now().Weekday()))
	default:
		fmt.Printf("Invalid day\n")

	}

	// TASK 4: Switch with Strings
	// Create a switch for weather conditions ("sunny", "rainy", "cloudy")
	switch "sunny" {
	case "sunny":
		fmt.Printf("Sunny\n")
	case "rainy":
		fmt.Printf("Rainy\n")
	case "cloudy":
		fmt.Printf("Cloudy\n")
	}
	// TASK 5: Switch with Multiple Cases
	// Use multiple cases in one switch: case 1, 2, 3: for weekdays
	switch int(time.Now().Weekday()) {
	case 1, 2, 3, 4, 5:
		fmt.Printf("Weekday\n")
	default:
		fmt.Printf("Weekend\n")

	}
}
