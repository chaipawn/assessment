package expense

import (
	"github.com/chaipawn/assessment/application/persistence"
	"github.com/chaipawn/assessment/domain"
)

type AddExpenseCommand struct {
	Title  string
	Amount float64
	Note   string
	Tags   []string
}

func NewAddExpenseCommand(title string, amount float64, note string, tags []string) AddExpenseCommand {
	return AddExpenseCommand{Title: title, Amount: amount, Note: note, Tags: tags}
}

type AddExpenseHandler struct {
	repository persistence.AddExpenseCommandRepository
}

func NewAddExpenseHandler(repository persistence.AddExpenseCommandRepository) AddExpenseHandler {
	return AddExpenseHandler{repository: repository}
}

func (handler AddExpenseHandler) Handle(command AddExpenseCommand) (*domain.Expense, error) {
	expense := domain.NewExpense(
		domain.NewExpenseId(0),
		domain.NewExpenseTitle(command.Title),
		domain.NewExpenseAmount(command.Amount),
		domain.NewExpenseNote(command.Note),
		domain.NewExpenseTags(command.Tags...),
	)
	newExpense, err := handler.repository.Create(expense)
	if err != nil {
		return nil, err
	}

	return newExpense, nil
}
