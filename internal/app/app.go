package app

import (
	"fmt"
	"github.com/rende11/weather/internal/client/http/geo"
	"github.com/rende11/weather/internal/client/http/ipify"
	"github.com/rende11/weather/internal/client/http/meteo"
	"log"
	"net/http"
	"time"
)

func Run(city string) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	ipifyClient := ipify.NewClient(httpClient)
	locationClient := geo.NewClient(httpClient)
	meteoClient := meteo.NewClient(httpClient)

	location, err := getLocation(ipifyClient, locationClient, city)

	if err != nil {
		log.Println(err)
		return
	}

	forecast, err := meteoClient.Forecast(location.Latitude, location.Longitude, location.Timezone)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("City: %s\nTemp: %.0f%s\nWind: %.1f%s\nHumidity: %d%s\n%d\n",
		location.City,
		forecast.Current.Temperature, forecast.CurrentUnits.TemperatureUnits,
		forecast.Current.WindSpeed, forecast.CurrentUnits.WindSpeedUnits,
		forecast.Current.Humidity, forecast.CurrentUnits.HumidityUnits,
		forecast.Current.WeatherCode,
	)
}
