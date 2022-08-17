package main

import "fmt"

func main() {
	m := make(map[string]map[int]int)
	marks := []string{"S", "H", "C", "D"}

	for i := 0; i < 4; i++ {
		m[marks[i]] = make(map[int]int)
		for j := 1; j <= 13; j++ {
			m[marks[i]][j] = 0
		}
	}

	var n int
	fmt.Scanf("%d", &n)

	var mark string
	var number int

	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &mark, &number)
		m[mark][number] = 1
	}

	for i := 0; i < 4; i++ {
		for j := 1; j <= 13; j++ {
			if m[marks[i]][j] == 0 {
				fmt.Println(marks[i], j)
			}
		}
	}
}
