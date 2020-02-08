package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.File("/", "public/index.html")
    e.GET("/todos", handler.List())

    e.Start(":8080")
}
