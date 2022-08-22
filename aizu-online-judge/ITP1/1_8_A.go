package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	tmp := sc.Text()
	for _, c := range tmp {
		if c >= rune('A') && c <= rune('Z') {
			fmt.Printf("%s", strings.ToLower(string(c)))
		} else if c >= rune('a') && c <= rune('z') {
			fmt.Printf("%s", strings.ToUpper(string(c)))
		} else {
			fmt.Printf("%s", string(c))
		}
	}
	fmt.Printf("\n")
}
