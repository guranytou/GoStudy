package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := createMap()
	m2 := createMap()
	m3 := createMap()
	m4 := createMap()

	var n int
	fmt.Scanf("%d", &n)

	var b, f, r, v int

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d", &b, &f, &r, &v)

		if b == 1 {
			m1[f][r] += v
		} else if b == 2 {
			m2[f][r] += v
		} else if b == 3 {
			m3[f][r] += v
		} else if b == 4 {
			m4[f][r] += v
		}
	}
	printMap(m1)
	fmt.Println(strings.Repeat("#", 20))
	printMap(m2)
	fmt.Println(strings.Repeat("#", 20))
	printMap(m3)
	fmt.Println(strings.Repeat("#", 20))
	printMap(m4)
}

func createMap() map[int]map[int]int {
	m := make(map[int]map[int]int)
	for i := 1; i < 4; i++ {
		m[i] = make(map[int]int)
		for j := 1; j < 11; j++ {
			m[i][j] = 0
		}
	}

	return m
}

func printMap(m map[int]map[int]int) {
	for i := 1; i < 4; i++ {
		for j := 1; j < 11; j++ {
			fmt.Printf(" %d", m[i][j])
		}
		fmt.Printf("\n")
	}
}
