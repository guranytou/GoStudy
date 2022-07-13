package main

import (
	"fmt"
)

func main() {
	var s string
	var count int

	fmt.Scanf("%s", &s)

	for _, c := range s {
		if string(c) == "1" {
			count++
		}
	}

	fmt.Printf("%d", count)
}
