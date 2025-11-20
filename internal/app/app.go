package app

import (
	"fmt"
	"github.com/rende11/weather/internal/client/http/ipify"
	"github.com/rende11/weather/internal/client/http/location"
	"github.com/rende11/weather/internal/client/http/meteo"
	"log"
	"net/http"
	"time"
)

func Run() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	ipifyClient := ipify.NewClient(httpClient)
	locationClient := location.NewClient(httpClient)
	meteoClient := meteo.NewClient(httpClient)

	ip, err := ipifyClient.GetIp()
	if err != nil {
		log.Println(err)
		return
	}

	loc, err := locationClient.GetLocation(ip.Ip)
	if err != nil {
		log.Println(err)
		return
	}

	forecast, err := meteoClient.Forecast(loc.Latitude, loc.Longitude, loc.Timezone)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Temp: %.0f%s\nWind: %.1f%s\nHumidity: %d%s\n%d\n",
		forecast.Current.Temperature, forecast.CurrentUnits.TemperatureUnits,
		forecast.Current.WindSpeed, forecast.CurrentUnits.WindSpeedUnits,
		forecast.Current.Humidity, forecast.CurrentUnits.HumidityUnits,
		forecast.Current.WeatherCode,
	)
}
