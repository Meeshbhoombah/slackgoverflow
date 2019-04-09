package listen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	s "github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	log "github.com/sirupsen/logrus"
)

func EventHandler(c echo.Context) error {
	r := c.Request()
	buf.ReadFrom(r.Body)
	body := buf.String()
	b := json.RawMessage(body)

	evt, err := slackevents.ParseEvent(b)
	if err != nil {
		log.Error(err)
		return err
	}

	switch evt.Type {
	case slackevents.URLVerification:
		var r *slackevents.ChallengeResponse

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
		return nil
	}
}
