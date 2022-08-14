package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scanf("%f", &r)

	fmt.Printf("%f %f", r*r*math.Pi, r*2*math.Pi)
}
