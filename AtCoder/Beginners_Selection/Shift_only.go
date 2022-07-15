package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}

	count := 0
	for {
		for i, v := range array {
			if v%2 == 1 {
				fmt.Printf("%d\n", count)
				return
			}
			array[i] = v / 2
		}
		count++
	}
}
