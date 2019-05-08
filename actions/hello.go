package actions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler sends hello to you!
func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
