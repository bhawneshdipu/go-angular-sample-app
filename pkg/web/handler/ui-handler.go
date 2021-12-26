package handler

import (
	config "github.com/bhawneshdipu/go-angular-sample-app/pkg/conf"
	"github.com/labstack/echo/v4"
	"net/http"
)

var AppConfig *config.AppConfig

func Init(config config.AppConfig) {
	AppConfig = &config
}

//Index e.GET("/", Index)
func UIIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}
