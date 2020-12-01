package main

import "fmt"

func sumZero(array []int) []int {
	left := 0
	right := len(array) - 1
	for left < right {
		var sum int = array[left] + array[right]
		if sum == 0 {
			return []int{array[left], array[right]}
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return nil
}

func main() {
	array := []int{-5, -2, 1, 2, 4}
	fmt.Println(sumZero(array))
}
