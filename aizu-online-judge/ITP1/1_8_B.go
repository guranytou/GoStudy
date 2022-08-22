package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	for {
		fmt.Scanf("%s", &s)
		if s == "0" {
			break
		} else {
			sum := 0
			for i := 0; i < len(s); i++ {
				n, _ := strconv.Atoi(string(s[i]))
				sum += n
			}
			fmt.Println(sum)
		}
	}
}
