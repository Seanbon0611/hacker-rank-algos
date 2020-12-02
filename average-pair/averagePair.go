package main

import "fmt"

func averagePair(arr []int, t float64) bool {
	start := 0
	end := len(arr) - 1
	for start < end {
		if (arr[start] + arr[end]/2) == int(t) {
			return true
		} else {
			start++
			end--
		}
	}
	return false
}

func main() {
	fmt.Println(averagePair([]int{1, 3, 3, 5, 6, 7, 10, 12, 19}, 8))
}
