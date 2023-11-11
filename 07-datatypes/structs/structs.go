package structs

// Named custom struct
type User struct {
	Name string
	Age  int
}

// Second named custom struct
type Admin struct {
	Name string
	Age  int
}

type Eyes struct {
	Color string
	Shape string
}

type Human struct {
	Eyes
}

type Dog struct {
	Eyes Eyes
}

type Suboptimal struct {
	bool1 bool
	int1  int
	bool2 bool
	int2  int
	bool3 bool
	int3  int
}

type Optimal struct {
	bool1 bool
	bool2 bool
	bool3 bool
	int1  int
	int2  int
	int3  int
}
