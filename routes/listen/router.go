package listen

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Group) {
	e.POST("/command", CommandHandler)
	e.POST("/event", EventHandler)
}
