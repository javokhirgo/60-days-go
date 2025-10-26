package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func readNumber(scanner *bufio.Scanner, prompt string) (float64, error) {
	fmt.Print(prompt)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}
	return num, nil
}

func readOperation(scanner *bufio.Scanner) (string, error) {
	fmt.Print("Enter operation (+, -, *, /): ")
	scanner.Scan()
	op := strings.TrimSpace(scanner.Text())
	if op != "+" && op != "-" && op != "*" && op != "/" {
		return "", fmt.Errorf("invalid operation")
	}
	return op, nil
}

func calculate(calc Calculator, a, b float64, op string) {
	var result float64
	var err error

	switch op {
	case "+":
		result = calc.Add(a, b)
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)
	case "-":
		result = calc.Subtract(a, b)
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, result)
	case "*":
		result = calc.Multiply(a, b)
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, result)
	case "/":
		result, err = calc.Divide(a, b)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	}
}

func main() {
	calc := Calculator{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Simple Calculator ===")

	for {
		a, err := readNumber(scanner, "Enter first number: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		op, err := readOperation(scanner)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		b, err := readNumber(scanner, "Enter second number: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		calculate(calc, a, b, op)

		fmt.Print("\nContinue? (y/n): ")
		scanner.Scan()
		choice := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if choice != "y" && choice != "yes" {
			fmt.Println("Goodbye!")
			break
		}
		fmt.Println()
	}
}
