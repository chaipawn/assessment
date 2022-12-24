package domain

type ExpenseNote struct {
	value string
}

func (note ExpenseNote) Value() string {
	return note.value
}

func NewExpenseNote(note string) ExpenseNote {
	return ExpenseNote{value: note}
}
