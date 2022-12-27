package expense

import (
	"github.com/chaipawn/assessment/application/persistence"
	"github.com/chaipawn/assessment/domain"
)

type UpdateExpenseCommand struct {
	Id     int
	Title  string
	Amount float64
	Note   string
	Tags   []string
}

func NewUpdateExpenseCommand(id int, title string, amount float64, note string, tags []string) UpdateExpenseCommand {
	return UpdateExpenseCommand{Id: id, Title: title, Amount: amount, Note: note, Tags: tags}
}

type UpdateExpenseHandler struct {
	repository persistence.UpdateExpenseCommandRepository
}

func NewUpdateExpenseHandler(repository persistence.UpdateExpenseCommandRepository) UpdateExpenseHandler {
	return UpdateExpenseHandler{repository: repository}
}

func (handler UpdateExpenseHandler) Handle(command UpdateExpenseCommand) (*domain.Expense, error) {
	expensId := domain.NewExpenseId(command.Id)
	existExpense, err := handler.repository.Read(expensId)
	if err != nil {
		return nil, err
	}
	if existExpense == nil {
		return nil, ErrorExpenseNotFound{Id: expensId}
	}

	newExpense := domain.NewExpense(
		expensId,
		domain.NewExpenseTitle(command.Title),
		domain.NewExpenseAmount(command.Amount),
		domain.NewExpenseNote(command.Note),
		domain.NewExpenseTags(command.Tags...),
	)
	updatedExpense, err := handler.repository.Update(newExpense)
	if err != nil {
		return nil, err
	}

	return updatedExpense, nil
}
