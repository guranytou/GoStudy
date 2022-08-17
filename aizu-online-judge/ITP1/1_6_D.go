package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	a := make(map[int]map[int]int)
	b := make([]int, m)

	var tmp int
	for i := 0; i < n; i++ {
		a[i] = make(map[int]int)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &tmp)
			a[i][j] = tmp
		}
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += a[i][j] * b[j]
		}
		fmt.Println(sum)
	}
}
