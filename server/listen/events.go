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

func RawMsg(r func() *http.Request) (json.RawMessage, error) {
	b := r().Body
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(b)
	if err != nil {
		return nil, err
	}

	bs := buf.String()

	rb := json.RawMessage(bs)
	return rb, nil
}

func EventHandler(c echo.Context) error {
	rb, err := RawMsg(c.Request)
        if err != nil {
		log.Error("Unable to parse JSON from body: ", err)
		return err
	}

	cfg := c.Get("0").(*config.Variables)

	opts := s.OptionVerifyToken(&s.TokenComparator{VerificationToken: cfg.SlackVerToken})

	evt, err := s.ParseEvent(rb, opts)
	if err != nil {
		log.Error("Unable to parse Slack event: ", err)
		return err
	}

	log.Infoln(evt.Type)

	switch evt.Type {
	case s.URLVerification:
		log.Println(evt.Token)
                c.Response().Header().Set("Context-Type", "Text")
	default:
	}

	return nil
}
