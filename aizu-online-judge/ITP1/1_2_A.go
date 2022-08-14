package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scanf("%f %f", &a, &b)

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a == b")
	}
}
