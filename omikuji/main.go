package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "omikuji"
	app.Usage = "Play omikuji"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		omikuji()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func omikuji() {
	rand.Seed(time.Now().UnixNano())
	luck := rand.Intn(12)
	switch luck {
	case 0:
		fmt.Println("大吉")
	case 1, 2:
		fmt.Println("中吉")
	case 3, 4, 5:
		fmt.Println("小吉")
	case 6, 7, 8:
		fmt.Println("末吉")
	case 9, 10:
		fmt.Println("吉")
	case 11:
		fmt.Println("凶")
	}
}
