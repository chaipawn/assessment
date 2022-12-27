package webapi

import "github.com/chaipawn/assessment/domain"

func createTagsResponse(tags domain.ExpenseTags) []string {
	responseTags := make([]string, 0, len(tags.Value()))
	for _, tag := range tags.Value() {
		responseTags = append(responseTags, tag.Value())
	}
	return responseTags
}

type CreateExpenseResponse struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func NewCreateExpenseRespons(expense domain.Expense) CreateExpenseResponse {
	tags := createTagsResponse(expense.Tags())

	return CreateExpenseResponse{
		Id:     expense.Id().Value(),
		Title:  expense.Title().Value(),
		Amount: expense.Amount().Value(),
		Note:   expense.Note().Value(),
		Tags:   tags,
	}
}

type GetExpenseResponse struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func NewGetExpenseResponse(expense domain.Expense) GetExpenseResponse {
	tags := createTagsResponse(expense.Tags())

	return GetExpenseResponse{
		Id:     expense.Id().Value(),
		Title:  expense.Title().Value(),
		Amount: expense.Amount().Value(),
		Note:   expense.Note().Value(),
		Tags:   tags,
	}
}
