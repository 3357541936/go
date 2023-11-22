package util

import (
	"fmt"
	"slices"
)

func RemoveA[T any](slice []T, index int) []T {
	var data = make([]T, 0)
	for i := 0; i < len(slice); i++ {
		if i == index {
			continue
		}
		data = append(data, slice[i])
	}
	slice = data
	data = nil
	return slice
}

func RemoveB[Q any](s []Q, index int) []Q {
	var count int = 0
	for i := 0; i < len(s); i++ {
		if i == index {
			continue
		}
		s[count] = s[i]
		count++

	}
	s = slices.Clip(s[:count])
	fmt.Printf("%v,%d,%d \n", s, len(s), cap(s))
	return s
}

func RemoveC[Q any](s []Q, index int) []Q {
	length := len(s)
	capacity := cap(s)
	if index < 0 || index+1 > capacity {
		println("Invalid Value: index 超出正常范围!")
		return s
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if i == index {
			continue
		}
		s[count] = s[i]
		count++
	}
	if float64(length-1)/float64(capacity) <= 0.3 {
		s = s[: count : capacity/2]
	} else {
		s = s[:count:capacity]
	}
	fmt.Printf("%v  %d  %d \n", s, len(s), cap(s))
	return s
}
