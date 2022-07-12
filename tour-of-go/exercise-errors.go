package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	prev_z := z
	var eps float64 = math.Nextafter(1, 2) - 1

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev_z) <= eps {
			break
		}
		fmt.Println(i, ":", z)
		prev_z = z
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-3))
}
