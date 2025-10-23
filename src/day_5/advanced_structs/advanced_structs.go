package main

import "fmt"

func main() {
	fmt.Println("=== Day 5: Advanced Structs and Methods ===")

	// TASK 1: Multiple Methods
	// Create a Calculator struct
	// Add methods: Add(a, b int) int, Subtract(a, b int) int
	c := Calculator{}
	fmt.Println(c.Add(5, 5))
	fmt.Println(c.Subtract(10, 5))

	// TASK 2: Methods with Return Values
	// Create a BankAccount struct with Balance float64
	// Add method GetBalance() float64 that returns the balance
	b := BankAccount{Balance: 1000000}
	fmt.Println(b.GetBalance())

	// TASK 3: Methods that Modify
	// Create a BankAccount struct with Balance float64
	// Add method Deposit(amount float64) that adds to balance
	// Add method Withdraw(amount float64) that subtracts from balance
	// Use pointer receivers
	b.Deposit(1000000)
	fmt.Println(b.Balance)
	b.Withdraw(2000000)
	fmt.Println(b.Balance)

	// TASK 4: Methods with Different Receiver Types
	// Create a Student struct with Name and Grades []int
	// Add method GetAverage() float64 (value receiver)
	// Add method AddGrade(grade int) (pointer receiver)
	s := Student{
		Name:   "Javokhir",
		Grades: []int{5, 5, 5, 5, 5},
	}
	fmt.Println(s.GetAverage())
	fmt.Println(s.Grades)
	s.AddGrade(5)
	fmt.Println(s.Grades)

	// TASK 5: Real-World Example
	// Create a Book struct with Title, Author, Pages
	// Add method IsLong() bool (returns true if Pages > 300)
	// Add method UpdatePages(pages int) (modifies Pages)
	book := Book{
		Title:  "Ikki eshik orasi",
		Author: "O'tkir Hoshimov",
		Pages:  700,
	}
	fmt.Println(book.isLong())
	fmt.Println(book.Pages)
	book.UpdatePages(800)
	fmt.Println(book.Pages)

}

type Calculator struct {
}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Subtract(a, b int) int {
	return a - b
}

type BankAccount struct {
	Balance float64
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	b.Balance -= amount
}

type Student struct {
	Name   string
	Grades []int
}

func (s Student) GetAverage() float64 {
	var sum int
	for _, v := range s.Grades {
		sum += v
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (book Book) isLong() bool {
	return book.Pages > 300
}

func (b *Book) UpdatePages(pages int) {
	b.Pages = pages
}
