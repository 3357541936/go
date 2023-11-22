package Slice

import "fmt"

func Slice() {
	s01 := []int{1, 2, 3, 4}
	fmt.Print("%v, %d, %d", s01, len(s01), cap(s01))

	s02 := make([]int, 3, 4)
	fmt.Print("%v, %d, %d", s02, len(s02), cap(s02))

	s02 = append(s02, 5)
	fmt.Print("%v, %d, %d", s02, len(s02), cap(s02))

	s02 = append(s02, 8)
	fmt.Print("%v, %d, %d", s02, len(s02), cap(s02))

	s03 := make([]int, 4)
	fmt.Print("%v, %d, %d", s03, len(s03), cap(s03))

	for index, value := range s01 {
		println(index, value)
	}
}
