package main

import "fmt"

func main() {
	fmt.Println("=== Day 9: Basic Pointers Practice ===")

	// TASK 1: Create a variable and get its pointer
	// Create an int variable x = 10
	// Create a pointer ptr that points to x (use &)
	// Print the value of x
	// Print the address stored in ptr
	// Print the value at ptr (use *)
	x := 10
	ptr := &x
	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// TASK 2: Modify value through pointer
	// Create an int variable num = 5
	// Create a pointer to num
	// Change num to 20 using the pointer (use *ptr = 20)
	// Print num to verify it changed
	num := 5
	ptrNum := &num
	*ptrNum = 20
	fmt.Println(num)

	// TASK 3: Pointer to string
	// Create a string variable name = "Alice"
	// Create a pointer to name
	// Change name to "Bob" using the pointer
	// Print the original variable to see the change
	name := "Alice"
	prtName := &name
	*prtName = "Bob"
	fmt.Println(name)

	// TASK 4: Check nil pointer
	// Declare a pointer to int without initializing it (var ptr *int)
	// Check if it's nil
	// Then assign it to point to a new variable
	// Check if it's still nil
	var ptr2 *int
	xPointer := 10
	if ptr2 == nil {
		ptr2 = &xPointer
	}
	fmt.Println(ptr2 == nil)
	xPointer = 25
	fmt.Println(*ptr2)

	// TASK 5: Pointer with multiple variables
	// Create two int variables a = 10, b = 20
	// Create a pointer ptr
	// Point it to a, print *ptr
	// Point it to b, print *ptr
	// Show that one pointer can point to different variables
	a := 10
	b := 20
	ptr3 := &a
	fmt.Println(*ptr3)
	ptr3 = &b
	fmt.Println(*ptr3)

}
