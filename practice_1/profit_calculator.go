package main

import "fmt"

func main() {
	var revenue float64
	var exenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&exenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - exenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
