package ipify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ipServiceUrl = "https://api.ipify.org?format=json"

type Response struct {
	Ip string `json:"ip"`
}

type Client struct {
	httpClient *http.Client
}

type IPClient interface {
	GetIP() (Response, error)
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

func (c *Client) GetIP() (Response, error) {
	resp, err := c.httpClient.Get(ipServiceUrl)
	if err != nil {
		return Response{}, fmt.Errorf("error while getting IP address: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("getting IP address status code %d", resp.StatusCode)
	}

	var ipResp Response
	err = json.NewDecoder(resp.Body).Decode(&ipResp)

	if err != nil {
		return Response{}, fmt.Errorf("error while parsing IP address response: %v", err)
	}

	return ipResp, nil
}
