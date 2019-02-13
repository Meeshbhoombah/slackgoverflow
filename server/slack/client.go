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
        sc := Client {
            app: s.New(c.SlackAuthToken),
            bot: s.New(c.SlackBotToken),
        }

        return sc, nil
}
