# Go Learning Journey - 60 Days

## Day 1: ✅ COMPLETED
**Topic:** Install Go, set GOPATH, run your first "Hello World"

### What I accomplished:
- ✅ Installed Go 1.25.3 (latest version)
- ✅ Set up GOPATH: `/media/javokhir/Local Disk/go_learning`
- ✅ Created proper workspace structure (src/, bin/, pkg/)
- ✅ Created and ran first "Hello World" program
- ✅ Learned basic Go program structure

### Key Concepts Learned:
1. **Package Declaration**: `package main` - defines the package (required for every Go file)
2. **Import Statement**: `import "fmt"` - imports the format package for printing
3. **Main Function**: `func main()` - entry point of Go programs
4. **Print Statement**: `fmt.Println()` - prints text to console with newline

### Go Commands Used:
- `go version` - Check Go version
- `go run main.go` - Run Go program directly
- `go env GOPATH` - Show Go workspace path

### Environment Setup:
- **GOPATH**: `/media/javokhir/Local Disk/go_learning`
- **GOROOT**: `/usr/local/go`
- **Go Version**: 1.25.3

### File Structure:
```
go_learning/
├── src/
│   └── hello/
│       └── main.go
|       └── README.md

├── bin/
├── pkg/

```

### First Program:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

### Important Learning Notes:
- **Go Modules**: Go 1.11+ with modules enabled ignores GOPATH (can create projects anywhere)
- **GOPATH vs GOROOT**: GOPATH is for your workspace, GOROOT is where Go is installed
- **Package Types**: `package main` for executables, other packages for libraries
- **Every Go file** must start with a package declaration



## Progress Tracker:
- **Day 1**: ✅ Install Go, set GOPATH, run "Hello World"
**Happy coding! 🚀**