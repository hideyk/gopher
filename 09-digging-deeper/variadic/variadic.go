package main

import "fmt"

func myVariadicFunc(name string, address ...string) {
	fmt.Printf("Hello %s\n", name)
	fmt.Println("Addresses:")
	if len(address) > 0 {
		for i, addr := range address {
			fmt.Printf("%d: %s\n", i+1, addr)
		}
	} else {
		fmt.Println("No address supplied")
	}
}

func main() {
	// single argument
	fmt.Println("Single argument")
	myVariadicFunc("Joe Bloggs", "Address 1")

	// multiple argument
	fmt.Println("\nMultiple arguments")
	myVariadicFunc("Hideyuki", "Address 1", "Address 2")

	// no arguments
	fmt.Println("\nNo argument")
	myVariadicFunc("Hideyuki")

	// passing a pre-built slice
	fmt.Println("\nPassing a slice as a variadic argument")
	addresses := []string{"Address 1", "Address 2"}
	myVariadicFunc("Hideyuki", addresses...)
}
