package expense

import (
	"fmt"

	"github.com/chaipawn/assessment/domain"
)

type ErrorExpenseNotFound struct {
	Id domain.ExpenseId
}

func (err ErrorExpenseNotFound) Error() string {
	return fmt.Sprintf("Expense Id %d not found", err.Id.Value())
}
