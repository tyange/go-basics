package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		showErrorMessage(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		showErrorMessage(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		showErrorMessage(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResult(ebt, profit, ratio)
}

func showErrorMessage(err error) {
	fmt.Println("ERROR")
	fmt.Println(err)
	fmt.Println("-----------")
}

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("invalid amount")
	}

	return userInput, nil
}
