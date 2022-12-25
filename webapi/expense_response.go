package webapi

import "github.com/chaipawn/assessment/domain"

type CreateExpenseResponse struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func NewCreateExpenseRespons(expense domain.Expense) CreateExpenseResponse {
	tags := make([]string, 0, len(expense.Tags().Value()))
	for _, tag := range expense.Tags().Value() {
		tags = append(tags, tag.Value())
	}

	return CreateExpenseResponse{
		Id:     expense.Id().Value(),
		Title:  expense.Title().Value(),
		Amount: expense.Amount().Value(),
		Note:   expense.Note().Value(),
		Tags:   tags,
	}
}
