package integrations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPGeolocationService struct {
	apiUrl     string
	apiKey     string
	httpClient http.Client
}

func (s *IPGeolocationService) Lookup(ip string) (IPGeolocationResponse, error) {
	requestUrl := fmt.Sprintf("%s?apiKey=%s&ip=%s", s.apiUrl, s.apiKey, ip)

	response, err := s.httpClient.Get(requestUrl)

	if err != nil {
		return IPGeolocationResponse{}, err
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return IPGeolocationResponse{}, err
	}

	var parsedResponse IPGeolocationResponse
	err = json.Unmarshal(bodyData, &parsedResponse)

	if err != nil {
		return IPGeolocationResponse{}, err
	}

	return parsedResponse, nil
}

func NewIPGeolocationService(apiKey string) *IPGeolocationService {
	return &IPGeolocationService{
		apiUrl: "https://api.ipgeolocation.io/ipgeo",
		apiKey: apiKey,
	}
}

type IPGeolocationResponse struct {
	IP            string `json:"ip"`
	ContinentCode string `json:"continent_code"`
	ContinentName string `json:"continent_name"`
	CountryCode   string `json:"country_code2"`
	CountryName   string `json:"country_name"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Languages     string `json:"languages"`
}
