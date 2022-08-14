package main

import (
	"fmt"
)

func main() {
	var x, y int
	for {
		fmt.Scanf("%d %d", &x, &y)
		if x == 0 && y == 0 {
			break
		}
		if x > y {
			fmt.Printf("%d %d\n", y, x)
		} else {
			fmt.Printf("%d %d\n", x, y)
		}
	}
}
