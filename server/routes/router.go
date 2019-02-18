package routes

import (
        "log"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/slack"
)

func Serve(cfg *config.Variables, e *echo.Echo, db *gorm.DB, sc *slack.Client) {
        log.Println(cfg)
        log.Println(db)
        log.Println(sc)
        log.Println(e)
}
