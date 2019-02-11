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

	SlackSecret    string `env:"SLACK_SIGNING_SECRET"`
	SlackClientId  string `env:"SLACK_CLIENT_ID"`
	SlackAuthToken string `env:"SLACK_AUTH_TOKEN"`
	SlackBotToken  string `env:"SLACK_BOT_TOKEN"`
}

func (v *Variables) Load() error {
	config := reflect.ValueOf(v).Elem()

        for lineNo := 0; lineNo < config.NumField(); lineNo++ {
		line := config.Type().Field(lineNo)
                val := config.FieldByName(line.Name)

		envVar := line.Tag.Get("env")
                envVal := os.Getenv(envVar)

                val.SetString(envVal)
	}

	return nil
}

