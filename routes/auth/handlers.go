package auth

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/archproj/slackoverflow/config"
)

func Authorize(c echo.Context) error {
	cfg := c.Get("0").(*config.Variables)

	url, err := GenerateURL(cfg)
	if err != nil {
		log.Error(err)
		return err
	}

	err = c.Redirect(http.StatusSeeOther, url)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func Integrate(c echo.Context) error {
	return nil
}
