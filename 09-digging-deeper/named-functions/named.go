package main

import "fmt"

func anonymousReturns(firstName, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func namedReturns(firstName, lastName string) (fullname string) {
	fullname = fmt.Sprintf("%s %s", firstName, lastName)
	return
}

func main() {
	fmt.Println(anonymousReturns("Hidey", "Kanazawa"))
	fmt.Println(namedReturns("Hidey", "Kanazawa"))
}
