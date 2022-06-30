package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev_z := z
	var eps float64 = math.Nextafter(1, 2) - 1

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev_z) <= eps {
			break
		}
		fmt.Println(i, ":", z)
		prev_z = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
