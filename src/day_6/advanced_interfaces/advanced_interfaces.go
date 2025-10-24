package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Day 6: Advanced Interfaces (Empty Interfaces & Type Assertions) ===")

	// TASK 1: Empty Interface Basics
	// Create a variable with empty interface type
	// Store different types in it (string, int, bool)
	// Print each value
	var variable interface{} = "Javokhir"
	fmt.Println(variable)
	variable = 25
	fmt.Println(variable)
	variable = true
	fmt.Println(variable)

	// TASK 2: Function with Empty Interface
	// Create a function that accepts interface{} parameter
	// Call it with different types
	// Print the values
	printValues("String")
	printValues(25)
	printValues(true)

	// TASK 3: Type Assertion Basics
	// Create an interface variable
	// Use type assertion to get the concrete type
	// Print the value
	var interValue interface{} = "Javokhir"
	name := interValue.(string)
	fmt.Printf(name + "\n")

	// TASK 4: Type Assertion with "ok" Check
	// Create an interface variable
	// Use type assertion with "ok" to safely check type
	// Handle both success and failure cases
	var interValue2 interface{} = "Javokhir"
	name2, ok := interValue2.(string)
	if ok {
		fmt.Println(name2)
	} else {
		fmt.Println("It is not string")
	}

	// TASK 5: Type Switch
	// Create a function that uses type switch
	// Check different types and print appropriate messages
	findType("Java")
	findType(25)
	findType(25.6)
	findType(true)

}
func printValues(i interface{}) {
	fmt.Println(i)
}

func findType(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	case bool:
		fmt.Println("Bool:", v)
	default:
		fmt.Println("Unknown type")
	}

}
