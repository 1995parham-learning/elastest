package actions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler sends hello to you!
// nolint: wrapcheck
func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
