package main

import (
	"fmt"
)

func main() {
	var H, W int
	for {
		fmt.Scanf("%d %d", &H, &W)

		if H == 0 && W == 0 {
			break
		}

		for i := 1; i <= H; i++ {
			for j := 1; j <= W; j++ {
				if i%2 != 0 && j%2 != 0 {
					fmt.Printf("#")
				} else if i%2 != 0 && j%2 == 0 {
					fmt.Printf(".")
				} else if i%2 == 0 && j%2 != 0 {
					fmt.Printf(".")
				} else if i%2 == 0 && j%2 == 0 {
					fmt.Printf("#")
				}

				if j == W {
					fmt.Printf("\n")
				}
			}
		}
		fmt.Printf("\n")
	}
}
