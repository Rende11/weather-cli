package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

type Response struct {
	Timezone     string       `json:"timezone"`
	CurrentUnits CurrentUnits `json:"current_units"`
	Current      Current      `json:"current"`
}

type CurrentUnits struct {
	TemperatureUnits string `json:"temperature_2m"`
	HumidityUnits    string `json:"relative_humidity_2m"`
	WindSpeedUnits   string `json:"wind_speed_10m"`
}

type Current struct {
	Temperature float64 `json:"temperature_2m"`
	Humidity    int     `json:"relative_humidity_2m"`
	WeatherCode int     `json:"weather_code"`
	WindSpeed   float64 `json:"wind_speed_10m"`
}

func (c *Client) Forecast(lat, long float64, timezone string) (Response, error) {
	params := url.Values{}
	params.Add("latitude", strconv.FormatFloat(lat, 'f', 2, 64))
	params.Add("longitude", strconv.FormatFloat(long, 'f', 2, 64))
	params.Add("current", "temperature_2m,relative_humidity_2m,apparent_temperature,is_day,wind_speed_10m,wind_direction_10m,wind_gusts_10m,precipitation,rain,showers,snowfall,weather_code,cloud_cover,surface_pressure")
	params.Add("wind_speed_unit", "ms")
	params.Add("timezone", timezone)

	reqUrl := url.URL{
		Scheme:   "https",
		Host:     "api.open-meteo.com",
		Path:     "v1/forecast",
		RawQuery: params.Encode(),
	}

	resp, err := c.httpClient.Get(reqUrl.String())

	if err != nil {
		return Response{}, fmt.Errorf("error while getting meteo: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("getting meteo status code %d", resp.StatusCode)
	}

	var forecastResponse Response

	err = json.NewDecoder(resp.Body).Decode(&forecastResponse)

	if err != nil {
		return Response{}, fmt.Errorf("error while parsing meteo response: %v", err)
	}

	return forecastResponse, nil
}
