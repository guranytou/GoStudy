package main

import "fmt"

func main() {
	i := 1
	var n int
	for {
		fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", i, n)
		i++
	}
}
