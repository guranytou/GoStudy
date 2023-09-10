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
		var b byte
		if str[i] >= 'A' && str[i] <= 'Z' {
			if rune(str[i])+rune(n) > 'Z' {
				b = 'A' + (str[i] + byte(n) - 'Z') - 1
			} else {
				b = str[i] + byte(n)
			}
		} else if str[i] >= 'a' && str[i] <= 'z' {
			if rune(str[i])+rune(n) > 'z' {
				b = 'a' + (str[i] + byte(n) - 'z') - 1
			} else {
				b = str[i] + byte(n)
			}
		} else {
			b = str[i]
		}
		convert[i] = string(b)
	}
	return strings.Join(convert, "")
}
