package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func UseMiddleware(e *echo.Echo) {
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "hello" && password == "world" {
			return true, nil
		}
		//return false, nil
		return true, nil
	}))
}

func Logger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logger Middleware")
}
