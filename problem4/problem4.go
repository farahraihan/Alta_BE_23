package main

import (
	"fmt"
)

func findMaxSumSubArray(k int, arr []int) int {
	var max int
	limit := len(arr)
	result := 0

	for i := 0; i < len(arr); i++ {
		max = 0
		if limit > k {
			for n := 0; n < k; n++ {
				max += arr[i+n]
			}
			if max > result {
				result = max
			}
		} else {
			break
		}
		limit--
	}

	return result
}

func main() {
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})) //9
	fmt.Println(findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))    //7
	fmt.Println(findMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))    //5
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))    //7
	fmt.Println(findMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))    //8
}
