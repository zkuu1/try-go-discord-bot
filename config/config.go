package config

import (
	"endcoding/json"
	"fmt"
	"io/ioutil"

)

var (
	Token string
	BotPrefix string

	config *config.Config 
)

type Config struct struct{
	Token string `json:"Token"`
	BotPrefix	string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.PrintIn("Reading Config File...")

	ioutil.ReadFile("config.json")
}