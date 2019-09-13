# wtw
wtw is a CLI tool to show the current weather on your terminal emulator using the [OpenWeatherMap API](https://openweathermap.org/current).

## Install
You can download one of the binaries available on the [releases page](https://github.com/efreitasn/wtw/releases) or install it using Go v1.12.9 or higher.

### Installing using go
```bash
go get github.com/efreitasn/wtw
$(go env GOPATH)/bin/wtw
```

## Prerequisites
* [OpenWeatherMap API key](https://openweathermap.org/appid).
* OpenWeatherMap city ID of the location you want to get weather information about.

### Getting the city ID
* Search for your city in [this URL](https://openweathermap.org/find).
* Once you find your city in the search results, get the city ID from the address bar (the digits after `https://openweathermap.org/city/`).

## How to use
First of all, you need to set your API key and your city ID. In order to do that, run the following command:

```bash
wtw set --city-id YOUR_CITY_ID --api-key YOUR_API_KEY
```

To get the current weather info, run:
```bash
wtw
```