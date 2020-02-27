package yandexweather

import (
	"fmt"
	"math"
)

var conditions = map[string]string {
	"clear": "ясно",
	"partly-cloudy": "малооблачно",
	"cloudy": "облачно с прояснениями",
	"overcast": "пасмурно",
	"partly-cloudy-and-light-rain": "небольшой дождь",
	"partly-cloudy-and-rain": "дождь",
	"overcast-and-rain": "сильный дождь",
	"overcast-thunderstorms-with-rain": "сильный дождь, гроза",
	"cloudy-and-light-rain": "небольшой дождь",
	"overcast-and-light-rain": "небольшой дождь",
	"cloudy-and-rain": "дождь",
	"overcast-and-wet-snow": "дождь со снегом",
	"partly-cloudy-and-light-snow": "небольшой снег",
	"partly-cloudy-and-snow": "снег",
	"overcast-and-snow": "снегопад",
	"cloudy-and-light-snow": "небольшой снег",
	"overcast-and-light-snow": "небольшой снег",
	"cloudy-and-snow": "снег",
}

var windDirs = map[string]string {
	"nw": "СЗ",
	"n": "С",
	"ne": "СВ",
	"e": "В",
	"se": "ЮВ",
	"s": "Ю",
	"sw": "ЮЗ",
	"w": "З",
	"c": "штиль",
}

//GetText returns simple text presentation for weather
func GetText(description WeatherDescription) string {
	condition := conditions[description.Fact.Condition]
	temp := getTemp(description.Fact.Temperature)
	wind := getWind(description.Fact.WindSpeed, description.Fact.WindDirection)
	return fmt.Sprintf("%s, %s, %s", condition, temp, wind)
}

func getTemp(temp float64) string {
	tempRound := int(math.Round(temp))
	return fmt.Sprintf("%d °C", tempRound)
}

func getWind(speed float64, dir string) string {
	speedRound := int(math.Round(speed))
	if dir == "c" || speedRound == 0 {
		return windDirs["c"]
	}
	windDir := windDirs[dir]
	return fmt.Sprintf("%d м/с %s", speedRound, windDir)
}




