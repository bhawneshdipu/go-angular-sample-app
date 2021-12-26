package main

import (
	config "github.com/bhawneshdipu/go-angular-sample-app/pkg/conf"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/web/middleware"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/web/routes"
	template2 "github.com/bhawneshdipu/go-angular-sample-app/pkg/web/template"
	"html/template"
)

func InitTemplate(config *config.AppConfig) {
	t := &template2.TemplateUI{
		Templates: template.Must(template.ParseGlob("./public/ui/dist/ui/*.html")),
	}
	config.TemplateUI = t
	config.Echo.Renderer = t
}

func main() {
	AppConfig := config.AppConfig{}
	AppConfig.Echo = routes.Routes()
	InitTemplate(&AppConfig)
	middleware.UseMiddleware(AppConfig.Echo)
	AppConfig.Echo.Logger.Fatal(AppConfig.Echo.Start(":8080"))
}
