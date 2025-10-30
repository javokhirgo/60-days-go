# Day 4 - Arrays, Slices, Maps

## what i learned

### arrays (fixed size)
- syntax: `var arr [5]int` or `arr := [5]int{1,2,3,4,5}`
- size is part of type
- value type (copied when passed to functions)
- can compare with `==`
- in java arrays are reference types, in go they're value types

### slices (dynamic arrays)
- syntax: `var slice []int` or `slice := []int{1,2,3}`
- no fixed size
- reference type (shared when passed to functions)
- can grow with `append()`
- use `make()` to create with capacity: `make([]int, 3, 10)` - length=3, capacity=10
- slices can share same underlying array

### maps (key-value)
- syntax: `m := make(map[string]int)`
- like HashMap in java
- use "ok" idiom to check if key exists: `value, ok := map[key]`
- delete with: `delete(map, key)`
- iterate with range: `for key, value := range map { ... }`

## value vs reference types
- value types (copied): int, float, bool, string, arrays
- reference types (shared): slices, maps, pointers
- this was confusing at first

## important about slices
- when you slice a slice `slice[1:3]`, it shares same memory
- if you change original, all slices pointing to it change too
- this is different from java

## tasks completed
- arrays/arrays.go - 5 tasks ✓
- slices/slices.go - 5 tasks ✓
- maps/maps.go - 5 tasks ✓

## java vs go
| java | go |
|------|-----|
| arrays are reference | arrays are value |
| ArrayList | slices |
| HashMap | maps |
| no "ok" idiom | map access with "ok" |

## time spent: ~2 hours
