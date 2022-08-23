package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var w string
	fmt.Scanf("%s", &w)

	sc := bufio.NewScanner(os.Stdin)
	cnt := 0
	for {
		sc.Scan()
		if sc.Text() == "END_OF_TEXT" {
			break
		}

		t := strings.Split(strings.ToLower(sc.Text()), " ")
		for i := 0; i < len(t); i++ {
			if t[i] == w {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
