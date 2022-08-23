package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, p string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &p)

	s = strings.Repeat(s, 2)
	if strings.Contains(s, p) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
