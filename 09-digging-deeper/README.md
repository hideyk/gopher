# Chapter 9 - Digging deeper

## 9.1 Developing with functions
We'll now focus on function signatures including variadic arguments, and return styles, including multiple returns

### 9.1.1 Function parameters
In Go, when defining function parameters we need to specify both the variable name and its type. To shorten the signature somewhat we can group parameters of the same type if it makes sense to do so.

```go
//ungrouped
func normalParams(name string, age int, houseNumber int, address1 string, address2 string){

}

//grouped
func groupedParams(name, address1, address2 string, houseNumber, age int) {

}
```

### 9.1.2 Variadic arguments
There can be only one variadic parameter in each function signature and it must be the last parameter. It is also optional, so we can omit the argument if we wish.

A parameter which accepts variadic arguments is denoted by the `spread` operatior, which precedes in type e.g. string.

```go
func myVariadicFunc(name string, address ...string) {
    ...
}
```

### 9.1.3 Multiple return values
While the flexibility to return more than two values is often useful, there's a balance to be found. If the no. of values we want from a function goes beyond three, it could mean our API would benefit from having those values as fields on a struct, or keys in a map. This way we could use a single value to return virtually unlimited data without complicating our codebase. Its also much easier to maintain.

Example:
```go
func Split(s, sep string) (string, string, bool){}
```

### 9.1.4 Function return styles
Functions can use either of two return styles, Anonymous or Named return parameters. We can't mix styles within the same function.

When using named return parameters, there is no need to return the parameters explicitly. A simple naked `return` statement is sufficient. Note also we don't need to declare the variables explicitly in the function body. 

Because named return parameters are declared in the signature, and because they are initialised to their zero value, they satisfy the function return automatically even if they are never assigned.


### 9.1.5 Functions are a type
This means functions can be assigned to variables, apssed as arguments to other functions, and returned as values from functions.

We call functions which accept and return other functions, *higher-order functions* and they offer a high degree of flexibility in how we arrange our code. Functions which accept other functions are also known as *callbacks*.

We should also mention *closures*. A *closure* is a function that remembers the values of the variables from the place where it was created, as if bound to them. Nothing here is unique to Go, but closures can be useful when we want to use a function in a different place, but we still want it to have access to the variables it needs. 

Observe how the `increment` function has access to the x variable even though it is defined outside the function. Possible because `increment` is a closure.

### 9.1.6 Pointer or value returns

When writing functions, we need to decide whether to return values or pointers. The decision can have an impact on both memory use and how memory is managed. This shouldn't be an initial consideration as always we should prioritise correctness and maintainability until we can verify we have performance issues.

The choice between returning a value or a pointer from a function should depend on intended use case. If we want to create a single value of something in a function and share it throughout out program, we should return a pointer to that value. On the other hand, if we only need to use the value within the function and do not need to mutate anything we share up the call stack, it may be more efficient to return a copy of the value instead.

A good example of when it is beneficial to return a pointer is when creating a database connection. By returning a pointer, multiple parts of the program can use the same connection simultaneously, rather than each part creating its own connection. This can help to improve the efficiency of the program by reducing the no of connections that need to be created and managed.


## 9.2 Memory management
Go performs memory management on our behalf. The compiler chooses where to put the values our program creates - on either the stack or the heap - when compiling our program, and the garbage collector manages heap memory during program execution.

`Stack` - Region of memory used for storing local variables and function parameters. Total size of the stack is limited in size but stack memory offers fast access. Each function call or goroutine starts with an initial stack of 2KB dedicated to it. This may grow as required.

`Heap` - Region of memory used for dynamically-allocated objects at runtime, or variables the compilers decides to put there instead of the stack. Heap memory is shared and not bound to any one function. There's more heap memory available than stack memory, but slower to access. Process of garbage collection on the heap adds further overhead. 

`Allocations` - Variables placed on the heap. The fewer allocations our programs make, the better. An allocation represents data in slower access memory, which has to be cleaned up by garbage collector. 

