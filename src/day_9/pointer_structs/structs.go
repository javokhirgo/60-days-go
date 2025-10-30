package main

import "fmt"

func main() {
	fmt.Println("=== Day 9: Pointers with Structs ===")

	// TASK 1: Create a struct with pointer
	// Define a Person struct with Name string and Age int
	// Create a person
	// Create a pointer to the person
	// Access fields through pointer (ptr.Name, ptr.Age)
	// Go automatically dereferences for you!
	person := Person{
		"Ali",
		&Address{
			"Tashkent",
		},
		25,
	}
	person_ptr := &person
	fmt.Println(person_ptr.Name)
	fmt.Println(person_ptr.Age)

	// TASK 2: Value receiver method (doesn't modify)
	// Add a method Birthday(p Person) that increments age
	// Create a person age 25
	// Call Birthday()
	// Print age - still 25 (not modified)
	Birthday(person)
	fmt.Println(person.Age)

	// TASK 3: Pointer receiver method (modifies)
	// Add a method BirthdayPointer(p *Person) that increments age
	// Create a person age 25
	// Call BirthdayPointer()
	// Print age - now 26 (modified!)
	BirthdayPointer(&person)
	fmt.Println(person.Age)

	// TASK 4: Nil pointer check
	// Create a pointer to Person without initializing (var ptr *Person)
	// Check if it's nil before using it
	// Then create a person and assign to ptr
	// Now use it safely
	var ptr *Person
	fmt.Println(ptr == nil)
	ptr = &person
	fmt.Println(ptr.Name)

	// TASK 5: Pointer to struct in another struct
	// Create an Address struct with City string
	// Create a Person struct with Name string and Address *Address (pointer!)
	// Create a person with an address
	// Modify the address through the person
	// This shows nested pointers
	person.Address = &Address{
		"Fergana",
	}
	fmt.Println(person.Address.City)

}

type Person struct {
	Name    string
	Address *Address
	Age     int
}

func Birthday(p Person) {
	p.Age++
}
func BirthdayPointer(p *Person) {
	p.Age++
}

type Address struct {
	City string
}
