package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	//sum = d
	//length = m
	var matchingSums int32
	var subSum int32
	var i int32 = 0
	for ; i <= m; i++ {
		subSum += s[i]
		if subSum == d {
			matchingSums++
		}
	}
	var j int32 = i
	for ; j < int32(len(s)); j++ {

	}
	fmt.Println(matchingSums)
	return matchingSums
}
func main() {
	birthday([]int32{1, 2, 1, 3, 2}, 3, 2)
}
