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
	ErrCouldNotVerify = errors.New("Could not verify token.")
)

func CommandHandler(c echo.Context) error {
	r, err := s.SlashCommandParse(c.Request())
	if err != nil {
		log.Error(err)
		return err
	}

	cfg := c.Get("0").(*config.Variables)
	if !r.ValidateToken(cfg.SlackVerToken) {
		log.Println(ErrCouldNotVerify)
		return err
	}

	sc := c.Get("1").(*slack.Client)

	log.Info(r.Command)

	switch r.Command {
	case "/ask":
		if str.Contains(r.Text, "?") {
			log.Println(r.Text)
			sc.Ask(r.Text, r.UserName)
		} else {
			txt := `Please rephrase as a question. E.g: What is love?`
			sc.NotifyUser(txt, r.UserID)
		}
	}

	return nil
}
