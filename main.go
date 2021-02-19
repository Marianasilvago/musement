package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/musement/entity"
	"github.com/musement/internal/config"
	"github.com/musement/internal/version"
	"github.com/musement/utils"
	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		handleError(err)
	}
}

func run() error {
	utils.InitLog()
	config, err := config.GetConfig()
	if err != nil {
		return errors.WithMessage(err, "get config")
	}
	utils.LogStart(version.VersionString("Musement-API"), config.Environment)

	urlString := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s,%s&days=%s", config.WeatherAPIKey, fmt.Sprintf("%f", config.Latitude), fmt.Sprintf("%f", config.Longitude), strconv.Itoa(config.ForcastDays))
	resp, err := http.Get(urlString)
	if err != nil {
		return errors.WithMessage(err, "HTTP GET request")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithMessage(err, "ReadAll")
	}
	resp.Body.Close()

	var v entity.ForecastData
	err = json.Unmarshal(data, &v)
	if err != nil {
		return errors.WithMessage(err, "Unmarshall")
	}

	weatherForcast := []string{}
	city := v.Location.Name
	for _, val2 := range v.Forecast.Forecastday {
		weatherForcast = append(weatherForcast, val2.Day.Condition.Text)
	}

	log.Println("Processed city ", city, "|", strings.Join(weatherForcast, " - "))
	return nil
}

func handleError(err error) {
	log.Fatalf("%s", err)
}
