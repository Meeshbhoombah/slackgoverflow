package slack

import (
	"errors"

	"github.com/archproj/slackoverflow/config"
	s "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

type Client struct {
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
		VerToken: cfg.SlackCredentials.SlackVerToken,
		App:      s.New(cfg.SlackCredentials.SlackAuthToken),
		Bot:      s.New(cfg.SlackCredentials.SlackBotToken),
	}

	return &sc, nil
}

func attachSlackoverflow(sc *Client) error {
	channels, err := sc.App.GetChannels(false)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		if channel.Name == "slackover" {
			sc.ChannelId = channel.ID
			log.Info("Channel ID: ", sc.ChannelId)
			return nil
		}

	}

	// Return error if channel is not found
	return errors.New("Channel not found")
}
