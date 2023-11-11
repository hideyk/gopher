package main

import (
	"example/hideyk/variables/user"
	"fmt"
)

const myConst = 1
const myConst2 int = 4

func SayHello(name *string) {
	fmt.Println(name)
	fmt.Printf("Hello %s\n", *name)
	*name = "Dave Blogs"
}

func main() {
	fmt.Println(myConst, myConst2)

	name := "Joe Blogs"
	SayHello(&name)
	fmt.Println(name)

	u1 := new(user.User)
	u2 := &user.User{
		Name: "Hideyuki",
	}
	fmt.Println("u1:", u1)

	user.ChangeName(u1, "Chrysan")
	user.ChangeName(u2, "Bob")

	fmt.Println("u1:", u1)
	fmt.Println("u2:", u2)
}
