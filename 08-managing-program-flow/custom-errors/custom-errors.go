package main

import (
	"fmt"
	"time"
)

type MyError struct {
	s         string
	signature string
}

func (e MyError) Error() string {
	return fmt.Sprintf("This is custom error generated at %s with signature %s and message: %s", time.Now(), e.signature, e.s)
}

func dummyFunction() error {
	return MyError{s: "Test error message", signature: "Hideyuki"}
}

func main() {
	if err := dummyFunction(); err != nil {
		fmt.Println(err)
	}
}
