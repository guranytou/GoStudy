package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	for bits := 0; bits < (1 << uint(n)); bits++ {
		if bits == x {
			for i := 0; i < n; i++ {
				if (bits>>uint(i))&1 == 1 {
					sum += a[i]
				}
			}
		}
	}
	fmt.Println(sum)
}
