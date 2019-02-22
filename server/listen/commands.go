package listen

import (
        str "strings"
        "errors"

	log "github.com/sirupsen/logrus"
        s "github.com/nlopes/slack"
	"github.com/labstack/echo"

        "github.com/archproj/slackoverflow/config"
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
                err := ErrCouldNotVerify
                log.Println(err)
                return err
        }

        switch r.Command {
        case "/ask":
                if r.Text[:4] == "anon" {
                        r.UserName = "anon"
                        r.Text = r.Text[5:]
                }

                if str.Contains(r.Text, "?") {
                        s.Ask(r.Text, r.UserName)
                } else {
                        // tell user they cannot post message
                        s.NotifyUser("Please rephrase as a question.")
                }
        }

        return nil
}
