package main

import (
	"fmt"
	"sort"
)

func main() {
	n := make([]int, 3)
	fmt.Scanf("%d %d %d", &n[0], &n[1], &n[2])
	sort.Slice(n, func(i int, j int) bool { return n[i] < n[j] })
	fmt.Printf("%d %d %d\n", n[0], n[1], n[2])
}
