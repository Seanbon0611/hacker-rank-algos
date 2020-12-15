package main

import (
	"fmt"
)

func maxSubarraySum(array []int, n int) int {
	if n > len(array) {
		return 0
	}
	maxSum := 0
	tempSum := 0
	//Loop through 1 time to get the sum of the first set of numbers
	for i := 0; i < n; i++ {
		maxSum += array[i]
	}
	//now that we have a max sum, assign temp sum to max sum
	tempSum = maxSum
	//loop through the length of the whole array starting from the index of the next number after the first set
	for i := n; i < len(array); i++ {

		//assigning new tempsum, we get the current tempSum, move down the number of indexes == to n, subtract that number, then add the number of the current loops index
		fmt.Println(tempSum, array[i-n], array[i])
		tempSum = tempSum - array[i-n] + array[i]
		//see which is larger the current maxSum or the tempSum we're holding, if tempSum is larger maxSum is now the value of TempSum
		maxSum = Max(maxSum, tempSum)
	}
	return maxSum
}

//helper function to get the higher of two integers
func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	// fmt.Println(maxSubarraySum([]int{-1, -1, -2, 4, 2, 3, 5, 1}, 4))
	// fmt.Println(maxSubarraySum([]int{}, 4))
	fmt.Println(maxSubarraySum([]int{100, 200, 300, 400}, 2))
}
