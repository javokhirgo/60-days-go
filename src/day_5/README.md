# Day 5 - Structs and Methods

## what i learned

### structs
- like classes in java but simpler
- syntax: `type Person struct { Name string; Age int }`
- create: `person := Person{Name: "Alice", Age: 25}`
- access fields: `person.Name`, `person.Age`
- no constructors, no inheritance

### methods
- functions attached to structs
- defined OUTSIDE the struct (not inside like java)
- syntax: `func (r Rectangle) Area() int { ... }`

### value receivers vs pointer receivers
- value receiver: `func (r Rectangle) Area() int`
  - works on copy
  - use when reading/calculating
  - doesnt modify original
  
- pointer receiver: `func (r *Rectangle) SetWidth(w int)`
  - works on original
  - use when modifying
  - changes the struct

## when to use which receiver
- reading data → value receiver
- modifying data → pointer receiver
- this took me a while to understand

## tasks completed
- basic_structs/structs.go - 5 tasks ✓
- advanced_structs/advanced_structs.go - 5 tasks ✓

## java vs go
| java | go |
|------|-----|
| methods inside class | methods outside struct |
| has constructors | no constructors |
| has inheritance | no inheritance |
| objects always reference | structs are value types |

## time spent: ~2 hours
