package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationServiceUrl = "http://ip-api.com/json/%s"

type Response struct {
	Status    string  `json:"status"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Timezone  string  `json:"timezone"`
}

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

func (c *Client) GetLocation(ip string) (Response, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf(locationServiceUrl, ip))

	if err != nil {
		return Response{}, fmt.Errorf("error while getting location: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("getting location status code %d", resp.StatusCode)
	}

	var locationResponse Response
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)

	if err != nil {
		return Response{}, fmt.Errorf("error while parsing location response: %v", err)
	}

	return locationResponse, nil
}
