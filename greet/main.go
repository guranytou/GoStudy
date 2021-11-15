package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "Give a goo greeting."
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
