package main

import "fmt"

func averagePair(arr []int, t float64) bool {
	start := 0          //point at 0 index
	end := len(arr) - 1 //point to last element of the array
	//while start is less than end, if the average of the two numbers is == the target average return true, if not continue from each end of the slice until the target is reached.
	for start < end {
		if (arr[start] + arr[end]/2) == int(t) {
			return true
		} else {
			start++
			end--
		}
	}
	//If there are no averages that match the target then return false
	return false
}

func main() {
	fmt.Println(averagePair([]int{1, 3, 3, 5, 6, 7, 10, 12, 19}, 8))
}
