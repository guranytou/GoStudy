package main

import (
	"fmt"
)

func main() {
	divide := []string{"dream", "dreamer", "erase", "eraser"}
	for i := 0; i < len(divide); i++ {
		divide[i] = reverse(divide[i])
	}

	var s string
	fmt.Scanf("%s", &s)
	s = reverse(s)

	can := true
	for i := 0; i < len(s); i++ {
		divCan := false
		for j := 0; j < len(divide); j++ {
			d := divide[j]
			if i+len(d) > len(s) {
				continue
			}
			if s[i:i+len(d)] == d {
				divCan = true
				i += len(d) - 1
			}
		}
		if !divCan {
			can = false
			break
		}
	}

	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
