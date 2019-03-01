package routes

import (
	"github.com/labstack/echo"

	"github.com/archproj/slackoverflow/routes/auth"
	"github.com/archproj/slackoverflow/routes/client"
	"github.com/archproj/slackoverflow/routes/listen"
)

func Bind(e *echo.Echo) {
	// render static files
	client.Routes(e)

	// Slack OAuth 2.0 to integrate app
	auth.Routes(e)

	// handle slash command
	l := e.Group("/listen")
	listen.Routes(l)
}
