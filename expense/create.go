package expense

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"fmt"
)


func (h *ExpenseHandler) CreateExpenseHandler(c echo.Context) error {
	expense := Expense{}
	err := c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := h.DB.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", 
	expense.Title, expense.Amount, expense.Note, pq.Array(&expense.Tags))
	fmt.Println(row)
	err = row.Scan(&expense.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, expense)
}