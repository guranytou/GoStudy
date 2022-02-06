package main

import (
	"fmt"
)

func main() {
	var n int
	sum := 0
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		sum = sum + a
	}

	fmt.Println(sum)
}
