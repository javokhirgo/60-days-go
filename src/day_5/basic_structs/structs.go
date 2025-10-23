package main

import "fmt"

func main() {
	fmt.Println("=== Day 5: Structs and Methods ===")

	// TASK 1: Create a Struct
	// Create a Person struct with Name, Age, and Email fields
	// Create an instance and print it

	javokhir := Person{
		Name:  "Javokhir",
		Age:   23,
		Email: "javokhirgo@gmail.com",
	}
	fmt.Println(javokhir)

	// TASK 2: Access Struct Fields
	// Create a Rectangle struct with Width and Height
	// Access and print the fields
	// Create an instance and print it

	rec1 := Rectangle{
		Width:  15,
		Height: 20,
	}
	fmt.Println(rec1.Height)
	fmt.Println(rec1.Width)

	// TASK 3: Modify Struct Fields
	// Create a Circle struct with Radius
	// Modify the Radius and print it

	type Circle struct {
		Radius int
	}

	circle := Circle{
		Radius: 30,
	}
	circle.Radius = 40
	fmt.Println(circle.Radius)

	// TASK 4: Value Receiver Method
	// Create a Rectangle struct with Width and Height
	// Add a method Area() that calculates area
	// The method should NOT modify the struct

	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	fmt.Println("Area:", area)

	// TASK 5: Pointer Receiver Method
	// Create a Person struct with Name and Age
	// Add a method SetAge(age int) that modifies the Age
	// The method SHOULD modify the struct

	person := Person{Name: "Alice", Age: 25}
	person.SetAge(30)
	fmt.Println("New age:", person.Age)
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p *Person) SetAge(age int) {
	p.Age = age
}
