package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := n - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d\n", a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
}
