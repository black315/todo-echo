package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "../model"
)

func List() echo.HandlerFunc{
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, model.GetTodos())
    }
}
