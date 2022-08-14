package main

import (
	"fmt"
	"strings"
)

func main() {
	var H, W int
	for {
		fmt.Scanf("%d %d", &H, &W)

		for i := 1; i <= H; i++ {
			fmt.Println(strings.Repeat("#", W))
		}
		if H == 0 && W == 0 {
			break
		}
		fmt.Printf("\n")
	}
}
