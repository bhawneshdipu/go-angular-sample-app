package main

import (
	config "github.com/bhawneshdipu/go-angular-sample-app/pkg/conf"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/model"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/repository"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/web/middleware"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/web/routes"
	template2 "github.com/bhawneshdipu/go-angular-sample-app/pkg/web/template"
	"github.com/labstack/gommon/log"
	"html/template"
	"time"
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
	AppConfig.Echo = routes.Routes(&AppConfig)
	InitTemplate(&AppConfig)
	middleware.UseMiddleware(AppConfig.Echo)
	userRepository := repository.UserRepository{}
	userRepository.Init()
	AppConfig.UserRepository = userRepository
	userRepository.Create(&model.User{
		Name:      "hello",
		IsActive:  true,
		Password:  "world",
		UserName:  "hello",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
	log.Info(userRepository.FindById(0))
	log.Info(userRepository.FindById(1))
	log.Info(userRepository.FindAll())
	AppConfig.Echo.Logger.Fatal(AppConfig.Echo.Start(":8080"))
}
