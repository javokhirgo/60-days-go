package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Day 2: Constants Practice ===")

	// TASK 1: Create mathematical constants
	// Pi, E, GoldenRatio
	const Pi = 3.14159
	const E = 2.71828
	const GoldenRatio = 1.61803

	// TASK 2: Create application constants
	// AppName, Version, MaxUsers (exported)
	const AppName = "MyApp"
	const Version = "1.0.0"
	const MaxUsers = 1000

	// TASK 3: Create internal constants
	// defaultPort, maxRetries, timeout (unexported)
	const defaultPort = 8080
	const maxRetries = 3
	const timeout = 10 * time.Second

	// TASK 4: Use constants in calculations
	// Calculate circle area using Pi constant
	const circleArea = Pi * 10 * 10

	// TASK 5: Print all constants with their types
	// Show that constants have types too
	fmt.Printf("Pi's type is %T\n", Pi)
	fmt.Printf("E's type is %T\n", E)
	fmt.Printf("GoldenRatio's type is %T\n", GoldenRatio)
	fmt.Printf("AppName's type is %T\n", AppName)
	fmt.Printf("Version's type is %T\n", Version)
	fmt.Printf("MaxUsers's type is %T\n", MaxUsers)
	fmt.Printf("defaultPort's type is %T\n", defaultPort)
	fmt.Printf("maxRetries's type is %T\n", maxRetries)
	fmt.Printf("timeout's type is %T\n", timeout)
	fmt.Printf("circleArea's type is %T\n", circleArea)
}
