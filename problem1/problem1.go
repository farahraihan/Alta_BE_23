package main

import (
	"fmt"
	"math"
)

func simpleEquations(a, b, c int) {
	// Batasi pencarian nilai x, y, dan z berdasarkan nilai c
	maxVal := int(math.Sqrt(float64(c))) + 1

	for x := -maxVal; x <= maxVal; x++ {
		for y := -maxVal; y <= maxVal; y++ {
			for z := -maxVal; z <= maxVal; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)
					return
				}
			}
		}
	}

	fmt.Println("No solution found")
}

func main() {
	simpleEquations(1, 2, 3)
	simpleEquations(6, 6, 14)
}
