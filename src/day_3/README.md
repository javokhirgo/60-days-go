# Day 3 - Functions, Multiple Returns, Error Handling

## what i learned

### function syntax
- go: `func name(params) returnType { ... }`
- java: `public static returnType name(params) { ... }`
- return type at the END in go, at the START in java

### multiple returns
- go can return multiple values: `func divide(a, b int) (int, int)`
- in java this is pain, need to create class or use array
- in go its built-in, just do `result, remainder := divide(10, 3)`

### error handling
- no exceptions in go like java
- functions return error as second value: `func square(x int) (int, error)`
- must check: `if err != nil { handle error }`
- return nil if no error: `return result, nil`
- return error: `return 0, errors.New("message")`

### variadic functions
- can take any number of args: `func sum(numbers ...int) int`
- use `...` before type
- inside function its just a slice

### named returns
- can name return values in function signature
- then use naked return (just `return`)
- still dont like it, prefer explicit returns

## tasks completed
- basic_functions/functions.go - 5 tasks ✓
- advanced_functions/advanced_functions.go - 5 tasks ✓

## key differences from java
- multiple returns built-in (java doesnt have this)
- explicit error handling (java uses try-catch)
- no function overloading in go
- return type at end (java at start)

## errors handling pattern
```go
result, err := someFunction()
if err != nil {
    fmt.Println("Error:", err)
    return
}
// use result here
```

## time spent: ~2 hours
