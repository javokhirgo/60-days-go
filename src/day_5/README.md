# Go Learning Journey - 60 Days

## Day 5: ✅ COMPLETED
**Topic:** Structs and Methods

### What I accomplished:
- ✅ Mastered struct creation and initialization
- ✅ Learned how to access and modify struct fields
- ✅ Understood value receivers vs pointer receivers
- ✅ Practiced attaching methods to structs
- ✅ Explored multiple methods on one struct
- ✅ Completed 10 practice tasks across 2 files
- ✅ Applied real-world examples (BankAccount, Student, Book)

### Key Concepts Learned:

#### 1. **Structs (Custom Data Types)**
- **Syntax**: `type Person struct { Name string; Age int }`
- **Initialization**: `person := Person{Name: "Alice", Age: 25}`
- **Purpose**: Group related data together (like classes in Java)
- **Key difference**: Structs are simpler than classes - just data, no inheritance

#### 2. **Accessing Struct Fields**
- **Dot notation**: `person.Name` or `person.Age`
- **Modification**: `person.Age = 30`
- **Direct access**: Can read and write fields directly

#### 3. **Value Receivers (Copies)**
- **Syntax**: `func (r Rectangle) Area() int { ... }`
- **Behavior**: Works on a COPY of the struct
- **Use when**: Reading data, not modifying
- **Example**: `rect.Area()` calculates area without modifying rect

#### 4. **Pointer Receivers (References)**
- **Syntax**: `func (r *Rectangle) SetWidth(width int) { ... }`
- **Behavior**: Works on the ORIGINAL struct
- **Use when**: Modifying the struct
- **Example**: `rect.SetWidth(15)` modifies the original rect

#### 5. **Methods Attached to Structs**
- **Placement**: Methods defined OUTSIDE the struct
- **Java difference**: Java methods inside class, Go methods outside struct
- **Flexibility**: Can attach multiple methods to one struct
- **Access**: Called on instances with dot notation

#### 6. **Methods with Return Values**
- **Syntax**: `func (b BankAccount) GetBalance() float64 { return b.Balance }`
- **Usage**: `balance := account.GetBalance()`
- **Purpose**: Retrieve data from struct

#### 7. **Methods that Modify Structs**
- **Pointer receivers**: Required for modifications
- **Syntax**: `func (b *BankAccount) Deposit(amount float64) { b.Balance += amount }`
- **Usage**: `account.Deposit(100)` modifies the account

### Practice Tasks Completed:

#### Basic Structs (5/5):
1. ✅ **Create a Struct** - Person with Name, Age, Email
2. ✅ **Access Struct Fields** - Rectangle with Width and Height
3. ✅ **Modify Struct Fields** - Circle with Radius modification
4. ✅ **Value Receiver Method** - Area() method that doesn't modify
5. ✅ **Pointer Receiver Method** - SetAge() method that modifies

#### Advanced Structs (5/5):
1. ✅ **Multiple Methods** - Calculator with Add and Subtract
2. ✅ **Methods with Return Values** - GetBalance() method
3. ✅ **Methods that Modify** - Deposit and Withdraw with pointer receivers
4. ✅ **Different Receiver Types** - GetAverage (value), AddGrade (pointer)
5. ✅ **Real-World Example** - Book with IsLong and UpdatePages

### Go Commands Used:
- `go run day_5/basic_structs/structs.go` - Run basic structs practice
- `go run day_5/advanced_structs/advanced_structs.go` - Run advanced structs practice

### File Structure Created:
```
day_5/
├── basic_structs/
│   └── structs.go (5 tasks completed)
├── advanced_structs/
│   └── advanced_structs.go (5 tasks completed)
└── README.md
```

### Important Learning Notes:

#### **Go vs Java Struct Differences:**
- **Java**: Methods inside class, constructors, inheritance
- **Go**: Methods outside struct, no constructors, no inheritance
- **Java objects**: Always reference types
- **Go structs**: Value types (copied) unless using pointers

#### **Value vs Pointer Receivers:**
- **Value receiver**: `func (r Rectangle) Area() int`
  - Works on copy
  - Use for reading/calculating
  - Safe, doesn't modify original
  
- **Pointer receiver**: `func (r *Rectangle) SetWidth(width int)`
  - Works on original
  - Use for modification
  - Modifies the struct

#### **Method Best Practices:**
- **Defined outside struct**: Methods attach to structs from outside
- **Called on instances**: `instance.MethodName()`
- **Choose receiver type**: Value for reading, pointer for modifying
- **Multiple methods**: One struct can have many methods

### Critical Insights Learned:

#### **1. Structs vs Classes:**
```go
// Go - Simple struct
type Person struct {
    Name string
    Age  int
}

// Methods outside struct
func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s\n", p.Name)
}
```

#### **2. Value Receiver Doesn't Modify:**
```go
func (r Rectangle) Area() int {
    return r.Width * r.Height  // Just calculates
}
// Original Rectangle unchanged
```

#### **3. Pointer Receiver Modifies:**
```go
func (r *Rectangle) SetWidth(width int) {
    r.Width = width  // Modifies original
}
// Original Rectangle changed
```

#### **4. Methods with Slices:**
```go
type Student struct {
    Name   string
    Grades []int
}

func (s *Student) AddGrade(grade int) {
    s.Grades = append(s.Grades, grade)  // Modifies slice
}
```

### Key Differences from Java:

#### **Java Classes:**
```java
public class Person {
    private String name;
    private int age;
    
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
    
    public int getAge() {
        return this.age;
    }
}
```

#### **Go Structs:**
```go
type Person struct {
    Name string
    Age  int
}

// No constructor - just initialize
person := Person{Name: "Alice", Age: 25}

// Methods outside struct
func (p Person) GetAge() int {
    return p.Age
}
```

### Time Spent: ~2 hours
- **Target**: 1.5-2 hours
- **Actual**: ~2 hours
- **Efficiency**: Good balance of learning and practice

## Progress Tracker:
- **Day 1**: ✅ Install Go, set GOPATH, run "Hello World"
- **Day 2**: ✅ Variables, Constants, Loops, If/Else
- **Day 3**: ✅ Functions, Multiple Returns, Error Handling
- **Day 4**: ✅ Arrays, Slices, Maps
- **Day 5**: ✅ Structs and Methods
- **Next**: Day 6 - Interfaces and Type Assertions

**Excellent progress! Ready for Day 6! 🚀**
