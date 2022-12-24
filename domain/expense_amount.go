package domain

type ExpenseAmount struct {
	value float64
}

func (amount ExpenseAmount) Value() float64 {
	return amount.value
}

func NewExpenseAmount(amount float64) ExpenseAmount {
	return ExpenseAmount{value: amount}
}
