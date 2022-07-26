package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d[i])
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	tmp := d[0]
	res := 1

	for i := 1; i < n; i++ {
		if tmp > d[i] {
			res++
		}
		tmp = d[i]
	}

	fmt.Printf("%d", res)
}
