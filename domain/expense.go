package domain

type Expense struct {
	id     ExpenseId
	title  ExpenseTitle
	amount ExpenseAmount
	note   ExpenseNote
	tags   ExpenseTags
}

func NewExpense(id ExpenseId, title ExpenseTitle, amount ExpenseAmount, note ExpenseNote, tags ExpenseTags) Expense {
	return Expense{
		id:     id,
		title:  title,
		amount: amount,
		note:   note,
		tags:   tags,
	}
}

func (expense Expense) Id() ExpenseId {
	return expense.id
}

func (expense Expense) Title() ExpenseTitle {
	return expense.title
}

func (expense Expense) Amount() ExpenseAmount {
	return expense.amount
}

func (expense Expense) Note() ExpenseNote {
	return expense.note
}

func (expense Expense) Tags() ExpenseTags {
	return expense.tags
}

func (expense Expense) SetTitle(title ExpenseTitle) Expense {
	return NewExpense(expense.id, title, expense.amount, expense.note, expense.tags)
}

func (expense Expense) SetAmount(amount ExpenseAmount) Expense {
	return NewExpense(expense.id, expense.title, amount, expense.note, expense.tags)
}

func (expense Expense) SetNote(note ExpenseNote) Expense {
	return NewExpense(expense.id, expense.title, expense.amount, note, expense.tags)
}

func (expense Expense) SetTags(tags ExpenseTags) Expense {
	return NewExpense(expense.id, expense.title, expense.amount, expense.note, tags)
}
