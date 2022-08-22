package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	keys := "abcdefghijklmnopqrstuvwxyz"
	cnt := make(map[string]int)
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(keys); i++ {
		cnt[string(keys[i])] = 0
	}

	for {
		sc.Scan()
		if sc.Text() == "" {
			break
		} else {
			input := sc.Text()
			for i := 0; i < len(input); i++ {
				for j := 0; j < len(keys); j++ {
					tmp := strings.ToLower(string(input[i]))
					if tmp == string(keys[j]) {
						cnt[string(keys[j])]++
					}
				}
			}
		}
	}

	for i := 0; i < len(keys); i++ {
		fmt.Printf("%s : %d\n", string(keys[i]), cnt[string(keys[i])])
	}
}
