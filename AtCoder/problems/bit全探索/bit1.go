package main

import "fmt"

func main() {
	n := 5
	a := []int{1, 2, 3, 6, 8}
	w := 10

	cnt := 0
	for bits := 0; bits < (1 << uint(n)); bits++ {
		sum := 0
		for i := 0; i < n; i++ {
			if (bits>>uint(i))&1 == 1 {
				sum += a[i]
			}
		}
		if sum == w {
			cnt++
		}
	}
	fmt.Println(cnt)
}
