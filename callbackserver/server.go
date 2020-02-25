package callbackserver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zoxan/bot/bot"
)

type vkEventData struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
}

type vkMessageEventData struct {
	Message message `json:"message"`
}

type message struct {
	Message struct {
		ID     int    `json:"id"`
		Date   int    `json:"date"`
		FromID int    `json:"from_id"`
		Text   string `json:"text"`
		PeerID int    `json:"peer_id"`
	} `json:"message"`
}

var (
	confirmationToken string
	botInstance       *bot.Bot
)

//Start ..
func Start(confToken string, port string) {
	confirmationToken = confToken
	botInstance = bot.NewBot()
	listenAndHandle(port)
}

func listenAndHandle(port string) {
	http.Handle("/", http.HandlerFunc(handleRoot))
	http.ListenAndServe(port, nil)
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var eventData vkEventData

	err := decoder.Decode(&eventData)
	if err != nil {
		log.Print("ERROR:", err)
		return
	}
	switch eventData.Type {
	case "confirmation":
		handleConfirmation(w, req, eventData)

	case "message_new":
		handleMessageNew(w, req, eventData)
	}
}

func handleConfirmation(w http.ResponseWriter, req *http.Request, eventData vkEventData) {
	log.Print("handleConfirmation")
	sendString(w, req, confirmationToken)
}

func handleMessageNew(w http.ResponseWriter, req *http.Request, eventData vkEventData) {
	log.Print("handleMessageNew")
	sendString(w, req, "ok")

	reader := bytes.NewReader(eventData.Object)
	decoder := json.NewDecoder(reader)
	var msg message
	err := decoder.Decode(&msg)
	if err != nil {
		log.Print("ERROR:", err)
		return
	}

	txt := msg.Message.Text
	botInstance.SendText(txt, msg.Message.FromID, msg.Message.PeerID)
}

func sendString(w http.ResponseWriter, eq *http.Request, data string) {
	w.Write([]byte(data))
}
