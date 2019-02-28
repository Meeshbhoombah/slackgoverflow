package routes

import (
	"github.com/labstack/echo"

	"github.com/archproj/slackoverflow/routes/listen"
)

func Bind(e *echo.Echo) {
	main.Routes(e)

	// Prefix all listen routes with
	l := e.Group("/listen")
	listen.Routes(l)
}
