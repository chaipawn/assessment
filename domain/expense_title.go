package domain

type ExpenseTitle struct {
	value string
}

func (title ExpenseTitle) Value() string {
	return title.value
}

func NewExpenseTitle(title string) ExpenseTitle {
	return ExpenseTitle{value: title}
}
