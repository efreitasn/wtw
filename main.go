package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// WeatherResponse is the data returned from a request to the OpenWeather API.
type WeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	fetchWeatherInfo := len(os.Args) == 1

	if !fetchWeatherInfo {
		if os.Args[1] != "set" {
			logError.Fatal("Error while parsing arguments.")
		}

		setFlag := flag.NewFlagSet("set", flag.ExitOnError)
		cityID := setFlag.String("city-id", "", "Set the city ID.")
		apiKey := setFlag.String("api-key", "", "Set the API key.")
		tempUnit := setFlag.String("unit", "", "Set the temperature unit.")

		setFlag.Parse(os.Args[2:])

		if *cityID == "" && *apiKey == "" && *tempUnit == "" {
			logError.Print("Error while parsing arguments.\n\n")
			logError.Println("Usage of set:")

			setFlag.PrintDefaults()
		}

		c, err := GetConfig()

		if os.IsNotExist(err) {
			c = NewConfig()
		} else if err != nil {
			logError.Fatal("Error while reading the config file.")
		}

		if *cityID != "" {
			c.CityID = *cityID
		}

		if *apiKey != "" {
			c.APIKey = *apiKey
		}

		if *tempUnit != "" {
			switch *tempUnit {
			case TempCelsius:
				fallthrough
			case TempFahrenheit:
				fallthrough
			case TempKelvin:
				c.Unit = *tempUnit
			default:
				logError.Fatal("Invalid temperature unit.")
			}
		}

		err = c.Write()

		if err != nil {
			logError.Fatal("Error while writing the config file")
		}

		return
	}

	c, err := GetConfig()

	if os.IsNotExist(err) {
		logError.Fatal("--city-id and --api-key are not set.")
	}

	if err != nil {
		logError.Fatal("Error while reading the config file")
	}

	if c.APIKey == "" {
		logError.Fatal("--api-key is not set.")
	}

	if c.CityID == "" {
		logError.Fatal("--city-id is not set.")
	}

	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?units=%v&id=%v&APPID=%v",
		configUnitToAPIUnit(c.Unit),
		c.CityID,
		c.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		logError.Fatal("Error while making HTTP request.")
	}
	defer resp.Body.Close()

	var respData WeatherResponse

	jsonDecoder := json.NewDecoder(resp.Body)

	jsonDecoder.Decode(&respData)

	if len(respData.Weather) > 0 {
		logSuccess.Printf("%v°%v - %v\n", respData.Main.Temp, c.Unit, respData.Weather[0].Main)
	} else {
		logSuccess.Printf("%v°%v\n", respData.Main.Temp, c.Unit)
	}
}
