package application

import (
	"github.com/chaipawn/assessment/application/persistence"
	"github.com/chaipawn/assessment/domain"
)

type AddExpenseHandler struct {
	repository persistence.AddExpenseCommandRepository
}

func NewAddExpenseHandler(repository persistence.AddExpenseCommandRepository) AddExpenseHandler {
	return AddExpenseHandler{repository: repository}
}

func (handler AddExpenseHandler) Handle(command AddExpenseCommand) domain.Expense {
	expense := domain.NewExpense(
		domain.NewExpenseId(0),
		domain.NewExpenseTitle(command.Title),
		domain.NewExpenseAmount(command.Amount),
		domain.NewExpenseNote(command.Note),
		domain.NewExpenseTags(command.Tags...),
	)
	newExpense := handler.repository.Create(expense)
	return newExpense
}
