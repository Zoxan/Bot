package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Zoxan/bot/bot"
	"github.com/Zoxan/bot/callbackserver"
	"github.com/Zoxan/bot/vkapi"
	"github.com/Zoxan/bot/yandexweather"
	"log"
)

type config struct {
	VkAccessToken       string `toml:"vk_access_token"`
	VkConfirmationToken string `toml:"vk_confirmation_token"`
	YandexWeatherToken  string `toml:"yandex_weather_token"`
	YandexWeatherTariff  string `toml:"yandex_tariff"`
	Port                string `toml:"port"`
	WeatherTargets      []struct {
		Caption   string `toml:"caption"`
		Latitude  float64 `toml:"latitude"`
		Longitude float64 `toml:"longitude"`
	} `toml:"bot_weather_targets"`
}

func main() {
	conf := &config{}
	_, err := toml.DecodeFile("main/config.toml", conf)

	if err != nil {
		log.Fatal(err)
	}
	weatherParams := fetchBotWeatherParams(*conf)
	bot := bot.NewBot(weatherParams)
	yandexweather.Init(conf.YandexWeatherToken, conf.YandexWeatherTariff)
	vkapi.Init(conf.VkAccessToken)
	callbackserver.Start(conf.VkConfirmationToken, conf.Port, bot)
}

func fetchBotWeatherParams(conf config)  (params []bot.WeatherParam) {
	for _, confParam := range conf.WeatherTargets {
		params = append(params, bot.WeatherParam{
			Caption:   confParam.Caption,
			Latitude:  confParam.Latitude,
			Longitude: confParam.Longitude,
		})
	}
	return
}
