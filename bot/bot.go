package bot

import (
	"fmt"
	"github.com/Zoxan/bot/yandexweather"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Zoxan/bot/vkapi"
)

//Bot ..
type Bot struct {
	botAppeal  string
	helloCmd   string
	weatherCmd string
	weatherParams []WeatherParam
}

type WeatherParam struct {
	Caption   string
	Latitude  float64
	Longitude float64
}

//NewBot ..
func NewBot(weatherParams []WeatherParam) *Bot {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Bot{
		botAppeal:  "бот",
		helloCmd:   "привет",
		weatherCmd: "погода",
		weatherParams: weatherParams,
	}
}

//SendText ..
func (bot *Bot) SendText(text string, fromID int, peerID int) {
	lowerText := strings.ToLower(text)
	if !strings.HasPrefix(lowerText, bot.botAppeal) {
		return
	}

	if strings.HasSuffix(lowerText, bot.helloCmd) {
		bot.execHello(fromID, peerID)
	}

	if strings.HasSuffix(lowerText, bot.weatherCmd) {
		bot.execWeather(peerID)
	}
}

func (bot *Bot) execHello(fromID int, peerID int) {
	user, err := vkapi.RequestUser(fromID)
	if err != nil {
		log.Print("ERROR request user:", err)
		return
	}
	message := fmt.Sprint("Привет, ", user.FirstName)

	err = vkapi.RequestSendToGroup(peerID, getRandromID(), message)
	if err != nil {
		log.Print("ERROR request send to group:", err)
	}
}

func (bot *Bot) execWeather(peerID int) {
	message := ""
	for _, param := range bot.weatherParams {
		res, _ := yandexweather.GetWeatherText(param.Latitude, param.Longitude)
		message += fmt.Sprintf("%s: %s\n", param.Caption, res)
	}

	err := vkapi.RequestSendToGroup(peerID, getRandromID(), message)
	if err != nil {
		log.Print("ERROR request send to group:", err)
	}
}

func getRandromID() int {
	return rand.Int()
}
