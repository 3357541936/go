package main

type User struct {
	name string
}

func main() {
	//Array.Array()
	//Slice.Slice()
	u1 := &User{}
	println(u1)

	u1.name = "Tom"
	println(u1.name)
}
