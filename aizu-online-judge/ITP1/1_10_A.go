package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scanf("%f %f %f %f", &x1, &y1, &x2, &y2)
	fmt.Println(math.Hypot(x2-x1, y2-y1))
}
