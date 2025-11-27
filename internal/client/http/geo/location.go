package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ipLocationServiceUrl = "http://ip-api.com/json/%s"
const cityLocationServiceUrl = "https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json"

type IPLocationResponse struct {
	Status    string  `json:"status"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Timezone  string  `json:"timezone"`
}

type CityLocationResponse struct {
	Results []CityLocationResult `json:"results"`
}

type CityLocationResult struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
	Timezone  string  `json:"timezone"`
}

type LocationInfo struct {
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Timezone  string  `json:"timezone"`
}

type LocationClient interface {
	GetLocationByIP(ip string) (LocationInfo, error)
	GetLocationByCity(city string) (LocationInfo, error)
}
type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

func (c *Client) GetLocationByIP(ip string) (LocationInfo, error) {
	op := "GetLocationByIP"
	resp, err := c.httpClient.Get(fmt.Sprintf(ipLocationServiceUrl, ip))

	if err != nil {
		return LocationInfo{}, fmt.Errorf("%s: error while getting location by ip: %v", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationInfo{}, fmt.Errorf("%s: getting location by ip status code %d", op, resp.StatusCode)
	}

	var locationResponse IPLocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)

	if err != nil {
		return LocationInfo{}, fmt.Errorf("%s: error while parsing location by ip response: %v", op, err)
	}

	return LocationInfo{
		Latitude:  locationResponse.Latitude,
		Longitude: locationResponse.Longitude,
		City:      locationResponse.City,
		Country:   locationResponse.Country,
		Timezone:  locationResponse.Timezone,
	}, nil
}

func (c *Client) GetLocationByCity(city string) (LocationInfo, error) {
	op := "GetLocationByCity"
	resp, err := c.httpClient.Get(fmt.Sprintf(cityLocationServiceUrl, city))

	if err != nil {
		return LocationInfo{}, fmt.Errorf("%s: error while getting location by city: %v", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationInfo{}, fmt.Errorf("%s: error while getting location by city status code %d", op, resp.StatusCode)
	}

	var locationResponse CityLocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)

	if err != nil {
		return LocationInfo{}, fmt.Errorf("%s: error while getting location by city response %d", op, err)
	}

	if locationResponse.Results == nil {
		return LocationInfo{},
			fmt.Errorf("%s: error while getting location by city response - city not found %s", op, city)
	}

	cityLocation := locationResponse.Results[0]

	fmt.Println("CITY", cityLocation)

	return LocationInfo{
		Latitude:  cityLocation.Latitude,
		Longitude: cityLocation.Longitude,
		City:      cityLocation.Name,
		Country:   cityLocation.Country,
		Timezone:  cityLocation.Timezone,
	}, nil
}
