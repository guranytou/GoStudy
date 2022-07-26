package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	res10000, res5000, res1000 := -1, -1, -1
	for i := 0; i <= x; i++ {
		for j := 0; j <= x-i; j++ {
			k := x - i - j
			total := 10000*i + 5000*j + 1000*k

			if y == total {
				res10000 = i
				res5000 = j
				res1000 = k
			}
		}
	}
	fmt.Printf("%d %d %d", res10000, res5000, res1000)
}
