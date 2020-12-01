package main

import "fmt"

/*
Implememnt a function called countUniqueValues which accepts a sorted array, and counts the unique values in the array. there can be negative numbers in the array but it will always be sorted
*/

func countUniqueValues(arr []int) int {
	count := make(map[int]int)

	for _, i := range arr {
		if count[i] > 0 {
			count[i]++
		} else {
			count[i] = 1
		}
	}

	return len(count)
}

func main() {
	fmt.Println(countUniqueValues([]int{-3, -2, -1, 1}))
}
