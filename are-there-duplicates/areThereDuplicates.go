package main

import (
	"fmt"
	"sort"
)

//pointers solution
func areThereDuplicates(a ...int) bool {
	sort.Ints(a)
	start := 0
	next := 1
	for next < len(a) {
		if a[start] == a[next] {
			return true
		} else {
			start++
			next++
		}
	}
	return false
}

func main() {
	fmt.Println(areThereDuplicates(1, 2, 3))
}
