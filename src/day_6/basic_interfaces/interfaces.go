package main

import "fmt"

func main() {
	fmt.Println("=== Day 6: Interfaces Practice ===")

	// TASK 1: Define an Interface
	// Create a Shape interface with Area() float64 method
	// Create Rectangle and Circle structs
	// Both should implement the Shape interface
	r := Rectangle{
		Width:  10,
		Height: 6,
	}
	c := Circle{
		Radius: 5,
	}
	fmt.Println(r.Area())
	fmt.Println(c.Area())

	// TASK 2: Implement Interface
	// Create a Writer interface with Write() method
	// Create ConsoleWriter struct that implements Writer
	// Call Write() method
	console := ConsoleWriter{
		Word: "Hello world ",
	}
	console.Write()

	// TASK 3: Multiple Methods in Interface
	// Create an Animal interface with MakeSound() and Move() methods
	// Create Dog and Cat structs that implement Animal
	// Call both methods on each

	dog := Dog{
		Sound:  "Wow",
		Moving: "Fast",
	}
	cat := Cat{
		Sound:  "Meow",
		Moving: "Slow",
	}

	dog.MakeSound()
	dog.Move()
	cat.MakeSound()
	cat.Move()

	// TASK 4: Interface Variables
	// Create a Payment interface with ProcessPayment() method
	// Create CreditCard and PayPal structs

	var p Payment
	p = CreditCard{}
	p.ProcessPayment()

	p = PayPal{}
	p.ProcessPayment()

	// TASK 5: Function That Takes Interface
	// Create a DisplayInfo() function that takes a Shape interface
	// Call it with different shapes
	DisplayInfo(r)
	DisplayInfo(c)
}

// task1
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return 3.14 * r.Radius * r.Radius
}

// task2
type Writer interface {
	Write()
}

type ConsoleWriter struct {
	Word string
}

func (console ConsoleWriter) Write() {
	fmt.Println(console.Word)
}

// task3
type Animal interface {
	MakeSound()
	Move()
}

type Dog struct {
	Sound  string
	Moving string
}

func (d Dog) MakeSound() {
	fmt.Println(d.Sound)
}
func (d Dog) Move() {
	fmt.Println(d.Moving)
}

type Cat struct {
	Sound  string
	Moving string
}

func (c Cat) MakeSound() {
	fmt.Println(c.Sound)
}
func (c Cat) Move() {
	fmt.Println(c.Moving)
}

// task4
type Payment interface {
	ProcessPayment()
}

type CreditCard struct {
}

func (c CreditCard) ProcessPayment() {
	fmt.Println("Payment done by credit card")
}

type PayPal struct {
}

func (p PayPal) ProcessPayment() {
	fmt.Println("Payment done by paypal")
}

// task5
func DisplayInfo(s Shape) {
	fmt.Println(s.Area())
}
