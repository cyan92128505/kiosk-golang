package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// WebConfig is setup web server
type WebConfig struct {
	URL int `json:"url"`
}

func (c WebConfig) toString() string {
	return toJSON(c)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Survival!")
	})
	e.POST("/Setup", func(c echo.Context) error {
		url := c.FormValue("url")
		return c.String(http.StatusOK, url)
	})
	e.Logger.Fatal(e.Start(":39268"))
}

func getConfig() WebConfig {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c WebConfig
	json.Unmarshal(raw, &c)
	return c
}
