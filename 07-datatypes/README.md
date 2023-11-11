# Chapter 7 - Data types

Go built-in datatypes are *primitives*

## 7.1 Basic tyoes

### 7.1.1 Number

Zero value is 0

*int*: Whole numbers (64 bytes)
*float64*: Decimals (bytes)
*uint*: Unsigned int64


### 7.1.2 Boolean
Zero value is *nil*

### 7.1.3 String
Actually a backing slice of bytes.

### 7.1.4 Byte
Zero value is 0
Alias for *uint8* which stores integers between 0 and 255, can represent any ASCII character.

### 7.1.5 Rune
To represent other numbers beyond ASCII alone we need to use a *rune*.

A *rune* is an alias afor *int32* which can represent a Unicode codepoint as a numerical value. As its an *int32* it is large enough to represent any Unicode character in a sequence of almost 150,000 at the time of writing.

A Unicode codepoint also has a UTF-8 hexadecimal representation. 

DO NOT USE `len()` by itself to determine the length of a stringm convert it to a slice of runes first. 

```go
package main

import "fmt"

func main() {
    myString := "9"
    fmt.Println("Length:", len([]rune(myString)))
}

```

## 7.2 Aggregate ttpes

### 7.2.1 Array
An array contains a fixed number of elements of a single type. Elements in an array are initialised to their *nil* value. 

An array cannot be resized once declared, length is fixed and a part of its type. A variable with type [5]int is not the same type as [6]int. An array has the same capacity as its length. 

```
var intArray [5]int
intArray2 := [5]int{1,2,3}
intArray3 := [...]int{4,5,6}

```

### 7.2.2 Struct

*Structs* are useful for representing collections of data in our programs.

*Structs* can be used anonymously - their properties defined at the point of use - or as a basis for building our own named custom types that represent collections.


#### 7.2.2.1 Composition and embedding
Structs can be embedded inside other structs. By taking smaller pieces of code we can combine their attributes and behaviour into a new piece of code and add extra attributes and behaviour we need.

In *Go*, this is called *composition*.

When we embed a struct inside another struct we can promote its fields and receivers so that they can be used as if part of that struct.


#### 7.2.2.2 Memory use, alignment & padding
Memory usage may be a consideration if struct has large no. of fields. It *may* be necessary to consider how struct fields are ordered, and possibly alternative integer types rather than the single *int*. The reason comes down to how structs and their fields are stored in memory.

In a 64-bit architecture, data is stored in 8-byte slots. An int64 variable sits perfectly inside one of these slots.

Two *int32* would occupy one slot, whereas an *int8* type occupies a single byte in memory (so does bool), we can hold 8 booleans or 8 *int8* variables in a slot.

Adjacent struct fields that cannot fit in the current 8-byte slot will spill over to the next 8-byte slot. This results in some level of *padding* in slots that do not utilise all eight bytes. 

Field ordering determines hwo much padding is required, so large structs with suboptimal field type ordering may use significantly more memory, becoming an issue.

```go
import (
    "fmt"
    "unsafe"
)

type Suboptimal struct {
    bool1 bool
    int1  int
    bool2 bool
    int2  int
    bool3 bool
    int3  int
}

type Optimal struct {
    bool1 bool
    bool2 bool
    bool3 bool
    int1  int
    int2  int
    int3  int
}

func main() {
    subOptimal := [1000]Suboptimal{}
    sizeSuboptimal := float64(unsafe.Sizeof(subOptimal))

    optimal := [1000]Optimal{}
    sizeOptimal := float64(unsafe.Sizeof(optimal))

    fmt.Printf("Size of suboptimal array = %v bytes\n", sizeSuboptimal)
    fmt.Printf("Size of optimal array = %v bytes\n", sizeOptimal)

    diff := (sizeSuboptimal - sizeOptimal) / sizeSuboptimal * 100
    fmt.Printf("Padding means we're using %d%% more memory...", int(diff))
}
```


## 7.3 Reference types
Values of these data types behave as if they *were* pointers:
- Maps
- Slices
- Channels

All are essentially descriptors which point to (or reference) some kind of backing data. 

### 7.3.2 Map
Keys can be associated with values and stored in memory. 

```go
func main() {
    // Option 1
    var myMap map[string]string     // Declaration
    myMap = make(map[string]string) // Initialisation (Required)

    // Option 2
    myMap2 := make(map[string]string)

    // Option 3
    myMap3 := map[string]string{}

    // Option 4
    myMap4 := map[string]string{
        "key": "value",
    }
}
```
Maps passed by value are still considered pass by reference. A function is able to mutate the attributes of a map as if they were pointers.


#### 7.3.2.1 Working with maps
When iterating over a map, it is impossible to predict in which order keys will be retrieved since map access is non-deterministic: keys are fetched in neither a LIFO or FIFO basis, pseudo-randomly. 

Removing keys from maps is performed with the `delete` keyword. Keys must be removed one by one, there is no way to delete an entire map.

