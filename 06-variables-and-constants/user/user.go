package user

type User struct {
	Name string
}

func ChangeName(user *User, name string) {
	user.Name = name
}
