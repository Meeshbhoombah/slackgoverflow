package main

import (
	"log"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/database"
	"github.com/archproj/slackoverflow/slack"
)

const (
	VERSION = "0.1.0"
)

func main() {
	cfg, err := config.Load() // from environment
	if err != nil {
		log.Panic(err)
	}

	db, err := database.Init(cfg)
	if err != nil {
		log.Panic(err)
	}

        log.Println(db)

	sc, err := slack.Init(cfg)
	if err != nil {
		log.Panic(err)
	}

        log.Println(sc.Chan)

        /*
        err = web.Serve(cfg, db, sc)
        if err != nil {
                log.Fatal(err)
        }
	*/
}
