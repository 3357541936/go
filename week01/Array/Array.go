package Array

func Array() {
	arr01 := [5]int{9, 8, 7, 6, 5}
	for index, value := range arr01 {
		println(index, value)
	}
}
