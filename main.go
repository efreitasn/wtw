package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type weatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	apiKey := os.Getenv("WTW_API_KEY")

	if apiKey == "" {
		panic("WTW_API_KEY env variable not set.\n")
	}

	cityID := os.Getenv("WTW_CITY_ID")

	if cityID == "" {
		panic("WTW_CITY_ID env variable not set.\n")
	}

	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?id=%v&units=metric&APPID=%v",
		cityID,
		apiKey,
	)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprint(os.Stderr, "Error while making HTTP request.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	var respData weatherResponse

	jsonDecoder := json.NewDecoder(resp.Body)

	jsonDecoder.Decode(&respData)

	fmt.Printf("%v ℃\n", respData.Main.Temp)
}
