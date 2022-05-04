package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.Bool("n", false, "line number")
	flag.Parse()
	t := 0

	for _, v := range flag.Args() {
		sf, err := os.Open(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ファイルを開けませんでした", err)
			os.Exit(1)
		}
		defer sf.Close()

		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			if *n == true {
				fmt.Println(t, scanner.Text())
			} else {
				fmt.Println(scanner.Text())
			}
			t++
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
			os.Exit(1)
		}

	}
}
