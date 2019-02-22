package listen

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	s "github.com/nlopes/slack/slackevents"
	log "github.com/sirupsen/logrus"
)

func rawMsg(r func() *http.Request) (json.RawMessage, error) {
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

func handle(evt s.EventsAPIEvent) error {
	switch evt.Type {
	case s.URLVerification:
		log.Println(evt.Token)
	default:
		log.Println(evt)
	}

	return nil
}

func EventHandler(c echo.Context) error {
	rb, err := rawMsg(c.Request)
	if err != nil {
		log.Error("Unable to parse JSON from body: ", err)
		return err
	}

	evt, err := s.ParseEvent(rb,
            // verify token in body (sent from Slack)
            s.OptionsVerifyToken(rb,
                &s.TokenComparator {
                    VertificationToken: cfg.SlackVerToken
                }
            )
        )

	if err != nil {
		log.Error("Unable to parse Slack event: ", err)
		return err
	}

	log.Infoln(evt.Type)

	err = handle(evt)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
