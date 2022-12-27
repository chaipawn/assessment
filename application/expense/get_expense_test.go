package expense_test

import (
	"testing"

	"github.com/chaipawn/assessment/application/expense"
	"github.com/chaipawn/assessment/domain"
)

type fakeGetExpenseQueryRepository struct {
	Expenses []domain.Expense
}

func (repository *fakeGetExpenseQueryRepository) Read(id domain.ExpenseId) (*domain.Expense, error) {
	for _, expense := range repository.Expenses {
		if expense.Id() == id {
			return &expense, nil
		}
	}
	return nil, nil
}

func TestGetExpense(t *testing.T) {
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
	repository := &fakeGetExpenseQueryRepository{Expenses: expenses}
	handler := expense.NewGetExpenseHandler(repository)
	expectId := 2
	query := expense.NewGetExpenseQuery(expectId)

	expenseEntity, _ := handler.Handle(query)

	if expenseEntity.Id().Value() != expectId {
		t.Errorf("Expense id expect %d, but got %d", expectId, expenseEntity.Id().Value())
	}

	if expenseEntity.Title().Value() != "Expense title 002" {
		t.Errorf("Expense title expect Expense title 0002, but got %s", expenseEntity.Title().Value())
	}

	if expenseEntity.Amount().Value() != float64(500) {
		t.Errorf("Expense amount expect %f, but got %f", float64(500), expenseEntity.Amount().Value())
	}

	if expenseEntity.Note().Value() != "Expense note 002" {
		t.Errorf("Expense note expect Expense note 002, but got %s", expenseEntity.Note().Value())
	}

	if len(expenseEntity.Tags().Value()) != 1 {
		t.Errorf("Expense tags count expect %d, but got %d", 1, len(expenseEntity.Tags().Value()))
	}

	if expenseEntity.Tags().Value()[0].Value() != "TagC" {
		t.Errorf("Expense tag expect TagC, but got %s", expenseEntity.Tags().Value()[0].Value())
	}
}

func TestGetExpenseNotFound(t *testing.T) {
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
	repository := &fakeGetExpenseQueryRepository{Expenses: expenses}
	handler := expense.NewGetExpenseHandler(repository)
	expectId := 4
	query := expense.NewGetExpenseQuery(expectId)
	expectError := expense.ErrorExpenseNotFound{Id: domain.NewExpenseId(expectId)}

	expenseEntity, err := handler.Handle(query)

	if expenseEntity != nil {
		t.Error("Expense expect nil")
	}

	if err == nil {
		t.Error("Expense expect error, but got nil")
	}

	if err.Error() != expectError.Error() {
		t.Errorf("Expense expect error not found, but got %s", err)
	}
}
