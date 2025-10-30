# Day 2 - Variables, Constants, Loops, If/Else

## what i learned today

### variables in go
- can declare variables 3 ways:
  - `var name string = "John"` - explicit type
  - `var name = "John"` - type inference
  - `name := "John"` - short declaration (only inside functions)
- at package level must use `var`, cant use `:=`
- zero values: int=0, float=0.0, bool=false, string=""

### constants
- use `const` keyword
- Capital letter = exported (public in java)
- lowercase = unexported (private in java)
- no keywords like public/private, just the first letter

### access control in go
- if name starts with capital letter = exported (like public in java)
- if starts with lowercase = unexported (like private in java)
- this applies to variables, functions, struct fields, everything

### named returns
- can name return values in function signature
- then just use `return` without values
- honestly looks weird, not a fan

### loops
- go only has `for` loop (no while, do-while like java)
- traditional: `for i := 0; i < 10; i++ { ... }`
- while-like: `for condition { ... }`
- range: `for i, char := range "hello" { ... }`
- break works same as java

### if/else and switch
- if/else works like java
- can do short if: `if x := getValue(); x > 0 { ... }`
- switch doesnt need break (auto breaks in go!)
- can have multiple cases: `case 1, 2, 3:`

## tasks completed
- variables.go - 5 tasks ✓
- constants.go - 5 tasks ✓
- loops.go - 5 tasks ✓
- cases.go - 5 tasks ✓

## errors fixed
- had "main redeclared" error when 2 files had main() in same dir
- fixed by putting each file in its own directory

## java vs go
| java | go |
|------|-----|
| public/private keywords | Capital/lowercase letter |
| for/while/do-while | only for |
| switch needs break | switch auto breaks |
| type name | name type |

## time spent: ~2.5 hours
