package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, C float64
	fmt.Scanf("%f %f %f", &a, &b, &C)

	rad := C * math.Pi / 180
	S := 0.5 * a * b * math.Sin(rad)
	h := 2 * S / a
	c := math.Sqrt(a*a + b*b - 2*a*b*math.Cos(rad))
	L := a + b + c

	fmt.Printf("%f\n%f\n%f", S, L, h)
}
