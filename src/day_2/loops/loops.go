package main

import "fmt"

func main() {
	fmt.Println("=== Go Loops Practice ===")

	// TASK 1: Traditional For Loop
	// Count from 1 to 10 and print each number
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// TASK 2: While-like Loop
	// Count down from 10 to 1 using for (like a while loop)
	i := 10
	for i >= 1 {
		fmt.Println(i)
		i--
	}

	// TASK 3: Range Loop
	// Iterate over a string and print each character
	name := "Javokhir"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%q\n", name[i])
	}

	// TASK 4: Loop with Break
	// Count from 1 to 100, but stop when you reach 7

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i == 7 {
			break
		}
	}

	// TASK 5: Nested Loops
	// Create a multiplication table (1x1=1, 1x2=2, etc.) up to 5x5

	for i := 1; i <= 5; i++ {
		for n := 1; n <= 5; n++ {
			fmt.Printf("%v x %v = %v\n", i, n, i*n)
		}
	}

}
