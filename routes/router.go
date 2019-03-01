package routes

import (
	"github.com/labstack/echo"

	"github.com/archproj/slackoverflow/routes/listen"
)

func Bind(e *echo.Echo) {
	// landing page to `/`
	main.Routes(e)

	// auth @ `authorize` and `integrate`
	auth.Routes(e)

	// prefix `/listen`
	l := e.Group("/listen")
	listen.Routes(l)
}
