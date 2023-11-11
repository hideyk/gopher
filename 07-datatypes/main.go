package main

import (
	"encoding/hex"
	"example/hideyk/datatypes/stringer"
	"example/hideyk/datatypes/structs"
	"fmt"
	"unsafe"
)

func send(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func read(ch chan int) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func main() {
	myString := "9"
	fmt.Println("Unicode codepoint represented by rune:", []rune(myString))
	fmt.Println("UTF-8 code represented by up to 4 bytes:", []byte(myString))

	fmt.Println("UTF-8 code represented as Hexadecimal:", hex.EncodeToString([]byte(myString)))
	fmt.Println("length:", len(myString))

	var intArray [5]int
	fmt.Println("Array length:", len(intArray))
	fmt.Println("Array capacity:", cap(intArray))

	for i := range intArray {
		fmt.Println("Array element index", i, "contains value", intArray[i])
	}
	intArray[0] = 1
	fmt.Println(intArray)

	intArray2 := [3]int{1, 2, 3}
	intArray3 := [...]int{4, 5, 6}
	fmt.Println(intArray2)
	fmt.Println(intArray3)

	// Structs
	var user structs.User

	// Anonymous struct created at point of use
	data := struct {
		Name string
		Age  int
	}{
		Name: "Chrysan",
		Age:  20,
	}

	user = data
	fmt.Println(user)

	user.Name = "Hideyk"
	fmt.Println(user)

	// Embedding and composition
	var human structs.Human
	human.Color = "Blue"

	var dog structs.Dog
	dog.Eyes.Color = "Red"

	subOptimal := [1000]structs.Suboptimal{}
	sizeSuboptimal := float64(unsafe.Sizeof(subOptimal))

	optimal := [1000]structs.Optimal{}
	sizeOptimal := float64(unsafe.Sizeof(optimal))

	fmt.Printf("Size of suboptimal array = %v bytes\n", sizeSuboptimal)
	fmt.Printf("Size of optimal array = %v bytes\n", sizeOptimal)

	diff := (sizeSuboptimal - sizeOptimal) / sizeSuboptimal * 100
	fmt.Printf("Padding means we're using %d%% more memory...", int(diff))

	ch := make(chan int)
	fmt.Println(ch)
	go read(ch)
	send(ch)

	// Stringer
	st := stringer.Person{"Chrysan"}
	fmt.Println(st)
}
