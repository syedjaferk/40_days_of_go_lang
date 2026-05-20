package main

import "fmt"

type Expense struct {
	Description string
	Amount      float64
	Category    string
	Date        string
}

var expenses = make(map[int]Expense)

func addExpense(id int, description string, amount float64,
	category string, date string) {

	expenses[id] = Expense{
		Description: description,
		Amount:      amount,
		Category:    category,
		Date:        date,
	}

	fmt.Println("Expense Added")

}

func viewExpenses() {

	for id, expense := range expenses {
		fmt.Println("ID ", id)
		fmt.Println(expense)
		fmt.Println("----------------------------")

	}
}

func updateExpense(id int, newAmount float64) {

	expense, exists := expenses[id]

	if !exists {
		fmt.Println("Expense Not Found")
		return
	}

	expense.Amount = newAmount
	fmt.Println("Expense Updated")
}

func deleteExpense(id int) {

	_, exists := expenses[id]

	if !exists {
		fmt.Println("Expense Not Found")
		return
	}

	delete(expenses, id)

	fmt.Println("Expense Deleted")
}

func totalExpense() {
	total := 0.0

	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Println("\n Total Expense ", total)
}

func categoryWiseExpense() {
	categoryTotals := make(map[string]float64)

	for _, expense := range expenses {
		category := expense.Category
		amount := expense.Amount

		categoryTotals[category] = categoryTotals[category] + amount
	}

	fmt.Println("------- Category Totals --------")

	for category, total := range categoryTotals {
		fmt.Println(category, "=>", total)
	}

}

func main() {
	// Runner
	addExpense(1, "Aavin Blue packet, Tea", 24, "Food", "2025-05-12")
	addExpense(2, "Dosa maavu", 30, "Food", "2025-05-14")
	addExpense(3, "Petrol tank fill", 300, "Travel", "2025-05-14")
	viewExpenses()
	updateExpense(1, 200)
	viewExpenses()
	// deleteExpense(1)
	viewExpenses()
	totalExpense()
	categoryWiseExpense()

}
