package main

import (
	"fmt"
	"strings"
)

func Formatter(str string, format func(in string) string) string {
	return format(str)
}

func main() {
	upper := func(in string) string {
		return strings.ToUpper(in)
	}

	lower := func(in string) string {
		return strings.ToLower(in)
	}

	myText := "Some rANDOm text TO foRMat"
	fmt.Println(Formatter(myText, upper))
	fmt.Println(Formatter(myText, lower))
	fmt.Println(Formatter(myText, strings.ToLower))
	fmt.Println(Formatter(myText, strings.ToUpper))
}
