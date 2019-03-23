package auth

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
)

func Authorize(c echo.Context) error {
	cfg := c.Get("0").(*config.Variables)

	url, err := GenerateOAuthURL(cfg)
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
	req := c.Request()

	code, err := ParseOAuthVerCode(req)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("INTEGRATING WORKSPACE WITH CODE: ", *code)

	cfg := c.Get("0").(*config.Variables)
	db := c.Get("1").(*gorm.DB)

	sc, err := slack.Init(cfg, db, code)
	if err != nil {
		log.Error(err)
		return err
	}

	url, err := sc.GenerateDeepLink()
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
