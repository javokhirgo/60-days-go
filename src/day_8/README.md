# Day 8: Packages, Imports, and Go Modules

## ✅ Completed Topics

### 1. **Package vs Import vs Module**
- **Package**: Namespace for organizing code (like Java packages)
- **Import**: Bringing packages into your code
- **Module**: Project root with `go.mod` file

### 2. **Standard Library Packages**
- `strings`: ToUpper, Contains, Split, Join
- `math`: Sqrt, Pow, Max
- `time`: Now, Hour, Minute, Second, Format

### 3. **Custom Packages**
- Created `calculator` package
- Exported functions: `Add`, `Subtract`, `Multiply`, `Divide` (capital letter)
- Unexported function: `validate` (lowercase)
- Exported constant: `MaxOperations`

### 4. **External Packages**
- Used `github.com/fatih/color` for colored terminal output
- Installed with `go get`
- Added to `go.mod` automatically

### 5. **Go Modules (go.mod & go.sum)**
- **go.mod**: Defines module name and dependencies
- **go.sum**: Stores checksums for external packages (security)
- **Direct dependency**: You import it in your code
- **Indirect dependency**: Brought in by direct deps (marked `// indirect`)

### 6. **Key Commands**
- `go get package`: Download external package and add to go.mod
- `go mod tidy`: Clean up dependencies based on actual imports
- `go mod init module_name`: Initialize new module
- `go run`: Automatically downloads dependencies and runs code

---

## 🎯 Tasks Completed

### Standard Library Practice
- ✅ Used `strings` package for text manipulation
- ✅ Used `math` package for calculations
- ✅ Used `time` package for date/time operations

### Custom Package
- ✅ Created `calculator` package with exported functions
- ✅ Implemented unexported helper function
- ✅ Defined exported constant
- ✅ Imported and used in `main.go`

### External Package
- ✅ Installed `github.com/fatih/color` with `go get`
- ✅ Used colored output functions
- ✅ Understood `go.mod` and `go.sum` changes

---

## 📊 Score: 8.5/10

**Strengths:**
- ✅ Perfect custom package implementation
- ✅ Correct exported/unexported usage
- ✅ Good understanding of modules

**Minor Issues:**
- Used `println()` instead of `fmt.Println()` in some places
- Small typo in print labels

---

## 🔑 Key Insights

### Java vs Go Comparison

| Java | Go |
|------|-----|
| `package com.example.calculator` | `package calculator` |
| `import java.util.ArrayList` | `import "fmt"` |
| Maven/Gradle (pom.xml/build.gradle) | go.mod |
| `public` keyword | Capital letter = exported |
| `private` keyword | lowercase = unexported |

### Best Practices
- Use `fmt.Println()` not `println()`
- Import then run `go mod tidy` (most common)
- Commit both `go.mod` and `go.sum`
- Use `go get` only for version management

---

## 📝 What You Learned Today
1. Go packages are simpler than Java (no `com.example.app`)
2. Capital letter = public, lowercase = private (no keywords!)
3. `go.mod` manages dependencies like Maven/Gradle
4. Internal packages don't appear in `go.mod`
5. External packages create `go.sum` for security

---

**Week 2, Day 1 Complete!** 🎉

