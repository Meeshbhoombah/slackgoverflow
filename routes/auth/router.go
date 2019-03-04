package auth

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/authorize", Authorize)
	e.GET("/integrate", Integrate)
}
