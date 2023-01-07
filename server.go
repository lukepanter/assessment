package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lukepanter/assessment/expense"
)

func main() {
	e := echo.New()
	h := expense.ExpenseHandler{}
	h.Initialize()

	e.POST("/expenses",h.CreateExpenseHandler)
	e.Logger.Fatal(e.Start(":2565"))
}