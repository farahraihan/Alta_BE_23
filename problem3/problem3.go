package main

import (
	"fmt"
	"sort"
)

func dragonOfLoowater(dragonHead, knightHeight []int) {
	n := len(dragonHead)
	m := len(knightHeight)
	var minSum int

	if n > m {
		fmt.Println("knight fall!")
	} else {

		sort.Ints(dragonHead)
		sort.Ints(knightHeight)

		j := 0
		for i := 0; i < n; i++ {
			for j < m && knightHeight[j] < dragonHead[i] {
				j++
			}
			if j == m {
				fmt.Println("knight fall!")
				return
			}
			minSum += knightHeight[j]
			j++
		}

		fmt.Println(minSum)

	}
}

func main() {
	dragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    //11
	dragonOfLoowater([]int{5, 10}, []int{5})         //knight fall
	dragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) //knight fall
	dragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) //10
}
