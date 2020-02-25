package bot

import (
	"fmt"
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
}

//NewBot ..
func NewBot() *Bot {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Bot{
		botAppeal:  "бот,",
		helloCmd:   "привет",
		weatherCmd: "погода",
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
	err := vkapi.RequestSendToGroup(peerID, getRandromID(), "погода отличная!")
	if err != nil {
		log.Print("ERROR request send to group:", err)
	}
}

func getRandromID() int {
	return rand.Int()
}
