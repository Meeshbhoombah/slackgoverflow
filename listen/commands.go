package listen

import (
        str "strings"
        "errors"

	log "github.com/sirupsen/logrus"
        s "github.com/nlopes/slack"
	"github.com/labstack/echo"

        "github.com/archproj/slackoverflow/slack"
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

        sc := c.Get("2").(*slack.Client)

        log.Info(r.Command)

        switch r.Command {
        case "/ask":
                if r.Text[:4] == "anon" {
                        r.UserName = "anon"
                        r.Text = r.Text[5:]
                }

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
