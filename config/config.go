package config

import (
	"os"
	"reflect"
)

type Variables struct {
	Host      string `env:"HOST"`
	Port      string `env:"PORT"`
	SecretKey string `env:"SECRET_KEY"`

	Dbuser string `env:"DBUSER"`
	Dbpass string `env:"DBPASS"`
	Dbhost string `env:"DBHOST"`
	Dbname string `env:"DBNAME"`
	Dbport string `env:"DBPORT"`
	Dburl  string `env:DATABASE_URL`

	SlackClientId     string `env:"SLACK_CLIENT_ID"`
	SlackClientSecret string `env:"SLACK_CLIENT_SECRET"`
	SlackRedirectURI  string `env:"SLACK_REDIRECT_URI"`
	SlackSecret       string `env:"SLACK_SIGNING_SECRET"`
	SlackUsrToken     string `env:"SLACK_USR_TOKEN"`
	SlackBotToken     string `env:"SLACK_BOT_TOKEN"`
	SlackVerToken     string `env:"SLACK_VER_TOKEN"`
}

func Load() (*Variables, error) {
	// TODO: errors, handle filepath, breakup func
	var cfg Variables
	config := reflect.ValueOf(&cfg).Elem()

	for lineNo := 0; lineNo < config.NumField(); lineNo++ {
		// TODO: handle nested config
		field := config.Type().Field(lineNo)
		val := config.FieldByName(field.Name)

		envKey := field.Tag.Get("env")
		envVal := os.Getenv(envKey)

		val.SetString(envVal)
	}

	return &cfg, nil
}
