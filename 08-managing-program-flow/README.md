# Chapter 8 - Managing program flow

## 8.1 Control Structures
We may employ three forms of control structure login in our programs to facilitate this:
- sequence: Linear statements are executed one after the other
- selection: Conditional flow - what to do when one or more of several possible outcomes is satisfied
- iteration: A section of code should be repeatedly run, and when that repetition should finish


### 8.1.1 Sequence logic
Runs from top to bottom
Exceptions include *GOTO* and *defer*
Defer statements are put close together for association and readability
Defer statements are run on a LIFO basis.

### 8.1.2 Selection logic
- if/else/elseif
- switch/case/default

### 8.1.3 Iteration logic
- Infinite
- Three component
- While equivalent
- Do, while equivalent
- For each
- Break and continue


## 8.2 Error handling

Go's error-handling capabilities have earned it a reputation as one of the most reliable languages for production-level applications.

In Go, there are no exceptions and no try/catch type operations.

We're able to create error values, and, decide how we handle the error values we receive.

```go
type error interface {
    Error() string
}
```

### 8.2.1 Error helpers
Most commonly used error helper package is `errors`. We can create an error using the `errors.New(string)` factory type function. We can see what is happening under the hood by inspecting the package:

```go
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

func New(text string) error {
    return &errorString(text)
}
```

#### 8.2.1.1 Predefined errors

We can create predefined errors (sentinel errors) in our code using either of the above approaches. 

See example in `./errors/errors.go`


### 8.2.2 Custom error types
Custom error types allow us to convey more specific information about the error

To create a custom error type, we need to create a new type that implements the error interfact, This can be achieved by simply defining a new type that has an `Error()` string receiver.

The `Error()` string receiver is called when the error is converted to a string, which is what the `fmt` package does when it prints an error.

### 8.2.3 Error wrapping
Errors may be wrapped by creating a new error value that includes the original error as part of its message or data.

The `fmt.Errorf()` function formats a new error message that includes the original error value. The `%w` format specifier is used to include the original error in the message.

Wrapping is useful when trying to understand where an error originated in our code

```go
wrappedErr := fmt.Errorf("Error occurred: %w, err")
```

### 8.2.4 Panic and recover
The `panic` and `recover` functions are used in exceptional situations, such as runtime errors which can't be handled following an unexpected event, makes it unsafe to continue execution of the program in an unknown state.

A panic is caused either by a runtime error or explicit call to the built-in `panic()` function, which is called with a simple argument. `recover()` is used to recover from a `panic` and continue with program execution.

When `panic()` is called, normal execution of the code is halted and begins unwinding the call stack. It will call any `deferred` functions and run any clean-up code. Once call stack has been completely unwound, runtime will look for a `recover()` function. If found, it will be called with the value that was passed to the `panic()` function.

`recover()` can only be used within a deferred function

### 8.3.1 Log package

```go
package main

import "log"

func main() {
    log.Println("Hello world!")
}
```