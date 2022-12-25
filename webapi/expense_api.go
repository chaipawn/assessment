package webapi

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/chaipawn/assessment/application/expense"
	"github.com/chaipawn/assessment/infrastructure"
	"github.com/labstack/echo/v4"
)

type ExpenseAPI struct {
	app     *echo.Echo
	Address string
}

func (api ExpenseAPI) Start() error {
	return api.app.Start(api.Address)
}

func (api ExpenseAPI) Shutdown(ctx context.Context) error {
	return api.app.Shutdown(ctx)
}

func NewExpenseAPI(address string, db *sql.DB) ExpenseAPI {
	e := echo.New()

	e.POST("/expenses", func(c echo.Context) error {
		var request CreateExpenseRequest
		err := c.Bind(&request)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		command := expense.NewAddExpenseCommand(request.Title, request.Amount, request.Note, request.Tags)
		repository := infrastructure.NewExpenseCommandRepository(db)
		handler := expense.NewAddExpenseHandler(repository)

		expense, err := handler.Handle(command)
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		response := NewCreateExpenseRespons(*expense)

		return c.JSON(http.StatusCreated, response)
	})

	return ExpenseAPI{app: e, Address: address}
}
