package config

import (
	template2 "github.com/bhawneshdipu/go-angular-sample-app/pkg/web/template"
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Echo          *echo.Echo
	TemplateUI    *template2.TemplateUI
}
