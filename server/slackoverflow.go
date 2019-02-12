package main

import (
	"log"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
)

var (
	Config config.Environment
	Slack  slack.Client
)

func init() {
	err := Config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = Slack.Init()
        if err != nil {
                log.Fatal(err)
        }
}

func main() {
}
