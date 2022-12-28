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

func createExpense(db *sql.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var request CreateExpenseRequest
		err := c.Bind(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Bad request"})
		}

		command := expense.NewAddExpenseCommand(request.Title, request.Amount, request.Note, request.Tags)
		repository := infrastructure.NewExpenseCommandRepository(db)
		handler := expense.NewAddExpenseHandler(repository)

		expenseEntity, err := handler.Handle(command)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		}
		response := NewCreateExpenseRespons(*expenseEntity)

		return c.JSON(http.StatusCreated, response)
	}
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

func updateExpense(db *sql.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var request UpdateExpenseRequest
		err := c.Bind(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Bad request"})
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("Expense id %s is invalid", c.Param("id"))})
		}

		command := expense.NewUpdateExpenseCommand(id, request.Title, request.Amount, request.Note, request.Tags)
		repository := infrastructure.NewExpenseCommandRepository(db)
		handler := expense.NewUpdateExpenseHandler(repository)

		expenseEntity, err := handler.Handle(command)
		if err != nil {
			c.Logger().Error(err)
			var errorNotFound expense.ErrorExpenseNotFound
			if errors.As(err, &errorNotFound) {
				return c.JSON(http.StatusNotFound, ErrorResponse{Message: errorNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		}
		response := NewUpdateExpenseResponse(*expenseEntity)

		return c.JSON(http.StatusOK, response)
	}
}

func getAllExpense(db *sql.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		query := expense.NewGetAllExpenseQuery()
		repository := infrastructure.NewExpenseQueryRepository(db)
		handler := expense.NewGetAllExpenseHandler(repository)

		expenseEntities, err := handler.Handle(query)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		}

		responses := make([]GetExpenseResponse, 0, len(expenseEntities))
		for _, entity := range expenseEntities {
			response := NewGetExpenseResponse(entity)
			responses = append(responses, response)
		}

		return c.JSON(http.StatusOK, responses)
	}
}

func NewExpenseAPI(address string, db *sql.DB) ExpenseAPI {
	e := echo.New()

	e.GET("/expenses", Authorize(getAllExpense(db)))
	e.POST("/expenses", Authorize(createExpense(db)))
	e.GET("/expenses/:id", Authorize(getExpenseById(db)))
	e.PUT("/expenses/:id", Authorize(updateExpense(db)))

	return ExpenseAPI{app: e, Address: address}
}
