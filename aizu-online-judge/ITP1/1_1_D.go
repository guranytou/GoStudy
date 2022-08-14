package main

import (
	"fmt"
)

func main() {
	var S int
	fmt.Scanf("%d", &S)

	h := S / 3600
	m := (S - h*3600) / 60
	s := S % 60
	fmt.Printf("%d:%d:%d\n", h, m, s)
}
