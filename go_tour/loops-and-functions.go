package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	delta := 1.0
	min_delta := 0.0000000001
	z := x / 3
	i := 1
	for ; math.Abs(delta) > min_delta; i++ {
		delta = (z*z - x) / (2 * z)
		z = z - delta
	}
	fmt.Println("iterations: ", i)
	return z
}

// func main() {
// 	fmt.Println(Sqrt(3))
// }
