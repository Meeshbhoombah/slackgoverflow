package config

import (
	"os"
        "fmt"
	"reflect"
)

type Variables struct {
	SecretKey string `env:"SECRET_KEY"`

	Dbuser string `env:"DBUSER"`
	Dbpass string `env:"DBPASS"`
	Dbhost string `env:"DBHOST"`
	Dbname string `env:"DBNAME"`

	SlackSecret    string `env:"SLACK_SECRET"`
	SlackClientId  string `env:"SLACK_CLIENT_ID"`
	SlackAuthToken string `env:"SLACK_AUTH_TOKEN"`
	SlackBotToken  string `env:"SLACK_BOT_TOKEN"`
}


func (c *Variables) Load(s interface{}) error {
        val := reflect.ValueOf(s).Elem()
        fmt.Println(val)
}

