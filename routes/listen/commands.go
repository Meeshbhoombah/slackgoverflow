package listen

import (
	"errors"
	str "strings"

	"github.com/labstack/echo"
	s "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
)

var (
	ErrUnverifiedRequest = errors.New("Request does not contain verified token.")
)

func CommandHandler(c echo.Context) error {
	r, err := s.SlashCommandParse(c.Request())
	if err != nil {
		log.Error(err)
		return err
	}

	cfg := c.Get("0").(*config.Variables)

	// authenticate request with using Verification TOken
	if !r.ValidateToken(cfg.SlackVerToken) {
		err := ErrUnverifiedRequest
		log.Error(err)
		return err
	}

	sc, err := slack.Init(cfg)
	if err != nil {
		log.Error(err)
	}

	switch r.Command {
	case "/ask":
		if str.Contains(r.Text, "?") {
			sc.Ask(r.Text, r.UserName)
		} else {
			txt := `Please rephrase as a question. E.g: What is love?`
			sc.NotifyUser(txt, r.UserID)
		}
	}

	return nil
}
