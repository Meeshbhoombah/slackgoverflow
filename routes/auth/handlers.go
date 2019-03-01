package auth

import (
	"fmt"
	//"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

func AuthHandler(c echo.Context) error {
	url := fmt.Sprintf(`slack.com/oauth/authorize`)

	err := c.Redirect(301, url)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
