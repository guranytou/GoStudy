package main

import (
	"fmt"
	"strings"
)

func main() {
	var H, W int
	for {
		fmt.Scanf("%d %d", &H, &W)

		if H == 0 && W == 0 {
			break
		}

		for i := 1; i <= H; i++ {
			if i == 1 || i == H {
				fmt.Println(strings.Repeat("#", W))
			} else {
				for j := 1; j <= W; j++ {
					if j == 1 {
						fmt.Printf("#")
					} else if j == W {
						fmt.Println("#")
					} else {
						fmt.Printf(".")
					}
				}
			}
		}
		fmt.Printf("\n")
	}
}
