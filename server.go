package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/lukepanter/assessment/expense"
)

func main() {
	expense.InitDB()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/expenses",expense.CreateExpenseHandler)
	e.Logger.Fatal(e.Start(":2565"))
}