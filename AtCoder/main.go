package main

import "fmt"

func main() {
	var n, a, b, ans int
	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n; i++ {
		v := i
		sum := 0

		for v > 0 {
			sum += v % 10
			v = v / 10
		}

		if sum >= a && sum <= b {
			ans += i
		}
	}

	fmt.Println(ans)
}
