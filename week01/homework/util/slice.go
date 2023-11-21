package util

func Remove[T any](slice []T, index int) []T {
	var data = make([]T, 0)
	for i := 0; i < len(slice); i++ {
		if i == index {
			continue
		}
		data = append(data, slice[i])
	}

	return data
}
