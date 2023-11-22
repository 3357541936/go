package main

import (
	"example.com/go/week01/homework/util"
	"fmt"
)

func main() {
	var slice01 []int = []int{1, 2, 3, 4, 5}
	var res = util.RemoveB(slice01, 4)
	fmt.Println(res)
}
