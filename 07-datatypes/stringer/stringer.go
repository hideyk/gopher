package stringer

import "fmt"

type Person struct {
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("Hey there I'm %v!", p.Name)
}
