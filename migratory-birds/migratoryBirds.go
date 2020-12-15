package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	dict := make(map[int32]int32)
	// var maxBirdCount int32
	var mostFrequentBird int32
	for _, v := range arr {
		fmt.Println(v)
		if _, ok := dict[arr[v]]; ok {
			dict[arr[v]]++
		} else {
			dict[arr[v]] = 1
		}
		fmt.Println(dict)
	}
	// for k, v := range dict {
	// 	if v > maxBirdCount {
	// 		mostFrequentBird = k
	// 	}
	// }
	return mostFrequentBird
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 5, 3}))
}
