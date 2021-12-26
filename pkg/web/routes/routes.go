package routes

import (
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/web/handler"
	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()
	e.GET("/", handler.Index)
	e.GET("/users/:id", handler.GetUser)
	e.POST("/users/save", handler.SaveUser)
	e.GET("/users/show", handler.ShowUser)

	//UI
	e.Static("/ui", "./public/ui/dist/ui/")

	return e
}
