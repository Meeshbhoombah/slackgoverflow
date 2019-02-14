package slack

import (
	s "github.com/nlopes/slack"

	"github.com/archproj/slackoverflow/config"
)

type Client struct {
	App       *s.Client
	Bot       *s.Client

        // #slackoverflow ChannelId
        Chan      string
}

func Init(cfg *config.Variables) (*Client, error) {
	// TODO: errors - Check Auth Token, Check Bot Token
	sc := Client{
		App: s.New(cfg.SlackAuthToken),
		Bot: s.New(cfg.SlackBotToken),
	}

        err := AttachSlackoverflow(&sc)
        if err != nil {
                return nil, err
        }

	return &sc, nil
}
