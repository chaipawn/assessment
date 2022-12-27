package persistence

import "github.com/chaipawn/assessment/domain"

type AddExpenseCommandRepository interface {
	Create(expense domain.Expense) (*domain.Expense, error)
}

type GetExpenseQueryRepository interface {
	Read(id domain.ExpenseId) (*domain.Expense, error)
}
