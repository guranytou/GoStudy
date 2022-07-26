package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &cards[i])
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += cards[i]
		} else {
			sum -= cards[i]
		}
	}

	fmt.Printf("%d", sum)
}
