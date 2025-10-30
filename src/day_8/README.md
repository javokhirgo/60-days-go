# Day 8 - Packages, Imports, Go Modules

## what i learned

### package vs import vs module
- **package**: like namespace in java, groups related code
- **import**: bringing packages into your code
- **module**: whole project with go.mod file

### standard library packages
used these today:
- `strings` - ToUpper, Contains, Split, Join
- `math` - Sqrt, Pow, Max
- `time` - Now, Hour, Minute, Second, Format

### custom packages
- created calculator package
- Capital letter = exported (Add, Subtract, Multiply, Divide)
- lowercase = unexported (validate)
- constants also follow this rule (MaxOperations)

### external packages
- used github.com/fatih/color for colored output
- installed with `go get`
- automatically added to go.mod

### go.mod and go.sum
- **go.mod**: like pom.xml in maven or build.gradle in gradle
- **go.sum**: stores checksums for security (only external packages)
- **direct dependency**: i import it in my code
- **indirect dependency**: brought by my dependencies (marked `// indirect`)

### commands learned
- `go get package` - download external package
- `go mod tidy` - clean up dependencies
- `go mod init module_name` - start new module
- `go run` - downloads dependencies automatically

### best practices
- import first, then run `go mod tidy` (most common way)
- use `go get` only for version management
- commit both go.mod and go.sum
- use `fmt.Println()` not `println()`

## tasks completed
- standart_library/standard_library.go - 5 tasks ✓
- calculator package - custom package ✓
- main.go - using custom package ✓
- external_packages/external.go - using external package ✓

## mistakes i made
- used `println()` instead of `fmt.Println()` in some places
- typo in print labels (said "hour" 3 times)
- created validate() function but didnt use it

## java vs go
| java | go |
|------|-----|
| package com.example.app | package app |
| import java.util.* | import "fmt" |
| Maven/Gradle | go.mod |
| public keyword | Capital letter |
| private keyword | lowercase |

## confusion i had
- thought go.sum includes my code (it doesnt, only external)
- wasnt sure when to use go get vs go mod tidy
- confused about direct vs indirect dependencies

## score from mentor: 8.5/10
lost points for println() and typos, but got the concepts right

## time spent: ~2 hours
