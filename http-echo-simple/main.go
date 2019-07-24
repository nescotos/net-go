package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//Creating Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hey, from Echo!")
	})
	//Using JWT Middleware
	config := middleware.JWTConfig{
		Claims:     &jwtLoginClaims{},
		SigningKey: []byte("my-secret"),
	}
	e.POST("/user", createUser)
	e.POST("/login", login)
	//Restricted Route
	e.GET("/protected", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtLoginClaims)
		username := claims.Username
		return c.String(http.StatusOK, "Hey, welcome to the application "+username+"!")
	}, middleware.JWTWithConfig(config))
	//Public Route
	e.GET("/public", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a public route!")
	})
	e.Logger.Fatal(e.Start(":3453"))
}
