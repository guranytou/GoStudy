package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	var q int
	fmt.Scanf("%d", &q)

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		sc.Scan()
		order := strings.Split(sc.Text(), " ")
		start, _ := strconv.Atoi(order[1])
		end, _ := strconv.Atoi(order[2])

		if order[0] == "print" {
			fmt.Println(str[start : end+1])
		} else if order[0] == "reverse" {
			word := str[start : end+1]
			var revS string
			for i := len(word) - 1; i >= 0; i-- {
				revS = revS + string(word[i])
			}
			str = str[:start] + revS + str[end+1:]
		} else if order[0] == "replace" {
			str = str[:start] + order[3] + str[end+1:]
		}
	}
}
