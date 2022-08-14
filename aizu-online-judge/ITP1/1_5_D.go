package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 || strings.Contains(strconv.Itoa(i), "3") {
			fmt.Printf(" %d", i)
		}
	}
}
