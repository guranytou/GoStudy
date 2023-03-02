package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	flag := false
	for bits := 0; bits < (1 << uint(n)); bits++ {
		sum := 0
		for i := 0; i < n; i++ {
			if (bits>>uint(i))&1 == 1 {
				sum += a[i]
			}
		}

		if sum == s {
			flag = true
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
