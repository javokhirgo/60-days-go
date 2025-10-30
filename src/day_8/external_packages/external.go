package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("=== External Package Practice ===")

	// TASK 1: Install the color package
	// Run in terminal: go get github.com/fatih/color

	// TASK 2: Import it above (uncomment after go get)
	// import "github.com/fatih/color"

	// TASK 3: Print red text
	// Use: color.Red("This is red text")
	color.Red("This is red text")

	// TASK 4: Print green text
	// Use: color.Green("This is green text")
	color.Green("This is green text")

	// TASK 5: Print blue text
	// Use: color.Blue("This is blue text")
	color.Blue("This is blue text")
	// TASK 6: Print yellow text with background
	// Use: color.New(color.FgYellow, color.BgBlack).Println("Yellow on black")
	color.New(color.FgYellow, color.BgBlack).Println("Yellow on black")
}
