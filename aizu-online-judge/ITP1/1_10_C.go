package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var n int
		fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}

		s := make([]int, n)
		sum := 0
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &s[i])
			sum += s[i]
		}

		var m float64
		m = float64(sum) / float64(n)

		dispSum := 0.0
		for i := 0; i < n; i++ {
			dispSum += math.Pow(float64(s[i])-m, 2)
		}

		fmt.Println(math.Pow(dispSum/float64(n), 0.5))
	}
}
