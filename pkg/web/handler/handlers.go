package handler

import (
	"encoding/json"
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

//Index e.GET("/", Index)
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

//GetUser e.GET("/users/:id", GetUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	idStr := c.Param("id")
	user := model.User{}
	id, _ := strconv.ParseInt(idStr, 10, 32)
	user = AppConfig.UserRepository.FindById(int(id))
	usr, _ := json.Marshal(user)
	return c.String(http.StatusOK, string(usr))
}

//ShowUser e.GET("/users/show", ShowUser)
func ShowUser(c echo.Context) error {
	users := AppConfig.UserRepository.FindAll()
	usr, _ := json.Marshal(users)
	return c.String(http.StatusOK, string(usr))
}

//SaveUser e.POST("/users/save", save)
func SaveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	password := c.FormValue("password")
	username := c.FormValue("username")
	user := model.User{
		Name:      name,
		UserName:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	AppConfig.UserRepository.Create(&user)
	usr, _ := json.Marshal(user)
	return c.String(http.StatusOK, string(usr))
}
