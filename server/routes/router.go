package routes

import (
        "log"

        "github.com/labstack/echo"
        "github.com/jinzhu/gorm"

        "github.com/archproj/slackoverflow/config"
        "github.com/archproj/slackoverflow/slack"
)

func Init(cfg *config.Variables, e *echo.Echo, db *gorm.DB, sc *slack.Client) {
}
