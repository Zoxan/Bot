package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Zoxan/bot/callbackserver"
	"github.com/Zoxan/bot/vkapi"
)

type config struct {
	AccessToken       string `toml:"access_token"`
	ConfirmationToken string `toml:"confirmation_token"`
	Port              string `toml:"port"`
}

func main() {
	conf := &config{}
	_, err := toml.DecodeFile("main/config.toml", conf)

	if err != nil {
		log.Fatal(err)
	}

	vkapi.Start(conf.AccessToken)
	callbackserver.Start(conf.ConfirmationToken, conf.Port)
}
