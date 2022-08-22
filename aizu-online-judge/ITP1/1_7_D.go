package main

import (
	"fmt"
)

func main() {
	var n, m, l int
	fmt.Scanf("%d %d %d", &n, &m, &l)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scanf("%d", &b[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			tmp := 0
			for k := 0; k < m; k++ {
				tmp += a[i][k] * b[k][j]
			}
			if j == l-1 {
				fmt.Println(tmp)
			} else {
				fmt.Printf("%d ", tmp)
			}
		}
	}
}
