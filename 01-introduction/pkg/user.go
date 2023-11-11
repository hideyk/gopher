// Package user implements functionality for working with users
//
// This content will be in a new paragraph
package user

type User struct {
	Name string
	Age  int32
}

// NewUser is a factory that allows us to configure the properties
// of a new User.
//
// See https://someurl.com for more information

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
