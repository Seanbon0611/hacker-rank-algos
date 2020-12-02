package main

import (
	"fmt"
	"sort"
)

//pointers solution
func areThereDuplicates(a ...int) bool {
	sort.Ints(a) //sort the slice
	start := 0   // start at 0 index
	next := 1    //the next element after the start index
	for next < len(a) {
		//because the array is sorted, if the index of start and index of next are the same, then we've found a duplicate
		if a[start] == a[next] {
			return true
		} else {
			//if the two elements are not the same then move onto the next set of elements
			start++
			next++
		}
	}
	//if there are no duplicates then we return false
	return false
}

func main() {
	fmt.Println(areThereDuplicates(1, 2, 3))
}
