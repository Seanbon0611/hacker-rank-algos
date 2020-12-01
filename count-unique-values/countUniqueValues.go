package main

import "fmt"

/*
Implememnt a function called countUniqueValues which accepts a sorted array, and counts the unique values in the array. there can be negative numbers in the array but it will always be sorted
*/

func countUniqueValues(arr []int) int {
	//set variable i at index of 0
	//go through a loop starting at index of 1 and on
	//if arr[i] != arr[j], incrememnt i by 1 and set arr[i] = arr[j]
	//because arrays are 0 indexed we return i + 1
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
		fmt.Println(i, arr[j])
	}
	return i + 1
}

func main() {
	fmt.Println(countUniqueValues([]int{1, 1, 1, 2, 2, 3, 4}))
}
