package main

import "fmt"

// TODO: Complete the tasks below
// Remove this comment and implement your solutions

func main() {
	fmt.Println("=== Day 2: Variables Practice ===")

	// TASK 1: Declare variables using different methods
	// Use var, :=, and multiple declaration
	var surname = "Foziljonov"
	var birthYear int = 2002
	var interest, favouriteSport = "Football", "Football"

	// TASK 2: Create variables for your personal info
	// name, age, city, country, isStudent
	var name, age, city, country, isStudent = "Javokhir", 23, "Tashkent", "Uzbekistan", false

	// TASK 3: Use type inference and explicit typing
	// Try both ways for the same variable
	var salary = 1000
	var salary2 float64 = 1000.50

	// TASK 4: Print all variables with their types
	// Use %T to show the type of each variable
	fmt.Printf("surname's type is : %T\n", surname)
	fmt.Printf("birthYear's type is : %T\n", birthYear)
	fmt.Printf("interest's type is : %T\n", interest)
	fmt.Printf("favouriteSport's type is : %T\n", favouriteSport)
	fmt.Printf("name's type is : %T\n", name)
	fmt.Printf("age's type is : %T\n", age)
	fmt.Printf("city's type is : %T\n", city)
	fmt.Printf("country's type is : %T\n", country)
	fmt.Printf("isStudent's type is : %T\n", isStudent)
	fmt.Printf("salary's type is : %T\n", salary)
	fmt.Printf("salary2's type is : %T\n", salary2)

	// TASK 5: Create zero value variables and print them
	// Show what Go's zero values look like
	var zeroString string
	var zeroInt int
	fmt.Printf("zeroString can be seen only using (quoteString) : %q\n", zeroString)
	fmt.Printf("ZeroInt value is %v\n", zeroInt)

}
