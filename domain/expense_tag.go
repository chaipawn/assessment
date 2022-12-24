package domain

type ExpenseTag struct {
	value string
}

func (tag ExpenseTag) Value() string {
	return tag.value
}

func NewExpenseTag(tag string) ExpenseTag {
	return ExpenseTag{value: tag}
}

type ExpenseTags struct {
	values []ExpenseTag
}

func (tags ExpenseTags) Value() []ExpenseTag {
	return tags.values
}

func NewExpenseTags(tags ...string) ExpenseTags {
	expense_tags := make([]ExpenseTag, len(tags))
	for _, tag := range tags {
		expense_tags = append(expense_tags, NewExpenseTag(tag))
	}
	return ExpenseTags{values: expense_tags}
}
