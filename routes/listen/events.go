package listen

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	s "github.com/nlopes/slack/slackevents"
	log "github.com/sirupsen/logrus"

	"github.com/archproj/slackoverflow/config"
)

func EventHandler(c echo.Context) error {
	r := c.Request()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	b := json.RawMessage(body)

	cfg := c.Get("0").(*config.Variables)

	evt, err := s.ParseEvent(b, s.OptionVerifyToken(&s.TokenComparator{VerificationToken: cfg.SlackVerToken}))
	if err != nil {
		log.Error(err)
		return err
	}

	switch evt.Type {
	case s.URLVerification:
		var r *s.ChallengeResponse

		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Error(err)
			return err
		}

		err = c.String(http.StatusOK, r.Challenge)
		if err != nil {
			log.Error(err)
			return err
		}
	default:
		log.Info(evt.Type)
	}

	return nil
}
