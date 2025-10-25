package main

import (
	"fmt"
)

type Calculator struct {
	Result float64
}

func (c *Calculator) Add(a, b float64) {
	c.Result = a + b
	fmt.Printf("%v + %v = %.2f\n", a, b, c.Result)
}

func (c *Calculator) Subtract(a, b float64) {
	c.Result = a - b
	fmt.Printf("%v - %v = %.2f\n", a, b, c.Result)
}

func (c *Calculator) Multiply(a, b float64) {
	c.Result = a * b
	fmt.Printf("%v * %v = %.2f\n", a, b, c.Result)
}

func (c *Calculator) Divide(a, b float64) error {
	if b == 0 {
		return fmt.Errorf("division by zero")
	}
	c.Result = a / b
	fmt.Printf("%v / %v = %.2f\n", a, b, c.Result)
	return nil
}

func getNumber(prompt string) float64 {
	var num float64
	for {
		fmt.Printf(prompt)
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Invalid input. Enter a number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return num
	}
}

func getOperation() int {
	var op int
	for {
		fmt.Println("\nChoose operation:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Print("Enter choice (1-4): ")
		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Invalid input. Enter 1-4.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if op >= 1 && op <= 4 {
			return op
		}
		fmt.Println("Please enter 1, 2, 3, or 4.")
	}
}

func menu(c *Calculator) {
	for {
		fmt.Println("\n=== Calculator ===")
		a := getNumber("Enter first number: ")
		op := getOperation()
		b := getNumber("Enter second number: ")

		switch op {
		case 1:
			c.Add(a, b)
		case 2:
			c.Subtract(a, b)
		case 3:
			c.Multiply(a, b)
		case 4:
			if err := c.Divide(a, b); err != nil {
				fmt.Println("Error:", err)
			}
		}

		fmt.Print("\nContinue? (1=yes, 0=no): ")
		var choice int
		fmt.Scan(&choice)
		if choice != 1 {
			fmt.Println("Goodbye!")
			return
		}
	}
}

func main() {
	c := &Calculator{Result: 0}
	menu(c)
}
