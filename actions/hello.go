package actions

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloHandler sends hello to you!
func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
