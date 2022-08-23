package main

import "fmt"

func main() {
	for {
		var s string
		fmt.Scanf("%s", &s)
		if s == "-" {
			break
		}

		var m int
		fmt.Scanf("%d", &m)

		var h int
		for i := 0; i < m; i++ {
			fmt.Scanf("%d", &h)
			f := s[:h]
			l := s[h:]
			s = l + f
		}

		fmt.Println(s)
	}
}
