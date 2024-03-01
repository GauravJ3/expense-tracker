// Simple Expense Tracker CLI Application
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Category string
	Value    float64
	Date     time.Time
}

// define an array of expenses
var expenses []Expense

func addExpense(category string, value float64, date time.Time) {
	expense := Expense{Category: category, Value: value, Date: date}
	expenses = append(expenses, expense)

	fmt.Print("Expense added !! ", expenses)
}

func showExpenseByCategory(category string) {
	// iterate through the expense array and showcase by category
}

func summarizeExpenseByDate(fromDate time.Time, toDate time.Time) {
	// Basically get all expenses fromDate - toDate
	// summarize it.
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter expense Type, Amount and date of expense to track.")
	input, _ := reader.ReadString('\n')

	parts := strings.Split(strings.TrimSpace(input), ",")
	category := parts[0]
	value, _ := strconv.ParseFloat(parts[1], 64)
	date, _ := time.Parse("2006-01-02", parts[2])

	addExpense(category, value, date)
}
