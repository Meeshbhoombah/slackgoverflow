package slack

import (
        "github.com/archproj/slackoverflow/config"
        s "github.com/nlopes/slack"
)

type Client struct {
        app     *s.Client
        bot     *s.Client
}

func CreateClient(c *config.Variables) (Client, error) {
        // TODO: errors - Check Auth Token, Check Bot Token
        var sc Client
        sc.app = s.New(c.SlackAuthToken)
        sc.bot = s.New(c.SlackBotToken)
        return sc, nil
}
