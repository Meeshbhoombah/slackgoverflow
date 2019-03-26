package config

import (
	"net"
	"net/url"
	"os"
	"reflect"
)

type Variables struct {
	Env string `env:"ENV"`

	Host string `env:"HOST"`
	Port string `env:"PORT"`

	SecretKey string `env:"SECRET_KEY"`

	TempChan string `env:"TEMP_CHAN"`

	// Heroku injects concatenated Postgres connection string as env var
	Dburl  string `env:"DATABASE_URL"`
	Dbuser string `env:"DBUSER"`
	Dbpass string `env:"DBPASS"`
	Dbhost string `env:"DBHOST"`
	Dbname string `env:"DBNAME"`
	Dbport string `env:"DBPORT"`

	SlackClientID     string `env:"SLACK_CLIENT_ID"`
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

	// Only parse connection string if in Production
	if cfg.Dburl != "" && cfg.Env != "DEVELOPMENT" {
		err := parseDburl(&cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func parseDburl(cfg *Variables) error {
	u, err := url.Parse(cfg.Dburl)
	if err != nil {
		return err
	}

	cfg.Dbuser = u.User.Username()
	cfg.Dbpass, _ = u.User.Password()

	h, p, _ := net.SplitHostPort(u.Host)

	cfg.Dbhost = h
	cfg.Dbport = p

	cfg.Dbname = u.Path

	return nil
}
