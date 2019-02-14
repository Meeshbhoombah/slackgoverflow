package main

import (
	"log"

	"github.com/archproj/slackoverflow/config"
        "github.com/archproj/slackoverflow/routes"
)

const (
        VERSION = "0.1.0"
)

func main() {
	cfg, err := config.Load() //from environment
	if err != nil {
		log.Panic(err)
	}

        routes.Serve(&cfg)
}
