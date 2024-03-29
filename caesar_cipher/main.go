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
	n := byte(r.Intn(25))

	convert := caesar_cipher(str, n)

	fmt.Println(str, "を", n, "文字ずらすと", convert)
}

func caesar_cipher(str string, n byte) string {
	convert := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		var b byte
		if str[i] >= 'A' && str[i] <= 'Z' {
			if str[i]+n > 'Z' {
				b = 'A' + (str[i] + n - 'Z') - 1
			} else {
				b = str[i] + n
			}
		} else if str[i] >= 'a' && str[i] <= 'z' {
			if str[i]+n > 'z' {
				b = 'a' + (str[i] + n - 'z') - 1
			} else {
				b = str[i] + n
			}
		} else {
			b = str[i]
		}
		convert[i] = string(b)
	}
	return strings.Join(convert, "")
}
