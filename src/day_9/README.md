# Day 9: Pointers and Memory

Today I learned about pointers in Go - how to work with memory addresses directly.

## What I Learned

### Basic Pointers
- `&x` - gets memory address of x
- `*ptr` - gets value at that address (dereferencing)
- `*int` - pointer type (pointer to int)
- Pointers can be `nil` (no address yet)

```go
x := 10
ptr := &x       // ptr holds address of x
fmt.Println(*ptr)  // prints 10
*ptr = 20       // changes x to 20
```

### Pointers with Functions
**Value parameters** - function gets a copy, can't modify original:
```go
func change(x int) {
    x = 99  // only changes the copy
}
```

**Pointer parameters** - function can modify original:
```go
func change(x *int) {
    *x = 99  // changes the original
}
```

This is how we do the classic swap:
```go
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}
```

### Pointers with Structs
Go automatically dereferences struct pointers - no need for `(*ptr).field`:
```go
p := &Person{Name: "Ali", Age: 25}
p.Name  // Go does (*p).Name automatically
```

**Value receiver** - doesn't modify original:
```go
func (p Person) Birthday() {
    p.Age++  // only changes copy
}
```

**Pointer receiver** - modifies original:
```go
func (p *Person) Birthday() {
    p.Age++  // changes original
}
```

### Nested Pointers
Can have pointers inside structs:
```go
type Person struct {
    Name    string
    Address *Address  // pointer to another struct
}
```

### Memory Safety
Go handles memory automatically - even if you return pointer to local variable, Go moves it to heap (escape analysis). No dangling pointers!

## Key Points
- Use `&` to get address, `*` to get value
- Pointer parameters let functions modify original values
- Arrays pass by value (copy), slices pass by reference
- Nil pointers cause panic - always check before use
- Go's automatic dereferencing makes struct pointers easy to use

---

**Practice Files:**
- `basic_pointers/pointers.go` - Basic pointer operations
- `pointer_functions/functions.go` - Pointers with functions
- `pointer_structs/structs.go` - Pointers with structs

