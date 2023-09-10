package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("文字列を入力してください: ")
	scanner.Scan()
	str := scanner.Text()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(25)

	convert := caesar_cipher(str, n)

	fmt.Println(str, "を", n, "文字ずらすと", convert)
}

func caesar_cipher(str string, n int) string {
	convert := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		var r rune
		if str[i] >= 'A' && str[i] <= 'Z' {
			if rune(str[i])+rune(n) > 'Z' {
				r = 'A' + (rune(str[i]) + rune(n) - 'Z') - 1
			} else {
				r = rune(str[i]) + rune(n)
			}
		} else if str[i] >= 'a' && str[i] <= 'z' {
			if rune(str[i])+rune(n) > 'z' {
				r = 'a' + (rune(str[i]) + rune(n) - 'z') - 1
			} else {
				r = rune(str[i]) + rune(n)
			}
		} else {
			r = rune(str[i])
		}
		convert[i] = string(r)
	}
	return strings.Join(convert, "")
}
