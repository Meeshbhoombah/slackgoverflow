package auth

import (
	"fmt"
	//"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

func OAuthHandler(c echo.Context) error {
	url := fmt.Sprintf(`http://www.slack.com/oauth/authorize`)

	err := c.Redirect(301, url)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
