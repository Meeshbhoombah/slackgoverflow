package routes

import (
	"github.com/labstack/echo"

	"github.com/archproj/slackoverflow/routes/auth"
	"github.com/archproj/slackoverflow/routes/static"
)

func Bind(e *echo.Echo) {
	// static files
	static.Routes(e)

	// Slack OAuth2.0
	auth.Routes(e)

	// slash command
	l := e.Group("/listen")
	listen.Routes(l)
}
