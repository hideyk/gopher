package main

import (
	"fmt"
	. "log"
	mat "math"

	"example/hideyk/helloworld/examplepackage"

	_ "github.com/golangatspeed/pkg/sideeffect"
)

func init() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("Hello World again!")
	Println("This is a log")
	Println(mat.Abs(23))

	examplepackage.ExampleFunction()
}
