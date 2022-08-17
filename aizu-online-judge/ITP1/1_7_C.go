package main

import (
	"fmt"
)

func main() {
	var r, c int
	fmt.Scanf("%d %d", &r, &c)
	var sum int
	w := make([]int, c)
	hSum := make([]int, c)
	for i := 0; i < r; i++ {
		sum = 0
		for j := 0; j < c; j++ {
			fmt.Scanf("%d", &w[j])
			sum += w[j]
			hSum[j] += w[j]

			if j == c-1 {
				fmt.Printf("%d %d\n", w[j], sum)
			} else {
				fmt.Printf("%d ", w[j])
			}
		}
	}

	sum = 0
	for i := range hSum {
		sum += hSum[i]

		if i == len(hSum)-1 {
			fmt.Printf("%d %d\n", hSum[i], sum)
		} else {
			fmt.Printf("%d ", hSum[i])
		}
	}
}
