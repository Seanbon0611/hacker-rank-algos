package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var i int32
	var pairs int32
	for ; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				pairs++
			}
		}
	}
	return pairs
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}
