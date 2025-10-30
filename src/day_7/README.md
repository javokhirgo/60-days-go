# Day 7 - CLI Calculator Project

## what i built
command-line calculator with add, subtract, multiply, divide operations

## what i used
- structs (Calculator type)
- methods with pointer receivers
- error handling (for divide by zero)
- loops (for menu)
- switch statement (for operation routing)
- fmt.Scan for user input

## problems i had
- first tried recursion for menu loop, but better to use single for loop
- had issue with invalid input causing infinite loop
- fixed with fmt.Scanln(&discard) to clear buffer

## files
- calculator/calculator.go - my version
- calculator_best/calculator_best.go - better version (made by mentor)
- calculator_pro/calculator.go - production style version

## what i learned from best practice version
- separate calculation from display
- return values from methods instead of printing inside
- use bufio for better input handling
- string operators not just numbers
- cleaner error handling

## improvements i need
- better input validation
- cleaner code structure
- more separation of concerns

week 1 done ✓

## time spent: ~2 hours
