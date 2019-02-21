package listen

import (
        "bytes"
        "net/http"
        "encoding/json"

        log "github.com/sirupsen/logrus"
	"github.com/labstack/echo"
        s "github.com/nlopes/slack/slackevents"
)

func EventHandler(c echo.Context) error {
        rb := c.Request().Body

        buf := new(bytes.Buffer)
        buf.ReadFrom(rb)
        b := buf.String()

        evt, err := s.ParseEvent(b)
        if err != nil {
                log.Error(err)
        }

        log.Println(evt.Type)

        switch evt.Type {
        case s.URLVerification:
                ver := new(s.EventsAPIURLVerificationEvent)
                if err = c.Bind(ver); err != nil {
                        log.Println(ver)
                        log.Println(ver.Token)
                }
        }

        return nil
}
