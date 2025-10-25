package main

import (
	"errors"
	"fmt"
)

func main() {
	c := Calculator{Result: 0}
	menu(&c)
}

type Calculator struct {
	Result float64
}

func (c Calculator) Add(a, b float64) {
	c.Result = a + b
	fmt.Printf("%v + %v = ", a, b)
	fmt.Println(c.Result)
	menu(&c)
}
func (c Calculator) Subtract(a, b float64) {
	c.Result = a - b
	fmt.Printf("%v - %v = ", a, b)
	fmt.Println(c.Result)
	menu(&c)
}
func (c Calculator) Multiply(a, b float64) {
	c.Result = a * b
	fmt.Printf("%v * %v = ", a, b)
	fmt.Println(c.Result)
	menu(&c)
}
func (c Calculator) Divide(a, b float64) {
	if b == 0 {
		fmt.Println(errors.New("Divisor can't be 0"))
	} else {
		c.Result = a / b
		fmt.Printf("%v / %v = ", a, b)
		fmt.Println(c.Result)
		menu(&c)
	}
}
func enterFirstNumber() float64 {
	var firstNumber float64
	for {
		fmt.Printf("Write the first number : ")
		_, err := fmt.Scan(&firstNumber)
		if err != nil {
			fmt.Println("Enter the actual number !!!")
			var discard string
			fmt.Scanln(&discard)
		} else {
			return firstNumber
		}
	}
}
func getOperation() int {
	var operator int
	for {
		fmt.Println("What are you going to do ?")
		fmt.Println("1.Add")
		fmt.Println("2.Subtract")
		fmt.Println("3.Multiply")
		fmt.Println("4.Divide")
		fmt.Printf("Just enter the number of operation  you want : ")
		_, err := fmt.Scan(&operator)
		if err != nil {
			fmt.Println("Enter the actual number !!!")
			var discard string
			fmt.Scanln(&discard)
		} else if operator < 1 || operator > 4 {
			fmt.Println("Enter the valid operator number !!!")
		} else {
			return operator
		}
	}
}
func enterSecondNumber() float64 {
	var secondNumber float64
	for {
		fmt.Printf("Write the second number : ")
		_, err := fmt.Scan(&secondNumber)
		if err != nil {
			fmt.Println("Enter the actual number !!!")
			var discard string
			fmt.Scanln(&discard)
		} else {
			return secondNumber
		}
	}
}
func menu(c *Calculator) {
	var menu int
	fmt.Printf("Tap 1 to open calculator or any symbol to quit : ")
	fmt.Scan(&menu)
	if menu != 1 {
		return
	}
	a := enterFirstNumber()
	operator := getOperation()
	b := enterSecondNumber()
	switch operator {
	case 1:
		c.Add(a, b)
	case 2:
		c.Subtract(a, b)
	case 3:
		c.Multiply(a, b)
	case 4:
		c.Divide(a, b)
	}
}
