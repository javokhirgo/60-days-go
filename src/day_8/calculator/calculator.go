package calculator

// TASK 1: Create exported functions
// Create Add(a, b int) int
// Create Subtract(a, b int) int
// Create Multiply(a, b int) int
// Create Divide(a, b int) float64
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

// TASK 2: Create unexported helper function
// Create validate(num int) bool (returns false if zero)
func validate(num int) bool {
	return num != 0
}

// TASK 3: Create exported constant
// Create MaxOperations = 100
const MaxOperations = 100

// TASK 4: Add all functions above
