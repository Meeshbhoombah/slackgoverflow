package router

import (
        "log"

        "github.com/archproj/slackoverflow/config"
        "github.com/archproj/slackoverflow/slack"
)

func Init(cfg *config.Variables, sc *slack.Client) {
        log.Println(cfg)
        log.Println(sc)
}
