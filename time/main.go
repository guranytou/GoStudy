package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "time"
	app.Usage = "Watch time"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println(time.Now())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
