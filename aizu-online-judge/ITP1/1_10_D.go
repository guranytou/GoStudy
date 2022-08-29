package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	x := createSlice(n)
	y := createSlice(n)

	p1 := calEuclideanDistance(x, y, n, 1)
	p2 := calEuclideanDistance(x, y, n, 2)
	p3 := calEuclideanDistance(x, y, n, 3)

	fmt.Printf("%f\n", p1)
	fmt.Printf("%f\n", p2)
	fmt.Printf("%f\n", p3)

	pInf := 0.0
	for i := 0; i < n; i++ {
		tmp := math.Abs(float64(x[i]) - float64(y[i]))
		if pInf < tmp {
			pInf = math.Abs(float64(x[i]) - float64(y[i]))
		}
	}
	fmt.Printf("%f", pInf)
}

func createSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	return s
}

func calEuclideanDistance(x []int, y []int, n int, p int) float64 {
	dSum := 0.0
	for i := 0; i < n; i++ {
		dSum += math.Pow(math.Abs(float64(x[i])-float64(y[i])), float64(p))
	}

	var f float64
	f = 1.0 / float64(p)
	d := math.Pow(dSum, f)

	return d
}
