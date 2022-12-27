package expense

import (
	"github.com/chaipawn/assessment/application/persistence"
	"github.com/chaipawn/assessment/domain"
)

type GetAllExpenseQuery struct{}

func NewGetAllExpenseQuery() GetAllExpenseQuery {
	return GetAllExpenseQuery{}
}

type GetAllExpenseHandler struct {
	repository persistence.GetAllExpenseQueryRepository
}

func NewGetAllExpenseHandler(repository persistence.GetAllExpenseQueryRepository) GetAllExpenseHandler {
	return GetAllExpenseHandler{repository: repository}
}

func (handler GetAllExpenseHandler) Handle(query GetAllExpenseQuery) ([]domain.Expense, error) {
	expenses, err := handler.repository.ReadAll()
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
