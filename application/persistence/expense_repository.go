package persistence

import "github.com/chaipawn/assessment/domain"

type AddExpenseCommandRepository interface {
	Create(expense domain.Expense) (*domain.Expense, error)
}

type GetExpenseQueryRepository interface {
	Read(id domain.ExpenseId) (*domain.Expense, error)
}

type UpdateExpenseCommandRepository interface {
	Read(id domain.ExpenseId) (*domain.Expense, error)
	Update(expense domain.Expense) (*domain.Expense, error)
}
