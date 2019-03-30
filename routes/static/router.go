package static

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.Static("/", "static/build/")
	e.Static("/static", "static/build/static/")
	// TODO: Fix path on client side
	e.Static("/Assets", "static/build/assets/Assets/")
}
