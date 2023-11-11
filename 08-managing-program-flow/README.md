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