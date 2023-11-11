package main

import "fmt"

func isEvenLesstThan(n int, c int) (bool, bool) {
	var even, less bool

	if n%2 == 0 {
		even = true
	}

	if n < c {
		less = true
	}

	return even, less
}

func main() {
	fmt.Println(isEvenLesstThan(2, 10))

	// Short form `if` statement
	if even, _ := isEvenLesstThan(2, 10); even {
		fmt.Println("Result was even")
	}

	// Switch statements
	switch myName := "Hideyuki"; myName {

	case "Hideyuki":
		fmt.Println("Hideyuki")
	case "Chrysan":
		fmt.Println("Chrysan")
	case "Dada":
		fmt.Println("Dada")
	}

	// Expressionless switch statement
	num := 12
	switch {
	case num >= 0 && num <= 10:
		fmt.Println("Between 0 and 10")
	case num >= 10:
		fmt.Println("Larger than 10")
	}

	// Multiple match tests per case statement
	switch letter := "z"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Letter was a vowel")
	default:
		fmt.Println("Letter was not a vowel")
	}

	// Fallthrough example
	day := "Mon"
	switch {
	case day == "Mon":
		fmt.Println("Monday")
		fallthrough
	case day == "Tue":
		fmt.Println("Tuesday")
	case day == "Wed":
		fmt.Println("Wednesday")
	}
}
