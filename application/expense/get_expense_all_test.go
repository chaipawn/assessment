package expense_test

import (
	"testing"

	"github.com/chaipawn/assessment/application/expense"
	"github.com/chaipawn/assessment/domain"
)

type fakeGetAllExpenseQueryRepository struct {
	Expenses []domain.Expense
}

func (repository *fakeGetAllExpenseQueryRepository) ReadAll() ([]domain.Expense, error) {
	return repository.Expenses, nil
}

func TestGetAllExpense(t *testing.T) {
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
	repository := &fakeGetAllExpenseQueryRepository{Expenses: expenses}
	handler := expense.NewGetAllExpenseHandler(repository)
	query := expense.NewGetAllExpenseQuery()

	expenseEntities, _ := handler.Handle(query)

	if len(expenseEntities) != len(repository.Expenses) {
		t.Errorf("expenses count expect %d, but got %d", len(repository.Expenses), len(expenseEntities))
	}
}
