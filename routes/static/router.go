package static

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.Static("/", "static/build/index.html")
}
