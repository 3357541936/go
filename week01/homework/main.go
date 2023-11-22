package main

import (
	"example.com/go/week01/homework/util"
)

func main() {
	var s []int = []int{0, 1, 2, 3, 4, 5, 6, 7}
	s = util.RemoveC(s, 0)
	s = util.RemoveC(s, 0)
	s = util.RemoveC(s, 0)
	s = util.RemoveC(s, 0)
	s = util.RemoveC(s, 7)
	s = util.RemoveC(s, 0)
	s = util.RemoveC(s, 8)

}
