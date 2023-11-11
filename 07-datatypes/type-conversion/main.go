package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myInt = 1
	var myFloat = 0.0
	var myString = "Hello World"
	var myBytes = []byte{240, 169, 154, 128}

	convertedInt := float64(myInt)
	convertedFloat := int(myFloat)
	convertedBytes := string(myBytes)
	convertedString := []byte(myString)

	fmt.Printf("After conversion, myInt is: %T, myFloat is %T, myString is %T, myByte is %T",
		convertedInt, convertedFloat, convertedString, convertedBytes)

	var mySecondInt = 128640
	stringifiedInt := fmt.Sprintf("%d", mySecondInt)
	fmt.Printf("The int is now a string: '%s", stringifiedInt)

	resStr := strconv.Itoa(myInt)
	fmt.Printf("resStr is %T of %v", resStr, resStr)
}
