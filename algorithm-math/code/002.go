package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	total := a + b + c
	fmt.Println(total)

}
