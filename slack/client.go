package slack

import (
	"github.com/jinzhu/gorm"
	s "github.com/nlopes/slack"

	"github.com/archproj/slackoverflow/config"
)

// Custom Client to encompass both the
type Client struct {
	*s.Client

	Ver string

	// Make School Product College Workspace Team ID
	TeamId string
	// `#slackoverflow` Channel ID
	ChanId string
}

func Init(cfg *config.Variables, db *gorm.DB, code *string) (*Client, error) {
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
		Ver: cfg.SlackVerToken,
		Usr: s.New(cfg.SlackUsrToken),
		Bot: s.New(cfg.SlackBotToken),
	}

	return &sc, nil
}

func attachSlackoverflow(sc *Client) error {
	channels, err := sc.Usr.GetChannels(false)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		if channel.Name == "slackover" {
			sc.ChanId = channel.ID
		}
	}

	// TODO: error if channel not found
	return nil
}
