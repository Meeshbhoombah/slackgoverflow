package slack

import (
	s "github.com/nlopes/slack"

	"github.com/archproj/slackoverflow/config"
)

type Client struct {
	App *s.Client
	Bot *s.Client

	// #slackoverflow ChannelId
	ChannelId string
}

func Init(cfg *config.Variables) (*Client, error) {
	sc, err := newClient(cfg)
	if err != nil {
		return nil, err
	}

	err = attachSlackoverflow(sc)
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func newClient(cfg *config.Variables) (*Client, error) {
	// TODO: errors - Check Auth Token, Check Bot Token
	sc := Client{
		App: s.New(cfg.SlackAuthToken),
		Bot: s.New(cfg.SlackBotToken),
	}

	return &sc, nil
}

func attachSlackoverflow(sc *Client) error {
	channels, err := sc.App.GetChannels(false)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		if channel.Name == "devp2p" {
			sc.ChannelId = channel.ID
		}
		// TODO: Add error if channel not found
	}

	return nil
}
