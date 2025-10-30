# Day 6 - Interfaces and Type Assertions

## what i learned

### interfaces
- defining behavior that types must have
- syntax: `type Shape interface { Area() float64 }`
- implementation is AUTOMATIC (no "implements" keyword like java)
- if a type has the required methods, it implements the interface
- this was very confusing at first

### why use interfaces?
- write one function that works with multiple types
- example: one `PrintArea(s Shape)` works for Rectangle, Circle, Triangle
- without interfaces need separate function for each type

### empty interface
- `interface{}` accepts ANY type
- like Object in java but more flexible
- use when dont care about type

### type assertions
- convert interface to concrete type
- syntax: `s := i.(string)`
- safe version: `s, ok := i.(string)` - check if ok before using

### type switches
- check what type interface holds
- syntax: `switch v := i.(type) { case int: ... case string: ... }`

## tasks completed
- basic_interfaces/interfaces.go - 5 tasks ✓
- advanced_interfaces/advanced_interfaces.go - 5 tasks ✓

## java vs go
| java | go |
|------|-----|
| implements keyword | automatic |
| explicit declaration | implicit (duck typing) |
| no empty interface | interface{} |

## honestly
interfaces were hardest topic so far. took me a while to understand why we even need them. but now i get it - polymorphism without inheritance.

## time spent: ~2 hours
