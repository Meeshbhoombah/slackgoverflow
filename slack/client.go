package slack

import (
	s "github.com/nlopes/slack"

	"github.com/archproj/slackoverflow/config"
)

// Encompasses both the App client and Bot client, as certain endpoints require
// one and not the other
type Client struct {
	// (deprecated) verification token
	VerToken string

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
		VerToken: cfg.SlackVerToken,
		App:      s.New(cfg.SlackAuthToken),
		Bot:      s.New(cfg.SlackBotToken),
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
	}

	// TODO: error if channel not found
	return nil
}
