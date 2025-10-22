package main

import "fmt"

func main() {
	fmt.Println("=== Day 4: Slices Practice ===")

	// TASK 1: Create and Initialize Slices
	// Create a slice with values 1,2,3,4,5
	// Print the slice and its length
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// TASK 2: Slice Operations
	// Create a slice of strings: "Go", "Java", "Python", "C++"
	// Use append to add "Rust" to the slice
	// Print the updated slice
	languages := []string{"Go", "Java", "Python", "C++"}
	languages = append(languages, "Rust")
	fmt.Println(languages)

	// TASK 3: Slice Slicing
	// Create a slice with numbers 1-10
	// Create new slices: first 5 elements, last 5 elements, middle 3 elements
	// Print all slices
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstFive := numbers[0:5]
	lastFive := numbers[5:10]
	fmt.Println(numbers)
	fmt.Println(firstFive)
	fmt.Println(lastFive)

	// TASK 4: Slice Capacity
	// Create a slice with make() specifying length=3, capacity=10
	// Add 7 elements to it
	// Print length and capacity
	slice := make([]int, 3, 10)
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(slice)) //10
	fmt.Println(cap(slice)) //10

	// TASK 5: Slice Modification
	// Create a slice and modify elements
	// Show how slices share underlying data
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	elements[0] = 11
	fmt.Println(elements)

}
