# Go Learning Journey - 60 Days

## Day 4: ✅ COMPLETED
**Topic:** Arrays, Slices, Maps

### What I accomplished:
- ✅ Mastered Go arrays (fixed-size collections)
- ✅ Learned slices (dynamic collections with shared underlying data)
- ✅ Understood maps (key-value pairs with "ok" idiom)
- ✅ Practiced array operations (creation, iteration, modification, comparison)
- ✅ Explored slice operations (append, slicing, capacity, shared data)
- ✅ Mastered map operations (CRUD, iteration, key checking)
- ✅ Completed 15 practice tasks across 3 files
- ✅ Fixed main function redeclaration issues with proper directory structure

### Key Concepts Learned:

#### 1. **Arrays (Fixed-Size Collections)**
- **Syntax**: `var arr [5]int` or `arr := [5]int{1, 2, 3, 4, 5}`
- **Characteristics**: Fixed size, value types (copied when passed to functions)
- **Operations**: Creation, initialization, iteration, modification, comparison
- **Comparison**: Arrays can be compared with `==` operator
- **Java difference**: Go arrays are value types, Java arrays are reference types

#### 2. **Slices (Dynamic Collections)**
- **Syntax**: `var slice []int` or `slice := []int{1, 2, 3}`
- **Characteristics**: Dynamic size, reference types (shared when passed to functions)
- **Operations**: `append()`, slicing `[start:end]`, `make()` with capacity
- **Key features**: Can grow beyond initial capacity, share underlying data
- **Capacity vs Length**: Length = used elements, Capacity = total available space

#### 3. **Maps (Key-Value Pairs)**
- **Syntax**: `var m map[string]int` or `m := make(map[string]int)`
- **Characteristics**: Reference types, unordered, fast lookups
- **Operations**: CRUD (Create, Read, Update, Delete), iteration with `range`
- **"ok" idiom**: `value, ok := map[key]` to check if key exists
- **Types**: `map[string]int`, `map[int]bool`, `map[string]string`

#### 4. **Value vs Reference Types**
- **Value types** (copied): `int`, `float64`, `bool`, `string`, arrays, structs
- **Reference types** (shared): slices, maps, pointers
- **Arrays**: Passed by value, changes don't affect original
- **Slices/Maps**: Passed by reference, changes affect original

#### 5. **Slice Slicing and Shared Data**
- **Slicing**: `slice[start:end]` creates new slice pointing to same data
- **Shared memory**: Multiple slices can point to same underlying array
- **Modifications**: Changes in one slice affect all slices sharing the data
- **Memory efficient**: No copying, just pointer manipulation

#### 6. **Map Operations and Patterns**
- **Creation**: `make(map[keyType]valueType)` or map literals
- **Access**: `value := map[key]` or `value, ok := map[key]`
- **Update**: `map[key] = newValue`
- **Delete**: `delete(map, key)`
- **Iteration**: `for key, value := range map`

### Practice Tasks Completed:

#### Arrays (5/5):
1. ✅ **Create and Initialize** - Array with values 1,2,3,4,5
2. ✅ **Array Operations** - String array with range loop
3. ✅ **Length and Indexing** - Float array with index-based loop
4. ✅ **Array Modification** - Boolean array modification
5. ✅ **Array Comparison** - Compare two arrays with `==`

#### Slices (5/5):
1. ✅ **Create and Initialize** - Slice with values 1,2,3,4,5
2. ✅ **Slice Operations** - Append "Rust" to string slice
3. ✅ **Slice Slicing** - Create sub-slices (first 5, last 5)
4. ✅ **Slice Capacity** - Use make() with length=3, capacity=10
5. ✅ **Slice Modification** - Show slice modification

#### Maps (5/5):
1. ✅ **Create and Initialize** - String to int map with make()
2. ✅ **Map Operations** - Add, update, delete entries
3. ✅ **Map Access** - Check if keys exist with "ok" idiom
4. ✅ **Map Iteration** - Loop over keys and values with range
5. ✅ **Map Types** - Different map types (string→string, int→bool)

### Go Commands Used:
- `go run day_4/arrays/arrays.go` - Run arrays practice
- `go run day_4/slices/slices.go` - Run slices practice
- `go run day_4/maps/maps.go` - Run maps practice

### File Structure Created:
```
day_4/
├── arrays/
│   └── arrays.go (5 tasks completed)
├── slices/
│   └── slices.go (5 tasks completed)
├── maps/
│   └── maps.go (5 tasks completed)
└── README.md
```

### Important Learning Notes:

#### **Go vs Java Collection Differences:**
- **Arrays**: Go arrays are value types, Java arrays are reference types
- **Slices**: Go's dynamic arrays, Java uses ArrayList
- **Maps**: Similar to Java HashMap, but with "ok" idiom for key checking
- **No generics**: Go doesn't have generics (yet), Java does

#### **Slice Power Features:**
- **Automatic growth**: Slices grow beyond capacity automatically
- **Shared data**: Multiple slices can share underlying array
- **Memory efficient**: No copying, just pointer manipulation
- **Flexible slicing**: `[start:end]` creates views of data

#### **Map Best Practices:**
- **Always check keys**: Use "ok" idiom to avoid zero values
- **Initialize with make()**: Prevents nil map panics
- **Use meaningful keys**: String keys are most common
- **Handle missing keys**: Always check if key exists before using

### Critical Insights Learned:

#### **1. Value vs Reference Types:**
```go
// Arrays (value type - copied)
func changeArray(arr [3]int) {
    arr[0] = 99  // Changes copy, not original
}

// Slices (reference type - shared)
func changeSlice(slc []int) {
    slc[0] = 99  // Changes original!
}
```

#### **2. Slice Capacity Magic:**
```go
slice := make([]int, 3, 10)  // length=3, capacity=10
slice = append(slice, 1, 2, 3, 4, 5, 6, 7)  // length=10, capacity=10
// Go automatically manages capacity!
```

#### **3. Map "ok" Idiom:**
```go
// Wrong - might get zero value
value := map[key]  // Could be 0, "", false

// Right - check if key exists
if value, ok := map[key]; ok {
    // Key exists, use value
} else {
    // Key doesn't exist
}
```

#### **4. Slice Shared Data:**
```go
original := []int{1, 2, 3, 4, 5}
view1 := original[1:4]  // [2, 3, 4]
view2 := original[2:5]  // [3, 4, 5]

original[2] = 999
// view1 = [2, 999, 4] - sees the change!
// view2 = [999, 4, 5] - sees the change!
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
- **Next**: Day 5 - Structs and Methods

**Excellent progress! Ready for Day 5! 🚀**
