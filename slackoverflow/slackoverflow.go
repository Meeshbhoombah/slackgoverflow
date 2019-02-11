package main

import (
	"github.com/archproj/slackoverflow/config"
	"log"
)

type Variables struct {
	SecretKey string `env:"SECRET_KEY"`

	Dbuser string `env:"DBUSER"`
	Dbpass string `env:"DBPASS"`
	Dbhost string `env:"DBHOST"`
	Dbname string `env:"DBNAME"`

	SlackSecret    string `env:"SLACK_SIGNING_SECRET"`
	SlackClientId  string `env:"SLACK_CLIENT_ID"`
	SlackAuthToken string `env:"SLACK_AUTH_TOKEN"`
	SlackBotToken  string `env:"SLACK_BOT_TOKEN"`
}

func main() {
	var c Variables
	err := config.Load(&c); if err != nil {
		log.Fatal(err)
	}
}
