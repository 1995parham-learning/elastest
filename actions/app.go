package actions

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
	"gopkg.in/go-playground/validator.v9"
)

// App creates new instance of Echo and configures it.
func App() *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())

	app.Use(apmechov4.Middleware())

	// Register validator
	app.Validator = &DefaultValidator{validator.New()}

	// Register routes
	app.GET("/hello", HelloHandler)

	return app
}
