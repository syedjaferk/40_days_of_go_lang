package main 

import "fmt"

var expenses = make(map[int]map[string]interface{})

func addExpense(id int, description string, amount float64,
category string, date string){


	expenses[id] = map[string]interface{}{
		"description": description,
		"amount": amount,
		"category": category,
		"date": date,
	}

	fmt.Println("Expense Added")

}

func viewExpenses(){

	for id, expense := range expenses {
		fmt.Println("ID ", id)
		fmt.Println(expense)
		fmt.Println("----------------------------")

	}
}

func updateExpense(id int, newAmount float64){

	expense, exists := expenses[id]

	if !exists {
		fmt.Println("Expense Not Found")
		return
	}

	expense["amount"] = newAmount
	fmt.Println("Expense Updated")
}

func deleteExpense(id int){

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
		total += expense["amount"].(float64)
	}

	fmt.Println("\n Total Expense ", total)
}

func categoryWiseExpense(){
	categoryTotals := make(map[string]float64)

	for _, expense := range expenses {
		category := expense["category"].(string)
		amount := expense["amount"].(float64)

		categoryTotals[category] = categoryTotals[category] + amount
	}

	fmt.Println("------- Category Totals --------")

	for category, total := range categoryTotals{
		fmt.Println(category, "=>", total)
	}


}

func main(){
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