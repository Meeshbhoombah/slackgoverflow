package client

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.Static("/", "static/public/index.html")
}
