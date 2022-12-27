package webapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func getExpenseById(db *sql.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("Expense id %s is invalid", c.Param("id"))})
		}

		query := expense.NewGetExpenseQuery(id)
		repository := infrastructure.NewExpenseQueryRepository(db)
		handler := expense.NewGetExpenseHandler(repository)

		expenseEntity, err := handler.Handle(query)
		if err != nil {
			c.Logger().Error(err)
			var errorNotFound expense.ErrorExpenseNotFound
			if errors.As(err, &errorNotFound) {
				return c.JSON(http.StatusNotFound, ErrorResponse{Message: errorNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		}

		response := NewGetExpenseResponse(*expenseEntity)

		return c.JSON(http.StatusOK, response)
	}
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

	e.GET("/expenses/:id", getExpenseById(db))

	return ExpenseAPI{app: e, Address: address}
}
