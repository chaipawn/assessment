package infrastructure

import (
	"database/sql"

	"github.com/chaipawn/assessment/domain"
	"github.com/lib/pq"
)

type ExpenseCommandRepository struct {
	db *sql.DB
}

func (repository ExpenseCommandRepository) Create(expense domain.Expense) (*domain.Expense, error) {
	tags := make([]string, 0, len(expense.Tags().Value()))
	for _, tag := range expense.Tags().Value() {
		tags = append(tags, tag.Value())
	}

	lastInsertId := 0
	err := repository.db.QueryRow(
		"INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id",
		expense.Title().Value(),
		expense.Amount().Value(),
		expense.Note().Value(),
		pq.Array(&tags),
	).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	newExpense := domain.NewExpense(
		domain.NewExpenseId(lastInsertId),
		expense.Title(),
		expense.Amount(),
		expense.Note(),
		expense.Tags(),
	)

	return &newExpense, nil
}

func NewExpenseCommandRepository(db *sql.DB) ExpenseCommandRepository {
	return ExpenseCommandRepository{db: db}
}
