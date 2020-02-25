package gismeteo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const (
	apiURL       = "api.gismeteo.net"
	apiCurrent   = "v2/weather/current"
	apiToken     = "56b30cb255.3443075"
	apiTokenName = "X-Gismeteo-Token"
)

//GetWeather requests weather by latitude and longitude
func GetWeather(latitude, longitude float64) (string, error) {
	query := url.Values{}
	query.Add("latitude", strconv.FormatFloat(latitude, 'f', -1, 64))
	query.Add("longitude", strconv.FormatFloat(latitude, 'f', -1, 64))

	url := buildURL(apiCurrent, query)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add(apiTokenName, apiToken)
	resp, _ := http.DefaultClient.Do(request)

	decoder := json.NewDecoder(resp.Body)
	var str string
	decoder.Decode(&str)
	return str, nil
}

func buildURL(method string, query url.Values) string {
	url := url.URL{
		Scheme:   "https",
		Host:     apiURL,
		Path:     method,
		RawQuery: query.Encode(),
	}

	return url.String()
}
