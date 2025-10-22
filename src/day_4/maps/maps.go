package main

import "fmt"

func main() {
	fmt.Println("=== Day 4: Maps Practice ===")

	// TASK 1: Create and Initialize Maps
	// Create a map with string keys and int values
	// Add some key-value pairs: "apple": 5, "banana": 3, "orange": 8
	// Print the map
	fruits := make(map[string]int)
	fruits["apple"] = 5
	fruits["banana"] = 3
	fruits["orange"] = 4
	fmt.Println(fruits)

	// TASK 2: Map Operations
	// Create a map of student names to grades
	// Add, update, and delete entries
	// Print the map after each operation

	students := make(map[string]int)
	students["Ali"] = 5
	students["Vali"] = 3
	students["Gani"] = 4
	fmt.Println(students)    //original
	students["Gani"] = 5     //update
	fmt.Println(students)    //after update
	delete(students, "Gani") //delete
	fmt.Println(students)    //after delete

	// TASK 3: Map Access and Checking
	// Create a map and access values
	// Check if keys exist using the "ok" idiom
	// Handle cases where keys don't exist
	v, ok := students["Gani"]
	if ok {
		fmt.Printf("%v\n", v)
	} else {
		fmt.Println("Has no value for this key")
	}

	// TASK 4: Map Iteration
	// Create a map and iterate over it
	// Print both keys and values

	for k, v := range students {
		fmt.Printf("Value is %v when key is %v\n", v, k)

	}

	// TASK 5: Map Types
	// Create different types of maps:
	// - map[string]string (name to email)
	// - map[int]bool (number to even/odd)
	// Print both maps
	nameEmail := map[string]string{
		"Ali":  "ali@gmail.com",
		"Vali": "Vali@gmail.com",
	}
	numberEven := map[int]bool{
		4: true,
		3: false,
	}

	fmt.Println(nameEmail)
	fmt.Println(numberEven)

}
