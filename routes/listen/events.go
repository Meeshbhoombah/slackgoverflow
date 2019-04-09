package listen

// Inconsistent API
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

func EventHandler(c echo.Context) error {
	eventsAPIEvent, e := slackevents.ParseEvent(json)
}
