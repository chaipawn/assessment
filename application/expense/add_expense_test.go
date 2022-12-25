package expense_test

import (
	"testing"

	"github.com/chaipawn/assessment/application/expense"
	"github.com/chaipawn/assessment/domain"
)

type fakeRepository struct {
	Expenses []domain.Expense
}

func (repository *fakeRepository) Create(expense domain.Expense) domain.Expense {
	repository.Expenses = append(repository.Expenses, expense)
	return domain.NewExpense(domain.NewExpenseId(len(repository.Expenses)), expense.Title(), expense.Amount(), expense.Note(), expense.Tags())
}

func TestAddExpense(t *testing.T) {
	repository := &fakeRepository{Expenses: make([]domain.Expense, 0, 1)}
	handler := expense.NewAddExpenseHandler(repository)
	command := expense.NewAddExpenseCommand("strawberry smoothie", 79, "night market promotion discount 10 bath", []string{"food", "beverage"})
	expectItemCount := 1
	expectExpenseId := 1

	expense := handler.Handle(command)

	if len(repository.Expenses) != expectItemCount {
		t.Errorf("expense count in repository expect %d, but got %d", expectItemCount, len(repository.Expenses))
	}

	if expense.Id().Value() != expectExpenseId {
		t.Errorf("expense id expect %d, but got %d", expectExpenseId, expense.Id().Value())
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
