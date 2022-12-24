package domain

type ExpenseId struct {
	value int
}

func (id ExpenseId) Value() int {
	return id.value
}

func NewExpenseId(id int) ExpenseId {
	return ExpenseId{value: id}
}
