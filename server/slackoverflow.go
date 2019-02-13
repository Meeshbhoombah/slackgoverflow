package main

import (
	"log"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
)

func main() {
        var cfg config.Variables

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

        sc, err := slack.CreateClient(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
