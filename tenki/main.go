package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type OWMAPIResponce struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Dt      int64     `json:"dt`
	Name    string    `json:"name"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

func main() {
	app := cli.NewApp()
	app.Name = "tenki"
	app.Usage = "Watch weather"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		city := c.Args().Get(0)
		tenki(city)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func tenki(city string) {
	token := "7cb2eee4863ef064326ec2d60935fa1b"
	endPoint := "https://api.openweathermap.org/data/2.5/weather"

	values := url.Values{}
	values.Set("APPID", token)

	if city == "" {
		values.Set("q", "tokyo,jp")
	} else {
		values.Set("q", city)
	}

	resp, err := http.Get(endPoint + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var apiRes OWMAPIResponce
	if err := json.Unmarshal(body, &apiRes); err != nil {
		panic(err)
	}

	fmt.Printf("時刻: %s\n", time.Unix(apiRes.Dt, 0))
	fmt.Printf("場所: %s\n", apiRes.Name)

	if apiRes.Weather[0].Main == "Clear" {
		fmt.Println("天気: 晴れ")
	} else if apiRes.Weather[0].Main == "Clouds" {
		fmt.Println("天気: くもり")
	} else if apiRes.Weather[0].Main == "Rain" {
		fmt.Println("天気: 雨")
	} else if apiRes.Weather[0].Main == "Snow" {
		fmt.Println("天気: 雪")
	} else {
		fmt.Printf("天気: %s\n", apiRes.Weather[0].Main)
	}

	fmt.Printf("気温: %g\n", math.Floor((apiRes.Main.Temp-273.15)*10)/10)
}
