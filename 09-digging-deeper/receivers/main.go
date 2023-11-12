package main

import "fmt"

type customer struct {
	Name string
}

// UpdateName is a pointer receiver
func (c *customer) UpdateName(newStr string) {
	c.Name = newStr
}

// PrintName is a value receiver
func (c customer) PrintName() {
	fmt.Println(c.Name)
	c.PrintLine()
}

func (customer) PrintLine() {
	fmt.Println("-------------------")
}

func main() {
	var cust = &customer{Name: "hideyk"}

	cust.PrintName()

	cust.UpdateName("Chrysan")
	cust.PrintName()
}