```go
func main() {
    var myMap = make(map[string]string)

    length := len(myMap)
    
    for k, v := range myMap {
        fmt.Println(k, v)
    }
    
    delete(myMap, "key3")
}
```

#### 7.3.2.2 Safely accessing keys in a map
Care should be taken to ensure program code cannot read and write to a map in a way that leaves the variable in an uncertain state. Known as a race condition, often a problem in asynchronous code.

If we request the value for a key which is absent, function will return the nil or zero value for the map value type. Make use of the second return value (bool).
```go
func main() {
    if active, ok := activeUser["Trevor"]; !ok {
        fmt.Println("User does not exist")
    } else {
        fmt.Println("Trevor's status:", active)
    }
}
```

### 7.3.3 Slice
```go
func main() {
    // Five options to create a slice
    var sl1 []int
    sl2 := []int{}
    sl3 := []int{1,2,3,4,5}
    sl4 := make([]int, 5)
    sl5 := make([]int, 1, 4)
}

``` 

Slices are a reference type because there's a backing data array, but we need to caveat that statement, because slices behave as a reference type only in certain circumstances.

Every slice has a descriptor - or slice header - which containers three important pieces of information about the slice. First, the pointer to the backing array where the actual slice data is stored. Second, the current length, and lastly the capacity of the backing array. 

We can add new data to a slice and not concern ourselves with the finite length of the backing array. When the backing array has no capacity left, Go will automatically create a new array on our behalf and copy old contents over. 

At this point, the slice header is changed. Address of the new array is stored and we have new length and capacity. 

General rule is that if an operation on a slice does not modify the slice header, it can be used as a reference type.

For example, if we alter existing elements in the slice and don't add new elements, we don't change the slice header. In such circumstances we're able to modify the original value's elements and the caller will see the change.

To safely add to a slice in a function, recognise that the append operation is creating a new value, we should return the new value to the caller.

```go
func addToSlice(s1 []int) []int {
    s1 = append(s1, 2)
    return s1
}

func main() {
    s1 := []int{1}
    s1 = addToSlice(s1)
}
```

#### 7.3.3.1 Working with slices
Take care when working with subsets of slices. What may appear to be a copy or a new slice value, may infact be pointing at the same back array of data as the original slice. It is possible to mutate elements in the original by mutating elements in what you may believe to be a copy, when it is not.

```go
func main() {
    sl := []int{1,2,3,4,5}
    sl2 := sl[1:5]
    
    sl2[2] = 20   // This changes both sl2 and the underlying array of sl

    sl3 := copy(sl[1:5])  // This copies the underlying 
}
```

#### 7.3.3.2 Append
```go
func main() {
    sl := []int{1,2}
    sl2 := []int{7,8,9}

    sl = append(sl, 3)
    sl = append(sl, 4, 5, 6)
    sl = append(sl, sl2...)
}
```

#### 7.3.3.3 Copy
```go
func main() {
    sl := []int{1,2,3,4,5}
    sl2 := make([]int, 4)

    copy(sl2, sl[1:5])
}
```

#### 7.3.3.4 Resizing
If a slice has no room i.e the backing array has no capacity, Go will handle creating additional capacity for us. 

At lower levels of capacity, the capacity doubled each time, then started to tail of, until eventually 25% of additional capacity was allocated each time.

#### 7.3.4 Channel
Channels are a mechanism for communicating between asynchronous code. We'll cover them in detail when we talk about concurrency, here we're going to briefly cover what they are, and how we create and use them, so we can demonstrate their *reference-type* characteristics. 

Channels are created using the `make()`. They can be *unbuffered* or *buffered* and are typed - each channel can send and receive values of a specific type.


## 7.4 Interface types
An interface represents behaviour, what something can do. Behaviour is defined by one or more *receiver* function signatures. 

In Go, interfaces are implemented implicitly. You don't see the *implements* keyword anywhere. Instead, in Go we use a principle known as *duck typing*. If something can walk like a duck and quack like a duck, it is a duck.

Additionally, Go has an empty interface type. This is an interface that defines no behaviour. Often represented by `interface{}`, you can substitute in any built-in or custom type. 
```go
type Stringer interface {
    String() string
}
```

Under the hood, `fmt.Println()` will output each variable using the *stringer* interface if it is implemented on the type. It does this by calling each type's `String()` receiver which returns a string to print. 

Both the built-in slice and map types implement the stringer interface. They both have a String() receiver bound to them.


## 7.5 Creating custom types
```go
import "fmt"

type numerator int
type denominator int

func divider(d denominator, n numerator) float64 {
    return float64(n) / float64(n)
}

func main() {
    var n numerator = 1
    var d denominator = 4
    result := divider(n, d)   // Throws an error 
}
```

## 7.6 Converting between types

### 7.6.1 Type conversion
Go does not support type casting, only type conversion.
Type casting can often be used with non-compatible data types as well as compatible data types, type conversion is restricted to compatible data types. You can't convert a string to an int using type conversion.

