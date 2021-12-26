package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//Index e.GET("/", Index)
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

//GetUser e.GET("/users/:id", GetUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//ShowUser e.GET("/users/show", ShowUser)
func ShowUser(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

//SaveUser e.POST("/users/save", save)
func SaveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
