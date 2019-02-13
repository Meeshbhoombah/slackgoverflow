package main

import (
	"log"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
)

const (
        VERISON = 0.1.0
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Panic(err)
	}

        sc, err := slack.NewClient(&cfg)
	if err != nil {
		log.Panic(err)
	}

        err = router.Init(&cfg, &sc)
        if err != nil {
                log.Panic(err)
        }
}
