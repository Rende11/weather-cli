package app

import (
	"fmt"
	"github.com/rende11/weather/internal/client/http/geo"
	"github.com/rende11/weather/internal/client/http/ipify"
)

func getLocation(ipifyClient ipify.IPClient, locationClient geo.LocationClient, city string) (geo.LocationInfo, error) {
	if city != "" {
		fmt.Println("provided city", city)
		return locationClient.GetLocationByCity(city)
	}

	fmt.Println("determine city by ip address")

	ip, err := ipifyClient.GetIP()
	if err != nil {
		return geo.LocationInfo{}, fmt.Errorf("error getting ip for current location: %v", err)
	}

	return locationClient.GetLocationByIP(ip.Ip)
}
