package actions

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.elastic.co/apm/module/apmecho"
	"gopkg.in/go-playground/validator.v9"
)

// App creates new instance of Echo and configures it
func App() *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())

	app.Use(apmecho.Middleware())

	// Validator
	app.Validator = &DefaultValidator{validator.New()}

	// Routes
	app.GET("/hello", HelloHandler)

	return app
}
