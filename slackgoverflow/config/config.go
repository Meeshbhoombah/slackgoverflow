
package config

import (
        "os"
        "reflect"
)

type Config struct {
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

type Env map[string]string

func (c *Config) LoadFrom(filename string) error {
        f, err := os.Open(filename)
        if err != nil {
            return err
        }
        defer f.Close()

        return nil
}

