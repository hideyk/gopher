# Chapter 6 - Variables and Constants

## 6.1 Variables

Variable declaration

```go
// Style 1. Declaration only in the global or function scope
var myVar int
var myVar, myVar2 int = 9, 10

// Style 2. Short form declaration and initialisation within the function scope only
myVar := 10
myVar, myVar2 := 9, 10

// Style 3. Multiple variables, omitting repeating var keyword, global or function scope
var (
    myVar int
    myVar2 int
)
```

## 6.2 Constants
Constants are immutable values declared with the `const` keyword and may be created in the same way as variables, with one exception, the short form `:=` is not valid.

```go
const myConst = 1
const myConst int
```

## 6.3 Scope
Only use constants in the global scope.

A variable of the same name can be declared in multiple scopes, as each scope has a separate namespace. Avoid this if possible, it can lead to *variable shadowing*.


## 6.4 Variable semantics. Pointers and values
In Go, we can share variables as *values* and *pointers*. A value is a copy of the contents of the variable, whereas a pointer is a copy of the address of the variable in memory. 

Pass by value: Sharing a copy of the variable
Pass by reference: Sharing a copy of the memory address of the variable (hexadecimal form: 0xc0000ac018)
- Large structs don't get duplicated in pass by reference 

## 6.5 Value initialisation

```go
package main

import "fmt"

type User struct {
    Name string
}

func main() {
    u1 := new(User)
    u2 := &User(
        Name: "Joe Blogs",
    )
}

```