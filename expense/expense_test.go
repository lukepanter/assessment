package expense

import (
 "testing"
 "github.com/DATA-DOG/go-sqlmock"
 "github.com/stretchr/testify/assert"
 "github.com/labstack/echo/v4"
 "net/http"
	"net/http/httptest"
	"strings"
	"github.com/lib/pq"
)

func TestCreateExpenseHandler(t *testing.T) {
	var expenseJSON = `{"title":"Car","amount":27,"note":"beautiful car","tags":["car"]}`
	var resultJSON = `{"id":0,"title":"Car","amount":27,"note":"beautiful car","tags":["car"]}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	
	db, mock, _ := sqlmock.New()
	h := &ExpenseHandler{db}
	
   
	rows := sqlmock.NewRows([]string{"id"}).AddRow(0)
	
	mock.ExpectQuery("INSERT INTO expenses").WithArgs("Car", 27, "beautiful car",pq.Array([]string{"car"})).WillReturnRows(rows)
	if assert.NoError(t, h.CreateExpenseHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, resultJSON, rec.Body.String()[:len(rec.Body.String())-1])
	}
   }