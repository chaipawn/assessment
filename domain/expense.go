package domain

type Expense struct {
	id     ExpenseId
	title  ExpenseTitle
	amount ExpenseAmount
	note   ExpenseNote
	tags   ExpenseTags
}

func NewExpense(id int, title string, amount float64, note string, tags []string) Expense {
	return Expense{
		id:     NewExpenseId(id),
		title:  NewExpenseTitle(title),
		amount: NewExpenseAmount(amount),
		note:   NewExpenseNote(note),
		tags:   NewExpenseTags(tags...),
	}
}
