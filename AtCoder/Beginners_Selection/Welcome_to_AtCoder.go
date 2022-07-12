package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var s string

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)

	ans := a + b + c

	fmt.Printf("%d %s\n", ans, s)
}
