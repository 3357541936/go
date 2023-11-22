package util

import (
	"fmt"
	"slices"
)

func Remove[T any](slice []T, index int) []T {
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
