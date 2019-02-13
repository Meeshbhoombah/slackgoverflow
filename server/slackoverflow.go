package main

import (
	"log"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
        "github.com/archproj/slackoverflow/router"
)

const (
        VERSION = 0.1
)

func main() {
	cfg, err := config.Load() //load variables from environment
	if err != nil {
		log.Panic(err)
	}

        sc, err := slack.Init(&cfg)
	if err != nil {
		log.Panic(err)
	}

        router.Init(&cfg, &sc)
}
