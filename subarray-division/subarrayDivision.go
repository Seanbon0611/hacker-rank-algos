package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	//sum = d
	//length = m
	if m > int32(len(s)) {
		return 0
	}
	var matchingSum int32
	var subSum int32
	var i int32

	for ; i < m; i++ {
		subSum += s[i]
		if subSum == d {
			matchingSum++
		}
	}
	for j := i; j < int32(len(s)); j++ {
		subSum = subSum - s[j-i] + s[j]
		if subSum == d {
			matchingSum++
		}
	}
	fmt.Println(matchingSum)
	return matchingSum
}
func main() {
	birthday([]int32{1, 2, 1, 3, 2}, 3, 2)
}
