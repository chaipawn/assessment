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

func (repository ExpenseCommandRepository) Read(id domain.ExpenseId) (*domain.Expense, error) {
	stmt, err := repository.db.Prepare("SELECT title, amount, note, tags FROM expenses WHERE id=$1")
	if err != nil {
		return nil, err
	}

	var (
		title  string
		amount float64
		note   string
		tags   []string
	)
	err = stmt.QueryRow(id.Value()).Scan(&title, &amount, &note, pq.Array(&tags))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	expense := domain.NewExpense(
		id,
		domain.NewExpenseTitle(title),
		domain.NewExpenseAmount(amount),
		domain.NewExpenseNote(note),
		domain.NewExpenseTags(tags...),
	)

	return &expense, nil
}

func (repository ExpenseCommandRepository) Update(expense domain.Expense) (*domain.Expense, error) {
	tags := make([]string, 0, len(expense.Tags().Value()))
	for _, tag := range expense.Tags().Value() {
		tags = append(tags, tag.Value())
	}

	_, err := repository.db.Exec(
		`
			UPDATE expenses 
			SET title = $2, amount = $3, note = $4, tags = $5
			WHERE id = $1
		`,
		expense.Id().Value(),
		expense.Title().Value(),
		expense.Amount().Value(),
		expense.Note().Value(),
		pq.Array(&tags),
	)
	if err != nil {
		return nil, err
	}

	return repository.Read(expense.Id())
}

func NewExpenseCommandRepository(db *sql.DB) ExpenseCommandRepository {
	return ExpenseCommandRepository{db: db}
}

type ExpenseQueryRepository struct {
	db *sql.DB
}

func (repository ExpenseQueryRepository) Read(id domain.ExpenseId) (*domain.Expense, error) {
	stmt, err := repository.db.Prepare("SELECT title, amount, note, tags FROM expenses WHERE id=$1")
	if err != nil {
		return nil, err
	}

	var (
		title  string
		amount float64
		note   string
		tags   []string
	)
	err = stmt.QueryRow(id.Value()).Scan(&title, &amount, &note, pq.Array(&tags))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	expense := domain.NewExpense(
		id,
		domain.NewExpenseTitle(title),
		domain.NewExpenseAmount(amount),
		domain.NewExpenseNote(note),
		domain.NewExpenseTags(tags...),
	)

	return &expense, nil
}

func NewExpenseQueryRepository(db *sql.DB) ExpenseQueryRepository {
	return ExpenseQueryRepository{db: db}
}
