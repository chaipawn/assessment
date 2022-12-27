package expense

import (
	"github.com/chaipawn/assessment/application/persistence"
	"github.com/chaipawn/assessment/domain"
)

type GetExpenseQuery struct {
	Id int
}

func NewGetExpenseQuery(id int) GetExpenseQuery {
	return GetExpenseQuery{Id: id}
}

type GetExpenseHandler struct {
	repository persistence.GetExpenseQueryRepository
}

func NewGetExpenseHandler(repository persistence.GetExpenseQueryRepository) GetExpenseHandler {
	return GetExpenseHandler{repository: repository}
}

func (handler GetExpenseHandler) Handle(query GetExpenseQuery) (*domain.Expense, error) {
	expenseId := domain.NewExpenseId(query.Id)
	expense, err := handler.repository.Read(expenseId)
	if err != nil {
		return nil, err
	}

	if expense == nil {
		return nil, ErrorExpenseNotFound{Id: expenseId}
	}

	return expense, nil
}
