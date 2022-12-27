package expense_test

import (
	"testing"

	"github.com/chaipawn/assessment/application/expense"
	"github.com/chaipawn/assessment/domain"
)

type fakeUpdateExpenseCommandRepository struct {
	Expenses []domain.Expense
}

func (repository *fakeUpdateExpenseCommandRepository) Read(id domain.ExpenseId) (*domain.Expense, error) {
	for _, expense := range repository.Expenses {
		if expense.Id() == id {
			return &expense, nil
		}
	}
	return nil, nil
}

func (repository *fakeUpdateExpenseCommandRepository) Update(expense domain.Expense) (*domain.Expense, error) {
	foundIndex := -1
	for index, existExpense := range repository.Expenses {
		if existExpense.Id() == expense.Id() {
			foundIndex = index
			break
		}
	}
	repository.Expenses[foundIndex] = expense
	return &expense, nil
}

func TestUpdateExpense(t *testing.T) {
	expenses := []domain.Expense{
		domain.NewExpense(
			domain.NewExpenseId(1),
			domain.NewExpenseTitle("Expense title 001"),
			domain.NewExpenseAmount(2000),
			domain.NewExpenseNote("Expense note 001"),
			domain.NewExpenseTags([]string{"TagA", "TagB"}...),
		),
		domain.NewExpense(
			domain.NewExpenseId(2),
			domain.NewExpenseTitle("Expense title 002"),
			domain.NewExpenseAmount(500),
			domain.NewExpenseNote("Expense note 002"),
			domain.NewExpenseTags([]string{"TagC"}...),
		),
		domain.NewExpense(
			domain.NewExpenseId(3),
			domain.NewExpenseTitle("Expense title 003"),
			domain.NewExpenseAmount(305),
			domain.NewExpenseNote("Expense note 003"),
			domain.NewExpenseTags([]string{"TagB", "TagC", "TagD"}...),
		),
	}
	repository := &fakeUpdateExpenseCommandRepository{Expenses: expenses}
	handler := expense.NewUpdateExpenseHandler(repository)
	command := expense.NewUpdateExpenseCommand(3, "Expense title 003 updated", 3000, "Expense note 003 updated", []string{"TagX", "TagY"})

	expense, _ := handler.Handle(command)

	if expense.Id().Value() != 3 {
		t.Errorf("expense id expect %d, but got %d", 3, expense.Id().Value())
	}

	if expense.Title().Value() != command.Title {
		t.Errorf("expense title expect %s, but got %s", command.Title, expense.Title().Value())
	}

	if expense.Amount().Value() != command.Amount {
		t.Errorf("expense amount expect %f, but got %f", command.Amount, expense.Amount().Value())
	}

	if expense.Note().Value() != command.Note {
		t.Errorf("expense note expect %s, but got %s", command.Note, expense.Note().Value())
	}

	if len(expense.Tags().Value()) != len(command.Tags) {
		t.Errorf("expense tags count exect %d, but got %d", len(command.Tags), len(expense.Tags().Value()))
	}

	for index, tag := range expense.Tags().Value() {
		if tag.Value() != command.Tags[index] {
			t.Errorf("expense tag %d expect %s, but got %s", index, command.Tags[index], tag.Value())
		}
	}
}
