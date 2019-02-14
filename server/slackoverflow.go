package main

import (
	"os"
        "log"
        "time"
        "context"
        "os/signal"

        "github.com/labstack/echo"

	"github.com/archproj/slackoverflow/slack"
	"github.com/archproj/slackoverflow/config"
        "github.com/archproj/slackoverflow/database"
)

const (
        VERSION = "0.1.0"
)

func main() {
	cfg, err := config.Load() // from environment
	if err != nil {
		log.Panic(err)
	}

        db, err := database.Init(&cfg)
        if err != nil {
                log.Panic(err)
        }

        sc, err := slack.Init(&cfg, &db)
	if err != nil {
		log.Panic(err)
	}

        err = web.Serve(&cfg, &db, &sc)
        if err != nil {
                log.Fatal(err)
        }
}
