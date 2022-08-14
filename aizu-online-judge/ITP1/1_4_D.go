package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
		sum += a[i]
	}
	sort.Slice(a, func(i int, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a[0], a[n-1], sum)
}
