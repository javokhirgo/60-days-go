package main

import "fmt"

func main() {
	fmt.Println("=== Day 4: Arrays Practice ===")

	// TASK 1: Create and Initialize Arrays
	// Create an array of 5 integers and initialize with values 1,2,3,4,5
	// Print the array
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// TASK 2: Array Operations
	// Create an array of 3 strings: "Go", "Java", "Python"
	// Print each element using a loop
	languages := [3]string{"Go", "Java", "Python"}
	for _, v := range languages {
		fmt.Println(v)
	}

	// TASK 3: Array Length and Indexing
	// Create an array of 4 floats: 1.1, 2.2, 3.3, 4.4
	// Print the length and access each element by index
	floatNums := [4]float64{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < len(floatNums); i++ {
		fmt.Println(floatNums[i])
	}

	// TASK 4: Array Modification
	// Create an array of 3 booleans: true, false, true
	// Change the second element to true and print the array
	booleans := [3]bool{true, false, true}
	booleans[1] = true
	for _, v := range booleans {
		fmt.Println(v)
	}

	// TASK 5: Array Comparison
	// Create two arrays with same values and compare them
	// Print whether they are equal
	same1 := [3]int{1, 2, 3}
	same2 := [3]int{1, 2, 3}
	fmt.Printf("%v\n", same1 == same2)

}
