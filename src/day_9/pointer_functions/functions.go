package main

import "fmt"

func main() {
	fmt.Println("=== Day 9: Pointers with Functions ===")

	// TASK 1: Value parameter (no change to original)
	// Create a function incrementValue(x int) that adds 1 to x
	// In main, create num = 10
	// Call incrementValue(num)
	// Print num - it should still be 10 (not changed)
	num := 10
	incrementValue(num)
	fmt.Println(num)

	// TASK 2: Pointer parameter (changes original)
	// Create a function incrementPointer(x *int) that adds 1 to *x
	// In main, create num = 10
	// Call incrementPointer(&num)
	// Print num - it should be 11 (changed!)
	incrementPointer(&num)
	fmt.Println(num)

	// TASK 3: Swap function with values (doesn't work)
	// Create a function swapValue(a, b int) that tries to swap them
	// In main, create x = 5, y = 10
	// Call swapValue(x, y)
	// Print x and y - they won't be swapped
	x := 5
	y := 10
	swapValue(x, y)
	fmt.Println(x)
	fmt.Println(y)

	// TASK 4: Swap function with pointers (works!)
	// Create a function swapPointer(a, b *int) that swaps *a and *b
	// In main, create x = 5, y = 10
	// Call swapPointer(&x, &y)
	// Print x and y - they should be swapped!

	swapPointer(&x, &y)
	fmt.Println(x)
	fmt.Println(y)

	// TASK 5: Function returning pointer
	// Create a function createInt(value int) *int that:
	//   - Creates a local variable with the value
	//   - Returns pointer to it
	// Call it and print the result
	// (This shows Go handles memory safely)
	fmt.Println(createInt(25))
}
func incrementValue(x int) {
	x = x + 1
}
func incrementPointer(x *int) {
	*x = *x + 1
}
func swapValue(a, b int) {
	temp := a
	a = b
	b = temp
}
func swapPointer(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func createInt(value int) *int {
	lv := value
	return &lv
}
