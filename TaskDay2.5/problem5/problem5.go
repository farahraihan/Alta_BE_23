package main

import (
	"fmt"
)

func RemoveDuplicates(array []int) int {
	var result int = len(array)
	var s []int

	for i := 0; i < len(array); i++ {
		duplicate := false
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				duplicate = true
			}
		}
		if !duplicate {
			s = append(s, array[i])
		}
	}

	result = len(s)

	return result
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) //4
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) //6
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))     //4
}
