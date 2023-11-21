package main

import (
	"example.com/go/week01/homework/util"
	"fmt"
)

func main() {
	var slice01 []int = []int{1, 2, 3, 4, 5}

	println("------------")
	println(len(slice01))
	println(cap(slice01))

	var res = util.Remove(slice01, 3)
	fmt.Println(res)

	println("------------")
	println(len(slice01))
	println(cap(slice01))

}
