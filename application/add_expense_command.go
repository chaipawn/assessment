package application

type AddExpenseCommand struct {
	Title  string
	Amount float64
	Note   string
	Tags   []string
}

func NewAddExpenseCommand(title string, amount float64, note string, tags []string) AddExpenseCommand {
	return AddExpenseCommand{Title: title, Amount: amount, Note: note, Tags: tags}
}
