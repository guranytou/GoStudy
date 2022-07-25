package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	res, num := 0, 0
	for i := 1; i <= n; i++ {
		num = findSumOfDigits(i)
		if a <= num && num <= b {
			res += i
		}
	}
	fmt.Printf("%d", res)
}

func findSumOfDigits(num int) int {
	res := 0
	for {
		res += num % 10
		num = num / 10
		if num == 0 {
			return res
		}
	}
}
