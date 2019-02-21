package middlewares

import (
        "log"

        "github.com/labstack/echo"
)

// Sets variables in Context, maintains order for retrival
func EmbedInContext(items ...interface{}) echo.MiddlewareFunc {
        return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
                    for pos, val := range items {
                        key := string(pos)
                        log.Println(key)
                        c.Set(key, val)
                    }

                    next(c)
                    return nil
		}
	}
}