`Escape analysis` - Process used by the compiler to determine what can be `inlined` and whether a value can sit on the stack or must be allocated to the heap. 

### 9.2.3 Garbage collcetion
Heap memory management is the responsibility of the garbage collector which will deallocate memory when it can no longer find any references to it. 

At runtime, the garbage collection process has a cost. Known as GC pause, historically it was a `stop-the-world` type operation in which all goroutines were paused while heap memory was deallocated. In most applications this pause was imperceptible, but in high-throughput, high-performance applications it could cause performance issues like `bottlenecks` and `throttling`.

In modern versions of Go, the garbage collection process is more sophisticated and doesn't need to suspend all goroutines at the same time and completes faster. 

### 9.2.4 Observing compiler escape analysis
Use escape analysis to observe how many allocations result from complex pieces of code. The snippet below shows how we use the `-gcflags` "-m" flag to view escape analysis output when running or building a program.
```bash
go build -gcflags="-m"
go run -gcflags="-m" ./main.go
```

Alongside escape analysis we can also use benchmarks, to determine when allocations are made. We'll create a benchmark file for each example and run it with this command:
```bash
go test -bench . -benchmen
```
The command instructs the `go test` tool to run all benchmarks, and include benchmark metrics for memory in the output. 

```go
package main

type Customer struct {
	Name  string
	Email string
	Age   int
}

//go:noinline
func NewCustomer(name string) Customer {
	cust := Customer{Name: name}
	return cust
}

//go:noinline
func NewCustomer2(name string) *Customer {
	cust2 := Customer{Name: name}
	return &cust2
}

//go:noinline
func NewCustomer3(name string) *Customer {
	return &Customer{Name: name}
}
func main() {
	input := "Joe Blogs"
	_ = NewCustomer(input)
	_ = NewCustomer2(input)
	_ = NewCustomer3(input)
}

```

The first function returns a value, which is a copy of the value created inside the function. Compiler knows everything can sit on the stack - when function is finished that copy does not reference anything inside the function. No allocations needed.

The second function creates a `Customer{}` value `cust2`, but then its address is taken and a pointer is returned. The compiler can't know how this value will be accessed and used later and it lives beyond the function's lifetime, so this value can't sit on the stack, it must be moved to the heap.

The final function is similar to the last, except an address is immediately taken and returned. Escape analysis indicates this escapes to the heap indicating there may be an allocation, and the benchmark results confirm that an allocation was made.


## 9.3 Using receivers with custom types
Receivers are essentially a different kind of function. They are similar to methods in OOP.

Unlike methods which exist on a class, receivers can be bound to any user-defined type, which allows the receiver to access the value of that type. If a receiver is bound to a custom struct, it can access the fields of the struct just like a method can access the properties of a class.

Receivers are useful when writing functionality intended to work with a value of a specific type, like getter and setter routines which read and update the value.

Receivers can either `receive` a copy of that value (`value receivers`), which as you expect can only be read, not mutated, or they can receive a pointer to the value (`pointer receivers`), which allows the receiver to mutate the value from within the receiver body. 


## 9.4 Working with interfaces

### 9.4.1 Recapping
In Go, an interface is a set of receiver signatures. A value of an interface type can hold any value that implements those receivers. Interfaces are a way to specify the required behaviour of a type: if a type has the receivers specified by an interface, then it is said to implement that interface. It does not need to be explicitly stated in the code that it implements the interface.

The empty interface, `interface{}` aliased `any`, specifies no behaviour. Consequently, all built-in and user-defined types implement this interface.

Interfaces are useful when we want to define generic behaviour that several types may need to implement. They allow us to write code using the interface. They allow us to write code using the interface - an abstraction of the implementaion - instead of the implementation itself.

Writing code that accepts interfaces in preference to concrete types is a common pattern in Go. For example, many functions accept the `io.writer` and `io.reader` interfaces, enabling them to call `Write` and `Read` implementations on the interface value passed.

### 9.4.2 Creating interfaces of our own
