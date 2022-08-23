package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	t, h := 0, 0
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		cards := strings.Split(sc.Text(), " ")

		if cards[0] == cards[1] {
			t++
			h++
		} else if cards[0] > cards[1] {
			t += 3
		} else {
			h += 3
		}
	}
	fmt.Println(t, h)
}
