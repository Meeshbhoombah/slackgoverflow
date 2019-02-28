package middlewares

import (
        "strconv"
	"github.com/labstack/echo"
)

// Sets passed variables in Context, maintains order for retrival
func EmbedInContext(items ...interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for pos, val := range items {
				key := strconv.Itoa(pos)
				c.Set(key, val)
			}

			next(c)
			return nil
		}
	}
}
