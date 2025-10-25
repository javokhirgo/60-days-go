# Day 7: CLI Calculator Project

## What I Built
A command-line calculator with 4 operations (Add, Subtract, Multiply, Divide).

## Key Concepts Applied
- **Structs**: Calculator type with Result field
- **Methods**: Attached operations to Calculator
- **Pointer Receivers**: Modify calculator state
- **Error Handling**: Return errors from Divide
- **Loops**: Input validation with `for`
- **Switch**: Route operations

## Best Practices Learned
1. **Pointer receivers** when modifying state
2. **Return errors** instead of printing inside methods
3. **Reusable functions** for repeated logic (input handling)
4. **Single loop** in menu instead of recursive calls
5. **Clear exit path** for user

## Files
- `calculator/calculator.go` - My version
- `calculator_best/calculator_best.go` - Best practices version

## Run
```bash
go run day_7/calculator/calculator.go
go run day_7/calculator_best/calculator_best.go
```


Week 1 Complete! ✅

