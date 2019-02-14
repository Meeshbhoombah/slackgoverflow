package config

import (
	"os"
	"reflect"
)

type Variables struct {
	SecretKey string `env:"SECRET_KEY"`

	Dbuser string `env:"DBUSER"`
	Dbpass string `env:"DBPASS"`
	Dbhost string `env:"DBHOST"`
	Dbname string `env:"DBNAME"`
	Dbport string `env:"DBPORT"`

	SlackSecret       string `env:"SLACK_SIGNING_SECRET"`
	SlackClientId     string `env:"SLACK_CLIENT_ID"`
	SlackClientSecret string `env:"SLACK_CLIENT_SECRET"`
	SlackAuthToken    string `env:"SLACK_AUTH_TOKEN"`
	SlackBotToken     string `env:"SLACK_BOT_TOKEN"`
}

func Load() (*Variables, error) {
	// TODO: errors, handle filepath
	var cfg Variables
	config := reflect.ValueOf(&cfg).Elem()

	for lineNo := 0; lineNo < config.NumField(); lineNo++ {
		field := config.Type().Field(lineNo)
		val := config.FieldByName(field.Name)

		envVar := field.Tag.Get("env")
		envVal := os.Getenv(envVar)

		val.SetString(envVal)
	}

	return &cfg, nil
}
