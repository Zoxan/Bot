package yandexweather

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const (
	apiURL       = "api.weather.yandex.ru"
	apiInformers   = "v1/informers"
	apiForecast   = "v1/forecast"
	apiTokenName = "X-Yandex-API-Key"
	apiLang = "ru_RU"
)

type WeatherDescription struct {
	Fact struct {
		Temperature float64 `json:"temp"`
		Condition string `json:"condition"`
		WindSpeed float64 `json:"wind_speed"`
		WindDirection string `json:"wind_dir"`
	} `json:"fact"`
}

var apiToken string
var yandexTariff string

func Init(token string, tariff string) {
	apiToken = token
	yandexTariff = tariff
}

func GetWeatherText(latitude, longitude float64) (string, error) {
	weather, err := GetWeather(latitude, longitude)
	if err != nil {
		return "", nil
	}

	return GetText(*weather), nil
}

//GetWeather calls Yandex.Weather api by latitude and longitude and return weather descr
func GetWeather(latitude, longitude float64) (*WeatherDescription, error) {
	query := url.Values{}
	query.Add("lat", strconv.FormatFloat(latitude, 'f', -1, 64))
	query.Add("lon", strconv.FormatFloat(longitude, 'f', -1, 64))

	url := buildURL(getApiMethod(), query)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add(apiTokenName, apiToken)
	resp, _ := http.DefaultClient.Do(request)

	decoder := json.NewDecoder(resp.Body)
	var weather *WeatherDescription
	err := decoder.Decode(&weather)
	if err != nil {
		return nil, err
	}
	return weather, nil
}

func getApiMethod() string {
	if yandexTariff == "weather_on_site" {
		return apiInformers
	} else {
		return apiForecast
	}
}

func buildURL(method string, query url.Values) string {
	query.Add("lang", apiLang)
	url := url.URL{
		Scheme:   "https",
		Host:     apiURL,
		Path:     method,
		RawQuery: query.Encode(),
	}

	return url.String()
}
