package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	ans := a * b

	if ans%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
