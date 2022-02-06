package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	count := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i+j <= s {
				count++
			}
		}
	}
	fmt.Println(count)
}
